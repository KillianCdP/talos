// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package system_test

import (
	"errors"
	"testing"
	"time"

	"github.com/siderolabs/go-retry/retry"
	"github.com/stretchr/testify/suite"

	"github.com/siderolabs/talos/internal/app/machined/pkg/system"
	"github.com/siderolabs/talos/internal/app/machined/pkg/system/events"
	"github.com/siderolabs/talos/pkg/conditions"
)

type ServiceRunnerSuite struct {
	suite.Suite
}

func (suite *ServiceRunnerSuite) assertStateSequence(expectedStates []events.ServiceState, sr *system.ServiceRunner) {
	states := []events.ServiceState{}

	for _, event := range sr.GetEventHistory(1000) {
		states = append(states, event.State)
	}

	suite.Assert().Equal(expectedStates, states)
}

func (suite *ServiceRunnerSuite) TestFullFlow() {
	sr := system.NewServiceRunner(&MockService{
		condition: conditions.None(),
	}, nil)

	errCh := make(chan error)

	go func() {
		errCh <- sr.Run()
	}()

	suite.Require().NoError(retry.Constant(time.Minute, retry.WithUnits(10*time.Millisecond)).Retry(func() error {
		state := sr.AsProto().State
		if state != events.StateRunning.String() {
			return retry.ExpectedErrorf("service should be running")
		}

		return nil
	}))

	select {
	case <-errCh:
		suite.Require().Fail("service running should be still running")
	default:
	}

	sr.Shutdown()

	suite.Assert().NoError(<-errCh)

	suite.assertStateSequence([]events.ServiceState{
		events.StateWaiting,
		events.StatePreparing,
		events.StatePreparing,
		events.StateRunning,
	}, sr)

	protoService := sr.AsProto()
	suite.Assert().Equal("MockRunner", protoService.Id)
	suite.Assert().Equal("Running", protoService.State)
	suite.Assert().True(protoService.Health.Unknown)
	suite.Assert().Len(protoService.Events.Events, 4)
}

func (suite *ServiceRunnerSuite) TestFullFlowHealthy() {
	sr := system.NewServiceRunner(&MockHealthcheckedService{}, nil)

	errCh := make(chan error)

	go func() {
		errCh <- sr.Run()
	}()

	suite.Require().NoError(retry.Constant(time.Minute, retry.WithUnits(10*time.Millisecond)).Retry(func() error {
		health := sr.AsProto().Health
		if health.Unknown || !health.Healthy {
			return retry.ExpectedErrorf("service should be healthy")
		}

		return nil
	}))

	select {
	case <-errCh:
		suite.Require().Fail("service running should be still running")
	default:
	}

	sr.Shutdown()

	suite.Assert().NoError(<-errCh)

	suite.assertStateSequence([]events.ServiceState{
		events.StatePreparing,
		events.StatePreparing,
		events.StateRunning,
		events.StateRunning, // one more notification when service is healthy
	}, sr)
}

func (suite *ServiceRunnerSuite) TestFullFlowHealthChanges() {
	m := MockHealthcheckedService{
		MockService: MockService{
			condition: conditions.None(),
		},
	}
	sr := system.NewServiceRunner(&m, nil)

	errCh := make(chan error)

	go func() {
		errCh <- sr.Run()
	}()

	suite.Require().NoError(retry.Constant(time.Minute, retry.WithUnits(10*time.Millisecond)).Retry(func() error {
		health := sr.AsProto().Health
		if health.Unknown || !health.Healthy {
			return retry.ExpectedErrorf("service should be healthy")
		}

		return nil
	}))

	m.SetHealthy(false)

	suite.Require().NoError(retry.Constant(time.Minute, retry.WithUnits(10*time.Millisecond)).Retry(func() error {
		health := sr.AsProto().Health
		if health.Unknown || health.Healthy {
			return retry.ExpectedErrorf("service should be not healthy")
		}

		return nil
	}))

	m.SetHealthy(true)

	suite.Require().NoError(retry.Constant(time.Minute, retry.WithUnits(10*time.Millisecond)).Retry(func() error {
		health := sr.AsProto().Health
		if health.Unknown || !health.Healthy {
			return retry.ExpectedErrorf("service should be healthy")
		}

		return nil
	}))

	sr.Shutdown()

	suite.Assert().NoError(<-errCh)

	suite.assertStateSequence([]events.ServiceState{
		events.StateWaiting,
		events.StatePreparing,
		events.StatePreparing,
		events.StateRunning,
		events.StateRunning, // initial: healthy
		events.StateRunning, // not healthy
		events.StateRunning, // once again healthy
	}, sr)
}

