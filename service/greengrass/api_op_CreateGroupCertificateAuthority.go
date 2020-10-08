// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a CA for the group. If a CA already exists, it will rotate the existing
// CA.
func (c *Client) CreateGroupCertificateAuthority(ctx context.Context, params *CreateGroupCertificateAuthorityInput, optFns ...func(*Options)) (*CreateGroupCertificateAuthorityOutput, error) {
	stack := middleware.NewStack("CreateGroupCertificateAuthority", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateGroupCertificateAuthorityMiddlewares(stack)
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
	addOpCreateGroupCertificateAuthorityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGroupCertificateAuthority(options.Region), middleware.Before)
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
			OperationName: "CreateGroupCertificateAuthority",
			Err:           err,
		}
	}
	out := result.(*CreateGroupCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGroupCertificateAuthorityInput struct {

	// The ID of the Greengrass group.
	//
	// This member is required.
	GroupId *string

	// A client token used to correlate requests and responses.
	AmznClientToken *string
}

type CreateGroupCertificateAuthorityOutput struct {

	// The ARN of the group certificate authority.
	GroupCertificateAuthorityArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateGroupCertificateAuthorityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateGroupCertificateAuthority{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGroupCertificateAuthority{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateGroupCertificateAuthority(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateGroupCertificateAuthority",
	}
}
