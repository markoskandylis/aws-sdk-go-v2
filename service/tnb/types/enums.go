// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DescriptorContentType string

// Enum values for DescriptorContentType
const (
	DescriptorContentTypeTextPlain DescriptorContentType = "text/plain"
)

// Values returns all known values for DescriptorContentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DescriptorContentType) Values() []DescriptorContentType {
	return []DescriptorContentType{
		"text/plain",
	}
}

type LcmOperationType string

// Enum values for LcmOperationType
const (
	LcmOperationTypeInstantiate LcmOperationType = "INSTANTIATE"
	LcmOperationTypeUpdate      LcmOperationType = "UPDATE"
	LcmOperationTypeTerminate   LcmOperationType = "TERMINATE"
)

// Values returns all known values for LcmOperationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LcmOperationType) Values() []LcmOperationType {
	return []LcmOperationType{
		"INSTANTIATE",
		"UPDATE",
		"TERMINATE",
	}
}

type NsdOnboardingState string

// Enum values for NsdOnboardingState
const (
	NsdOnboardingStateCreated   NsdOnboardingState = "CREATED"
	NsdOnboardingStateOnboarded NsdOnboardingState = "ONBOARDED"
	NsdOnboardingStateError     NsdOnboardingState = "ERROR"
)

// Values returns all known values for NsdOnboardingState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NsdOnboardingState) Values() []NsdOnboardingState {
	return []NsdOnboardingState{
		"CREATED",
		"ONBOARDED",
		"ERROR",
	}
}

type NsdOperationalState string

// Enum values for NsdOperationalState
const (
	NsdOperationalStateEnabled  NsdOperationalState = "ENABLED"
	NsdOperationalStateDisabled NsdOperationalState = "DISABLED"
)

// Values returns all known values for NsdOperationalState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NsdOperationalState) Values() []NsdOperationalState {
	return []NsdOperationalState{
		"ENABLED",
		"DISABLED",
	}
}

type NsdUsageState string

// Enum values for NsdUsageState
const (
	NsdUsageStateInUse    NsdUsageState = "IN_USE"
	NsdUsageStateNotInUse NsdUsageState = "NOT_IN_USE"
)

// Values returns all known values for NsdUsageState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NsdUsageState) Values() []NsdUsageState {
	return []NsdUsageState{
		"IN_USE",
		"NOT_IN_USE",
	}
}

type NsLcmOperationState string

// Enum values for NsLcmOperationState
const (
	NsLcmOperationStateProcessing NsLcmOperationState = "PROCESSING"
	NsLcmOperationStateCompleted  NsLcmOperationState = "COMPLETED"
	NsLcmOperationStateFailed     NsLcmOperationState = "FAILED"
	NsLcmOperationStateCancelling NsLcmOperationState = "CANCELLING"
	NsLcmOperationStateCancelled  NsLcmOperationState = "CANCELLED"
)

// Values returns all known values for NsLcmOperationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NsLcmOperationState) Values() []NsLcmOperationState {
	return []NsLcmOperationState{
		"PROCESSING",
		"COMPLETED",
		"FAILED",
		"CANCELLING",
		"CANCELLED",
	}
}

type NsState string

// Enum values for NsState
const (
	NsStateInstantiated          NsState = "INSTANTIATED"
	NsStateNotInstantiated       NsState = "NOT_INSTANTIATED"
	NsStateImpaired              NsState = "IMPAIRED"
	NsStateStopped               NsState = "STOPPED"
	NsStateDeleted               NsState = "DELETED"
	NsStateInstantiateInProgress NsState = "INSTANTIATE_IN_PROGRESS"
	NsStateUpdateInProgress      NsState = "UPDATE_IN_PROGRESS"
	NsStateTerminateInProgress   NsState = "TERMINATE_IN_PROGRESS"
)

