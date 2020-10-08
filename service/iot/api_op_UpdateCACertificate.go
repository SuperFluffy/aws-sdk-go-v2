// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a registered CA certificate.
func (c *Client) UpdateCACertificate(ctx context.Context, params *UpdateCACertificateInput, optFns ...func(*Options)) (*UpdateCACertificateOutput, error) {
	stack := middleware.NewStack("UpdateCACertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateCACertificateMiddlewares(stack)
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
	addOpUpdateCACertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCACertificate(options.Region), middleware.Before)
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
			OperationName: "UpdateCACertificate",
			Err:           err,
		}
	}
	out := result.(*UpdateCACertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input to the UpdateCACertificate operation.
type UpdateCACertificateInput struct {

	// The CA certificate identifier.
	//
	// This member is required.
	CertificateId *string

	// The new value for the auto registration status. Valid values are: "ENABLE" or
	// "DISABLE".
	NewAutoRegistrationStatus types.AutoRegistrationStatus

	// The updated status of the CA certificate. Note: The status value
	// REGISTER_INACTIVE is deprecated and should not be used.
	NewStatus types.CACertificateStatus

	// Information about the registration configuration.
	RegistrationConfig *types.RegistrationConfig

	// If true, removes auto registration.
	RemoveAutoRegistration *bool
}

type UpdateCACertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateCACertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCACertificate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCACertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateCACertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateCACertificate",
	}
}
