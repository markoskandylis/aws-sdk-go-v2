// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// A remediation exception is when a specified resource is no longer considered for
// auto-remediation. This API adds a new exception or updates an existing exception
// for a specified resource with a specified Config rule. Config generates a
// remediation exception when a problem occurs running a remediation action for a
// specified resource. Remediation exceptions blocks auto-remediation until the
// exception is cleared. When placing an exception on an Amazon Web Services
// resource, it is recommended that remediation is set as manual remediation until
// the given Config rule for the specified resource evaluates the resource as
// NON_COMPLIANT. Once the resource has been evaluated as NON_COMPLIANT, you can
// add remediation exceptions and change the remediation type back from Manual to
// Auto if you want to use auto-remediation. Otherwise, using auto-remediation
// before a NON_COMPLIANT evaluation result can delete resources before the
// exception is applied. Placing an exception can only be performed on resources
// that are NON_COMPLIANT. If you use this API for COMPLIANT resources or resources
// that are NOT_APPLICABLE, a remediation exception will not be generated. For more
// information on the conditions that initiate the possible Config evaluation
// results, see Concepts | Config Rules
// (https://docs.aws.amazon.com/config/latest/developerguide/config-concepts.html#aws-config-rules)
// in the Config Developer Guide.
func (c *Client) PutRemediationExceptions(ctx context.Context, params *PutRemediationExceptionsInput, optFns ...func(*Options)) (*PutRemediationExceptionsOutput, error) {
	if params == nil {
		params = &PutRemediationExceptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRemediationExceptions", params, optFns, c.addOperationPutRemediationExceptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRemediationExceptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRemediationExceptionsInput struct {

	// The name of the Config rule for which you want to create remediation exception.
	//
	// This member is required.
	ConfigRuleName *string

	// An exception list of resource exception keys to be processed with the current
	// request. Config adds exception for each resource key. For example, Config adds 3
	// exceptions for 3 resource keys.
	//
	// This member is required.
	ResourceKeys []types.RemediationExceptionResourceKey

	// The exception is automatically deleted after the expiration date.
	ExpirationTime *time.Time

	// The message contains an explanation of the exception.
	Message *string

	noSmithyDocumentSerde
}

type PutRemediationExceptionsOutput struct {

	// Returns a list of failed remediation exceptions batch objects. Each object in
	// the batch consists of a list of failed items and failure messages.
	FailedBatches []types.FailedRemediationExceptionBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRemediationExceptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRemediationExceptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRemediationExceptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutRemediationExceptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRemediationExceptions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutRemediationExceptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutRemediationExceptions",
	}
}
