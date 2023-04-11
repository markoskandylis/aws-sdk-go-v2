// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You do not have sufficient access to perform this action.
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

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

// The request contains invalid parameters for the ARN or tags. This exception also
// occurs when you call a tagging API on a cancelled signing profile.
type BadRequestException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "BadRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource encountered a conflicting state.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

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

// An internal error occurred.
type InternalServiceErrorException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *InternalServiceErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServiceErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServiceErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The signing profile was not found.
type NotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A specified resource could not be found.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

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

// The client is making a request that exceeds service limits.
type ServiceLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *ServiceLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was denied due to request throttling. Instead of this error,
// TooManyRequestsException should be used.
type ThrottlingException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ThrottlingException"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The allowed number of job-signing requests has been exceeded. This error
// supersedes the error ThrottlingException.
type TooManyRequestsException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyRequestsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You signing certificate could not be validated.
type ValidationException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

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
