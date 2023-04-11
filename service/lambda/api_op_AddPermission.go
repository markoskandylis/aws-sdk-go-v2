// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Grants an Amazon Web Service, Amazon Web Services account, or Amazon Web
// Services organization permission to use a function. You can apply the policy at
// the function level, or specify a qualifier to restrict access to a single
// version or alias. If you use a qualifier, the invoker must use the full Amazon
// Resource Name (ARN) of that version or alias to invoke the function. Note:
// Lambda does not support adding policies to version $LATEST. To grant permission
// to another account, specify the account ID as the Principal. To grant permission
// to an organization defined in Organizations, specify the organization ID as the
// PrincipalOrgID. For Amazon Web Services, the principal is a domain-style
// identifier that the service defines, such as s3.amazonaws.com or
// sns.amazonaws.com. For Amazon Web Services, you can also specify the ARN of the
// associated resource as the SourceArn. If you grant permission to a service
// principal without specifying the source, other accounts could potentially
// configure resources in their account to invoke your Lambda function. This
// operation adds a statement to a resource-based permissions policy for the
// function. For more information about function policies, see Using resource-based
// policies for Lambda
// (https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html).
func (c *Client) AddPermission(ctx context.Context, params *AddPermissionInput, optFns ...func(*Options)) (*AddPermissionOutput, error) {
	if params == nil {
		params = &AddPermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddPermission", params, optFns, c.addOperationAddPermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddPermissionInput struct {

	// The action that the principal can use on the function. For example,
	// lambda:InvokeFunction or lambda:GetFunction.
	//
	// This member is required.
	Action *string

	// The name of the Lambda function, version, or alias. Name formats
	//
	// * Function
	// name – my-function (name-only), my-function:v1 (with alias).
	//
	// * Function ARN –
	// arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// * Partial ARN –
	// 123456789012:function:my-function.
	//
	// You can append a version number or alias to
	// any of the formats. The length constraint applies only to the full ARN. If you
	// specify only the function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// The Amazon Web Service or Amazon Web Services account that invokes the function.
	// If you specify a service, use SourceArn or SourceAccount to limit who can invoke
	// the function through that service.
	//
	// This member is required.
	Principal *string

	// A statement identifier that differentiates the statement from others in the same
	// policy.
	//
	// This member is required.
	StatementId *string

	// For Alexa Smart Home functions, a token that the invoker must supply.
	EventSourceToken *string

	// The type of authentication that your function URL uses. Set to AWS_IAM if you
	// want to restrict access to authenticated users only. Set to NONE if you want to
	// bypass IAM authentication to create a public endpoint. For more information, see
	// Security and auth model for Lambda function URLs
	// (https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).
	FunctionUrlAuthType types.FunctionUrlAuthType

	// The identifier for your organization in Organizations. Use this to grant
	// permissions to all the Amazon Web Services accounts under this organization.
	PrincipalOrgID *string

	// Specify a version or alias to add permissions to a published version of the
	// function.
	Qualifier *string

	// Update the policy only if the revision ID matches the ID that's specified. Use
	// this option to avoid modifying a policy that has changed since you last read it.
	RevisionId *string

	// For Amazon Web Service, the ID of the Amazon Web Services account that owns the
	// resource. Use this together with SourceArn to ensure that the specified account
	// owns the resource. It is possible for an Amazon S3 bucket to be deleted by its
	// owner and recreated by another account.
	SourceAccount *string

	// For Amazon Web Services, the ARN of the Amazon Web Services resource that
	// invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic. Note
	// that Lambda configures the comparison using the StringLike operator.
	SourceArn *string

	noSmithyDocumentSerde
}

type AddPermissionOutput struct {

	// The permission statement that's added to the function policy.
	Statement *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddPermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAddPermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAddPermission{}, middleware.After)
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
	if err = addOpAddPermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddPermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddPermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "AddPermission",
	}
}