// Values returns all known values for NsState. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (NsState) Values() []NsState {
	return []NsState{
		"INSTANTIATED",
		"NOT_INSTANTIATED",
		"IMPAIRED",
		"STOPPED",
		"DELETED",
		"INSTANTIATE_IN_PROGRESS",
		"UPDATE_IN_PROGRESS",
		"TERMINATE_IN_PROGRESS",
	}
}

type OnboardingState string

// Enum values for OnboardingState
const (
	OnboardingStateCreated   OnboardingState = "CREATED"
	OnboardingStateOnboarded OnboardingState = "ONBOARDED"
	OnboardingStateError     OnboardingState = "ERROR"
)

// Values returns all known values for OnboardingState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OnboardingState) Values() []OnboardingState {
	return []OnboardingState{
		"CREATED",
		"ONBOARDED",
		"ERROR",
	}
}

type OperationalState string

// Enum values for OperationalState
const (
	OperationalStateEnabled  OperationalState = "ENABLED"
	OperationalStateDisabled OperationalState = "DISABLED"
)

// Values returns all known values for OperationalState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationalState) Values() []OperationalState {
	return []OperationalState{
		"ENABLED",
		"DISABLED",
	}
}

type PackageContentType string

// Enum values for PackageContentType
const (
	PackageContentTypeApplicationZip PackageContentType = "application/zip"
)

// Values returns all known values for PackageContentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PackageContentType) Values() []PackageContentType {
	return []PackageContentType{
		"application/zip",
	}
}

type TaskStatus string

// Enum values for TaskStatus
const (
	TaskStatusScheduled  TaskStatus = "SCHEDULED"
	TaskStatusStarted    TaskStatus = "STARTED"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusCompleted  TaskStatus = "COMPLETED"
	TaskStatusError      TaskStatus = "ERROR"
	TaskStatusSkipped    TaskStatus = "SKIPPED"
	TaskStatusCancelled  TaskStatus = "CANCELLED"
)

// Values returns all known values for TaskStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TaskStatus) Values() []TaskStatus {
	return []TaskStatus{
		"SCHEDULED",
		"STARTED",
		"IN_PROGRESS",
		"COMPLETED",
		"ERROR",
		"SKIPPED",
		"CANCELLED",
	}
}

type UpdateSolNetworkType string

// Enum values for UpdateSolNetworkType
const (
	UpdateSolNetworkTypeModifyVnfInformation UpdateSolNetworkType = "MODIFY_VNF_INFORMATION"
)

// Values returns all known values for UpdateSolNetworkType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UpdateSolNetworkType) Values() []UpdateSolNetworkType {
	return []UpdateSolNetworkType{
		"MODIFY_VNF_INFORMATION",
	}
}

type UsageState string

// Enum values for UsageState
const (
	UsageStateInUse    UsageState = "IN_USE"
	UsageStateNotInUse UsageState = "NOT_IN_USE"
)

// Values returns all known values for UsageState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UsageState) Values() []UsageState {
	return []UsageState{
		"IN_USE",
		"NOT_IN_USE",
	}
}

type VnfInstantiationState string

// Enum values for VnfInstantiationState
const (
	VnfInstantiationStateInstantiated    VnfInstantiationState = "INSTANTIATED"
	VnfInstantiationStateNotInstantiated VnfInstantiationState = "NOT_INSTANTIATED"
)

// Values returns all known values for VnfInstantiationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VnfInstantiationState) Values() []VnfInstantiationState {
	return []VnfInstantiationState{
		"INSTANTIATED",
		"NOT_INSTANTIATED",
	}
}

type VnfOperationalState string

// Enum values for VnfOperationalState
const (
	VnfOperationalStateStarted VnfOperationalState = "STARTED"
	VnfOperationalStateStopped VnfOperationalState = "STOPPED"
)

// Values returns all known values for VnfOperationalState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VnfOperationalState) Values() []VnfOperationalState {
	return []VnfOperationalState{
		"STARTED",
		"STOPPED",
	}
}
