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

// This operation enables you to delete multiple objects from a bucket using a
// single HTTP request. If you know the object keys that you want to delete, then
// this operation provides a suitable alternative to sending individual delete
// requests, reducing per-request overhead.  <p>The request contains a list of up
// to 1000 keys that you want to delete. In the XML, you provide the object key
// names, and optionally, version IDs if you want to delete a specific version of
// the object from a versioning-enabled bucket. For each key, Amazon S3 performs a
// delete operation and returns the result of that delete, success, or failure, in
// the response. Note that if the object specified in the request is not found,
// Amazon S3 returns the result as deleted.</p> <p> The operation supports two
// modes for the response: verbose and quiet. By default, the operation uses
// verbose mode in which the response includes the result of deletion of each key
// in your request. In quiet mode the response includes only keys where the delete
// operation encountered an error. For a successful deletion, the operation does
// not return any information about the delete in the response body.</p> <p>When
// performing this operation on an MFA Delete enabled bucket, that attempts to
// delete any versioned objects, you must include an MFA token. If you do not
// provide one, the entire request will fail, even if there are non-versioned
// objects you are trying to delete. If you provide an invalid token, whether there
// are versioned keys in the request or not, the entire Multi-Object Delete request
// will fail. For information about MFA Delete, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html#MultiFactorAuthenticationDelete">
// MFA Delete</a>.</p> <p>Finally, the Content-MD5 header is required for all
// Multi-Object Delete requests. Amazon S3 uses the header value to ensure that
// your request body has not been altered in transit.</p> <p>The following
// operations are related to <code>DeleteObjects</code>:</p> <ul> <li> <p>
// <a>CreateMultipartUpload</a> </p> </li> <li> <p> <a>UploadPart</a> </p> </li>
// <li> <p> <a>CompleteMultipartUpload</a> </p> </li> <li> <p> <a>ListParts</a>
// </p> </li> <li> <p> <a>AbortMultipartUpload</a> </p> </li> </ul>
func (c *Client) DeleteObjects(ctx context.Context, params *DeleteObjectsInput, optFns ...func(*Options)) (*DeleteObjectsOutput, error) {
	stack := middleware.NewStack("DeleteObjects", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteObjectsMiddlewares(stack)
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
	addOpDeleteObjectsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteObjects(options.Region), middleware.Before)
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
			OperationName: "DeleteObjects",
			Err:           err,
		}
	}
	out := result.(*DeleteObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteObjectsInput struct {

	// The bucket name containing the objects to delete. When using this API with an
	// access point, you must direct requests to the access point hostname. The access
	// point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// Container for the request.
	//
	// This member is required.
	Delete *types.Delete

	// Specifies whether you want to delete this object even if it has a
	// Governance-type Object Lock in place. You must have sufficient permissions to
	// perform this operation.
	BypassGovernanceRetention *bool

	// The concatenation of the authentication device's serial number, a space, and the
	// value that is displayed on your authentication device. Required to permanently
	// delete a versioned object if versioning is configured with MFA delete enabled.
	MFA *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer
}

type DeleteObjectsOutput struct {

	// Container element for a successful delete. It identifies the object that was
	// successfully deleted.
	Deleted []*types.DeletedObject

	// Container for a failed delete operation that describes the object that Amazon S3
	// attempted to delete and the error it encountered.
	Errors []*types.Error

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteObjectsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteObjects{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteObjects{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteObjects(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteObjects",
	}
}
