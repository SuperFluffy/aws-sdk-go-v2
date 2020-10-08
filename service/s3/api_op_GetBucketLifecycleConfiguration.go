// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Bucket lifecycle configuration now supports specifying a lifecycle rule using an
// object key name prefix, one or more object tags, or a combination of both.
// Accordingly, this section describes the latest API. The response describes the
// new filter element that you can use to specify a filter to select a subset of
// objects to which the rule applies. If you are still using previous version of
// the lifecycle configuration, it works. For the earlier API description, see
// GetBucketLifecycle ().  <p>Returns the lifecycle configuration information set
// on the bucket. For information about lifecycle configuration, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html">Object
// Lifecycle Management</a>.</p> <p>To use this operation, you must have permission
// to perform the <code>s3:GetLifecycleConfiguration</code> action. The bucket
// owner has this permission, by default. The bucket owner can grant this
// permission to others. For more information about permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources">Permissions
// Related to Bucket Subresource Operations</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html">Managing
// Access Permissions to Your Amazon S3 Resources</a>.</p> <p>
// <code>GetBucketLifecycleConfiguration</code> has the following special
// error:</p> <ul> <li> <p>Error code: <code>NoSuchLifecycleConfiguration</code>
// </p> <ul> <li> <p>Description: The lifecycle configuration does not exist.</p>
// </li> <li> <p>HTTP Status Code: 404 Not Found</p> </li> <li> <p>SOAP Fault Code
// Prefix: Client</p> </li> </ul> </li> </ul> <p>The following operations are
// related to <code>GetBucketLifecycleConfiguration</code>:</p> <ul> <li> <p>
// <a>GetBucketLifecycle</a> </p> </li> <li> <p> <a>PutBucketLifecycle</a> </p>
// </li> <li> <p> <a>DeleteBucketLifecycle</a> </p> </li> </ul>
func (c *Client) GetBucketLifecycleConfiguration(ctx context.Context, params *GetBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*GetBucketLifecycleConfigurationOutput, error) {
	stack := middleware.NewStack("GetBucketLifecycleConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetBucketLifecycleConfigurationMiddlewares(stack)
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
	addOpGetBucketLifecycleConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketLifecycleConfiguration(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)

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
			OperationName: "GetBucketLifecycleConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetBucketLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketLifecycleConfigurationInput struct {

	// The name of the bucket for which to get the lifecycle information.
	//
	// This member is required.
	Bucket *string
}

type GetBucketLifecycleConfigurationOutput struct {

	// Container for a lifecycle rule.
	Rules []*types.LifecycleRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetBucketLifecycleConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetBucketLifecycleConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketLifecycleConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBucketLifecycleConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketLifecycleConfiguration",
	}
}
