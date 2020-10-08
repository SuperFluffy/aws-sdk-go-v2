// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve a JSON array of up to twenty of your job templates. This will return
// the templates themselves, not just a list of them. To retrieve the next twenty
// templates, use the nextToken string returned with the array
func (c *Client) ListJobTemplates(ctx context.Context, params *ListJobTemplatesInput, optFns ...func(*Options)) (*ListJobTemplatesOutput, error) {
	stack := middleware.NewStack("ListJobTemplates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListJobTemplatesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListJobTemplates(options.Region), middleware.Before)
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
			OperationName: "ListJobTemplates",
			Err:           err,
		}
	}
	out := result.(*ListJobTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobTemplatesInput struct {

	// Optionally, specify a job template category to limit responses to only job
	// templates from that category.
	Category *string

	// Optional. When you request a list of job templates, you can choose to list them
	// alphabetically by NAME or chronologically by CREATION_DATE. If you don't
	// specify, the service will list them by name.
	ListBy types.JobTemplateListBy

	// Optional. Number of job templates, up to twenty, that will be returned at one
	// time.
	MaxResults *int32

	// Use this string, provided with the response to a previous request, to request
	// the next batch of job templates.
	NextToken *string

	// Optional. When you request lists of resources, you can specify whether they are
	// sorted in ASCENDING or DESCENDING order. Default varies by resource.
	Order types.Order
}

type ListJobTemplatesOutput struct {

	// List of Job templates.
	JobTemplates []*types.JobTemplate

	// Use this string to request the next batch of job templates.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListJobTemplatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListJobTemplates{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobTemplates{}, middleware.After)
}

func newServiceMetadataMiddleware_opListJobTemplates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "ListJobTemplates",
	}
}
