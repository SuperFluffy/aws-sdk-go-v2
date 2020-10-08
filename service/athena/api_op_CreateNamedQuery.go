// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a named query in the specified workgroup. Requires that you have access
// to the workgroup. For code samples using the AWS SDK for Java, see Examples and
// Code Samples (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in
// the Amazon Athena User Guide.
func (c *Client) CreateNamedQuery(ctx context.Context, params *CreateNamedQueryInput, optFns ...func(*Options)) (*CreateNamedQueryOutput, error) {
	stack := middleware.NewStack("CreateNamedQuery", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateNamedQueryMiddlewares(stack)
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
	addIdempotencyToken_opCreateNamedQueryMiddleware(stack, options)
	addOpCreateNamedQueryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNamedQuery(options.Region), middleware.Before)
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
			OperationName: "CreateNamedQuery",
			Err:           err,
		}
	}
	out := result.(*CreateNamedQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNamedQueryInput struct {

	// The database to which the query belongs.
	//
	// This member is required.
	Database *string

	// The query name.
	//
	// This member is required.
	Name *string

	// The contents of the query with all query statements.
	//
	// This member is required.
	QueryString *string

	// A unique case-sensitive string used to ensure the request to create the query is
	// idempotent (executes only once). If another CreateNamedQuery request is
	// received, the same response is returned and another query is not created. If a
	// parameter has changed, for example, the QueryString, an error is returned. This
	// token is listed as not required because AWS SDKs (for example the AWS SDK for
	// Java) auto-generate the token for users. If you are not using the AWS SDK or the
	// AWS CLI, you must provide this token or the action will fail.
	ClientRequestToken *string

	// The query description.
	Description *string

	// The name of the workgroup in which the named query is being created.
	WorkGroup *string
}

type CreateNamedQueryOutput struct {

	// The unique ID of the query.
	NamedQueryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateNamedQueryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNamedQuery{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNamedQuery{}, middleware.After)
}

type idempotencyToken_initializeOpCreateNamedQuery struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNamedQuery) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNamedQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNamedQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNamedQueryInput ")
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
func addIdempotencyToken_opCreateNamedQueryMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateNamedQuery{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNamedQuery(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "CreateNamedQuery",
	}
}
