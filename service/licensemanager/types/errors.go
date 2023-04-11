// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Access to resource denied.
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Amazon Web Services user account does not have permission to perform the
// action. Check the IAM policy associated with this account.
type AuthorizationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AuthorizationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AuthorizationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AuthorizationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AuthorizationException"
	}
	return *e.ErrorCodeOverride
}
func (e *AuthorizationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There was a conflict processing the request. Try your request again.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The entitlement is not allowed.
type EntitlementNotAllowedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *EntitlementNotAllowedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntitlementNotAllowedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntitlementNotAllowedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EntitlementNotAllowedException"
	}
	return *e.ErrorCodeOverride
}
func (e *EntitlementNotAllowedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A dependency required to run the API is missing.
type FailedDependencyException struct {
	Message *string

	ErrorCodeOverride *string

	ErrorCode_ *string

	noSmithyDocumentSerde
}

func (e *FailedDependencyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FailedDependencyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FailedDependencyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "FailedDependencyException"
	}
	return *e.ErrorCodeOverride
}
func (e *FailedDependencyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request uses too many filters or too many filter values.
type FilterLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *FilterLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FilterLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FilterLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "FilterLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *FilterLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more parameter values are not valid.
type InvalidParameterValueException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterValueException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// License Manager cannot allocate a license to a resource because of its state.
// For example, you cannot allocate a license to an instance in the process of
// shutting down.
type InvalidResourceStateException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidResourceStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidResourceStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidResourceStateException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidResourceStateException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidResourceStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You do not have enough licenses available to support a new resource launch.
type LicenseUsageException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LicenseUsageException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LicenseUsageException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LicenseUsageException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LicenseUsageException"
	}
	return *e.ErrorCodeOverride
}
func (e *LicenseUsageException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There are no entitlements found for this license, or the entitlement maximum
// count is reached.
type NoEntitlementsAllowedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NoEntitlementsAllowedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoEntitlementsAllowedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoEntitlementsAllowedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NoEntitlementsAllowedException"
	}
	return *e.ErrorCodeOverride
}
func (e *NoEntitlementsAllowedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Too many requests have been submitted. Try again after a brief wait.
type RateLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RateLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RateLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RateLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RateLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *RateLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This is not the correct Region for the resource. Try again.
type RedirectException struct {
	Message *string

	ErrorCodeOverride *string

	Location *string

	noSmithyDocumentSerde
}

func (e *RedirectException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RedirectException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RedirectException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RedirectException"
	}
	return *e.ErrorCodeOverride
}
func (e *RedirectException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your resource limits have been exceeded.
type ResourceLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource cannot be found.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The server experienced an internal error. Try again.
type ServerInternalException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServerInternalException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServerInternalException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServerInternalException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServerInternalException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServerInternalException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The digital signature method is unsupported. Try your request again.
type UnsupportedDigitalSignatureMethodException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UnsupportedDigitalSignatureMethodException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedDigitalSignatureMethodException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedDigitalSignatureMethodException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedDigitalSignatureMethodException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedDigitalSignatureMethodException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The provided input is not valid. Try your request again.
type ValidationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
