// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified resources from the specified group.
func (c *Client) UngroupResources(ctx context.Context, params *UngroupResourcesInput, optFns ...func(*Options)) (*UngroupResourcesOutput, error) {
	stack := middleware.NewStack("UngroupResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUngroupResourcesMiddlewares(stack)
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
	addOpUngroupResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUngroupResources(options.Region), middleware.Before)
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
			OperationName: "UngroupResources",
			Err:           err,
		}
	}
	out := result.(*UngroupResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UngroupResourcesInput struct {

	// The name or the ARN of the resource group from which to remove the resources.
	//
	// This member is required.
	Group *string

	// The ARNs of the resources to be removed from the group.
	//
	// This member is required.
	ResourceArns []*string
}

type UngroupResourcesOutput struct {

	// The resources that failed to be removed from the group.
	Failed []*types.FailedResource

	// The ARNs of the resources that were successfully removed from the group.
	Succeeded []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUngroupResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUngroupResources{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUngroupResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opUngroupResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "UngroupResources",
	}
}
