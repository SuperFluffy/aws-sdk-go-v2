// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stores customer feedback for a CodeGuru Reviewer recommendation. When this API
// is called again with different reactions the previous feedback is overwritten.
func (c *Client) PutRecommendationFeedback(ctx context.Context, params *PutRecommendationFeedbackInput, optFns ...func(*Options)) (*PutRecommendationFeedbackOutput, error) {
	stack := middleware.NewStack("PutRecommendationFeedback", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutRecommendationFeedbackMiddlewares(stack)
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
	addOpPutRecommendationFeedbackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutRecommendationFeedback(options.Region), middleware.Before)
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
			OperationName: "PutRecommendationFeedback",
			Err:           err,
		}
	}
	out := result.(*PutRecommendationFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRecommendationFeedbackInput struct {

	// The Amazon Resource Name (ARN) of the CodeReview
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html)
	// object.
	//
	// This member is required.
	CodeReviewArn *string

	// List for storing reactions. Reactions are utf-8 text code for emojis. If you
	// send an empty list it clears all your feedback.
	//
	// This member is required.
	Reactions []types.Reaction

	// The recommendation ID that can be used to track the provided recommendations and
	// then to collect the feedback.
	//
	// This member is required.
	RecommendationId *string
}

type PutRecommendationFeedbackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutRecommendationFeedbackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutRecommendationFeedback{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutRecommendationFeedback{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutRecommendationFeedback(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "PutRecommendationFeedback",
	}
}
