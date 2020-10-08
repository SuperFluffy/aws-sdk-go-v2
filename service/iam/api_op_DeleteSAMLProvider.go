// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a SAML provider resource in IAM. Deleting the provider resource from IAM
// does not update any roles that reference the SAML provider resource's ARN as a
// principal in their trust policies. Any attempt to assume a role that references
// a non-existent provider resource ARN fails. This operation requires Signature
// Version 4
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
func (c *Client) DeleteSAMLProvider(ctx context.Context, params *DeleteSAMLProviderInput, optFns ...func(*Options)) (*DeleteSAMLProviderOutput, error) {
	stack := middleware.NewStack("DeleteSAMLProvider", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteSAMLProviderMiddlewares(stack)
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
	addOpDeleteSAMLProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSAMLProvider(options.Region), middleware.Before)
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
			OperationName: "DeleteSAMLProvider",
			Err:           err,
		}
	}
	out := result.(*DeleteSAMLProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSAMLProviderInput struct {

	// The Amazon Resource Name (ARN) of the SAML provider to delete.
	//
	// This member is required.
	SAMLProviderArn *string
}

type DeleteSAMLProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteSAMLProviderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteSAMLProvider{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteSAMLProvider{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSAMLProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteSAMLProvider",
	}
}
