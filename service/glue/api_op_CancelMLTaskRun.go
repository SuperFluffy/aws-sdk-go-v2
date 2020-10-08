// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels (stops) a task run. Machine learning task runs are asynchronous tasks
// that AWS Glue runs on your behalf as part of various machine learning workflows.
// You can cancel a machine learning task run at any time by calling
// CancelMLTaskRun with a task run's parent transform's TransformID and the task
// run's TaskRunId.
func (c *Client) CancelMLTaskRun(ctx context.Context, params *CancelMLTaskRunInput, optFns ...func(*Options)) (*CancelMLTaskRunOutput, error) {
	stack := middleware.NewStack("CancelMLTaskRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCancelMLTaskRunMiddlewares(stack)
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
	addOpCancelMLTaskRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelMLTaskRun(options.Region), middleware.Before)
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
			OperationName: "CancelMLTaskRun",
			Err:           err,
		}
	}
	out := result.(*CancelMLTaskRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelMLTaskRunInput struct {

	// A unique identifier for the task run.
	//
	// This member is required.
	TaskRunId *string

	// The unique identifier of the machine learning transform.
	//
	// This member is required.
	TransformId *string
}

type CancelMLTaskRunOutput struct {

	// The status for this run.
	Status types.TaskStatusType

	// The unique identifier for the task run.
	TaskRunId *string

	// The unique identifier of the machine learning transform.
	TransformId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCancelMLTaskRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCancelMLTaskRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelMLTaskRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelMLTaskRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CancelMLTaskRun",
	}
}
