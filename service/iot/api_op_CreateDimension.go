// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a dimension that you can use to limit the scope of a metric used in a
// security profile for AWS IoT Device Defender. For example, using a TOPIC_FILTER
// dimension, you can narrow down the scope of the metric only to MQTT topics whose
// name match the pattern specified in the dimension.
func (c *Client) CreateDimension(ctx context.Context, params *CreateDimensionInput, optFns ...func(*Options)) (*CreateDimensionOutput, error) {
	stack := middleware.NewStack("CreateDimension", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDimensionMiddlewares(stack)
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
	addIdempotencyToken_opCreateDimensionMiddleware(stack, options)
	addOpCreateDimensionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDimension(options.Region), middleware.Before)
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
			OperationName: "CreateDimension",
			Err:           err,
		}
	}
	out := result.(*CreateDimensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDimensionInput struct {

	// Each dimension must have a unique client request token. If you try to create a
	// new dimension with the same token as a dimension that already exists, an
	// exception occurs. If you omit this value, AWS SDKs will automatically generate a
	// unique client request.
	//
	// This member is required.
	ClientRequestToken *string

	// A unique identifier for the dimension. Choose something that describes the type
	// and value to make it easy to remember what it does.
	//
	// This member is required.
	Name *string

	// Specifies the value or list of values for the dimension. For TOPIC_FILTER
	// dimensions, this is a pattern used to match the MQTT topic (for example,
	// "admin/#").
	//
	// This member is required.
	StringValues []*string

	// Specifies the type of dimension. Supported types: TOPIC_FILTER.
	//
	// This member is required.
	Type types.DimensionType

	// Metadata that can be used to manage the dimension.
	Tags []*types.Tag
}

type CreateDimensionOutput struct {

	// The ARN (Amazon resource name) of the created dimension.
	Arn *string

	// A unique identifier for the dimension.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDimensionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDimension{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDimension{}, middleware.After)
}

type idempotencyToken_initializeOpCreateDimension struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDimension) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDimension) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDimensionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDimensionInput ")
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
func addIdempotencyToken_opCreateDimensionMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateDimension{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDimension(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateDimension",
	}
}
