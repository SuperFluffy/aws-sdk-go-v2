// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubrefactorspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all Amazon Web Services Migration Hub Refactor Spaces service virtual
// private clouds (VPCs) that are part of the environment.
func (c *Client) ListEnvironmentVpcs(ctx context.Context, params *ListEnvironmentVpcsInput, optFns ...func(*Options)) (*ListEnvironmentVpcsOutput, error) {
	if params == nil {
		params = &ListEnvironmentVpcsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnvironmentVpcs", params, optFns, c.addOperationListEnvironmentVpcsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnvironmentVpcsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnvironmentVpcsInput struct {

	// The ID of the environment.
	//
	// This member is required.
	EnvironmentIdentifier *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEnvironmentVpcsOutput struct {

	// The list of EnvironmentVpc objects.
	EnvironmentVpcList []types.EnvironmentVpc

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnvironmentVpcsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEnvironmentVpcs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEnvironmentVpcs{}, middleware.After)
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
	if err = addOpListEnvironmentVpcsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnvironmentVpcs(options.Region), middleware.Before); err != nil {
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

// ListEnvironmentVpcsAPIClient is a client that implements the ListEnvironmentVpcs
// operation.
type ListEnvironmentVpcsAPIClient interface {
	ListEnvironmentVpcs(context.Context, *ListEnvironmentVpcsInput, ...func(*Options)) (*ListEnvironmentVpcsOutput, error)
}

var _ ListEnvironmentVpcsAPIClient = (*Client)(nil)

// ListEnvironmentVpcsPaginatorOptions is the paginator options for
// ListEnvironmentVpcs
type ListEnvironmentVpcsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnvironmentVpcsPaginator is a paginator for ListEnvironmentVpcs
type ListEnvironmentVpcsPaginator struct {
	options   ListEnvironmentVpcsPaginatorOptions
	client    ListEnvironmentVpcsAPIClient
	params    *ListEnvironmentVpcsInput
	nextToken *string
	firstPage bool
}

// NewListEnvironmentVpcsPaginator returns a new ListEnvironmentVpcsPaginator
func NewListEnvironmentVpcsPaginator(client ListEnvironmentVpcsAPIClient, params *ListEnvironmentVpcsInput, optFns ...func(*ListEnvironmentVpcsPaginatorOptions)) *ListEnvironmentVpcsPaginator {
	if params == nil {
		params = &ListEnvironmentVpcsInput{}
	}

	options := ListEnvironmentVpcsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnvironmentVpcsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnvironmentVpcsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnvironmentVpcs page.
func (p *ListEnvironmentVpcsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnvironmentVpcsOutput, error) {
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

	result, err := p.client.ListEnvironmentVpcs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEnvironmentVpcs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "refactor-spaces",
		OperationName: "ListEnvironmentVpcs",
	}
}
