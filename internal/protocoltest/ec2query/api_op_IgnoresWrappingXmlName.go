// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The xmlName trait on the output structure is ignored in AWS Query. The wrapping
// element is always operation name + "Response".
func (c *Client) IgnoresWrappingXmlName(ctx context.Context, params *IgnoresWrappingXmlNameInput, optFns ...func(*Options)) (*IgnoresWrappingXmlNameOutput, error) {
	stack := middleware.NewStack("IgnoresWrappingXmlName", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpIgnoresWrappingXmlNameMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opIgnoresWrappingXmlName(options.Region), middleware.Before)
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
			OperationName: "IgnoresWrappingXmlName",
			Err:           err,
		}
	}
	out := result.(*IgnoresWrappingXmlNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IgnoresWrappingXmlNameInput struct {
}

type IgnoresWrappingXmlNameOutput struct {
	Foo *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpIgnoresWrappingXmlNameMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpIgnoresWrappingXmlName{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpIgnoresWrappingXmlName{}, middleware.After)
}

func newServiceMetadataMiddleware_opIgnoresWrappingXmlName(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "IgnoresWrappingXmlName",
	}
}
