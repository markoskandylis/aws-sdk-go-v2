// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// An error occurred while processing the request.
type BaseException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *BaseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BaseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BaseException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "BaseException"
	}
	return *e.ErrorCodeOverride
}
func (e *BaseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because it attempted an operation which is not enabled.
type DisabledOperationException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *DisabledOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DisabledOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DisabledOperationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DisabledAction"
	}
	return *e.ErrorCodeOverride
}
func (e *DisabledOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal error occurred while processing the request. If this problem
// persists, report an issue from the Service Health Dashboard
// (http://status.aws.amazon.com/).
type InternalException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *InternalException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The request was rejected because it specified an invalid type definition.
type InvalidTypeException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *InvalidTypeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTypeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTypeException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidType"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidTypeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because a resource limit has already been met.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because it attempted to create a resource that already
// exists.
type ResourceAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	Code *string

	noSmithyDocumentSerde
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceAlreadyExists"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because it attempted to reference a resource that does
// not exist.
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
		return "ResourceNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because it has invalid parameters.
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
