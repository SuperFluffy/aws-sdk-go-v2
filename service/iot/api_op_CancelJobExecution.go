// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels the execution of a job for a given thing.
func (c *Client) CancelJobExecution(ctx context.Context, params *CancelJobExecutionInput, optFns ...func(*Options)) (*CancelJobExecutionOutput, error) {
	stack := middleware.NewStack("CancelJobExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCancelJobExecutionMiddlewares(stack)
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
	addOpCancelJobExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelJobExecution(options.Region), middleware.Before)
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
			OperationName: "CancelJobExecution",
			Err:           err,
		}
	}
	out := result.(*CancelJobExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelJobExecutionInput struct {

	// The ID of the job to be canceled.
	//
	// This member is required.
	JobId *string

	// The name of the thing whose execution of the job will be canceled.
	//
	// This member is required.
	ThingName *string

	// (Optional) The expected current version of the job execution. Each time you
	// update the job execution, its version is incremented. If the version of the job
	// execution stored in Jobs does not match, the update is rejected with a
	// VersionMismatch error, and an ErrorResponse that contains the current job
	// execution status data is returned. (This makes it unnecessary to perform a
	// separate DescribeJobExecution request in order to obtain the job execution
	// status data.)
	ExpectedVersion *int64

	// (Optional) If true the job execution will be canceled if it has status
	// IN_PROGRESS or QUEUED, otherwise the job execution will be canceled only if it
	// has status QUEUED. If you attempt to cancel a job execution that is IN_PROGRESS,
	// and you do not set force to true, then an InvalidStateTransitionException will
	// be thrown. The default is false. Canceling a job execution which is
	// "IN_PROGRESS", will cause the device to be unable to update the job execution
	// status. Use caution and ensure that the device is able to recover to a valid
	// state.
	Force *bool

	// A collection of name/value pairs that describe the status of the job execution.
	// If not specified, the statusDetails are unchanged. You can specify at most 10
	// name/value pairs.
	StatusDetails map[string]*string
}

type CancelJobExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCancelJobExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCancelJobExecution{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelJobExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelJobExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CancelJobExecution",
	}
}
