// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// A policy revision id was provided which does not match the latest policy
// revision. This exception is also if a policy revision id of 0 is provided via
// PutResourcePolicy and a policy with the same name already exists.
type InvalidPolicyRevisionIdException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidPolicyRevisionIdException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPolicyRevisionIdException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPolicyRevisionIdException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidPolicyRevisionIdException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidPolicyRevisionIdException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request is missing required parameters or has invalid parameters.
type InvalidRequestException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided resource policy would prevent the caller of this request from
// calling PutResourcePolicy in the future.
type LockoutPreventionException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LockoutPreventionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LockoutPreventionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LockoutPreventionException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LockoutPreventionException"
	}
	return *e.ErrorCodeOverride
}
func (e *LockoutPreventionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Invalid policy document provided in request.
type MalformedPolicyDocumentException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *MalformedPolicyDocumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MalformedPolicyDocumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MalformedPolicyDocumentException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MalformedPolicyDocumentException"
	}
	return *e.ErrorCodeOverride
}
func (e *MalformedPolicyDocumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exceeded the maximum number of resource policies for a target Amazon Web
// Services account.
type PolicyCountLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *PolicyCountLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PolicyCountLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PolicyCountLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PolicyCountLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *PolicyCountLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exceeded the maximum size for a resource policy.
type PolicySizeLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *PolicySizeLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PolicySizeLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PolicySizeLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PolicySizeLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *PolicySizeLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource was not found. Verify that the name or Amazon Resource Name (ARN)
// of the resource is correct.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceName *string

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

// You have reached the maximum number of sampling rules.
type RuleLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *RuleLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RuleLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RuleLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RuleLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *RuleLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request exceeds the maximum number of requests per second.
type ThrottledException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ThrottledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ThrottledException"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded the maximum number of tags you can apply to this resource.
type TooManyTagsException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceName *string

	noSmithyDocumentSerde
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyTagsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
