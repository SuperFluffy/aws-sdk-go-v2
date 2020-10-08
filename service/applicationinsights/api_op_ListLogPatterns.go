// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the log patterns in the specific log LogPatternSet.
func (c *Client) ListLogPatterns(ctx context.Context, params *ListLogPatternsInput, optFns ...func(*Options)) (*ListLogPatternsOutput, error) {
	stack := middleware.NewStack("ListLogPatterns", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListLogPatternsMiddlewares(stack)
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
	addOpListLogPatternsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListLogPatterns(options.Region), middleware.Before)
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
			OperationName: "ListLogPatterns",
			Err:           err,
		}
	}
	out := result.(*ListLogPatternsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLogPatternsInput struct {

	// The name of the resource group.
	//
	// This member is required.
	ResourceGroupName *string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The name of the log pattern set.
	PatternSetName *string
}

type ListLogPatternsOutput struct {

	// The list of log patterns.
	LogPatterns []*types.LogPattern

	// The token used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The name of the resource group.
	ResourceGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListLogPatternsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListLogPatterns{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLogPatterns{}, middleware.After)
}

func newServiceMetadataMiddleware_opListLogPatterns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "ListLogPatterns",
	}
}
