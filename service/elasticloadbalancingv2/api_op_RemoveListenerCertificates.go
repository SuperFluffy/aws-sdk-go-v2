// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified certificate from the certificate list for the specified
// HTTPS or TLS listener. You can't remove the default certificate for a listener.
// To replace the default certificate, call ModifyListener (). To list the
// certificates for your listener, use DescribeListenerCertificates ().
func (c *Client) RemoveListenerCertificates(ctx context.Context, params *RemoveListenerCertificatesInput, optFns ...func(*Options)) (*RemoveListenerCertificatesOutput, error) {
	stack := middleware.NewStack("RemoveListenerCertificates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRemoveListenerCertificatesMiddlewares(stack)
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
	addOpRemoveListenerCertificatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveListenerCertificates(options.Region), middleware.Before)
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
			OperationName: "RemoveListenerCertificates",
			Err:           err,
		}
	}
	out := result.(*RemoveListenerCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveListenerCertificatesInput struct {

	// The certificate to remove. You can specify one certificate per call. Set
	// CertificateArn to the certificate ARN but do not set IsDefault.
	//
	// This member is required.
	Certificates []*types.Certificate

	// The Amazon Resource Name (ARN) of the listener.
	//
	// This member is required.
	ListenerArn *string
}

type RemoveListenerCertificatesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRemoveListenerCertificatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRemoveListenerCertificates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveListenerCertificates{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveListenerCertificates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "RemoveListenerCertificates",
	}
}
