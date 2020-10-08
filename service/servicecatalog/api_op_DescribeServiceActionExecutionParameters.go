// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Finds the default parameters for a specific self-service action on a specific
// provisioned product and returns a map of the results to the user.
func (c *Client) DescribeServiceActionExecutionParameters(ctx context.Context, params *DescribeServiceActionExecutionParametersInput, optFns ...func(*Options)) (*DescribeServiceActionExecutionParametersOutput, error) {
	stack := middleware.NewStack("DescribeServiceActionExecutionParameters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeServiceActionExecutionParametersMiddlewares(stack)
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
	addOpDescribeServiceActionExecutionParametersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServiceActionExecutionParameters(options.Region), middleware.Before)
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
			OperationName: "DescribeServiceActionExecutionParameters",
			Err:           err,
		}
	}
	out := result.(*DescribeServiceActionExecutionParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeServiceActionExecutionParametersInput struct {

	// The identifier of the provisioned product.
	//
	// This member is required.
	ProvisionedProductId *string

	// The self-service action identifier.
	//
	// This member is required.
	ServiceActionId *string

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
}

type DescribeServiceActionExecutionParametersOutput struct {

	// The parameters of the self-service action.
	ServiceActionParameters []*types.ExecutionParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeServiceActionExecutionParametersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeServiceActionExecutionParameters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeServiceActionExecutionParameters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeServiceActionExecutionParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DescribeServiceActionExecutionParameters",
	}
}
