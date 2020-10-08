// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS CloudFormation template.
func (c *Client) CreateCloudFormationTemplate(ctx context.Context, params *CreateCloudFormationTemplateInput, optFns ...func(*Options)) (*CreateCloudFormationTemplateOutput, error) {
	stack := middleware.NewStack("CreateCloudFormationTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateCloudFormationTemplateMiddlewares(stack)
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
	addOpCreateCloudFormationTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCloudFormationTemplate(options.Region), middleware.Before)
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
			OperationName: "CreateCloudFormationTemplate",
			Err:           err,
		}
	}
	out := result.(*CreateCloudFormationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCloudFormationTemplateInput struct {

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationId *string

	// The semantic version of the application: https://semver.org/
	// (https://semver.org/)
	SemanticVersion *string
}

type CreateCloudFormationTemplateOutput struct {

	// The application Amazon Resource Name (ARN).
	ApplicationId *string

	// The date and time this resource was created.
	CreationTime *string

	// The date and time this template expires. Templates expire 1 hour after creation.
	ExpirationTime *string

	// The semantic version of the application: https://semver.org/
	// (https://semver.org/)
	SemanticVersion *string

	// Status of the template creation workflow.Possible values: PREPARING | ACTIVE |
	// EXPIRED
	Status types.Status

	// The UUID returned by CreateCloudFormationTemplate.Pattern:
	// [0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}
	TemplateId *string

	// A link to the template that can be used to deploy the application using AWS
	// CloudFormation.
	TemplateUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateCloudFormationTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateCloudFormationTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCloudFormationTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCloudFormationTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "CreateCloudFormationTemplate",
	}
}
