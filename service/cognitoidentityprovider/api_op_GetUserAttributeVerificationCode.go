// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the user attribute verification code for the specified attribute name.
func (c *Client) GetUserAttributeVerificationCode(ctx context.Context, params *GetUserAttributeVerificationCodeInput, optFns ...func(*Options)) (*GetUserAttributeVerificationCodeOutput, error) {
	stack := middleware.NewStack("GetUserAttributeVerificationCode", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetUserAttributeVerificationCodeMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetUserAttributeVerificationCodeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUserAttributeVerificationCode(options.Region), middleware.Before)
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
			OperationName: "GetUserAttributeVerificationCode",
			Err:           err,
		}
	}
	out := result.(*GetUserAttributeVerificationCodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to get user attribute verification.
type GetUserAttributeVerificationCodeInput struct {

	// The access token returned by the server response to get the user attribute
	// verification code.
	//
	// This member is required.
	AccessToken *string

	// The attribute name returned by the server response to get the user attribute
	// verification code.
	//
	// This member is required.
	AttributeName *string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the
	// GetUserAttributeVerificationCode API action, Amazon Cognito invokes the function
	// that is assigned to the custom message trigger. When Amazon Cognito invokes this
	// function, it passes a JSON payload, which the function receives as input. This
	// payload contains a clientMetadata attribute, which provides the data that you
	// assigned to the ClientMetadata parameter in your
	// GetUserAttributeVerificationCode request. In your function code in AWS Lambda,
	// you can process the clientMetadata value to enhance your workflow for your
	// specific needs. For more information, see Customizing User Pool Workflows with
	// Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. Take the following limitations into
	// consideration when you use the ClientMetadata parameter:
	//
	//     * Amazon Cognito
	// does not store the ClientMetadata value. This data is available only to AWS
	// Lambda triggers that are assigned to a user pool to support custom workflows. If
	// your user pool configuration does not include triggers, the ClientMetadata
	// parameter serves no purpose.
	//
	//     * Amazon Cognito does not validate the
	// ClientMetadata value.
	//
	//     * Amazon Cognito does not encrypt the the
	// ClientMetadata value, so don't use it to provide sensitive information.
	ClientMetadata map[string]*string
}

// The verification code response returned by the server response to get the user
// attribute verification code.
type GetUserAttributeVerificationCodeOutput struct {

	// The code delivery details returned by the server in response to the request to
	// get the user attribute verification code.
	CodeDeliveryDetails *types.CodeDeliveryDetailsType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetUserAttributeVerificationCodeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetUserAttributeVerificationCode{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUserAttributeVerificationCode{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUserAttributeVerificationCode(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUserAttributeVerificationCode",
	}
}
