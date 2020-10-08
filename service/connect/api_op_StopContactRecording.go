// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// When a contact is being recorded, this API stops recording the call.
// StopContactRecording is a one-time action. If you use StopContactRecording to
// stop recording an ongoing call, you can't use StartContactRecording to restart
// it. For scenarios where the recording has started and you want to suspend it for
// sensitive information (for example, to collect a credit card number), and then
// restart it, use SuspendContactRecording and ResumeContactRecording.  <p>Only
// voice recordings are supported at this time.</p>
func (c *Client) StopContactRecording(ctx context.Context, params *StopContactRecordingInput, optFns ...func(*Options)) (*StopContactRecordingOutput, error) {
	stack := middleware.NewStack("StopContactRecording", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStopContactRecordingMiddlewares(stack)
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
	addOpStopContactRecordingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopContactRecording(options.Region), middleware.Before)
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
			OperationName: "StopContactRecording",
			Err:           err,
		}
	}
	out := result.(*StopContactRecordingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopContactRecordingInput struct {

	// The identifier of the contact.
	//
	// This member is required.
	ContactId *string

	// The identifier of the contact. This is the identifier of the contact associated
	// with the first interaction with the contact center.
	//
	// This member is required.
	InitialContactId *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string
}

type StopContactRecordingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStopContactRecordingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStopContactRecording{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStopContactRecording{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopContactRecording(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "StopContactRecording",
	}
}
