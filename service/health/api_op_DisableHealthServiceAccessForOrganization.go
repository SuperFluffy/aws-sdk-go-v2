// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Calling this operation disables Health from working with AWS Organizations. This
// does not remove the Service Linked Role (SLR) from the the master account in
// your organization. Use the IAM console, API, or AWS CLI to remove the SLR if
// desired. To call this operation, you must sign in as an IAM user, assume an IAM
// role, or sign in as the root user (not recommended) in the organization's master
// account.
func (c *Client) DisableHealthServiceAccessForOrganization(ctx context.Context, params *DisableHealthServiceAccessForOrganizationInput, optFns ...func(*Options)) (*DisableHealthServiceAccessForOrganizationOutput, error) {
	stack := middleware.NewStack("DisableHealthServiceAccessForOrganization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisableHealthServiceAccessForOrganizationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableHealthServiceAccessForOrganization(options.Region), middleware.Before)
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
			OperationName: "DisableHealthServiceAccessForOrganization",
			Err:           err,
		}
	}
	out := result.(*DisableHealthServiceAccessForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableHealthServiceAccessForOrganizationInput struct {
}

type DisableHealthServiceAccessForOrganizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisableHealthServiceAccessForOrganizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisableHealthServiceAccessForOrganization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableHealthServiceAccessForOrganization{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableHealthServiceAccessForOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DisableHealthServiceAccessForOrganization",
	}
}
