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

// Deletes a list of specified images within a repository. Images are specified
// with either an imageTag or imageDigest. You can remove a tag from an image by
// specifying the image's tag in your request. When you remove the last tag from an
// image, the image is deleted from your repository. You can completely delete an
// image (and all of its tags) by specifying the image's digest in your request.
func (c *Client) BatchDeleteImage(ctx context.Context, params *BatchDeleteImageInput, optFns ...func(*Options)) (*BatchDeleteImageOutput, error) {
	stack := middleware.NewStack("BatchDeleteImage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchDeleteImageMiddlewares(stack)
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
	addOpBatchDeleteImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteImage(options.Region), middleware.Before)
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
			OperationName: "BatchDeleteImage",
			Err:           err,
		}
	}
	out := result.(*BatchDeleteImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Deletes specified images within a specified repository. Images are specified
// with either the imageTag or imageDigest.
type BatchDeleteImageInput struct {

	// A list of image ID references that correspond to images to delete. The format of
	// the imageIds reference is imageTag=tag or imageDigest=digest.
	//
	// This member is required.
	ImageIds []*types.ImageIdentifier

	// The repository that contains the image to delete.
	//
	// This member is required.
	RepositoryName *string

	// The AWS account ID associated with the registry that contains the image to
	// delete. If you do not specify a registry, the default registry is assumed.
	RegistryId *string
}

type BatchDeleteImageOutput struct {

	// Any failures associated with the call.
	Failures []*types.ImageFailure

	// The image IDs of the deleted images.
	ImageIds []*types.ImageIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchDeleteImageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeleteImage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeleteImage{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchDeleteImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "BatchDeleteImage",
	}
}
