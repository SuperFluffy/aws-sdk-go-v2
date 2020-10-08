// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or overwrites one or more tags for the specified Amazon SageMaker resource.
// You can add tags to notebook instances, training jobs, hyperparameter tuning
// jobs, batch transform jobs, models, labeling jobs, work teams, endpoint
// configurations, and endpoints. Each tag consists of a key and an optional value.
// Tag keys must be unique per resource. For more information about tags, see For
// more information, see AWS Tagging Strategies
// (https://aws.amazon.com/answers/account-management/aws-tagging-strategies/).
// Tags that you add to a hyperparameter tuning job by calling this API are also
// added to any training jobs that the hyperparameter tuning job launches after you
// call this API, but not to training jobs that the hyperparameter tuning job
// launched before you called this API. To make sure that the tags associated with
// a hyperparameter tuning job are also added to all training jobs that the
// hyperparameter tuning job launches, add the tags when you first create the
// tuning job by specifying them in the Tags parameter of
// CreateHyperParameterTuningJob ()
func (c *Client) AddTags(ctx context.Context, params *AddTagsInput, optFns ...func(*Options)) (*AddTagsOutput, error) {
	stack := middleware.NewStack("AddTags", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddTagsMiddlewares(stack)
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
	addOpAddTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddTags(options.Region), middleware.Before)
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
			OperationName: "AddTags",
			Err:           err,
		}
	}
	out := result.(*AddTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddTagsInput struct {

	// The Amazon Resource Name (ARN) of the resource that you want to tag.
	//
	// This member is required.
	ResourceArn *string

	// An array of Tag objects. Each tag is a key-value pair. Only the key parameter is
	// required. If you don't specify a value, Amazon SageMaker sets the value to an
	// empty string.
	//
	// This member is required.
	Tags []*types.Tag
}

type AddTagsOutput struct {

	// A list of tags associated with the Amazon SageMaker resource.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddTagsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddTags{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddTags{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "AddTags",
	}
}
