// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches a policy to a root, an organizational unit (OU), or an individual
// account. How the policy affects accounts depends on the type of policy. Refer to
// the AWS Organizations User Guide for information about each policy type:
//
//     *
// AISERVICES_OPT_OUT_POLICY
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
//
//
// * BACKUP_POLICY
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html)
//
//
// * SERVICE_CONTROL_POLICY
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html)
//
//
// * TAG_POLICY
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
//
// This
// operation can be called only from the organization's master account.
func (c *Client) AttachPolicy(ctx context.Context, params *AttachPolicyInput, optFns ...func(*Options)) (*AttachPolicyOutput, error) {
	stack := middleware.NewStack("AttachPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAttachPolicyMiddlewares(stack)
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
	addOpAttachPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachPolicy(options.Region), middleware.Before)
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
			OperationName: "AttachPolicy",
			Err:           err,
		}
	}
	out := result.(*AttachPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachPolicyInput struct {

	// The unique identifier (ID) of the policy that you want to attach to the target.
	// You can get the ID for the policy by calling the ListPolicies () operation. The
	// regex pattern (http://wikipedia.org/wiki/regex) for a policy ID string requires
	// "p-" followed by from 8 to 128 lowercase or uppercase letters, digits, or the
	// underscore character (_).
	//
	// This member is required.
	PolicyId *string

	// The unique identifier (ID) of the root, OU, or account that you want to attach
	// the policy to. You can get the ID by calling the ListRoots (),
	// ListOrganizationalUnitsForParent (), or ListAccounts () operations. The regex
	// pattern (http://wikipedia.org/wiki/regex) for a target ID string requires one of
	// the following:
	//
	//     * Root - A string that begins with "r-" followed by from 4
	// to 32 lowercase letters or digits.
	//
	//     * Account - A string that consists of
	// exactly 12 digits.
	//
	//     * Organizational unit (OU) - A string that begins with
	// "ou-" followed by from 4 to 32 lowercase letters or digits (the ID of the root
	// that the OU is in). This string is followed by a second "-" dash and from 8 to
	// 32 additional lowercase letters or digits.
	//
	// This member is required.
	TargetId *string
}

type AttachPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAttachPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAttachPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "AttachPolicy",
	}
}
