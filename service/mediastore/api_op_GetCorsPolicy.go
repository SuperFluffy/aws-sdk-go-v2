// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastore

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediastore/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the cross-origin resource sharing (CORS) configuration information that
// is set for the container. To use this operation, you must have permission to
// perform the MediaStore:GetCorsPolicy action. By default, the container owner has
// this permission and can grant it to others.
func (c *Client) GetCorsPolicy(ctx context.Context, params *GetCorsPolicyInput, optFns ...func(*Options)) (*GetCorsPolicyOutput, error) {
	stack := middleware.NewStack("GetCorsPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetCorsPolicyMiddlewares(stack)
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
	addOpGetCorsPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCorsPolicy(options.Region), middleware.Before)
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
			OperationName: "GetCorsPolicy",
			Err:           err,
		}
	}
	out := result.(*GetCorsPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCorsPolicyInput struct {

	// The name of the container that the policy is assigned to.
	//
	// This member is required.
	ContainerName *string
}

type GetCorsPolicyOutput struct {

	// The CORS policy assigned to the container.
	//
	// This member is required.
	CorsPolicy []*types.CorsRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetCorsPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetCorsPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCorsPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCorsPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediastore",
		OperationName: "GetCorsPolicy",
	}
}
