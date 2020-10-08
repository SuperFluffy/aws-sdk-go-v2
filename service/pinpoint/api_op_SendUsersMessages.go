// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates and sends a message to a list of users.
func (c *Client) SendUsersMessages(ctx context.Context, params *SendUsersMessagesInput, optFns ...func(*Options)) (*SendUsersMessagesOutput, error) {
	stack := middleware.NewStack("SendUsersMessages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSendUsersMessagesMiddlewares(stack)
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
	addOpSendUsersMessagesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendUsersMessages(options.Region), middleware.Before)
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
			OperationName: "SendUsersMessages",
			Err:           err,
		}
	}
	out := result.(*SendUsersMessagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendUsersMessagesInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// Specifies the configuration and other settings for a message to send to all the
	// endpoints that are associated with a list of users.
	//
	// This member is required.
	SendUsersMessageRequest *types.SendUsersMessageRequest
}

type SendUsersMessagesOutput struct {

	// Provides information about which users and endpoints a message was sent to.
	//
	// This member is required.
	SendUsersMessageResponse *types.SendUsersMessageResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSendUsersMessagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSendUsersMessages{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSendUsersMessages{}, middleware.After)
}

func newServiceMetadataMiddleware_opSendUsersMessages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "SendUsersMessages",
	}
}
