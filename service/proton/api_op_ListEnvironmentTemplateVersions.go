// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List major or minor versions of an environment template with detail data.
func (c *Client) ListEnvironmentTemplateVersions(ctx context.Context, params *ListEnvironmentTemplateVersionsInput, optFns ...func(*Options)) (*ListEnvironmentTemplateVersionsOutput, error) {
	if params == nil {
		params = &ListEnvironmentTemplateVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnvironmentTemplateVersions", params, optFns, addOperationListEnvironmentTemplateVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnvironmentTemplateVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnvironmentTemplateVersionsInput struct {

	// The name of the environment template.
	//
	// This member is required.
	TemplateName *string

	// To view a list of minor of versions under a major version of an environment
	// template, include majorVersion. To view a list of major versions of an
	// environment template, exclude majorVersion.
	MajorVersion *string

	// The maximum number of major or minor versions of an environment template to
	// list.
	MaxResults *int32

	// A token to indicate the location of the next major or minor version in the array
	// of major or minor versions of an environment template, after the list of major
	// or minor versions that was previously requested.
	NextToken *string
}

type ListEnvironmentTemplateVersionsOutput struct {

	// An array of major or minor versions of an environment template detail data.
	//
	// This member is required.
	TemplateVersions []types.EnvironmentTemplateVersionSummary

	// A token to indicate the location of the next major or minor version in the array
	// of major or minor versions of an environment template, after the list of major
	// or minor versions that was previously requested.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEnvironmentTemplateVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListEnvironmentTemplateVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListEnvironmentTemplateVersions{}, middleware.After)
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
	if err = addOpListEnvironmentTemplateVersionsValidationMiddleware(stack); err != nil {
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

// ListEnvironmentTemplateVersionsAPIClient is a client that implements the
// ListEnvironmentTemplateVersions operation.
type ListEnvironmentTemplateVersionsAPIClient interface {
	ListEnvironmentTemplateVersions(context.Context, *ListEnvironmentTemplateVersionsInput, ...func(*Options)) (*ListEnvironmentTemplateVersionsOutput, error)
}

var _ ListEnvironmentTemplateVersionsAPIClient = (*Client)(nil)

// ListEnvironmentTemplateVersionsPaginatorOptions is the paginator options for
// ListEnvironmentTemplateVersions
type ListEnvironmentTemplateVersionsPaginatorOptions struct {
	// The maximum number of major or minor versions of an environment template to
	// list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnvironmentTemplateVersionsPaginator is a paginator for
// ListEnvironmentTemplateVersions
type ListEnvironmentTemplateVersionsPaginator struct {
	options   ListEnvironmentTemplateVersionsPaginatorOptions
	client    ListEnvironmentTemplateVersionsAPIClient
	params    *ListEnvironmentTemplateVersionsInput
	nextToken *string
	firstPage bool
}

// NewListEnvironmentTemplateVersionsPaginator returns a new
// ListEnvironmentTemplateVersionsPaginator
func NewListEnvironmentTemplateVersionsPaginator(client ListEnvironmentTemplateVersionsAPIClient, params *ListEnvironmentTemplateVersionsInput, optFns ...func(*ListEnvironmentTemplateVersionsPaginatorOptions)) *ListEnvironmentTemplateVersionsPaginator {
	if params == nil {
		params = &ListEnvironmentTemplateVersionsInput{}
	}

	options := ListEnvironmentTemplateVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnvironmentTemplateVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnvironmentTemplateVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListEnvironmentTemplateVersions page.
func (p *ListEnvironmentTemplateVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnvironmentTemplateVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListEnvironmentTemplateVersions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}
