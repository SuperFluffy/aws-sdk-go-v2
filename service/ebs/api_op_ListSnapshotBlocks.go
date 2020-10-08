// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the block indexes and block tokens for blocks in an Amazon Elastic Block
// Store snapshot.
func (c *Client) ListSnapshotBlocks(ctx context.Context, params *ListSnapshotBlocksInput, optFns ...func(*Options)) (*ListSnapshotBlocksOutput, error) {
	stack := middleware.NewStack("ListSnapshotBlocks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListSnapshotBlocksMiddlewares(stack)
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
	addOpListSnapshotBlocksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSnapshotBlocks(options.Region), middleware.Before)
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
			OperationName: "ListSnapshotBlocks",
			Err:           err,
		}
	}
	out := result.(*ListSnapshotBlocksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSnapshotBlocksInput struct {

	// The ID of the snapshot from which to get block indexes and block tokens.
	//
	// This member is required.
	SnapshotId *string

	// The number of results to return.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The block index from which the list should start. The list in the response will
	// start from this block index or the next valid block index in the snapshot.
	StartingBlockIndex *int32
}

type ListSnapshotBlocksOutput struct {

	// The size of the block.
	BlockSize *int32

	// An array of objects containing information about the blocks.
	Blocks []*types.Block

	// The time when the BlockToken expires.
	ExpiryTime *time.Time

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The size of the volume in GB.
	VolumeSize *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListSnapshotBlocksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListSnapshotBlocks{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListSnapshotBlocks{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSnapshotBlocks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ebs",
		OperationName: "ListSnapshotBlocks",
	}
}
