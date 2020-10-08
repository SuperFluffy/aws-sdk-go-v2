// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a member from a chat room in an Amazon Chime Enterprise account.
func (c *Client) DeleteRoomMembership(ctx context.Context, params *DeleteRoomMembershipInput, optFns ...func(*Options)) (*DeleteRoomMembershipOutput, error) {
	stack := middleware.NewStack("DeleteRoomMembership", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteRoomMembershipMiddlewares(stack)
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
	addOpDeleteRoomMembershipValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRoomMembership(options.Region), middleware.Before)
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
			OperationName: "DeleteRoomMembership",
			Err:           err,
		}
	}
	out := result.(*DeleteRoomMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRoomMembershipInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The member ID (user ID or bot ID).
	//
	// This member is required.
	MemberId *string

	// The room ID.
	//
	// This member is required.
	RoomId *string
}

type DeleteRoomMembershipOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteRoomMembershipMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteRoomMembership{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteRoomMembership{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteRoomMembership(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DeleteRoomMembership",
	}
}
