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

// Sets the default authorizer. This will be used if a websocket connection is made
// without specifying an authorizer.
func (c *Client) SetDefaultAuthorizer(ctx context.Context, params *SetDefaultAuthorizerInput, optFns ...func(*Options)) (*SetDefaultAuthorizerOutput, error) {
	stack := middleware.NewStack("SetDefaultAuthorizer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSetDefaultAuthorizerMiddlewares(stack)
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
	addOpSetDefaultAuthorizerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetDefaultAuthorizer(options.Region), middleware.Before)
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
			OperationName: "SetDefaultAuthorizer",
			Err:           err,
		}
	}
	out := result.(*SetDefaultAuthorizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetDefaultAuthorizerInput struct {

	// The authorizer name.
	//
	// This member is required.
	AuthorizerName *string
}

type SetDefaultAuthorizerOutput struct {

	// The authorizer ARN.
	AuthorizerArn *string

	// The authorizer name.
	AuthorizerName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSetDefaultAuthorizerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSetDefaultAuthorizer{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSetDefaultAuthorizer{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetDefaultAuthorizer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "SetDefaultAuthorizer",
	}
}
