// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list that describes the streaming sessions for a specified stack and
// fleet. If a UserId is provided for the stack and fleet, only streaming sessions
// for that user are described. If an authentication type is not provided, the
// default is to authenticate users using a streaming URL.
func (c *Client) DescribeSessions(ctx context.Context, params *DescribeSessionsInput, optFns ...func(*Options)) (*DescribeSessionsOutput, error) {
	stack := middleware.NewStack("DescribeSessions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeSessionsMiddlewares(stack)
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
	addOpDescribeSessionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSessions(options.Region), middleware.Before)
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
			OperationName: "DescribeSessions",
			Err:           err,
		}
	}
	out := result.(*DescribeSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSessionsInput struct {

	// The name of the fleet. This value is case-sensitive.
	//
	// This member is required.
	FleetName *string

	// The name of the stack. This value is case-sensitive.
	//
	// This member is required.
	StackName *string

	// The authentication method. Specify API for a user authenticated using a
	// streaming URL or SAML for a SAML federated user. The default is to authenticate
	// users using a streaming URL.
	AuthenticationType types.AuthenticationType

	// The size of each page of results. The default value is 20 and the maximum value
	// is 50.
	Limit *int32

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	// The user identifier.
	UserId *string
}

type DescribeSessionsOutput struct {

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// Information about the streaming sessions.
	Sessions []*types.Session

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeSessionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSessions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSessions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSessions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "DescribeSessions",
	}
}
