// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Changes information about an ClientCertificate () resource.
func (c *Client) UpdateClientCertificate(ctx context.Context, params *UpdateClientCertificateInput, optFns ...func(*Options)) (*UpdateClientCertificateOutput, error) {
	stack := middleware.NewStack("UpdateClientCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateClientCertificateMiddlewares(stack)
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
	addOpUpdateClientCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateClientCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "UpdateClientCertificate",
			Err:           err,
		}
	}
	out := result.(*UpdateClientCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change information about an ClientCertificate () resource.
type UpdateClientCertificateInput struct {

	// [Required] The identifier of the ClientCertificate () resource to be updated.
	//
	// This member is required.
	ClientCertificateId *string

	Name *string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a client certificate used to configure client-side SSL authentication
// while sending requests to the integration endpoint. Client certificates are used
// to authenticate an API by the backend server. To authenticate an API client (or
// user), use IAM roles and policies, a custom Authorizer () or an Amazon Cognito
// user pool. Use Client-Side Certificate
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-client-side-ssl-authentication.html)
type UpdateClientCertificateOutput struct {

	// The identifier of the client certificate.
	ClientCertificateId *string

	// The timestamp when the client certificate was created.
	CreatedDate *time.Time

	// The description of the client certificate.
	Description *string

	// The timestamp when the client certificate will expire.
	ExpirationDate *time.Time

	// The PEM-encoded public key of the client certificate, which can be used to
	// configure certificate authentication in the integration endpoint .
	PemEncodedCertificate *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateClientCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateClientCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateClientCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateClientCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateClientCertificate",
	}
}
