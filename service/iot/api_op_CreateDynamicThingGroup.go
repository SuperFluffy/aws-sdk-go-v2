// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a dynamic thing group.
func (c *Client) CreateDynamicThingGroup(ctx context.Context, params *CreateDynamicThingGroupInput, optFns ...func(*Options)) (*CreateDynamicThingGroupOutput, error) {
	stack := middleware.NewStack("CreateDynamicThingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDynamicThingGroupMiddlewares(stack)
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
	addOpCreateDynamicThingGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDynamicThingGroup(options.Region), middleware.Before)
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
			OperationName: "CreateDynamicThingGroup",
			Err:           err,
		}
	}
	out := result.(*CreateDynamicThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDynamicThingGroupInput struct {

	// The dynamic thing group search query string. See Query Syntax
	// (https://docs.aws.amazon.com/iot/latest/developerguide/query-syntax.html) for
	// information about query string syntax.
	//
	// This member is required.
	QueryString *string

	// The dynamic thing group name to create.
	//
	// This member is required.
	ThingGroupName *string

	// The dynamic thing group index name. Currently one index is supported:
	// "AWS_Things".
	IndexName *string

	// The dynamic thing group query version. Currently one query version is supported:
	// "2017-09-30". If not specified, the query version defaults to this value.
	QueryVersion *string

	// Metadata which can be used to manage the dynamic thing group.
	Tags []*types.Tag

	// The dynamic thing group properties.
	ThingGroupProperties *types.ThingGroupProperties
}

type CreateDynamicThingGroupOutput struct {

	// The dynamic thing group index name.
	IndexName *string

	// The dynamic thing group search query string.
	QueryString *string

	// The dynamic thing group query version.
	QueryVersion *string

	// The dynamic thing group ARN.
	ThingGroupArn *string

	// The dynamic thing group ID.
	ThingGroupId *string

	// The dynamic thing group name.
	ThingGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDynamicThingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDynamicThingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDynamicThingGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDynamicThingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateDynamicThingGroup",
	}
}
