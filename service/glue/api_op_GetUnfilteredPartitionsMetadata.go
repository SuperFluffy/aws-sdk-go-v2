// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) GetUnfilteredPartitionsMetadata(ctx context.Context, params *GetUnfilteredPartitionsMetadataInput, optFns ...func(*Options)) (*GetUnfilteredPartitionsMetadataOutput, error) {
	if params == nil {
		params = &GetUnfilteredPartitionsMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUnfilteredPartitionsMetadata", params, optFns, c.addOperationGetUnfilteredPartitionsMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUnfilteredPartitionsMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUnfilteredPartitionsMetadataInput struct {

	// This member is required.
	CatalogId *string

	// This member is required.
	DatabaseName *string

	// This member is required.
	SupportedPermissionTypes []types.PermissionType

	// This member is required.
	TableName *string

	AuditContext *types.AuditContext

	Expression *string

	MaxResults *int32

	NextToken *string

	// Defines a non-overlapping region of a table's partitions, allowing multiple
	// requests to be run in parallel.
	Segment *types.Segment

	noSmithyDocumentSerde
}

type GetUnfilteredPartitionsMetadataOutput struct {
	NextToken *string

	UnfilteredPartitions []types.UnfilteredPartition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUnfilteredPartitionsMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUnfilteredPartitionsMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUnfilteredPartitionsMetadata{}, middleware.After)
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
	if err = addOpGetUnfilteredPartitionsMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUnfilteredPartitionsMetadata(options.Region), middleware.Before); err != nil {
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

// GetUnfilteredPartitionsMetadataAPIClient is a client that implements the
// GetUnfilteredPartitionsMetadata operation.
type GetUnfilteredPartitionsMetadataAPIClient interface {
	GetUnfilteredPartitionsMetadata(context.Context, *GetUnfilteredPartitionsMetadataInput, ...func(*Options)) (*GetUnfilteredPartitionsMetadataOutput, error)
}

var _ GetUnfilteredPartitionsMetadataAPIClient = (*Client)(nil)

// GetUnfilteredPartitionsMetadataPaginatorOptions is the paginator options for
// GetUnfilteredPartitionsMetadata
type GetUnfilteredPartitionsMetadataPaginatorOptions struct {
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetUnfilteredPartitionsMetadataPaginator is a paginator for
// GetUnfilteredPartitionsMetadata
type GetUnfilteredPartitionsMetadataPaginator struct {
	options   GetUnfilteredPartitionsMetadataPaginatorOptions
	client    GetUnfilteredPartitionsMetadataAPIClient
	params    *GetUnfilteredPartitionsMetadataInput
	nextToken *string
	firstPage bool
}

// NewGetUnfilteredPartitionsMetadataPaginator returns a new
// GetUnfilteredPartitionsMetadataPaginator
func NewGetUnfilteredPartitionsMetadataPaginator(client GetUnfilteredPartitionsMetadataAPIClient, params *GetUnfilteredPartitionsMetadataInput, optFns ...func(*GetUnfilteredPartitionsMetadataPaginatorOptions)) *GetUnfilteredPartitionsMetadataPaginator {
	if params == nil {
		params = &GetUnfilteredPartitionsMetadataInput{}
	}

	options := GetUnfilteredPartitionsMetadataPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetUnfilteredPartitionsMetadataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetUnfilteredPartitionsMetadataPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetUnfilteredPartitionsMetadata page.
func (p *GetUnfilteredPartitionsMetadataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetUnfilteredPartitionsMetadataOutput, error) {
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

	result, err := p.client.GetUnfilteredPartitionsMetadata(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetUnfilteredPartitionsMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetUnfilteredPartitionsMetadata",
	}
}
