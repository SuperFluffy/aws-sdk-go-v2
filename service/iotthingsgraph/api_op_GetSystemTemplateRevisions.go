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

// Gets revisions made to the specified system template. Only the previous 100
// revisions are stored. If the system has been deprecated, this action will return
// the revisions that occurred before its deprecation. This action won't work with
// systems that have been deleted.
func (c *Client) GetSystemTemplateRevisions(ctx context.Context, params *GetSystemTemplateRevisionsInput, optFns ...func(*Options)) (*GetSystemTemplateRevisionsOutput, error) {
	stack := middleware.NewStack("GetSystemTemplateRevisions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetSystemTemplateRevisionsMiddlewares(stack)
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
	addOpGetSystemTemplateRevisionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSystemTemplateRevisions(options.Region), middleware.Before)
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
			OperationName: "GetSystemTemplateRevisions",
			Err:           err,
		}
	}
	out := result.(*GetSystemTemplateRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSystemTemplateRevisionsInput struct {

	// The ID of the system template. The ID should be in the following format.
	// urn:tdm:REGION/ACCOUNT ID/default:system:SYSTEMNAME
	//
	// This member is required.
	Id *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string
}

type GetSystemTemplateRevisionsOutput struct {

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// An array of objects that contain summary data about the system template
	// revisions.
	Summaries []*types.SystemTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetSystemTemplateRevisionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetSystemTemplateRevisions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSystemTemplateRevisions{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSystemTemplateRevisions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "GetSystemTemplateRevisions",
	}
}
