// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of ChannelInfo objects. Each object describes a signaling
// channel. To retrieve only those channels that satisfy a specific condition, you
// can specify a ChannelNameCondition.
func (c *Client) ListSignalingChannels(ctx context.Context, params *ListSignalingChannelsInput, optFns ...func(*Options)) (*ListSignalingChannelsOutput, error) {
	stack := middleware.NewStack("ListSignalingChannels", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListSignalingChannelsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSignalingChannels(options.Region), middleware.Before)
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
			OperationName: "ListSignalingChannels",
			Err:           err,
		}
	}
	out := result.(*ListSignalingChannelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSignalingChannelsInput struct {

	// Optional: Returns only the channels that satisfy a specific condition.
	ChannelNameCondition *types.ChannelNameCondition

	// The maximum number of channels to return in the response. The default is 500.
	MaxResults *int32

	// If you specify this parameter, when the result of a ListSignalingChannels
	// operation is truncated, the call returns the NextToken in the response. To get
	// another batch of channels, provide this token in your next request.
	NextToken *string
}

type ListSignalingChannelsOutput struct {

	// An array of ChannelInfo objects.
	ChannelInfoList []*types.ChannelInfo

	// If the response is truncated, the call returns this element with a token. To get
	// the next batch of streams, use this token in your next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListSignalingChannelsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListSignalingChannels{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListSignalingChannels{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSignalingChannels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "ListSignalingChannels",
	}
}
