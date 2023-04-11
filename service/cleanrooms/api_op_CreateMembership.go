// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a membership for a specific collaboration identifier and joins the
// collaboration.
func (c *Client) CreateMembership(ctx context.Context, params *CreateMembershipInput, optFns ...func(*Options)) (*CreateMembershipOutput, error) {
	if params == nil {
		params = &CreateMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMembership", params, optFns, c.addOperationCreateMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMembershipInput struct {

	// The unique ID for the associated collaboration.
	//
	// This member is required.
	CollaborationIdentifier *string

	// An indicator as to whether query logging has been enabled or disabled for the
	// collaboration.
	//
	// This member is required.
	QueryLogStatus types.MembershipQueryLogStatus

	// An optional label that you can assign to a resource when you create it. Each tag
	// consists of a key and an optional value, both of which you define. When you use
	// tagging, you can also use tag-based access control in IAM policies to control
	// access to this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateMembershipOutput struct {

	// The membership that was created.
	//
	// This member is required.
	Membership *types.Membership

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMembership{}, middleware.After)
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
	if err = addOpCreateMembershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMembership(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMembership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cleanrooms",
		OperationName: "CreateMembership",
	}
}
