// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new origin access identity. If you're using Amazon S3 for your origin,
// you can use an origin access identity to require users to access your content
// using a CloudFront URL instead of the Amazon S3 URL. For more information about
// how to use origin access identities, see Serving Private Content through
// CloudFront
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html)
// in the Amazon CloudFront Developer Guide.
func (c *Client) CreateCloudFrontOriginAccessIdentity(ctx context.Context, params *CreateCloudFrontOriginAccessIdentityInput, optFns ...func(*Options)) (*CreateCloudFrontOriginAccessIdentityOutput, error) {
	stack := middleware.NewStack("CreateCloudFrontOriginAccessIdentity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpCreateCloudFrontOriginAccessIdentityMiddlewares(stack)
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
	addOpCreateCloudFrontOriginAccessIdentityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCloudFrontOriginAccessIdentity(options.Region), middleware.Before)
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
			OperationName: "CreateCloudFrontOriginAccessIdentity",
			Err:           err,
		}
	}
	out := result.(*CreateCloudFrontOriginAccessIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to create a new origin access identity (OAI). An origin access
// identity is a special CloudFront user that you can associate with Amazon S3
// origins, so that you can secure all or just some of your Amazon S3 content. For
// more information, see  Restricting Access to Amazon S3 Content by Using an
// Origin Access Identity
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html)
// in the Amazon CloudFront Developer Guide.
type CreateCloudFrontOriginAccessIdentityInput struct {

	// The current configuration information for the identity.
	//
	// This member is required.
	CloudFrontOriginAccessIdentityConfig *types.CloudFrontOriginAccessIdentityConfig
}

// The returned result of the corresponding request.
type CreateCloudFrontOriginAccessIdentityOutput struct {

	// The origin access identity's information.
	CloudFrontOriginAccessIdentity *types.CloudFrontOriginAccessIdentity

	// The current version of the origin access identity created.
	ETag *string

	// The fully qualified URI of the new origin access identity just created.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpCreateCloudFrontOriginAccessIdentityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpCreateCloudFrontOriginAccessIdentity{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpCreateCloudFrontOriginAccessIdentity{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCloudFrontOriginAccessIdentity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateCloudFrontOriginAccessIdentity",
	}
}
