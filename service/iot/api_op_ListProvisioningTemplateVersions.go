// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// A list of fleet provisioning template versions.
func (c *Client) ListProvisioningTemplateVersions(ctx context.Context, params *ListProvisioningTemplateVersionsInput, optFns ...func(*Options)) (*ListProvisioningTemplateVersionsOutput, error) {
	stack := middleware.NewStack("ListProvisioningTemplateVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListProvisioningTemplateVersionsMiddlewares(stack)
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
	addOpListProvisioningTemplateVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProvisioningTemplateVersions(options.Region), middleware.Before)
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
			OperationName: "ListProvisioningTemplateVersions",
			Err:           err,
		}
	}
	out := result.(*ListProvisioningTemplateVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProvisioningTemplateVersionsInput struct {

	// The name of the fleet provisioning template.
	//
	// This member is required.
	TemplateName *string

	// The maximum number of results to return at one time.
	MaxResults *int32

	// A token to retrieve the next set of results.
	NextToken *string
}

type ListProvisioningTemplateVersionsOutput struct {

	// A token to retrieve the next set of results.
	NextToken *string

	// The list of fleet provisioning template versions.
	Versions []*types.ProvisioningTemplateVersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListProvisioningTemplateVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListProvisioningTemplateVersions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListProvisioningTemplateVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListProvisioningTemplateVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListProvisioningTemplateVersions",
	}
}
