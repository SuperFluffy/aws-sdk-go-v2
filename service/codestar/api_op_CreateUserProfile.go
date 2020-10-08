// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a profile for a user that includes user preferences, such as the display
// name and email address assocciated with the user, in AWS CodeStar. The user
// profile is not project-specific. Information in the user profile is displayed
// wherever the user's information appears to other users in AWS CodeStar.
func (c *Client) CreateUserProfile(ctx context.Context, params *CreateUserProfileInput, optFns ...func(*Options)) (*CreateUserProfileOutput, error) {
	stack := middleware.NewStack("CreateUserProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateUserProfileMiddlewares(stack)
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
	addOpCreateUserProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserProfile(options.Region), middleware.Before)
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
			OperationName: "CreateUserProfile",
			Err:           err,
		}
	}
	out := result.(*CreateUserProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserProfileInput struct {

	// The name that will be displayed as the friendly name for the user in AWS
	// CodeStar.
	//
	// This member is required.
	DisplayName *string

	// The email address that will be displayed as part of the user's profile in AWS
	// CodeStar.
	//
	// This member is required.
	EmailAddress *string

	// The Amazon Resource Name (ARN) of the user in IAM.
	//
	// This member is required.
	UserArn *string

	// The SSH public key associated with the user in AWS CodeStar. If a project owner
	// allows the user remote access to project resources, this public key will be used
	// along with the user's private key for SSH access.
	SshPublicKey *string
}

type CreateUserProfileOutput struct {

	// The Amazon Resource Name (ARN) of the user in IAM.
	//
	// This member is required.
	UserArn *string

	// The date the user profile was created, in timestamp format.
	CreatedTimestamp *time.Time

	// The name that is displayed as the friendly name for the user in AWS CodeStar.
	DisplayName *string

	// The email address that is displayed as part of the user's profile in AWS
	// CodeStar.
	EmailAddress *string

	// The date the user profile was last modified, in timestamp format.
	LastModifiedTimestamp *time.Time

	// The SSH public key associated with the user in AWS CodeStar. This is the public
	// portion of the public/private keypair the user can use to access project
	// resources if a project owner allows the user remote access to those resources.
	SshPublicKey *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateUserProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUserProfile{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUserProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateUserProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar",
		OperationName: "CreateUserProfile",
	}
}