func (suite *ServiceRunnerSuite) TestWaitingDescriptionChange() {
	oldWaitConditionCheckInterval := system.WaitConditionCheckInterval
	system.WaitConditionCheckInterval = 10 * time.Millisecond

	defer func() {
		system.WaitConditionCheckInterval = oldWaitConditionCheckInterval
	}()

	cond1 := NewMockCondition("cond1")
	cond2 := NewMockCondition("cond2")
	sr := system.NewServiceRunner(&MockService{
		condition: conditions.WaitForAll(cond1, cond2),
	}, nil)

	errCh := make(chan error)

	go func() {
		errCh <- sr.Run()
	}()

	suite.Require().NoError(retry.Constant(time.Minute, retry.WithUnits(10*time.Millisecond)).Retry(func() error {
		state := sr.AsProto().State
		if state != events.StateWaiting.String() {
			return retry.ExpectedErrorf("service should be waiting")
		}

		return nil
	}))

	select {
	case <-errCh:
		suite.Require().Fail("service running should be still running")
	default:
	}

	close(cond1.done)

	suite.Require().NoError(retry.Constant(time.Minute, retry.WithUnits(10*time.Millisecond)).Retry(func() error {
		events := sr.AsProto().Events.Events

		lastMsg := events[len(events)-1].Msg
		if lastMsg != "Waiting for cond2" {
			return retry.ExpectedErrorf("service should be waiting on 2nd condition")
		}

		return nil
	}))

	select {
	case <-errCh:
		suite.Require().Fail("service running should be still running")
	default:
	}

	close(cond2.done)

	suite.Require().NoError(retry.Constant(time.Minute, retry.WithUnits(10*time.Millisecond)).Retry(func() error {
		state := sr.AsProto().State
		if state != events.StateRunning.String() {
			return retry.ExpectedErrorf("service should be running")
		}

		return nil
	}))

	sr.Shutdown()

	suite.Assert().NoError(<-errCh)

	suite.assertStateSequence([]events.ServiceState{
		events.StateWaiting,
		events.StateWaiting,
		events.StatePreparing,
		events.StatePreparing,
		events.StateRunning,
	}, sr)

	events := sr.GetEventHistory(10000)
	suite.Assert().Equal("Waiting for cond1, cond2", events[0].Message)
	suite.Assert().Equal("Waiting for cond2", events[1].Message)
}

func (suite *ServiceRunnerSuite) TestPreStageFail() {
	svc := &MockService{
		preError: errors.New("pre failed"),
	}
	sr := system.NewServiceRunner(svc, nil)
	err := sr.Run()

	suite.assertStateSequence([]events.ServiceState{
		events.StatePreparing,
	}, sr)
	suite.Assert().EqualError(err, "failed to run pre stage: pre failed")
}

func (suite *ServiceRunnerSuite) TestRunnerStageFail() {
	svc := &MockService{
		runnerError: errors.New("runner failed"),
	}
	sr := system.NewServiceRunner(svc, nil)
	err := sr.Run()

	suite.assertStateSequence([]events.ServiceState{
		events.StatePreparing,
		events.StatePreparing,
	}, sr)
	suite.Assert().EqualError(err, "failed to create runner: runner failed")
}

