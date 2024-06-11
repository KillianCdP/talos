// Code generated by "enumer -type=VolumeType,VolumePhase,FilesystemType,EncryptionKeyType,EncryptionProviderType -linecomment -text"; DO NOT EDIT.

package block

import (
	"fmt"
	"strings"
)

const _VolumeTypeName = "partitiondisktmpfs"

var _VolumeTypeIndex = [...]uint8{0, 9, 13, 18}

const _VolumeTypeLowerName = "partitiondisktmpfs"

func (i VolumeType) String() string {
	if i < 0 || i >= VolumeType(len(_VolumeTypeIndex)-1) {
		return fmt.Sprintf("VolumeType(%d)", i)
	}
	return _VolumeTypeName[_VolumeTypeIndex[i]:_VolumeTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _VolumeTypeNoOp() {
	var x [1]struct{}
	_ = x[VolumeTypePartition-(0)]
	_ = x[VolumeTypeDisk-(1)]
	_ = x[VolumeTypeTmpfs-(2)]
}

var _VolumeTypeValues = []VolumeType{VolumeTypePartition, VolumeTypeDisk, VolumeTypeTmpfs}

var _VolumeTypeNameToValueMap = map[string]VolumeType{
	_VolumeTypeName[0:9]:        VolumeTypePartition,
	_VolumeTypeLowerName[0:9]:   VolumeTypePartition,
	_VolumeTypeName[9:13]:       VolumeTypeDisk,
	_VolumeTypeLowerName[9:13]:  VolumeTypeDisk,
	_VolumeTypeName[13:18]:      VolumeTypeTmpfs,
	_VolumeTypeLowerName[13:18]: VolumeTypeTmpfs,
}

var _VolumeTypeNames = []string{
	_VolumeTypeName[0:9],
	_VolumeTypeName[9:13],
	_VolumeTypeName[13:18],
}

// VolumeTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func VolumeTypeString(s string) (VolumeType, error) {
	if val, ok := _VolumeTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _VolumeTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to VolumeType values", s)
}

// VolumeTypeValues returns all values of the enum
func VolumeTypeValues() []VolumeType {
	return _VolumeTypeValues
}

// VolumeTypeStrings returns a slice of all String values of the enum
func VolumeTypeStrings() []string {
	strs := make([]string, len(_VolumeTypeNames))
	copy(strs, _VolumeTypeNames)
	return strs
}

// IsAVolumeType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i VolumeType) IsAVolumeType() bool {
	for _, v := range _VolumeTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for VolumeType
func (i VolumeType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for VolumeType
func (i *VolumeType) UnmarshalText(text []byte) error {
	var err error
	*i, err = VolumeTypeString(string(text))
	return err
}

const _VolumePhaseName = "waitingfailedmissinglocatedprovisionedpreparedreadyclosed"

var _VolumePhaseIndex = [...]uint8{0, 7, 13, 20, 27, 38, 46, 51, 57}

const _VolumePhaseLowerName = "waitingfailedmissinglocatedprovisionedpreparedreadyclosed"

func (i VolumePhase) String() string {
	if i < 0 || i >= VolumePhase(len(_VolumePhaseIndex)-1) {
		return fmt.Sprintf("VolumePhase(%d)", i)
	}
	return _VolumePhaseName[_VolumePhaseIndex[i]:_VolumePhaseIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _VolumePhaseNoOp() {
	var x [1]struct{}
	_ = x[VolumePhaseWaiting-(0)]
	_ = x[VolumePhaseFailed-(1)]
	_ = x[VolumePhaseMissing-(2)]
	_ = x[VolumePhaseLocated-(3)]
	_ = x[VolumePhaseProvisioned-(4)]
	_ = x[VolumePhasePrepared-(5)]
	_ = x[VolumePhaseReady-(6)]
	_ = x[VolumePhaseClosed-(7)]
}

var _VolumePhaseValues = []VolumePhase{VolumePhaseWaiting, VolumePhaseFailed, VolumePhaseMissing, VolumePhaseLocated, VolumePhaseProvisioned, VolumePhasePrepared, VolumePhaseReady, VolumePhaseClosed}

var _VolumePhaseNameToValueMap = map[string]VolumePhase{
	_VolumePhaseName[0:7]:        VolumePhaseWaiting,
	_VolumePhaseLowerName[0:7]:   VolumePhaseWaiting,
	_VolumePhaseName[7:13]:       VolumePhaseFailed,
	_VolumePhaseLowerName[7:13]:  VolumePhaseFailed,
	_VolumePhaseName[13:20]:      VolumePhaseMissing,
	_VolumePhaseLowerName[13:20]: VolumePhaseMissing,
	_VolumePhaseName[20:27]:      VolumePhaseLocated,
	_VolumePhaseLowerName[20:27]: VolumePhaseLocated,
	_VolumePhaseName[27:38]:      VolumePhaseProvisioned,
	_VolumePhaseLowerName[27:38]: VolumePhaseProvisioned,
	_VolumePhaseName[38:46]:      VolumePhasePrepared,
	_VolumePhaseLowerName[38:46]: VolumePhasePrepared,
	_VolumePhaseName[46:51]:      VolumePhaseReady,
	_VolumePhaseLowerName[46:51]: VolumePhaseReady,
	_VolumePhaseName[51:57]:      VolumePhaseClosed,
	_VolumePhaseLowerName[51:57]: VolumePhaseClosed,
}

var _VolumePhaseNames = []string{
	_VolumePhaseName[0:7],
	_VolumePhaseName[7:13],
	_VolumePhaseName[13:20],
	_VolumePhaseName[20:27],
	_VolumePhaseName[27:38],
	_VolumePhaseName[38:46],
	_VolumePhaseName[46:51],
	_VolumePhaseName[51:57],
}

// VolumePhaseString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func VolumePhaseString(s string) (VolumePhase, error) {
	if val, ok := _VolumePhaseNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _VolumePhaseNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to VolumePhase values", s)
}

// VolumePhaseValues returns all values of the enum
func VolumePhaseValues() []VolumePhase {
	return _VolumePhaseValues
}

// VolumePhaseStrings returns a slice of all String values of the enum
func VolumePhaseStrings() []string {
	strs := make([]string, len(_VolumePhaseNames))
	copy(strs, _VolumePhaseNames)
	return strs
}

// IsAVolumePhase returns "true" if the value is listed in the enum definition. "false" otherwise
func (i VolumePhase) IsAVolumePhase() bool {
	for _, v := range _VolumePhaseValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for VolumePhase
func (i VolumePhase) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for VolumePhase
func (i *VolumePhase) UnmarshalText(text []byte) error {
	var err error
	*i, err = VolumePhaseString(string(text))
	return err
}

const _FilesystemTypeName = "nonexfs"

var _FilesystemTypeIndex = [...]uint8{0, 4, 7}

const _FilesystemTypeLowerName = "nonexfs"

func (i FilesystemType) String() string {
	if i < 0 || i >= FilesystemType(len(_FilesystemTypeIndex)-1) {
		return fmt.Sprintf("FilesystemType(%d)", i)
	}
	return _FilesystemTypeName[_FilesystemTypeIndex[i]:_FilesystemTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _FilesystemTypeNoOp() {
	var x [1]struct{}
	_ = x[FilesystemTypeNone-(0)]
	_ = x[FilesystemTypeXFS-(1)]
}

var _FilesystemTypeValues = []FilesystemType{FilesystemTypeNone, FilesystemTypeXFS}

var _FilesystemTypeNameToValueMap = map[string]FilesystemType{
	_FilesystemTypeName[0:4]:      FilesystemTypeNone,
	_FilesystemTypeLowerName[0:4]: FilesystemTypeNone,
	_FilesystemTypeName[4:7]:      FilesystemTypeXFS,
	_FilesystemTypeLowerName[4:7]: FilesystemTypeXFS,
}

var _FilesystemTypeNames = []string{
	_FilesystemTypeName[0:4],
	_FilesystemTypeName[4:7],
}

// FilesystemTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FilesystemTypeString(s string) (FilesystemType, error) {
	if val, ok := _FilesystemTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _FilesystemTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to FilesystemType values", s)
}

// FilesystemTypeValues returns all values of the enum
func FilesystemTypeValues() []FilesystemType {
	return _FilesystemTypeValues
}

// FilesystemTypeStrings returns a slice of all String values of the enum
func FilesystemTypeStrings() []string {
	strs := make([]string, len(_FilesystemTypeNames))
	copy(strs, _FilesystemTypeNames)
	return strs
}

// IsAFilesystemType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i FilesystemType) IsAFilesystemType() bool {
	for _, v := range _FilesystemTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for FilesystemType
func (i FilesystemType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for FilesystemType
func (i *FilesystemType) UnmarshalText(text []byte) error {
	var err error
	*i, err = FilesystemTypeString(string(text))
	return err
}

const _EncryptionKeyTypeName = "staticnodeIDkmstpm"

var _EncryptionKeyTypeIndex = [...]uint8{0, 6, 12, 15, 18}

const _EncryptionKeyTypeLowerName = "staticnodeidkmstpm"

func (i EncryptionKeyType) String() string {
	if i < 0 || i >= EncryptionKeyType(len(_EncryptionKeyTypeIndex)-1) {
		return fmt.Sprintf("EncryptionKeyType(%d)", i)
	}
	return _EncryptionKeyTypeName[_EncryptionKeyTypeIndex[i]:_EncryptionKeyTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _EncryptionKeyTypeNoOp() {
	var x [1]struct{}
	_ = x[EncryptionKeyStatic-(0)]
	_ = x[EncryptionKeyNodeID-(1)]
	_ = x[EncryptionKeyKMS-(2)]
	_ = x[EncryptionKeyTPM-(3)]
}

var _EncryptionKeyTypeValues = []EncryptionKeyType{EncryptionKeyStatic, EncryptionKeyNodeID, EncryptionKeyKMS, EncryptionKeyTPM}

var _EncryptionKeyTypeNameToValueMap = map[string]EncryptionKeyType{
	_EncryptionKeyTypeName[0:6]:        EncryptionKeyStatic,
	_EncryptionKeyTypeLowerName[0:6]:   EncryptionKeyStatic,
	_EncryptionKeyTypeName[6:12]:       EncryptionKeyNodeID,
	_EncryptionKeyTypeLowerName[6:12]:  EncryptionKeyNodeID,
	_EncryptionKeyTypeName[12:15]:      EncryptionKeyKMS,
	_EncryptionKeyTypeLowerName[12:15]: EncryptionKeyKMS,
	_EncryptionKeyTypeName[15:18]:      EncryptionKeyTPM,
	_EncryptionKeyTypeLowerName[15:18]: EncryptionKeyTPM,
}

var _EncryptionKeyTypeNames = []string{
	_EncryptionKeyTypeName[0:6],
	_EncryptionKeyTypeName[6:12],
	_EncryptionKeyTypeName[12:15],
	_EncryptionKeyTypeName[15:18],
}

// EncryptionKeyTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func EncryptionKeyTypeString(s string) (EncryptionKeyType, error) {
	if val, ok := _EncryptionKeyTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _EncryptionKeyTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to EncryptionKeyType values", s)
}

// EncryptionKeyTypeValues returns all values of the enum
func EncryptionKeyTypeValues() []EncryptionKeyType {
	return _EncryptionKeyTypeValues
}

// EncryptionKeyTypeStrings returns a slice of all String values of the enum
func EncryptionKeyTypeStrings() []string {
	strs := make([]string, len(_EncryptionKeyTypeNames))
	copy(strs, _EncryptionKeyTypeNames)
	return strs
}

// IsAEncryptionKeyType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i EncryptionKeyType) IsAEncryptionKeyType() bool {
	for _, v := range _EncryptionKeyTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for EncryptionKeyType
func (i EncryptionKeyType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for EncryptionKeyType
func (i *EncryptionKeyType) UnmarshalText(text []byte) error {
	var err error
	*i, err = EncryptionKeyTypeString(string(text))
	return err
}

const _EncryptionProviderTypeName = "noneluks2"

var _EncryptionProviderTypeIndex = [...]uint8{0, 4, 9}

const _EncryptionProviderTypeLowerName = "noneluks2"

func (i EncryptionProviderType) String() string {
	if i < 0 || i >= EncryptionProviderType(len(_EncryptionProviderTypeIndex)-1) {
		return fmt.Sprintf("EncryptionProviderType(%d)", i)
	}
	return _EncryptionProviderTypeName[_EncryptionProviderTypeIndex[i]:_EncryptionProviderTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _EncryptionProviderTypeNoOp() {
	var x [1]struct{}
	_ = x[EncryptionProviderNone-(0)]
	_ = x[EncryptionProviderLUKS2-(1)]
}

var _EncryptionProviderTypeValues = []EncryptionProviderType{EncryptionProviderNone, EncryptionProviderLUKS2}

var _EncryptionProviderTypeNameToValueMap = map[string]EncryptionProviderType{
	_EncryptionProviderTypeName[0:4]:      EncryptionProviderNone,
	_EncryptionProviderTypeLowerName[0:4]: EncryptionProviderNone,
	_EncryptionProviderTypeName[4:9]:      EncryptionProviderLUKS2,
	_EncryptionProviderTypeLowerName[4:9]: EncryptionProviderLUKS2,
}

var _EncryptionProviderTypeNames = []string{
	_EncryptionProviderTypeName[0:4],
	_EncryptionProviderTypeName[4:9],
}

// EncryptionProviderTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func EncryptionProviderTypeString(s string) (EncryptionProviderType, error) {
	if val, ok := _EncryptionProviderTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _EncryptionProviderTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to EncryptionProviderType values", s)
}

// EncryptionProviderTypeValues returns all values of the enum
func EncryptionProviderTypeValues() []EncryptionProviderType {
	return _EncryptionProviderTypeValues
}

// EncryptionProviderTypeStrings returns a slice of all String values of the enum
func EncryptionProviderTypeStrings() []string {
	strs := make([]string, len(_EncryptionProviderTypeNames))
	copy(strs, _EncryptionProviderTypeNames)
	return strs
}

// IsAEncryptionProviderType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i EncryptionProviderType) IsAEncryptionProviderType() bool {
	for _, v := range _EncryptionProviderTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for EncryptionProviderType
func (i EncryptionProviderType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for EncryptionProviderType
func (i *EncryptionProviderType) UnmarshalText(text []byte) error {
	var err error
	*i, err = EncryptionProviderTypeString(string(text))
	return err
}
