// Code generated by smithy-go-codegen DO NOT EDIT.

package mobile

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mobile/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates customized software development kit (SDK) and or tool packages used to
// integrate mobile web or mobile app clients with backend AWS resources.
func (c *Client) ExportBundle(ctx context.Context, params *ExportBundleInput, optFns ...func(*Options)) (*ExportBundleOutput, error) {
	stack := middleware.NewStack("ExportBundle", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpExportBundleMiddlewares(stack)
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
	addOpExportBundleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExportBundle(options.Region), middleware.Before)
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
			OperationName: "ExportBundle",
			Err:           err,
		}
	}
	out := result.(*ExportBundleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request structure used to request generation of custom SDK and tool packages
// required to integrate mobile web or app clients with backed AWS resources.
type ExportBundleInput struct {

	// Unique bundle identifier.
	//
	// This member is required.
	BundleId *string

	// Developer desktop or target application platform.
	Platform types.Platform

	// Unique project identifier.
	ProjectId *string
}

// Result structure which contains link to download custom-generated SDK and tool
// packages used to integrate mobile web or app clients with backed AWS resources.
type ExportBundleOutput struct {

	// URL which contains the custom-generated SDK and tool packages used to integrate
	// the client mobile app or web app with the AWS resources created by the AWS
	// Mobile Hub project.
	DownloadUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpExportBundleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpExportBundle{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpExportBundle{}, middleware.After)
}

func newServiceMetadataMiddleware_opExportBundle(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "awsmobilehubservice",
		OperationName: "ExportBundle",
	}
}
