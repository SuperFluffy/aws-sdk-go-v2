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
	"time"
)

// Lists the parts that have been uploaded for a specific multipart upload. This
// operation must include the upload ID, which you obtain by sending the initiate
// multipart upload request (see CreateMultipartUpload ()). This request returns a
// maximum of 1,000 uploaded parts. The default number of parts returned is 1,000
// parts. You can restrict the number of parts returned by specifying the max-parts
// request parameter. If your multipart upload consists of more than 1,000 parts,
// the response returns an IsTruncated field with the value of true, and a
// NextPartNumberMarker element. In subsequent ListParts requests you can include
// the part-number-marker query string parameter and set its value to the
// NextPartNumberMarker field value from the previous response.  <p>For more
// information on multipart uploads, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html">Uploading
// Objects Using Multipart Upload</a>.</p> <p>For information on permissions
// required to use the multipart upload API, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html">Multipart
// Upload API and Permissions</a>.</p> <p>The following operations are related to
// <code>ListParts</code>:</p> <ul> <li> <p> <a>CreateMultipartUpload</a> </p>
// </li> <li> <p> <a>UploadPart</a> </p> </li> <li> <p>
// <a>CompleteMultipartUpload</a> </p> </li> <li> <p> <a>AbortMultipartUpload</a>
// </p> </li> <li> <p> <a>ListMultipartUploads</a> </p> </li> </ul>
func (c *Client) ListParts(ctx context.Context, params *ListPartsInput, optFns ...func(*Options)) (*ListPartsOutput, error) {
	stack := middleware.NewStack("ListParts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListPartsMiddlewares(stack)
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
	addOpListPartsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListParts(options.Region), middleware.Before)
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
			OperationName: "ListParts",
			Err:           err,
		}
	}
	out := result.(*ListPartsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPartsInput struct {

	// Name of the bucket to which the parts are being uploaded. When using this API
	// with an access point, you must direct requests to the access point hostname. The
	// access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// Object key for which the multipart upload was initiated.
	//
	// This member is required.
	Key *string

	// Upload ID identifying the multipart upload whose parts are being listed.
	//
	// This member is required.
	UploadId *string

	// Sets the maximum number of parts to return.
	MaxParts *int32

	// Specifies the part after which listing should begin. Only parts with higher part
	// numbers will be listed.
	PartNumberMarker *int32

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer
}

type ListPartsOutput struct {

	// If the bucket has a lifecycle rule configured with an action to abort incomplete
	// multipart uploads and the prefix in the lifecycle rule matches the object name
	// in the request, then the response includes this header indicating when the
	// initiated multipart upload will become eligible for abort operation. For more
	// information, see Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle
	// Policy
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config).
	// <p>The response will also include the <code>x-amz-abort-rule-id</code> header
	// that will provide the ID of the lifecycle configuration rule that defines this
	// action.</p>
	AbortDate *time.Time

	// This header is returned along with the x-amz-abort-date header. It identifies
	// applicable lifecycle configuration rule that defines the action to abort
	// incomplete multipart uploads.
	AbortRuleId *string

	// Name of the bucket to which the multipart upload was initiated.
	Bucket *string

	// Container element that identifies who initiated the multipart upload. If the
	// initiator is an AWS account, this element provides the same information as the
	// Owner element. If the initiator is an IAM User, this element provides the user
	// ARN and display name.
	Initiator *types.Initiator

	// Indicates whether the returned list of parts is truncated. A true value
	// indicates that the list was truncated. A list can be truncated if the number of
	// parts exceeds the limit returned in the MaxParts element.
	IsTruncated *bool

	// Object key for which the multipart upload was initiated.
	Key *string

	// Maximum number of parts that were allowed in the response.
	MaxParts *int32

	// When a list is truncated, this element specifies the last part in the list, as
	// well as the value to use for the part-number-marker request parameter in a
	// subsequent request.
	NextPartNumberMarker *int32

	// Container element that identifies the object owner, after the object is created.
	// If multipart upload is initiated by an IAM user, this element provides the
	// parent account ID and display name.
	Owner *types.Owner

	// When a list is truncated, this element specifies the last part in the list, as
	// well as the value to use for the part-number-marker request parameter in a
	// subsequent request.
	PartNumberMarker *int32

	// Container for elements related to a particular part. A response can contain zero
	// or more Part elements.
	Parts []*types.Part

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// Class of storage (STANDARD or REDUCED_REDUNDANCY) used to store the uploaded
	// object.
	StorageClass types.StorageClass

	// Upload ID identifying the multipart upload whose parts are being listed.
	UploadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListPartsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListParts{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListParts{}, middleware.After)
}

func newServiceMetadataMiddleware_opListParts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListParts",
	}
}
