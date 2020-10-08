// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a target with a maintenance window.
func (c *Client) RegisterTargetWithMaintenanceWindow(ctx context.Context, params *RegisterTargetWithMaintenanceWindowInput, optFns ...func(*Options)) (*RegisterTargetWithMaintenanceWindowOutput, error) {
	stack := middleware.NewStack("RegisterTargetWithMaintenanceWindow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRegisterTargetWithMaintenanceWindowMiddlewares(stack)
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
	addIdempotencyToken_opRegisterTargetWithMaintenanceWindowMiddleware(stack, options)
	addOpRegisterTargetWithMaintenanceWindowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterTargetWithMaintenanceWindow(options.Region), middleware.Before)
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
			OperationName: "RegisterTargetWithMaintenanceWindow",
			Err:           err,
		}
	}
	out := result.(*RegisterTargetWithMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTargetWithMaintenanceWindowInput struct {

	// The type of target being registered with the maintenance window.
	//
	// This member is required.
	ResourceType types.MaintenanceWindowResourceType

	// The targets to register with the maintenance window. In other words, the
	// instances to run commands on when the maintenance window runs.  <p>You can
	// specify targets using instance IDs, resource group names, or tags that have been
	// applied to instances. Example 1: Specify instance IDs
	// Key=InstanceIds,Values=instance-id-1,instance-id-2,instance-id-3  Example 2: Use
	// tag key-pairs applied to instances
	// Key=tag:my-tag-key,Values=my-tag-value-1,my-tag-value-2  Example 3: Use tag-keys
	// applied to instances Key=tag-key,Values=my-tag-key-1,my-tag-key-2  <p>
	// <b>Example 4</b>: Use resource group names</p> <p>
	// <code>Key=resource-groups:Name,Values=<i>resource-group-name</i> </code> </p>
	// <p> <b>Example 5</b>: Use filters for resource group types</p> <p>
	// <code>Key=resource-groups:ResourceTypeFilters,Values=<i>resource-type-1</i>,<i>resource-type-2</i>
	// </code> </p> <note> <p>For <code>Key=resource-groups:ResourceTypeFilters</code>,
	// specify resource types in the following format</p> <p>
	// <code>Key=resource-groups:ResourceTypeFilters,Values=<i>AWS::EC2::INSTANCE</i>,<i>AWS::EC2::VPC</i>
	// </code> </p> </note> <p>For more information about these examples formats,
	// including the best use case for each one,  see Examples: Register targets with a
	// maintenance window
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	// in the AWS Systems Manager User Guide.
	//
	// This member is required.
	Targets []*types.Target

	// The ID of the maintenance window the target should be registered with.
	//
	// This member is required.
	WindowId *string

	// User-provided idempotency token.
	ClientToken *string

	// An optional description for the target.
	Description *string

	// An optional name for the target.
	Name *string

	// User-provided value that will be included in any CloudWatch events raised while
	// running tasks for these targets in this maintenance window.
	OwnerInformation *string
}

type RegisterTargetWithMaintenanceWindowOutput struct {

	// The ID of the target definition in this maintenance window.
	WindowTargetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRegisterTargetWithMaintenanceWindowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterTargetWithMaintenanceWindow{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterTargetWithMaintenanceWindow{}, middleware.After)
}

type idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RegisterTargetWithMaintenanceWindowInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RegisterTargetWithMaintenanceWindowInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opRegisterTargetWithMaintenanceWindowMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpRegisterTargetWithMaintenanceWindow{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRegisterTargetWithMaintenanceWindow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "RegisterTargetWithMaintenanceWindow",
	}
}
