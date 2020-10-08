// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This API operation deletes an Amazon S3 on Outposts bucket. To delete an S3
// bucket, see DeleteBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html) in the
// Amazon Simple Storage Service API. Deletes the Amazon S3 on Outposts bucket. All
// objects (including all object versions and delete markers) in the bucket must be
// deleted before the bucket itself can be deleted. For more information, see Using
// Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in Amazon
// Simple Storage Service Developer Guide. All Amazon S3 on Outposts REST API
// requests for this action require an additional parameter of outpost-id to be
// passed with the request and an S3 on Outposts endpoint hostname prefix instead
// of s3-control. For an example of the request syntax for Amazon S3 on Outposts
// that uses the S3 on Outposts endpoint hostname prefix and the outpost-id derived
// using the access point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_DeleteBucket.html#API_control_DeleteBucket_Examples)
// section below. Related Resources
//
//     * CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html)
//
//
// * GetBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_GetBucket.html)
//
//
// * DeleteObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html)
func (c *Client) DeleteBucket(ctx context.Context, params *DeleteBucketInput, optFns ...func(*Options)) (*DeleteBucketOutput, error) {
	stack := middleware.NewStack("DeleteBucket", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteBucketMiddlewares(stack)
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
	addOpDeleteBucketValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucket(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)

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
			OperationName: "DeleteBucket",
			Err:           err,
		}
	}
	out := result.(*DeleteBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketInput struct {

	// The account ID that owns the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// Specifies the bucket being deleted. For Amazon S3 on Outposts specify the ARN of
	// the bucket accessed in the format arn:aws:s3-outposts:::outpost//bucket/. For
	// example, to access the bucket reports through outpost my-outpost owned by
	// account 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string
}

type DeleteBucketOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteBucketMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucket{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucket{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBucket(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucket",
	}
}
