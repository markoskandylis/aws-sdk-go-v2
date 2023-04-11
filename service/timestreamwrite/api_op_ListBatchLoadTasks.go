// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamwrite

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of batch load tasks, along with the name, status, when the task
// is resumable until, and other details. See code sample
// (https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.list-batch-load-tasks.html)
// for details.
func (c *Client) ListBatchLoadTasks(ctx context.Context, params *ListBatchLoadTasksInput, optFns ...func(*Options)) (*ListBatchLoadTasksOutput, error) {
	if params == nil {
		params = &ListBatchLoadTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBatchLoadTasks", params, optFns, c.addOperationListBatchLoadTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBatchLoadTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBatchLoadTasksInput struct {

	// The total number of items to return in the output. If the total number of items
	// available is more than the value specified, a NextToken is provided in the
	// output. To resume pagination, provide the NextToken value as argument of a
	// subsequent API invocation.
	MaxResults *int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	// Status of the batch load task.
	TaskStatus types.BatchLoadStatus

	noSmithyDocumentSerde
}

type ListBatchLoadTasksOutput struct {

	// A list of batch load task details.
	BatchLoadTasks []types.BatchLoadTask

	// A token to specify where to start paginating. Provide the next
	// ListBatchLoadTasksRequest.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBatchLoadTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListBatchLoadTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListBatchLoadTasks{}, middleware.After)
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
	if err = addOpListBatchLoadTasksDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBatchLoadTasks(options.Region), middleware.Before); err != nil {
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

func addOpListBatchLoadTasksDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Serialize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
				opt.EndpointResolverUsedForDiscovery = o.EndpointDiscovery.EndpointResolverUsedForDiscovery
			},
		},
		DiscoverOperation:            c.fetchOpListBatchLoadTasksDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    true,
	}, "ResolveEndpoint", middleware.After)
}

func (c *Client) fetchOpListBatchLoadTasksDiscoverEndpoint(ctx context.Context, input interface{}, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	in, ok := input.(*ListBatchLoadTasksInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)

	key := fmt.Sprintf("Timestream Write.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	endpoint, err := c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, key, opt)
	if err != nil {
		return internalEndpointDiscovery.WeightedAddress{}, err
	}

	weighted, ok := endpoint.GetValidAddress()
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("no valid endpoint address returned by the endpoint discovery api")
	}
	return weighted, nil
}

// ListBatchLoadTasksAPIClient is a client that implements the ListBatchLoadTasks
// operation.
type ListBatchLoadTasksAPIClient interface {
	ListBatchLoadTasks(context.Context, *ListBatchLoadTasksInput, ...func(*Options)) (*ListBatchLoadTasksOutput, error)
}

var _ ListBatchLoadTasksAPIClient = (*Client)(nil)

// ListBatchLoadTasksPaginatorOptions is the paginator options for
// ListBatchLoadTasks
type ListBatchLoadTasksPaginatorOptions struct {
	// The total number of items to return in the output. If the total number of items
	// available is more than the value specified, a NextToken is provided in the
	// output. To resume pagination, provide the NextToken value as argument of a
	// subsequent API invocation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBatchLoadTasksPaginator is a paginator for ListBatchLoadTasks
type ListBatchLoadTasksPaginator struct {
	options   ListBatchLoadTasksPaginatorOptions
	client    ListBatchLoadTasksAPIClient
	params    *ListBatchLoadTasksInput
	nextToken *string
	firstPage bool
}

// NewListBatchLoadTasksPaginator returns a new ListBatchLoadTasksPaginator
func NewListBatchLoadTasksPaginator(client ListBatchLoadTasksAPIClient, params *ListBatchLoadTasksInput, optFns ...func(*ListBatchLoadTasksPaginatorOptions)) *ListBatchLoadTasksPaginator {
	if params == nil {
		params = &ListBatchLoadTasksInput{}
	}

	options := ListBatchLoadTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBatchLoadTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBatchLoadTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBatchLoadTasks page.
func (p *ListBatchLoadTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBatchLoadTasksOutput, error) {
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

	result, err := p.client.ListBatchLoadTasks(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListBatchLoadTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "timestream",
		OperationName: "ListBatchLoadTasks",
	}
}
