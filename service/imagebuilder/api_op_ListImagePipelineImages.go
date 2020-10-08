// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of images created by the specified pipeline.
func (c *Client) ListImagePipelineImages(ctx context.Context, params *ListImagePipelineImagesInput, optFns ...func(*Options)) (*ListImagePipelineImagesOutput, error) {
	stack := middleware.NewStack("ListImagePipelineImages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListImagePipelineImagesMiddlewares(stack)
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
	addOpListImagePipelineImagesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListImagePipelineImages(options.Region), middleware.Before)
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
			OperationName: "ListImagePipelineImages",
			Err:           err,
		}
	}
	out := result.(*ListImagePipelineImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImagePipelineImagesInput struct {

	// The Amazon Resource Name (ARN) of the image pipeline whose images you want to
	// view.
	//
	// This member is required.
	ImagePipelineArn *string

	// The filters.
	Filters []*types.Filter

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string
}

type ListImagePipelineImagesOutput struct {

	// The list of images built by this pipeline.
	ImageSummaryList []*types.ImageSummary

	// The next token used for paginated responses. When this is not empty, there are
	// additional elements that the service has not included in this request. Use this
	// token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListImagePipelineImagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListImagePipelineImages{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListImagePipelineImages{}, middleware.After)
}

func newServiceMetadataMiddleware_opListImagePipelineImages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListImagePipelineImages",
	}
}
