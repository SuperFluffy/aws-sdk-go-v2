// Code generated by smithy-go-codegen DO NOT EDIT.

package mobile

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Exports project configuration to a snapshot which can be downloaded and shared.
// Note that mobile app push credentials are encrypted in exported projects, so
// they can only be shared successfully within the same AWS account.
func (c *Client) ExportProject(ctx context.Context, params *ExportProjectInput, optFns ...func(*Options)) (*ExportProjectOutput, error) {
	stack := middleware.NewStack("ExportProject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpExportProjectMiddlewares(stack)
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
	addOpExportProjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExportProject(options.Region), middleware.Before)
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
			OperationName: "ExportProject",
			Err:           err,
		}
	}
	out := result.(*ExportProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request structure used in requests to export project configuration details.
type ExportProjectInput struct {

	// Unique project identifier.
	//
	// This member is required.
	ProjectId *string
}

// Result structure used for requests to export project configuration details.
type ExportProjectOutput struct {

	// URL which can be used to download the exported project configuation file(s).
	DownloadUrl *string

	// URL which can be shared to allow other AWS users to create their own project in
	// AWS Mobile Hub with the same configuration as the specified project. This URL
	// pertains to a snapshot in time of the project configuration that is created when
	// this API is called. If you want to share additional changes to your project
	// configuration, then you will need to create and share a new snapshot by calling
	// this method again.
	ShareUrl *string

	// Unique identifier for the exported snapshot of the project configuration. This
	// snapshot identifier is included in the share URL.
	SnapshotId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpExportProjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpExportProject{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpExportProject{}, middleware.After)
}

func newServiceMetadataMiddleware_opExportProject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "awsmobilehubservice",
		OperationName: "ExportProject",
	}
}
