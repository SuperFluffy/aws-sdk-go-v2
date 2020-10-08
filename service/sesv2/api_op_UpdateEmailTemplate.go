// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an email template. Email templates enable you to send personalized email
// to one or more destinations in a single API operation. For more information, see
// the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html).
// You can execute this operation no more than once per second.
func (c *Client) UpdateEmailTemplate(ctx context.Context, params *UpdateEmailTemplateInput, optFns ...func(*Options)) (*UpdateEmailTemplateOutput, error) {
	stack := middleware.NewStack("UpdateEmailTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateEmailTemplateMiddlewares(stack)
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
	addOpUpdateEmailTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEmailTemplate(options.Region), middleware.Before)
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
			OperationName: "UpdateEmailTemplate",
			Err:           err,
		}
	}
	out := result.(*UpdateEmailTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to update an email template. For more information, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html).
type UpdateEmailTemplateInput struct {

	// The content of the email template, composed of a subject line, an HTML part, and
	// a text-only part.
	//
	// This member is required.
	TemplateContent *types.EmailTemplateContent

	// The name of the template you want to update.
	//
	// This member is required.
	TemplateName *string
}

// If the action is successful, the service sends back an HTTP 200 response with an
// empty HTTP body.
type UpdateEmailTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateEmailTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEmailTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEmailTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateEmailTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "UpdateEmailTemplate",
	}
}
