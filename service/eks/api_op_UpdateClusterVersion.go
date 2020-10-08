// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an Amazon EKS cluster to the specified Kubernetes version. Your cluster
// continues to function during the update. The response output includes an update
// ID that you can use to track the status of your cluster update with the
// DescribeUpdate () API operation. Cluster updates are asynchronous, and they
// should finish within a few minutes. During an update, the cluster status moves
// to UPDATING (this status transition is eventually consistent). When the update
// is complete (either Failed or Successful), the cluster status moves to Active.
// If your cluster has managed node groups attached to it, all of your node groups’
// Kubernetes versions must match the cluster’s Kubernetes version in order to
// update the cluster to a new Kubernetes version.
func (c *Client) UpdateClusterVersion(ctx context.Context, params *UpdateClusterVersionInput, optFns ...func(*Options)) (*UpdateClusterVersionOutput, error) {
	stack := middleware.NewStack("UpdateClusterVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateClusterVersionMiddlewares(stack)
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
	addIdempotencyToken_opUpdateClusterVersionMiddleware(stack, options)
	addOpUpdateClusterVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateClusterVersion(options.Region), middleware.Before)
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
			OperationName: "UpdateClusterVersion",
			Err:           err,
		}
	}
	out := result.(*UpdateClusterVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClusterVersionInput struct {

	// The name of the Amazon EKS cluster to update.
	//
	// This member is required.
	Name *string

	// The desired Kubernetes version following a successful update.
	//
	// This member is required.
	Version *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string
}

type UpdateClusterVersionOutput struct {

	// The full description of the specified update
	Update *types.Update

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateClusterVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateClusterVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateClusterVersion{}, middleware.After)
}

type idempotencyToken_initializeOpUpdateClusterVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateClusterVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateClusterVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateClusterVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateClusterVersionInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateClusterVersionMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpUpdateClusterVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateClusterVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "UpdateClusterVersion",
	}
}
