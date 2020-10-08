// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Forgets the device, as an administrator. Calling this action requires developer
// credentials.
func (c *Client) AdminForgetDevice(ctx context.Context, params *AdminForgetDeviceInput, optFns ...func(*Options)) (*AdminForgetDeviceOutput, error) {
	stack := middleware.NewStack("AdminForgetDevice", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminForgetDeviceMiddlewares(stack)
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
	addOpAdminForgetDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminForgetDevice(options.Region), middleware.Before)
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
			OperationName: "AdminForgetDevice",
			Err:           err,
		}
	}
	out := result.(*AdminForgetDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Sends the forgot device request, as an administrator.
type AdminForgetDeviceInput struct {

	// The device key.
	//
	// This member is required.
	DeviceKey *string

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The user name.
	//
	// This member is required.
	Username *string
}

type AdminForgetDeviceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminForgetDeviceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminForgetDevice{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminForgetDevice{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminForgetDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminForgetDevice",
	}
}
