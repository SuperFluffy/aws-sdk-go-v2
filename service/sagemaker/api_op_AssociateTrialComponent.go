// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a trial component with a trial. A trial component can be associated
// with multiple trials. To disassociate a trial component from a trial, call the
// DisassociateTrialComponent () API.
func (c *Client) AssociateTrialComponent(ctx context.Context, params *AssociateTrialComponentInput, optFns ...func(*Options)) (*AssociateTrialComponentOutput, error) {
	stack := middleware.NewStack("AssociateTrialComponent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateTrialComponentMiddlewares(stack)
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
	addOpAssociateTrialComponentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateTrialComponent(options.Region), middleware.Before)
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
			OperationName: "AssociateTrialComponent",
			Err:           err,
		}
	}
	out := result.(*AssociateTrialComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateTrialComponentInput struct {

	// The name of the component to associated with the trial.
	//
	// This member is required.
	TrialComponentName *string

	// The name of the trial to associate with.
	//
	// This member is required.
	TrialName *string
}

type AssociateTrialComponentOutput struct {

	// The Amazon Resource Name (ARN) of the trial.
	TrialArn *string

	// The ARN of the trial component.
	TrialComponentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateTrialComponentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateTrialComponent{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateTrialComponent{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateTrialComponent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "AssociateTrialComponent",
	}
}
