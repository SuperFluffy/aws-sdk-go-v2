// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the state of fast snapshot restores for your snapshots.
func (c *Client) DescribeFastSnapshotRestores(ctx context.Context, params *DescribeFastSnapshotRestoresInput, optFns ...func(*Options)) (*DescribeFastSnapshotRestoresOutput, error) {
	stack := middleware.NewStack("DescribeFastSnapshotRestores", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeFastSnapshotRestoresMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFastSnapshotRestores(options.Region), middleware.Before)
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
			OperationName: "DescribeFastSnapshotRestores",
			Err:           err,
		}
	}
	out := result.(*DescribeFastSnapshotRestoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFastSnapshotRestoresInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters. The possible values are:
	//
	//     * availability-zone: The Availability
	// Zone of the snapshot.
	//
	//     * owner-id: The ID of the AWS account that enabled
	// fast snapshot restore on the snapshot.
	//
	//     * snapshot-id: The ID of the
	// snapshot.
	//
	//     * state: The state of fast snapshot restores for the snapshot
	// (enabling | optimizing | enabled | disabling | disabled).
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type DescribeFastSnapshotRestoresOutput struct {

	// Information about the state of fast snapshot restores.
	FastSnapshotRestores []*types.DescribeFastSnapshotRestoreSuccessItem

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeFastSnapshotRestoresMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeFastSnapshotRestores{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeFastSnapshotRestores{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeFastSnapshotRestores(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeFastSnapshotRestores",
	}
}
