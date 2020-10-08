// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns metadata about the images in a repository. Beginning with Docker version
// 1.9, the Docker client compresses image layers before pushing them to a V2
// Docker registry. The output of the docker images command shows the uncompressed
// image size, so it may return a larger image size than the image sizes returned
// by DescribeImages ().
func (c *Client) DescribeImages(ctx context.Context, params *DescribeImagesInput, optFns ...func(*Options)) (*DescribeImagesOutput, error) {
	stack := middleware.NewStack("DescribeImages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeImagesMiddlewares(stack)
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
	addOpDescribeImagesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImages(options.Region), middleware.Before)
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
			OperationName: "DescribeImages",
			Err:           err,
		}
	}
	out := result.(*DescribeImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImagesInput struct {

	// The repository that contains the images to describe.
	//
	// This member is required.
	RepositoryName *string

	// The filter key and value with which to filter your DescribeImages results.
	Filter *types.DescribeImagesFilter

	// The list of image IDs for the requested repository.
	ImageIds []*types.ImageIdentifier

	// The maximum number of repository results returned by DescribeImages in paginated
	// output. When this parameter is used, DescribeImages only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeImages
	// request with the returned nextToken value. This value can be between 1 and 1000.
	// If this parameter is not used, then DescribeImages returns up to 100 results and
	// a nextToken value, if applicable. This option cannot be used when you specify
	// images with imageIds.
	MaxResults *int32

	// The nextToken value returned from a previous paginated DescribeImages request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	// This option cannot be used when you specify images with imageIds.
	NextToken *string

	// The AWS account ID associated with the registry that contains the repository in
	// which to describe images. If you do not specify a registry, the default registry
	// is assumed.
	RegistryId *string
}

type DescribeImagesOutput struct {

	// A list of ImageDetail () objects that contain data about the image.
	ImageDetails []*types.ImageDetail

	// The nextToken value to include in a future DescribeImages request. When the
	// results of a DescribeImages request exceed maxResults, this value can be used to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeImagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImages{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImages{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeImages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "DescribeImages",
	}
}
