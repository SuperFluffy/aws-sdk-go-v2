// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete a list of parameters.
func (c *Client) DeleteParameters(ctx context.Context, params *DeleteParametersInput, optFns ...func(*Options)) (*DeleteParametersOutput, error) {
	stack := middleware.NewStack("DeleteParameters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteParametersMiddlewares(stack)
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
	addOpDeleteParametersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteParameters(options.Region), middleware.Before)
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
			OperationName: "DeleteParameters",
			Err:           err,
		}
	}
	out := result.(*DeleteParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteParametersInput struct {

	// The names of the parameters to delete.
	//
	// This member is required.
	Names []*string
}

type DeleteParametersOutput struct {

	// The names of the deleted parameters.
	DeletedParameters []*string

	// The names of parameters that weren't deleted because the parameters are not
	// valid.
	InvalidParameters []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteParametersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteParameters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteParameters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DeleteParameters",
	}
}
