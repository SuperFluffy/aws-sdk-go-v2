// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Terminates the specified WorkSpaces. Terminating a WorkSpace is a permanent
// action and cannot be undone. The user's data is destroyed. If you need to
// archive any user data, contact Amazon Web Services before terminating the
// WorkSpace. You can terminate a WorkSpace that is in any state except SUSPENDED.
// This operation is asynchronous and returns before the WorkSpaces have been
// completely terminated.
func (c *Client) TerminateWorkspaces(ctx context.Context, params *TerminateWorkspacesInput, optFns ...func(*Options)) (*TerminateWorkspacesOutput, error) {
	stack := middleware.NewStack("TerminateWorkspaces", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTerminateWorkspacesMiddlewares(stack)
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
	addOpTerminateWorkspacesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateWorkspaces(options.Region), middleware.Before)
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
			OperationName: "TerminateWorkspaces",
			Err:           err,
		}
	}
	out := result.(*TerminateWorkspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateWorkspacesInput struct {

	// The WorkSpaces to terminate. You can specify up to 25 WorkSpaces.
	//
	// This member is required.
	TerminateWorkspaceRequests []*types.TerminateRequest
}

type TerminateWorkspacesOutput struct {

	// Information about the WorkSpaces that could not be terminated.
	FailedRequests []*types.FailedWorkspaceChangeRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTerminateWorkspacesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTerminateWorkspaces{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTerminateWorkspaces{}, middleware.After)
}

func newServiceMetadataMiddleware_opTerminateWorkspaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "TerminateWorkspaces",
	}
}
