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

// Get the configuration information about an origin access identity.
func (c *Client) GetCloudFrontOriginAccessIdentityConfig(ctx context.Context, params *GetCloudFrontOriginAccessIdentityConfigInput, optFns ...func(*Options)) (*GetCloudFrontOriginAccessIdentityConfigOutput, error) {
	stack := middleware.NewStack("GetCloudFrontOriginAccessIdentityConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetCloudFrontOriginAccessIdentityConfigMiddlewares(stack)
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
	addOpGetCloudFrontOriginAccessIdentityConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCloudFrontOriginAccessIdentityConfig(options.Region), middleware.Before)
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
			OperationName: "GetCloudFrontOriginAccessIdentityConfig",
			Err:           err,
		}
	}
	out := result.(*GetCloudFrontOriginAccessIdentityConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The origin access identity's configuration information. For more information,
// see CloudFrontOriginAccessIdentityConfig
// (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CloudFrontOriginAccessIdentityConfig.html).
type GetCloudFrontOriginAccessIdentityConfigInput struct {

	// The identity's ID.
	//
	// This member is required.
	Id *string
}

// The returned result of the corresponding request.
type GetCloudFrontOriginAccessIdentityConfigOutput struct {

	// The origin access identity's configuration information.
	CloudFrontOriginAccessIdentityConfig *types.CloudFrontOriginAccessIdentityConfig

	// The current version of the configuration. For example: E2QWRUHAPOMQZL.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetCloudFrontOriginAccessIdentityConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetCloudFrontOriginAccessIdentityConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetCloudFrontOriginAccessIdentityConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCloudFrontOriginAccessIdentityConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetCloudFrontOriginAccessIdentityConfig",
	}
}
