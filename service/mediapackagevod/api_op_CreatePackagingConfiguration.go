// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new MediaPackage VOD PackagingConfiguration resource.
func (c *Client) CreatePackagingConfiguration(ctx context.Context, params *CreatePackagingConfigurationInput, optFns ...func(*Options)) (*CreatePackagingConfigurationOutput, error) {
	if params == nil {
		params = &CreatePackagingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePackagingConfiguration", params, optFns, c.addOperationCreatePackagingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePackagingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A new MediaPackage VOD PackagingConfiguration resource configuration.
type CreatePackagingConfigurationInput struct {

	// The ID of the PackagingConfiguration.
	//
	// This member is required.
	Id *string

	// The ID of a PackagingGroup.
	//
	// This member is required.
	PackagingGroupId *string

	// A CMAF packaging configuration.
	CmafPackage *types.CmafPackage

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage

	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *types.MssPackage

	// A collection of tags associated with a resource
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreatePackagingConfigurationOutput struct {

	// The ARN of the PackagingConfiguration.
	Arn *string

	// A CMAF packaging configuration.
	CmafPackage *types.CmafPackage

	// The time the PackagingConfiguration was created.
	CreatedAt *string

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage

	// The ID of the PackagingConfiguration.
	Id *string

	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *types.MssPackage

	// The ID of a PackagingGroup.
	PackagingGroupId *string

	// A collection of tags associated with a resource
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePackagingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePackagingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePackagingConfiguration{}, middleware.After)
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
	if err = addOpCreatePackagingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePackagingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePackagingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "CreatePackagingConfiguration",
	}
}
