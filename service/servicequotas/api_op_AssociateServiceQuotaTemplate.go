// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates the Service Quotas template with your organization so that when new
// accounts are created in your organization, the template submits increase
// requests for the specified service quotas. Use the Service Quotas template to
// request an increase for any adjustable quota value. After you define the Service
// Quotas template, use this operation to associate, or enable, the template.
func (c *Client) AssociateServiceQuotaTemplate(ctx context.Context, params *AssociateServiceQuotaTemplateInput, optFns ...func(*Options)) (*AssociateServiceQuotaTemplateOutput, error) {
	stack := middleware.NewStack("AssociateServiceQuotaTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateServiceQuotaTemplateMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateServiceQuotaTemplate(options.Region), middleware.Before)
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
			OperationName: "AssociateServiceQuotaTemplate",
			Err:           err,
		}
	}
	out := result.(*AssociateServiceQuotaTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateServiceQuotaTemplateInput struct {
}

type AssociateServiceQuotaTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateServiceQuotaTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateServiceQuotaTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateServiceQuotaTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateServiceQuotaTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "AssociateServiceQuotaTemplate",
	}
}
