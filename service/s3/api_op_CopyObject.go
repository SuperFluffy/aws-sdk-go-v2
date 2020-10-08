// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a copy of an object that is already stored in Amazon S3. You can store
// individual objects of up to 5 TB in Amazon S3. You create a copy of your object
// up to 5 GB in size in a single atomic operation using this API. However, to copy
// an object greater than 5 GB, you must use the multipart upload Upload Part -
// Copy API. For more information, see Copy Object Using the REST Multipart Upload
// API
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/CopyingObjctsUsingRESTMPUapi.html).
// All copy requests must be authenticated. Additionally, you must have read access
// to the source object and write access to the destination bucket. For more
// information, see REST Authentication
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html). Both
// the Region that you want to copy the object from and the Region that you want to
// copy the object to must be enabled for your account. A copy request might return
// an error when Amazon S3 receives the copy request or while Amazon S3 is copying
// the files. If the error occurs before the copy operation starts, you receive a
// standard Amazon S3 error. If the error occurs during the copy operation, the
// error response is embedded in the <code>200 OK</code> response. This means that
// a <code>200 OK</code> response can contain either a success or an error. Design
// your application to parse the contents of the response and handle it
// appropriately.</p> <p>If the copy is successful, you receive a response with
// information about the copied object.</p> <note> <p>If the request is an HTTP 1.1
// request, the response is chunk encoded. If it were not, it would not contain the
// content-length, and you would need to read the entire body.</p> </note> <p>The
// copy request charge is based on the storage class and Region that you specify
// for the destination object. For pricing information, see <a
// href="https://aws.amazon.com/s3/pricing/">Amazon S3 pricing</a>.</p> <important>
// <p>Amazon S3 transfer acceleration does not support cross-Region copies. If you
// request a cross-Region copy using a transfer acceleration endpoint, you get a
// 400 <code>Bad Request</code> error. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html">Transfer
// Acceleration</a>.</p> </important> <p> <b>Metadata</b> </p> <p>When copying an
// object, you can preserve all metadata (default) or specify new metadata.
// However, the ACL is not preserved and is set to private for the user making the
// request. To override the default ACL setting, specify a new ACL when generating
// a copy request. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html">Using
// ACLs</a>. </p> <p>To specify whether you want the object metadata copied from
// the source object or replaced with metadata provided in the request, you can
// optionally add the <code>x-amz-metadata-directive</code> header. When you grant
// permissions, you can use the <code>s3:x-amz-metadata-directive</code> condition
// key to enforce certain metadata behavior when objects are uploaded. For more
// information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/amazon-s3-policy-keys.html">Specifying
// Conditions in a Policy</a> in the <i>Amazon S3 Developer Guide</i>. For a
// complete list of Amazon S3-specific condition keys, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/list_amazons3.html">Actions,
// Resources, and Condition Keys for Amazon S3</a>.</p> <p> <b>
// <code>x-amz-copy-source-if</code> Headers</b> </p> <p>To only copy an object
// under certain conditions, such as whether the <code>Etag</code> matches or
// whether the object was modified before or after a specified date, use the
// following request parameters:</p> <ul> <li> <p>
// <code>x-amz-copy-source-if-match</code> </p> </li> <li> <p>
// <code>x-amz-copy-source-if-none-match</code> </p> </li> <li> <p>
// <code>x-amz-copy-source-if-unmodified-since</code> </p> </li> <li> <p>
// <code>x-amz-copy-source-if-modified-since</code> </p> </li> </ul> <p> If both
// the <code>x-amz-copy-source-if-match</code> and
// <code>x-amz-copy-source-if-unmodified-since</code> headers are present in the
// request and evaluate as follows, Amazon S3 returns <code>200 OK</code> and
// copies the data:</p> <ul> <li> <p> <code>x-amz-copy-source-if-match</code>
// condition evaluates to true</p> </li> <li> <p>
// <code>x-amz-copy-source-if-unmodified-since</code> condition evaluates to
// false</p> </li> </ul> <p>If both the
// <code>x-amz-copy-source-if-none-match</code> and
// <code>x-amz-copy-source-if-modified-since</code> headers are present in the
// request and evaluate as follows, Amazon S3 returns the <code>412 Precondition
// Failed</code> response code:</p> <ul> <li> <p>
// <code>x-amz-copy-source-if-none-match</code> condition evaluates to false</p>
// </li> <li> <p> <code>x-amz-copy-source-if-modified-since</code> condition
// evaluates to true</p> </li> </ul> <note> <p>All headers with the
// <code>x-amz-</code> prefix, including <code>x-amz-copy-source</code>, must be
// signed.</p> </note> <p> <b>Encryption</b> </p> <p>The source object that you are
// copying can be encrypted or unencrypted. The source object can be encrypted with
// server-side encryption using AWS managed encryption keys (SSE-S3 or SSE-KMS) or
// by using a customer-provided encryption key. With server-side encryption, Amazon
// S3 encrypts your data as it writes it to disks in its data centers and decrypts
// the data when you access it. </p> <p>You can optionally use the appropriate
// encryption-related headers to request server-side encryption for the target
// object. You have the option to provide your own encryption key or use SSE-S3 or
// SSE-KMS, regardless of the form of server-side encryption that was used to
// encrypt the source object. You can even request encryption if the source object
// was not encrypted. For more information about server-side encryption, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/serv-side-encryption.html">Using
// Server-Side Encryption</a>.</p> <p> <b>Access Control List (ACL)-Specific
// Request Headers</b> </p> <p>When copying an object, you can optionally use
// headers to grant ACL-based permissions. By default, all objects are private.
// Only the owner has full access control. When adding a new object, you can grant
// permissions to individual AWS accounts or to predefined groups defined by Amazon
// S3. These permissions are then added to the ACL on the object. For more
// information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html">Access
// Control List (ACL) Overview</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-using-rest-api.html">Managing
// ACLs Using the REST API</a>. </p> <p> <b>Storage Class Options</b> </p> <p>You
// can use the <code>CopyObject</code> operation to change the storage class of an
// object that is already stored in Amazon S3 using the <code>StorageClass</code>
// parameter. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html">Storage
// Classes</a> in the <i>Amazon S3 Service Developer Guide</i>.</p> <p>
// <b>Versioning</b> </p> <p>By default, <code>x-amz-copy-source</code> identifies
// the current version of an object to copy. If the current version is a delete
// marker, Amazon S3 behaves as if the object was deleted. To copy a different
// version, use the <code>versionId</code> subresource.</p> <p>If you enable
// versioning on the target bucket, Amazon S3 generates a unique version ID for the
// object being copied. This version ID is different from the version ID of the
// source object. Amazon S3 returns the version ID of the copied object in the
// <code>x-amz-version-id</code> response header in the response.</p> <p>If you do
// not enable versioning or suspend it on the target bucket, the version ID that
// Amazon S3 generates is always null.</p> <p>If the source object's storage class
// is GLACIER, you must restore a copy of this object before you can use it as a
// source object for the copy operation. For more information, see .</p> <p>The
// following operations are related to <code>CopyObject</code>:</p> <ul> <li> <p>
// <a>PutObject</a> </p> </li> <li> <p> <a>GetObject</a> </p> </li> </ul> <p>For
// more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/CopyingObjectsExamples.html">Copying
// Objects</a>.</p>
func (c *Client) CopyObject(ctx context.Context, params *CopyObjectInput, optFns ...func(*Options)) (*CopyObjectOutput, error) {
	stack := middleware.NewStack("CopyObject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpCopyObjectMiddlewares(stack)
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
	addOpCopyObjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyObject(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	s3cust.HandleResponseErrorWith200Status(stack)

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
			OperationName: "CopyObject",
			Err:           err,
		}
	}
	out := result.(*CopyObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyObjectInput struct {

	// The name of the destination bucket.
	//
	// This member is required.
	Bucket *string

	// The name of the source bucket and key name of the source object, separated by a
	// slash (/). Must be URL-encoded.
	//
	// This member is required.
	CopySource *string

	// The key of the destination object.
	//
	// This member is required.
	Key *string

	// The canned ACL to apply to the object.
	ACL types.ObjectCannedACL

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string

	// Specifies presentational information for the object.
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus what
	// decoding mechanisms must be applied to obtain the media-type referenced by the
	// Content-Type header field.
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// A standard MIME type describing the format of the object data.
	ContentType *string

	// Copies the object if its entity tag (ETag) matches the specified tag.
	CopySourceIfMatch *string

	// Copies the object if it has been modified since the specified time.
	CopySourceIfModifiedSince *time.Time

	// Copies the object if its entity tag (ETag) is different than the specified ETag.
	CopySourceIfNoneMatch *string

	// Copies the object if it hasn't been modified since the specified time.
	CopySourceIfUnmodifiedSince *time.Time

	// Specifies the algorithm to use when decrypting the source object (for example,
	// AES256).
	CopySourceSSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use to decrypt
	// the source object. The encryption key provided in this header must be one that
	// was used when the source object was created.
	CopySourceSSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	CopySourceSSECustomerKeyMD5 *string

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	GrantFullControl *string

	// Allows grantee to read the object data and its metadata.
	GrantRead *string

	// Allows grantee to read the object ACL.
	GrantReadACP *string

	// Allows grantee to write the ACL for the applicable object.
	GrantWriteACP *string

	// A map of metadata to store with the object in S3.
	Metadata map[string]*string

	// Specifies whether the metadata is copied from the source object or replaced with
	// metadata provided in the request.
	MetadataDirective types.MetadataDirective

	// Specifies whether you want to apply a Legal Hold to the copied object.
	ObjectLockLegalHoldStatus types.ObjectLockLegalHoldStatus

	// The Object Lock mode that you want to apply to the copied object.
	ObjectLockMode types.ObjectLockMode

	// The date and time when you want the copied object's Object Lock to expire.
	ObjectLockRetainUntilDate *time.Time

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string

	// Specifies the AWS KMS Encryption Context to use for object encryption. The value
	// of this header is a base64-encoded UTF-8 string holding JSON with the encryption
	// context key-value pairs.
	SSEKMSEncryptionContext *string

	// Specifies the AWS KMS key ID to use for object encryption. All GET and PUT
	// requests for an object protected by AWS KMS will fail if not made via SSL or
	// using SigV4. For information about configuring using any of the officially
	// supported AWS SDKs and AWS CLI, see Specifying the Signature Version in Request
	// Authentication
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version)
	// in the Amazon S3 Developer Guide.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// The type of storage to use for the object. Defaults to 'STANDARD'.
	StorageClass types.StorageClass

	// The tag-set for the object destination object this value must be used in
	// conjunction with the TaggingDirective. The tag-set must be encoded as URL Query
	// parameters.
	Tagging *string

	// Specifies whether the object tag-set are copied from the source object or
	// replaced with tag-set provided in the request.
	TaggingDirective types.TaggingDirective

	// If the bucket is configured as a website, redirects requests for this object to
	// another object in the same bucket or to an external URL. Amazon S3 stores the
	// value of this header in the object metadata.
	WebsiteRedirectLocation *string
}

type CopyObjectOutput struct {

	// Container for all response elements.
	CopyObjectResult *types.CopyObjectResult

	// Version of the copied object in the destination bucket.
	CopySourceVersionId *string

	// If the object expiration is configured, the response includes this header.
	Expiration *string

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm used.
	SSECustomerAlgorithm *string

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string

	// If present, specifies the AWS KMS Encryption Context to use for object
	// encryption. The value of this header is a base64-encoded UTF-8 string holding
	// JSON with the encryption context key-value pairs.
	SSEKMSEncryptionContext *string

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// Version ID of the newly created copy.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpCopyObjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpCopyObject{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpCopyObject{}, middleware.After)
}

func newServiceMetadataMiddleware_opCopyObject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CopyObject",
	}
}
