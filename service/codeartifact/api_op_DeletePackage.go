// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a package and all associated package versions. A deleted package cannot
// be restored. To delete one or more package versions, use the
// DeletePackageVersions
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DeletePackageVersions.html)
// API.
func (c *Client) DeletePackage(ctx context.Context, params *DeletePackageInput, optFns ...func(*Options)) (*DeletePackageOutput, error) {
	if params == nil {
		params = &DeletePackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePackage", params, optFns, c.addOperationDeletePackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePackageInput struct {

	// The name of the domain that contains the package to delete.
	//
	// This member is required.
	Domain *string

	// The format of the requested package to delete.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package to delete.
	//
	// This member is required.
	Package *string

	// The name of the repository that contains the package to delete.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The namespace of the package to delete. The package component that specifies its
	// namespace depends on its type. For example:
	//
	// * The namespace of a Maven package
	// is its groupId. The namespace is required when deleting Maven package
	// versions.
	//
	// * The namespace of an npm package is its scope.
	//
	// * Python and NuGet
	// packages do not contain corresponding components, packages of those formats do
	// not have a namespace.
	//
	// * The namespace of a generic package is its namespace.
	Namespace *string

	noSmithyDocumentSerde
}

type DeletePackageOutput struct {

	// Details about a package, including its format, namespace, and name.
	DeletedPackage *types.PackageSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeletePackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeletePackage{}, middleware.After)
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
	if err = addOpDeletePackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DeletePackage",
	}
}
