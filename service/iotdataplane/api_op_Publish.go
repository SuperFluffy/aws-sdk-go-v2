// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdataplane

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Publishes state information. For more information, see HTTP Protocol
// (http://docs.aws.amazon.com/iot/latest/developerguide/protocols.html#http) in
// the AWS IoT Developer Guide.
func (c *Client) Publish(ctx context.Context, params *PublishInput, optFns ...func(*Options)) (*PublishOutput, error) {
	stack := middleware.NewStack("Publish", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPublishMiddlewares(stack)
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
	addOpPublishValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPublish(options.Region), middleware.Before)
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
			OperationName: "Publish",
			Err:           err,
		}
	}
	out := result.(*PublishOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the Publish operation.
type PublishInput struct {

	// The name of the MQTT topic.
	//
	// This member is required.
	Topic *string

	// The state information, in JSON format.
	Payload []byte

	// The Quality of Service (QoS) level.
	Qos *int32
}

type PublishOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPublishMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPublish{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPublish{}, middleware.After)
}

func newServiceMetadataMiddleware_opPublish(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotdata",
		OperationName: "Publish",
	}
}
