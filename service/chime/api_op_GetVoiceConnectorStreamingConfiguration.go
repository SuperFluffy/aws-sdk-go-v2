// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the streaming configuration details for the specified Amazon Chime
// Voice Connector. Shows whether media streaming is enabled for sending to Amazon
// Kinesis. It also shows the retention period, in hours, for the Amazon Kinesis
// data.
func (c *Client) GetVoiceConnectorStreamingConfiguration(ctx context.Context, params *GetVoiceConnectorStreamingConfigurationInput, optFns ...func(*Options)) (*GetVoiceConnectorStreamingConfigurationOutput, error) {
	stack := middleware.NewStack("GetVoiceConnectorStreamingConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetVoiceConnectorStreamingConfigurationMiddlewares(stack)
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
	addOpGetVoiceConnectorStreamingConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetVoiceConnectorStreamingConfiguration(options.Region), middleware.Before)
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
			OperationName: "GetVoiceConnectorStreamingConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetVoiceConnectorStreamingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVoiceConnectorStreamingConfigurationInput struct {

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string
}

type GetVoiceConnectorStreamingConfigurationOutput struct {

	// The streaming configuration details.
	StreamingConfiguration *types.StreamingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetVoiceConnectorStreamingConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetVoiceConnectorStreamingConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVoiceConnectorStreamingConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetVoiceConnectorStreamingConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetVoiceConnectorStreamingConfiguration",
	}
}
