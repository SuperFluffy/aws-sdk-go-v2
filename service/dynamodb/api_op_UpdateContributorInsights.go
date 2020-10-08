// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the status for contributor insights for a specific table or index.
func (c *Client) UpdateContributorInsights(ctx context.Context, params *UpdateContributorInsightsInput, optFns ...func(*Options)) (*UpdateContributorInsightsOutput, error) {
	stack := middleware.NewStack("UpdateContributorInsights", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpUpdateContributorInsightsMiddlewares(stack)
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
	addOpUpdateContributorInsightsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContributorInsights(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)

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
			OperationName: "UpdateContributorInsights",
			Err:           err,
		}
	}
	out := result.(*UpdateContributorInsightsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContributorInsightsInput struct {

	// Represents the contributor insights action.
	//
	// This member is required.
	ContributorInsightsAction types.ContributorInsightsAction

	// The name of the table.
	//
	// This member is required.
	TableName *string

	// The global secondary index name, if applicable.
	IndexName *string
}

type UpdateContributorInsightsOutput struct {

	// The status of contributor insights
	ContributorInsightsStatus types.ContributorInsightsStatus

	// The name of the global secondary index, if applicable.
	IndexName *string

	// The name of the table.
	TableName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpUpdateContributorInsightsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateContributorInsights{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateContributorInsights{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateContributorInsights(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "UpdateContributorInsights",
	}
}
