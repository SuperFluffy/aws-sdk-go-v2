// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a system.
func (c *Client) GetSystemTemplate(ctx context.Context, params *GetSystemTemplateInput, optFns ...func(*Options)) (*GetSystemTemplateOutput, error) {
	stack := middleware.NewStack("GetSystemTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetSystemTemplateMiddlewares(stack)
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
	addOpGetSystemTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSystemTemplate(options.Region), middleware.Before)
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
			OperationName: "GetSystemTemplate",
			Err:           err,
		}
	}
	out := result.(*GetSystemTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSystemTemplateInput struct {

	// The ID of the system to get. This ID must be in the user's namespace. The ID
	// should be in the following format. urn:tdm:REGION/ACCOUNT
	// ID/default:system:SYSTEMNAME
	//
	// This member is required.
	Id *string

	// The number that specifies the revision of the system to get.
	RevisionNumber *int64
}

type GetSystemTemplateOutput struct {

	// An object that contains summary data about the system.
	Description *types.SystemTemplateDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetSystemTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetSystemTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSystemTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSystemTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "GetSystemTemplate",
	}
}
