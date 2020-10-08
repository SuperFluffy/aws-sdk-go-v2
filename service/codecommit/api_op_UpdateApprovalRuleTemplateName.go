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

// Updates the name of a specified approval rule template.
func (c *Client) UpdateApprovalRuleTemplateName(ctx context.Context, params *UpdateApprovalRuleTemplateNameInput, optFns ...func(*Options)) (*UpdateApprovalRuleTemplateNameOutput, error) {
	stack := middleware.NewStack("UpdateApprovalRuleTemplateName", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateApprovalRuleTemplateNameMiddlewares(stack)
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
	addOpUpdateApprovalRuleTemplateNameValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApprovalRuleTemplateName(options.Region), middleware.Before)
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
			OperationName: "UpdateApprovalRuleTemplateName",
			Err:           err,
		}
	}
	out := result.(*UpdateApprovalRuleTemplateNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApprovalRuleTemplateNameInput struct {

	// The new name you want to apply to the approval rule template.
	//
	// This member is required.
	NewApprovalRuleTemplateName *string

	// The current name of the approval rule template.
	//
	// This member is required.
	OldApprovalRuleTemplateName *string
}

type UpdateApprovalRuleTemplateNameOutput struct {

	// The structure and content of the updated approval rule template.
	//
	// This member is required.
	ApprovalRuleTemplate *types.ApprovalRuleTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateApprovalRuleTemplateNameMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateApprovalRuleTemplateName{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateApprovalRuleTemplateName{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateApprovalRuleTemplateName(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "UpdateApprovalRuleTemplateName",
	}
}
