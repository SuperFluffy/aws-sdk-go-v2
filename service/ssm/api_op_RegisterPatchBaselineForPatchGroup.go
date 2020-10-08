// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a patch baseline for a patch group.
func (c *Client) RegisterPatchBaselineForPatchGroup(ctx context.Context, params *RegisterPatchBaselineForPatchGroupInput, optFns ...func(*Options)) (*RegisterPatchBaselineForPatchGroupOutput, error) {
	stack := middleware.NewStack("RegisterPatchBaselineForPatchGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRegisterPatchBaselineForPatchGroupMiddlewares(stack)
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
	addOpRegisterPatchBaselineForPatchGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterPatchBaselineForPatchGroup(options.Region), middleware.Before)
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
			OperationName: "RegisterPatchBaselineForPatchGroup",
			Err:           err,
		}
	}
	out := result.(*RegisterPatchBaselineForPatchGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterPatchBaselineForPatchGroupInput struct {

	// The ID of the patch baseline to register the patch group with.
	//
	// This member is required.
	BaselineId *string

	// The name of the patch group that should be registered with the patch baseline.
	//
	// This member is required.
	PatchGroup *string
}

type RegisterPatchBaselineForPatchGroupOutput struct {

	// The ID of the patch baseline the patch group was registered with.
	BaselineId *string

	// The name of the patch group registered with the patch baseline.
	PatchGroup *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRegisterPatchBaselineForPatchGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterPatchBaselineForPatchGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterPatchBaselineForPatchGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterPatchBaselineForPatchGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "RegisterPatchBaselineForPatchGroup",
	}
}
