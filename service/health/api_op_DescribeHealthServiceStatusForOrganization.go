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

// This operation provides status information on enabling or disabling AWS Health
// to work with your organization. To call this operation, you must sign in as an
// IAM user, assume an IAM role, or sign in as the root user (not recommended) in
// the organization's master account.
func (c *Client) DescribeHealthServiceStatusForOrganization(ctx context.Context, params *DescribeHealthServiceStatusForOrganizationInput, optFns ...func(*Options)) (*DescribeHealthServiceStatusForOrganizationOutput, error) {
	stack := middleware.NewStack("DescribeHealthServiceStatusForOrganization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeHealthServiceStatusForOrganizationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHealthServiceStatusForOrganization(options.Region), middleware.Before)
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
			OperationName: "DescribeHealthServiceStatusForOrganization",
			Err:           err,
		}
	}
	out := result.(*DescribeHealthServiceStatusForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHealthServiceStatusForOrganizationInput struct {
}

type DescribeHealthServiceStatusForOrganizationOutput struct {

	// Information about the status of enabling or disabling AWS Health Organizational
	// View in your organization. Valid values are ENABLED | DISABLED | PENDING.
	HealthServiceAccessStatusForOrganization *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeHealthServiceStatusForOrganizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeHealthServiceStatusForOrganization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeHealthServiceStatusForOrganization{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeHealthServiceStatusForOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DescribeHealthServiceStatusForOrganization",
	}
}
