// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information on a specific signing profile.
func (c *Client) GetSigningProfile(ctx context.Context, params *GetSigningProfileInput, optFns ...func(*Options)) (*GetSigningProfileOutput, error) {
	stack := middleware.NewStack("GetSigningProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSigningProfileMiddlewares(stack)
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
	addOpGetSigningProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSigningProfile(options.Region), middleware.Before)
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
			OperationName: "GetSigningProfile",
			Err:           err,
		}
	}
	out := result.(*GetSigningProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSigningProfileInput struct {

	// The name of the target signing profile.
	//
	// This member is required.
	ProfileName *string
}

type GetSigningProfileOutput struct {

	// The Amazon Resource Name (ARN) for the signing profile.
	Arn *string

	// A list of overrides applied by the target signing profile for signing
	// operations.
	Overrides *types.SigningPlatformOverrides

	// The ID of the platform that is used by the target signing profile.
	PlatformId *string

	// The name of the target signing profile.
	ProfileName *string

	// The ARN of the certificate that the target profile uses for signing operations.
	SigningMaterial *types.SigningMaterial

	// A map of key-value pairs for signing operations that is attached to the target
	// signing profile.
	SigningParameters map[string]*string

	// The status of the target signing profile.
	Status types.SigningProfileStatus

	// A list of tags associated with the signing profile.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSigningProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSigningProfile{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSigningProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSigningProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "signer",
		OperationName: "GetSigningProfile",
	}
}
