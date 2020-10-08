// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get information about an OpsItem by using the ID. You must have permission in
// AWS Identity and Access Management (IAM) to view information about an OpsItem.
// For more information, see Getting started with OpsCenter
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-getting-started.html)
// in the AWS Systems Manager User Guide. Operations engineers and IT professionals
// use OpsCenter to view, investigate, and remediate operational issues impacting
// the performance and health of their AWS resources. For more information, see AWS
// Systems Manager OpsCenter
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html) in
// the AWS Systems Manager User Guide.
func (c *Client) GetOpsItem(ctx context.Context, params *GetOpsItemInput, optFns ...func(*Options)) (*GetOpsItemOutput, error) {
	stack := middleware.NewStack("GetOpsItem", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetOpsItemMiddlewares(stack)
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
	addOpGetOpsItemValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOpsItem(options.Region), middleware.Before)
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
			OperationName: "GetOpsItem",
			Err:           err,
		}
	}
	out := result.(*GetOpsItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOpsItemInput struct {

	// The ID of the OpsItem that you want to get.
	//
	// This member is required.
	OpsItemId *string
}

type GetOpsItemOutput struct {

	// The OpsItem.
	OpsItem *types.OpsItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetOpsItemMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetOpsItem{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOpsItem{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetOpsItem(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetOpsItem",
	}
}
