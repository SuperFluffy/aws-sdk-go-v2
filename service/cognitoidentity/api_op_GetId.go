// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates (or retrieves) a Cognito ID. Supplying multiple logins will create an
// implicit linked account. This is a public API. You do not need any credentials
// to call this API.
func (c *Client) GetId(ctx context.Context, params *GetIdInput, optFns ...func(*Options)) (*GetIdOutput, error) {
	stack := middleware.NewStack("GetId", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetIdMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetIdValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetId(options.Region), middleware.Before)
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
			OperationName: "GetId",
			Err:           err,
		}
	}
	out := result.(*GetIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the GetId action.
type GetIdInput struct {

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string

	// A standard AWS account ID (9+ digits).
	AccountId *string

	// A set of optional name-value pairs that map provider names to provider tokens.
	// The available provider names for Logins are as follows:
	//
	//     * Facebook:
	// graph.facebook.com
	//
	//     * Amazon Cognito user pool: cognito-idp..amazonaws.com/,
	// for example, cognito-idp.us-east-1.amazonaws.com/us-east-1_123456789.
	//
	//     *
	// Google: accounts.google.com
	//
	//     * Amazon: www.amazon.com
	//
	//     * Twitter:
	// api.twitter.com
	//
	//     * Digits: www.digits.com
	Logins map[string]*string
}

// Returned in response to a GetId request.
type GetIdOutput struct {

	// A unique identifier in the format REGION:GUID.
	IdentityId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetIdMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetId{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetId{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetId(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "GetId",
	}
}
