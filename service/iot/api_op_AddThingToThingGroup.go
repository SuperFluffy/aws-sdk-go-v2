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

// Adds a thing to a thing group.
func (c *Client) AddThingToThingGroup(ctx context.Context, params *AddThingToThingGroupInput, optFns ...func(*Options)) (*AddThingToThingGroupOutput, error) {
	stack := middleware.NewStack("AddThingToThingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAddThingToThingGroupMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddThingToThingGroup(options.Region), middleware.Before)
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
			OperationName: "AddThingToThingGroup",
			Err:           err,
		}
	}
	out := result.(*AddThingToThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddThingToThingGroupInput struct {

	// Override dynamic thing groups with static thing groups when 10-group limit is
	// reached. If a thing belongs to 10 thing groups, and one or more of those groups
	// are dynamic thing groups, adding a thing to a static group removes the thing
	// from the last dynamic group.
	OverrideDynamicGroups *bool

	// The ARN of the thing to add to a group.
	ThingArn *string

	// The ARN of the group to which you are adding a thing.
	ThingGroupArn *string

	// The name of the group to which you are adding a thing.
	ThingGroupName *string

	// The name of the thing to add to a group.
	ThingName *string
}

type AddThingToThingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAddThingToThingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAddThingToThingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAddThingToThingGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddThingToThingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "AddThingToThingGroup",
	}
}
