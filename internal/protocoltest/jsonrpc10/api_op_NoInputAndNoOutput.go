// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The example tests how requests and responses are serialized when there's no
// request or response payload because the operation has no input or output. While
// this should be rare, code generators must support this.
func (c *Client) NoInputAndNoOutput(ctx context.Context, params *NoInputAndNoOutputInput, optFns ...func(*Options)) (*NoInputAndNoOutputOutput, error) {
	stack := middleware.NewStack("NoInputAndNoOutput", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpNoInputAndNoOutputMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opNoInputAndNoOutput(options.Region), middleware.Before)
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
			OperationName: "NoInputAndNoOutput",
			Err:           err,
		}
	}
	out := result.(*NoInputAndNoOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NoInputAndNoOutputInput struct {
}

type NoInputAndNoOutputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpNoInputAndNoOutputMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpNoInputAndNoOutput{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpNoInputAndNoOutput{}, middleware.After)
}

func newServiceMetadataMiddleware_opNoInputAndNoOutput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NoInputAndNoOutput",
	}
}
