// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the name of an inbound or an outbound resolver endpoint.
func (c *Client) UpdateResolverEndpoint(ctx context.Context, params *UpdateResolverEndpointInput, optFns ...func(*Options)) (*UpdateResolverEndpointOutput, error) {
	stack := middleware.NewStack("UpdateResolverEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateResolverEndpointMiddlewares(stack)
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
	addOpUpdateResolverEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResolverEndpoint(options.Region), middleware.Before)
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
			OperationName: "UpdateResolverEndpoint",
			Err:           err,
		}
	}
	out := result.(*UpdateResolverEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResolverEndpointInput struct {

	// The ID of the resolver endpoint that you want to update.
	//
	// This member is required.
	ResolverEndpointId *string

	// The name of the resolver endpoint that you want to update.
	Name *string
}

type UpdateResolverEndpointOutput struct {

	// The response to an UpdateResolverEndpoint request.
	ResolverEndpoint *types.ResolverEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateResolverEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateResolverEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateResolverEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateResolverEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "UpdateResolverEndpoint",
	}
}