func (suite *ServiceRunnerSuite) TestRunnerStageSkipped() {
	svc := &MockService{
		nilRunner: true,
	}
	sr := system.NewServiceRunner(svc, nil)
	err := sr.Run()

	suite.assertStateSequence([]events.ServiceState{
		events.StatePreparing,
		events.StatePreparing,
	}, sr)
	suite.Assert().ErrorIs(err, system.ErrSkip)
}

func (suite *ServiceRunnerSuite) TestAbortOnCondition() {
	svc := &MockService{
		condition: conditions.WaitForFileToExist("/doesntexistever"),
	}
	sr := system.NewServiceRunner(svc, nil)

	errCh := make(chan error)

	go func() {
		errCh <- sr.Run()
	}()

	suite.Require().NoError(retry.Constant(time.Minute, retry.WithUnits(10*time.Millisecond)).Retry(func() error {
		state := sr.AsProto().State
		if state != events.StateWaiting.String() {
			return retry.ExpectedErrorf("service should be waiting")
		}

		return nil
	}))

	select {
	case <-errCh:
		suite.Require().Fail("service running should be still running")
	default:
	}

	sr.Shutdown()

	suite.Assert().EqualError(<-errCh, "condition failed: context canceled")

	suite.assertStateSequence([]events.ServiceState{
		events.StateWaiting,
	}, sr)
}

func (suite *ServiceRunnerSuite) TestPostStateFail() {
	svc := &MockService{
		condition: conditions.None(),
		postError: errors.New("post failed"),
	}
	sr := system.NewServiceRunner(svc, nil)

	errCh := make(chan error)

	go func() {
		errCh <- sr.Run()
	}()

	sr.Shutdown()

	suite.Assert().NoError(<-errCh)

	suite.assertStateSequence([]events.ServiceState{
		events.StateWaiting,
		events.StatePreparing,
		events.StatePreparing,
		events.StateRunning,
		events.StateFailed,
	}, sr)
}

func (suite *ServiceRunnerSuite) TestRunFail() {
	runner := &MockRunner{exitCh: make(chan error)}
	svc := &MockService{runner: runner}
	sr := system.NewServiceRunner(svc, nil)

	errCh := make(chan error)

	go func() {
		errCh <- sr.Run()
	}()

	runner.exitCh <- errors.New("run failed")

	suite.Assert().EqualError(<-errCh, "failed running service: error running service: run failed")

	suite.assertStateSequence([]events.ServiceState{
		events.StatePreparing,
		events.StatePreparing,
		events.StateRunning,
	}, sr)
}

func (suite *ServiceRunnerSuite) TestFullFlowRestart() {
	sr := system.NewServiceRunner(&MockService{
		condition: conditions.None(),
	}, nil)

	errCh := make(chan error)

	go func() {
		errCh <- sr.Run()
	}()

	suite.Require().NoError(retry.Constant(time.Minute, retry.WithUnits(10*time.Millisecond)).Retry(func() error {
		state := sr.AsProto().State
		if state != events.StateRunning.String() {
			return retry.ExpectedErrorf("service should be running")
		}

		return nil
	}))

	select {
	case <-errCh:
		suite.Require().Fail("service running should be still running")
	default:
	}

	sr.Shutdown()

	suite.Assert().NoError(<-errCh)

	go func() {
		errCh <- sr.Run()
	}()

	suite.Require().NoError(retry.Constant(time.Minute, retry.WithUnits(10*time.Millisecond)).Retry(func() error {
		state := sr.AsProto().State
		if state != events.StateRunning.String() {
			return retry.ExpectedErrorf("service should be running")
		}

		return nil
	}))

	select {
	case <-errCh:
		suite.Require().Fail("service running should be still running")
	default:
	}

	sr.Shutdown()

	suite.Assert().NoError(<-errCh)

	suite.assertStateSequence([]events.ServiceState{
		events.StateWaiting,
		events.StatePreparing,
		events.StatePreparing,
		events.StateRunning,
		events.StateWaiting,
		events.StatePreparing,
		events.StatePreparing,
		events.StateRunning,
	}, sr)
}

func TestServiceRunnerSuite(t *testing.T) {
	suite.Run(t, new(ServiceRunnerSuite))
}
