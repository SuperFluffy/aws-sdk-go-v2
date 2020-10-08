// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Deletes a snapshot of application state.
func (c *Client) DeleteApplicationSnapshot(ctx context.Context, params *DeleteApplicationSnapshotInput, optFns ...func(*Options)) (*DeleteApplicationSnapshotOutput, error) {
	stack := middleware.NewStack("DeleteApplicationSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteApplicationSnapshotMiddlewares(stack)
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
	addOpDeleteApplicationSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationSnapshot(options.Region), middleware.Before)
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
			OperationName: "DeleteApplicationSnapshot",
			Err:           err,
		}
	}
	out := result.(*DeleteApplicationSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationSnapshotInput struct {

	// The name of an existing application.
	//
	// This member is required.
	ApplicationName *string

	// The creation timestamp of the application snapshot to delete. You can retrieve
	// this value using or .
	//
	// This member is required.
	SnapshotCreationTimestamp *time.Time

	// The identifier for the snapshot delete.
	//
	// This member is required.
	SnapshotName *string
}

type DeleteApplicationSnapshotOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteApplicationSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteApplicationSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DeleteApplicationSnapshot",
	}
}
