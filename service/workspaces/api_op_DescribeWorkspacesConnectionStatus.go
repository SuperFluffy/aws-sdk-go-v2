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

// Describes the connection status of the specified WorkSpaces.
func (c *Client) DescribeWorkspacesConnectionStatus(ctx context.Context, params *DescribeWorkspacesConnectionStatusInput, optFns ...func(*Options)) (*DescribeWorkspacesConnectionStatusOutput, error) {
	stack := middleware.NewStack("DescribeWorkspacesConnectionStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeWorkspacesConnectionStatusMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspacesConnectionStatus(options.Region), middleware.Before)
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
			OperationName: "DescribeWorkspacesConnectionStatus",
			Err:           err,
		}
	}
	out := result.(*DescribeWorkspacesConnectionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspacesConnectionStatusInput struct {

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string

	// The identifiers of the WorkSpaces. You can specify up to 25 WorkSpaces.
	WorkspaceIds []*string
}

type DescribeWorkspacesConnectionStatusOutput struct {

	// The token to use to retrieve the next set of results, or null if no more results
	// are available.
	NextToken *string

	// Information about the connection status of the WorkSpace.
	WorkspacesConnectionStatus []*types.WorkspaceConnectionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeWorkspacesConnectionStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspacesConnectionStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspacesConnectionStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeWorkspacesConnectionStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspacesConnectionStatus",
	}
}
