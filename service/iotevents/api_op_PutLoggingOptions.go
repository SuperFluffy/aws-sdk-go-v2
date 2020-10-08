// Code generated by smithy-go-codegen DO NOT EDIT.

package iotevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotevents/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets or updates the AWS IoT Events logging options. If you update the value of
// any loggingOptions field, it takes up to one minute for the change to take
// effect. If you change the policy attached to the role you specified in the
// roleArn field (for example, to correct an invalid policy), it takes up to five
// minutes for that change to take effect.
func (c *Client) PutLoggingOptions(ctx context.Context, params *PutLoggingOptionsInput, optFns ...func(*Options)) (*PutLoggingOptionsOutput, error) {
	stack := middleware.NewStack("PutLoggingOptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutLoggingOptionsMiddlewares(stack)
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
	addOpPutLoggingOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutLoggingOptions(options.Region), middleware.Before)
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
			OperationName: "PutLoggingOptions",
			Err:           err,
		}
	}
	out := result.(*PutLoggingOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLoggingOptionsInput struct {

	// The new values of the AWS IoT Events logging options.
	//
	// This member is required.
	LoggingOptions *types.LoggingOptions
}

type PutLoggingOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutLoggingOptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutLoggingOptions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLoggingOptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutLoggingOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotevents",
		OperationName: "PutLoggingOptions",
	}
}
