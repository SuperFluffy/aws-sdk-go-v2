// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an Amazon Pinpoint configuration for a recommender model.
func (c *Client) DeleteRecommenderConfiguration(ctx context.Context, params *DeleteRecommenderConfigurationInput, optFns ...func(*Options)) (*DeleteRecommenderConfigurationOutput, error) {
	stack := middleware.NewStack("DeleteRecommenderConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteRecommenderConfigurationMiddlewares(stack)
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
	addOpDeleteRecommenderConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRecommenderConfiguration(options.Region), middleware.Before)
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
			OperationName: "DeleteRecommenderConfiguration",
			Err:           err,
		}
	}
	out := result.(*DeleteRecommenderConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRecommenderConfigurationInput struct {

	// The unique identifier for the recommender model configuration. This identifier
	// is displayed as the Recommender ID on the Amazon Pinpoint console.
	//
	// This member is required.
	RecommenderId *string
}

type DeleteRecommenderConfigurationOutput struct {

	// Provides information about Amazon Pinpoint configuration settings for retrieving
	// and processing data from a recommender model.
	//
	// This member is required.
	RecommenderConfigurationResponse *types.RecommenderConfigurationResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteRecommenderConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteRecommenderConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteRecommenderConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteRecommenderConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "DeleteRecommenderConfiguration",
	}
}
