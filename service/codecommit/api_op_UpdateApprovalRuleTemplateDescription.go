// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the description for a specified approval rule template.
func (c *Client) UpdateApprovalRuleTemplateDescription(ctx context.Context, params *UpdateApprovalRuleTemplateDescriptionInput, optFns ...func(*Options)) (*UpdateApprovalRuleTemplateDescriptionOutput, error) {
	stack := middleware.NewStack("UpdateApprovalRuleTemplateDescription", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateApprovalRuleTemplateDescriptionMiddlewares(stack)
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
	addOpUpdateApprovalRuleTemplateDescriptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApprovalRuleTemplateDescription(options.Region), middleware.Before)
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
			OperationName: "UpdateApprovalRuleTemplateDescription",
			Err:           err,
		}
	}
	out := result.(*UpdateApprovalRuleTemplateDescriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApprovalRuleTemplateDescriptionInput struct {

	// The updated description of the approval rule template.
	//
	// This member is required.
	ApprovalRuleTemplateDescription *string

	// The name of the template for which you want to update the description.
	//
	// This member is required.
	ApprovalRuleTemplateName *string
}

type UpdateApprovalRuleTemplateDescriptionOutput struct {

	// The structure and content of the updated approval rule template.
	//
	// This member is required.
	ApprovalRuleTemplate *types.ApprovalRuleTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateApprovalRuleTemplateDescriptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateApprovalRuleTemplateDescription{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateApprovalRuleTemplateDescription{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateApprovalRuleTemplateDescription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "UpdateApprovalRuleTemplateDescription",
	}
}
