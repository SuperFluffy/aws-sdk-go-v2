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

// Enables fast snapshot restores for the specified snapshots in the specified
// Availability Zones. You get the full benefit of fast snapshot restores after
// they enter the enabled state. To get the current state of fast snapshot
// restores, use DescribeFastSnapshotRestores (). To disable fast snapshot
// restores, use DisableFastSnapshotRestores (). For more information, see Amazon
// EBS Fast Snapshot Restore
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-fast-snapshot-restore.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) EnableFastSnapshotRestores(ctx context.Context, params *EnableFastSnapshotRestoresInput, optFns ...func(*Options)) (*EnableFastSnapshotRestoresOutput, error) {
	stack := middleware.NewStack("EnableFastSnapshotRestores", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpEnableFastSnapshotRestoresMiddlewares(stack)
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
	addOpEnableFastSnapshotRestoresValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableFastSnapshotRestores(options.Region), middleware.Before)
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
			OperationName: "EnableFastSnapshotRestores",
			Err:           err,
		}
	}
	out := result.(*EnableFastSnapshotRestoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableFastSnapshotRestoresInput struct {

	// One or more Availability Zones. For example, us-east-2a.
	//
	// This member is required.
	AvailabilityZones []*string

	// The IDs of one or more snapshots. For example, snap-1234567890abcdef0. You can
	// specify a snapshot that was shared with you from another AWS account.
	//
	// This member is required.
	SourceSnapshotIds []*string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type EnableFastSnapshotRestoresOutput struct {

	// Information about the snapshots for which fast snapshot restores were
	// successfully enabled.
	Successful []*types.EnableFastSnapshotRestoreSuccessItem

	// Information about the snapshots for which fast snapshot restores could not be
	// enabled.
	Unsuccessful []*types.EnableFastSnapshotRestoreErrorItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpEnableFastSnapshotRestoresMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpEnableFastSnapshotRestores{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpEnableFastSnapshotRestores{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableFastSnapshotRestores(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "EnableFastSnapshotRestores",
	}
}
