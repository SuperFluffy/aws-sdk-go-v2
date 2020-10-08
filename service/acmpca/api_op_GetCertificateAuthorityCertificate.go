// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the certificate and certificate chain for your private certificate
// authority (CA). Both the certificate and the chain are base64 PEM-encoded. The
// chain does not include the CA certificate. Each certificate in the chain signs
// the one before it.
func (c *Client) GetCertificateAuthorityCertificate(ctx context.Context, params *GetCertificateAuthorityCertificateInput, optFns ...func(*Options)) (*GetCertificateAuthorityCertificateOutput, error) {
	stack := middleware.NewStack("GetCertificateAuthorityCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetCertificateAuthorityCertificateMiddlewares(stack)
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
	addOpGetCertificateAuthorityCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCertificateAuthorityCertificate(options.Region), middleware.Before)
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
			OperationName: "GetCertificateAuthorityCertificate",
			Err:           err,
		}
	}
	out := result.(*GetCertificateAuthorityCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCertificateAuthorityCertificateInput struct {

	// The Amazon Resource Name (ARN) of your private CA. This is of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string
}

type GetCertificateAuthorityCertificateOutput struct {

	// Base64-encoded certificate authority (CA) certificate.
	Certificate *string

	// Base64-encoded certificate chain that includes any intermediate certificates and
	// chains up to root on-premises certificate that you used to sign your private CA
	// certificate. The chain does not include your private CA certificate. If this is
	// a root CA, the value will be null.
	CertificateChain *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetCertificateAuthorityCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetCertificateAuthorityCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCertificateAuthorityCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCertificateAuthorityCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "GetCertificateAuthorityCertificate",
	}
}
