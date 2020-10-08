// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels an existing Amazon FSx for Lustre data repository task if that task is
// in either the PENDING or EXECUTING state. When you cancel a task, Amazon FSx
// does the following.
//
//     * Any files that FSx has already exported are not
// reverted.
//
//     * FSx continues to export any files that are "in-flight" when the
// cancel operation is received.
//
//     * FSx does not export any files that have not
// yet been exported.
func (c *Client) CancelDataRepositoryTask(ctx context.Context, params *CancelDataRepositoryTaskInput, optFns ...func(*Options)) (*CancelDataRepositoryTaskOutput, error) {
	stack := middleware.NewStack("CancelDataRepositoryTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCancelDataRepositoryTaskMiddlewares(stack)
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
	addOpCancelDataRepositoryTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelDataRepositoryTask(options.Region), middleware.Before)
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
			OperationName: "CancelDataRepositoryTask",
			Err:           err,
		}
	}
	out := result.(*CancelDataRepositoryTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Cancels a data repository task.
type CancelDataRepositoryTaskInput struct {

	// Specifies the data repository task to cancel.
	//
	// This member is required.
	TaskId *string
}

type CancelDataRepositoryTaskOutput struct {

	// The lifecycle status of the data repository task, as follows:
	//
	//     * PENDING -
	// Amazon FSx has not started the task.
	//
	//     * EXECUTING - Amazon FSx is processing
	// the task.
	//
	//     * FAILED - Amazon FSx was not able to complete the task. For
	// example, there may be files the task failed to process. The
	// DataRepositoryTaskFailureDetails () property provides more information about
	// task failures.
	//
	//     * SUCCEEDED - FSx completed the task successfully.
	//
	//     *
	// CANCELED - Amazon FSx canceled the task and it did not complete.
	//
	//     *
	// CANCELING - FSx is in process of canceling the task.
	Lifecycle types.DataRepositoryTaskLifecycle

	// The ID of the task being canceled.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCancelDataRepositoryTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCancelDataRepositoryTask{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelDataRepositoryTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelDataRepositoryTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "CancelDataRepositoryTask",
	}
}
