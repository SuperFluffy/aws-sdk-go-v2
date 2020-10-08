// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of journal export job descriptions for a specified ledger. This
// action returns a maximum of MaxResults items, and is paginated so that you can
// retrieve all the items by calling ListJournalS3ExportsForLedger multiple times.
// This action does not return any expired export jobs. For more information, see
// Export Job Expiration
// (https://docs.aws.amazon.com/qldb/latest/developerguide/export-journal.request.html#export-journal.request.expiration)
// in the Amazon QLDB Developer Guide.
func (c *Client) ListJournalS3ExportsForLedger(ctx context.Context, params *ListJournalS3ExportsForLedgerInput, optFns ...func(*Options)) (*ListJournalS3ExportsForLedgerOutput, error) {
	stack := middleware.NewStack("ListJournalS3ExportsForLedger", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListJournalS3ExportsForLedgerMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListJournalS3ExportsForLedgerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListJournalS3ExportsForLedger(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListJournalS3ExportsForLedger",
			Err:           err,
		}
	}
	out := result.(*ListJournalS3ExportsForLedgerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJournalS3ExportsForLedgerInput struct {

	// The name of the ledger.
	//
	// This member is required.
	Name *string

	// The maximum number of results to return in a single
	// ListJournalS3ExportsForLedger request. (The actual number of results returned
	// might be fewer.)
	MaxResults *int32

	// A pagination token, indicating that you want to retrieve the next page of
	// results. If you received a value for NextToken in the response from a previous
	// ListJournalS3ExportsForLedger call, then you should use that value as input
	// here.
	NextToken *string
}

type ListJournalS3ExportsForLedgerOutput struct {

	// The array of journal export job descriptions that are associated with the
	// specified ledger.
	JournalS3Exports []*types.JournalS3ExportDescription

	// * If NextToken is empty, then the last page of results has been processed and
	// there are no more results to be retrieved.
	//
	//     * If NextToken is not empty,
	// then there are more results available. To retrieve the next page of results, use
	// the value of NextToken in a subsequent ListJournalS3ExportsForLedger call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListJournalS3ExportsForLedgerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListJournalS3ExportsForLedger{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListJournalS3ExportsForLedger{}, middleware.After)
}

func newServiceMetadataMiddleware_opListJournalS3ExportsForLedger(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ListJournalS3ExportsForLedger",
	}
}
