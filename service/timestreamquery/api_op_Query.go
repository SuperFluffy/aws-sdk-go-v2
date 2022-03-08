// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Query is a synchronous operation that enables you to run a query against your
// Amazon Timestream data. Query will time out after 60 seconds. You must update
// the default timeout in the SDK to support a timeout of 60 seconds. See the code
// sample
// (https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.run-query.html)
// for details. Your query request will fail in the following cases:
//
// * If you
// submit a Query request with the same client token outside of the 5-minute
// idempotency window.
//
// * If you submit a Query request with the same client token,
// but change other parameters, within the 5-minute idempotency window.
//
// * If the
// size of the row (including the query metadata) exceeds 1 MB, then the query will
// fail with the following error message: Query aborted as max page response size
// has been exceeded by the output result row
//
// * If the IAM principal of the query
// initiator and the result reader are not the same and/or the query initiator and
// the result reader do not have the same query string in the query requests, the
// query will fail with an Invalid pagination token error.
func (c *Client) Query(ctx context.Context, params *QueryInput, optFns ...func(*Options)) (*QueryOutput, error) {
	if params == nil {
		params = &QueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Query", params, optFns, c.addOperationQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryInput struct {

	// The query to be run by Timestream.
	//
	// This member is required.
	QueryString *string

	// Unique, case-sensitive string of up to 64 ASCII characters specified when a
	// Query request is made. Providing a ClientToken makes the call to Query
	// idempotent. This means that running the same query repeatedly will produce the
	// same result. In other words, making multiple identical Query requests has the
	// same effect as making a single request. When using ClientToken in a query, note
	// the following:
	//
	// * If the Query API is instantiated without a ClientToken, the
	// Query SDK generates a ClientToken on your behalf.
	//
	// * If the Query invocation
	// only contains the ClientToken but does not include a NextToken, that invocation
	// of Query is assumed to be a new query run.
	//
	// * If the invocation contains
	// NextToken, that particular invocation is assumed to be a subsequent invocation
	// of a prior call to the Query API, and a result set is returned.
	//
	// * After 4
	// hours, any request with the same ClientToken is treated as a new request.
	ClientToken *string

	// The total number of rows to be returned in the Query output. The initial run of
	// Query with a MaxRows value specified will return the result set of the query in
	// two cases:
	//
	// * The size of the result is less than 1MB.
	//
	// * The number of rows in
	// the result set is less than the value of maxRows.
	//
	// Otherwise, the initial
	// invocation of Query only returns a NextToken, which can then be used in
	// subsequent calls to fetch the result set. To resume pagination, provide the
	// NextToken value in the subsequent command. If the row size is large (e.g. a row
	// has many columns), Timestream may return fewer rows to keep the response size
	// from exceeding the 1 MB limit. If MaxRows is not provided, Timestream will send
	// the necessary number of rows to meet the 1 MB limit.
	MaxRows *int32

	// A pagination token used to return a set of results. When the Query API is
	// invoked using NextToken, that particular invocation is assumed to be a
	// subsequent invocation of a prior call to Query, and a result set is returned.
	// However, if the Query invocation only contains the ClientToken, that invocation
	// of Query is assumed to be a new query run. Note the following when using
	// NextToken in a query:
	//
	// * A pagination token can be used for up to five Query
	// invocations, OR for a duration of up to 1 hour – whichever comes first.
	//
	// * Using
	// the same NextToken will return the same set of records. To keep paginating
	// through the result set, you must to use the most recent nextToken.
	//
	// * Suppose a
	// Query invocation returns two NextToken values, TokenA and TokenB. If TokenB is
	// used in a subsequent Query invocation, then TokenA is invalidated and cannot be
	// reused.
	//
	// * To request a previous result set from a query after pagination has
	// begun, you must re-invoke the Query API.
	//
	// * The latest NextToken should be used
	// to paginate until null is returned, at which point a new NextToken should be
	// used.
	//
	// * If the IAM principal of the query initiator and the result reader are
	// not the same and/or the query initiator and the result reader do not have the
	// same query string in the query requests, the query will fail with an Invalid
	// pagination token error.
	NextToken *string

	noSmithyDocumentSerde
}

type QueryOutput struct {

	// The column data types of the returned result set.
	//
	// This member is required.
	ColumnInfo []types.ColumnInfo

	// A unique ID for the given query.
	//
	// This member is required.
	QueryId *string

	// The result set rows returned by the query.
	//
	// This member is required.
	Rows []types.Row

	// A pagination token that can be used again on a Query call to get the next set of
	// results.
	NextToken *string

	// Information about the status of the query, including progress and bytes scanned.
	QueryStatus *types.QueryStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpQuery{}, middleware.After)
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
	if err = addOpQueryDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addIdempotencyToken_opQueryMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQuery(options.Region), middleware.Before); err != nil {
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

func addOpQueryDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Serialize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
				opt.EndpointResolverUsedForDiscovery = o.EndpointDiscovery.EndpointResolverUsedForDiscovery
			},
		},
		DiscoverOperation:            c.fetchOpQueryDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    true,
	}, "ResolveEndpoint", middleware.After)
}

func (c *Client) fetchOpQueryDiscoverEndpoint(ctx context.Context, input interface{}, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	in, ok := input.(*QueryInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)

	key := fmt.Sprintf("Timestream Query.%v", identifierMap)

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

// QueryAPIClient is a client that implements the Query operation.
type QueryAPIClient interface {
	Query(context.Context, *QueryInput, ...func(*Options)) (*QueryOutput, error)
}

var _ QueryAPIClient = (*Client)(nil)

// QueryPaginatorOptions is the paginator options for Query
type QueryPaginatorOptions struct {
	// The total number of rows to be returned in the Query output. The initial run of
	// Query with a MaxRows value specified will return the result set of the query in
	// two cases:
	//
	// * The size of the result is less than 1MB.
	//
	// * The number of rows in
	// the result set is less than the value of maxRows.
	//
	// Otherwise, the initial
	// invocation of Query only returns a NextToken, which can then be used in
	// subsequent calls to fetch the result set. To resume pagination, provide the
	// NextToken value in the subsequent command. If the row size is large (e.g. a row
	// has many columns), Timestream may return fewer rows to keep the response size
	// from exceeding the 1 MB limit. If MaxRows is not provided, Timestream will send
	// the necessary number of rows to meet the 1 MB limit.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// QueryPaginator is a paginator for Query
type QueryPaginator struct {
	options   QueryPaginatorOptions
	client    QueryAPIClient
	params    *QueryInput
	nextToken *string
	firstPage bool
}

// NewQueryPaginator returns a new QueryPaginator
func NewQueryPaginator(client QueryAPIClient, params *QueryInput, optFns ...func(*QueryPaginatorOptions)) *QueryPaginator {
	if params == nil {
		params = &QueryInput{}
	}

	options := QueryPaginatorOptions{}
	if params.MaxRows != nil {
		options.Limit = *params.MaxRows
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &QueryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *QueryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next Query page.
func (p *QueryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*QueryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRows = limit

	result, err := p.client.Query(ctx, &params, optFns...)
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

type idempotencyToken_initializeOpQuery struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpQuery) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*QueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *QueryInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opQueryMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpQuery{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "timestream",
		OperationName: "Query",
	}
}
