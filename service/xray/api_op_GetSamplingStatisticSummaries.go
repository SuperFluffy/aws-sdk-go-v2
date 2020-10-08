// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about recent sampling results for all sampling rules.
func (c *Client) GetSamplingStatisticSummaries(ctx context.Context, params *GetSamplingStatisticSummariesInput, optFns ...func(*Options)) (*GetSamplingStatisticSummariesOutput, error) {
	stack := middleware.NewStack("GetSamplingStatisticSummaries", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSamplingStatisticSummariesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSamplingStatisticSummaries(options.Region), middleware.Before)
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
			OperationName: "GetSamplingStatisticSummaries",
			Err:           err,
		}
	}
	out := result.(*GetSamplingStatisticSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSamplingStatisticSummariesInput struct {

	// Pagination token.
	NextToken *string
}

type GetSamplingStatisticSummariesOutput struct {

	// Pagination token.
	NextToken *string

	// Information about the number of requests instrumented for each sampling rule.
	SamplingStatisticSummaries []*types.SamplingStatisticSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSamplingStatisticSummariesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSamplingStatisticSummaries{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSamplingStatisticSummaries{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSamplingStatisticSummaries(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetSamplingStatisticSummaries",
	}
}
