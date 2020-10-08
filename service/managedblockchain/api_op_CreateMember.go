// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a member within a Managed Blockchain network.
func (c *Client) CreateMember(ctx context.Context, params *CreateMemberInput, optFns ...func(*Options)) (*CreateMemberOutput, error) {
	stack := middleware.NewStack("CreateMember", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateMemberMiddlewares(stack)
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
	addIdempotencyToken_opCreateMemberMiddleware(stack, options)
	addOpCreateMemberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMember(options.Region), middleware.Before)
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
			OperationName: "CreateMember",
			Err:           err,
		}
	}
	out := result.(*CreateMemberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMemberInput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the operation. An idempotent operation completes no more than one time. This
	// identifier is required only if you make a service request directly using an HTTP
	// client. It is generated automatically if you use an AWS SDK or the AWS CLI.
	//
	// This member is required.
	ClientRequestToken *string

	// The unique identifier of the invitation that is sent to the member to join the
	// network.
	//
	// This member is required.
	InvitationId *string

	// Member configuration parameters.
	//
	// This member is required.
	MemberConfiguration *types.MemberConfiguration

	// The unique identifier of the network in which the member is created.
	//
	// This member is required.
	NetworkId *string
}

type CreateMemberOutput struct {

	// The unique identifier of the member.
	MemberId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateMemberMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateMember{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMember{}, middleware.After)
}

type idempotencyToken_initializeOpCreateMember struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMember) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMember) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMemberInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMemberInput ")
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
func addIdempotencyToken_opCreateMemberMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateMember{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMember(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "CreateMember",
	}
}
