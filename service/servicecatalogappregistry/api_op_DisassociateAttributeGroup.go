// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates an attribute group from an application to remove the extra
// attributes contained in the attribute group from the application's metadata.
// This operation reverts AssociateAttributeGroup.
func (c *Client) DisassociateAttributeGroup(ctx context.Context, params *DisassociateAttributeGroupInput, optFns ...func(*Options)) (*DisassociateAttributeGroupOutput, error) {
	if params == nil {
		params = &DisassociateAttributeGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateAttributeGroup", params, optFns, c.addOperationDisassociateAttributeGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateAttributeGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateAttributeGroupInput struct {

	// The name, ID, or ARN of the application.
	//
	// This member is required.
	Application *string

	// The name, ID, or ARN of the attribute group that holds the attributes to
	// describe the application.
	//
	// This member is required.
	AttributeGroup *string

	noSmithyDocumentSerde
}

type DisassociateAttributeGroupOutput struct {

	// The Amazon resource name (ARN) that specifies the application.
	ApplicationArn *string

	// The Amazon resource name (ARN) that specifies the attribute group.
	AttributeGroupArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateAttributeGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateAttributeGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateAttributeGroup{}, middleware.After)
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
	if err = addOpDisassociateAttributeGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateAttributeGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateAttributeGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DisassociateAttributeGroup",
	}
}
