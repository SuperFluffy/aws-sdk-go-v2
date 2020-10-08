// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides the feedback for an authentication event whether it was from a valid
// user or not. This feedback is used for improving the risk evaluation decision
// for the user pool as part of Amazon Cognito advanced security.
func (c *Client) UpdateAuthEventFeedback(ctx context.Context, params *UpdateAuthEventFeedbackInput, optFns ...func(*Options)) (*UpdateAuthEventFeedbackOutput, error) {
	stack := middleware.NewStack("UpdateAuthEventFeedback", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateAuthEventFeedbackMiddlewares(stack)
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
	addOpUpdateAuthEventFeedbackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAuthEventFeedback(options.Region), middleware.Before)
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
			OperationName: "UpdateAuthEventFeedback",
			Err:           err,
		}
	}
	out := result.(*UpdateAuthEventFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAuthEventFeedbackInput struct {

	// The event ID.
	//
	// This member is required.
	EventId *string

	// The feedback token.
	//
	// This member is required.
	FeedbackToken *string

	// The authentication event feedback value.
	//
	// This member is required.
	FeedbackValue types.FeedbackValueType

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The user pool username.
	//
	// This member is required.
	Username *string
}

type UpdateAuthEventFeedbackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateAuthEventFeedbackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAuthEventFeedback{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAuthEventFeedback{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateAuthEventFeedback(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "UpdateAuthEventFeedback",
	}
}
