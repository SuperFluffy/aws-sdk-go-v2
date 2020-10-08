// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Attaches an IAM policy to the specified resource. The only supported
// use for this action is to share a RuleGroup across accounts. The
// PutPermissionPolicy is subject to the following restrictions:
//
//     * You can
// attach only one policy with each PutPermissionPolicy request.
//
//     * The policy
// must include an Effect, Action and Principal.
//
//     * <p> <code>Effect</code>
// must specify <code>Allow</code>.</p> </li> <li> <p>The <code>Action</code> in
// the policy must be <code>waf:UpdateWebACL</code>,
// <code>waf-regional:UpdateWebACL</code>, <code>waf:GetRuleGroup</code> and
// <code>waf-regional:GetRuleGroup</code> . Any extra or wildcard actions in the
// policy will be rejected.</p> </li> <li> <p>The policy cannot include a
// <code>Resource</code> parameter.</p> </li> <li> <p>The ARN in the request must
// be a valid WAF RuleGroup ARN and the RuleGroup must exist in the same
// region.</p> </li> <li> <p>The user making the request must be the owner of the
// RuleGroup.</p> </li> <li> <p>Your policy must be composed using IAM Policy
// version 2012-10-17.</p> </li> </ul> <p>For more information, see <a
// href="https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html">IAM
// Policies</a>. </p> <p>An example of a valid policy parameter is shown in the
// Examples section below.</p>
func (c *Client) PutPermissionPolicy(ctx context.Context, params *PutPermissionPolicyInput, optFns ...func(*Options)) (*PutPermissionPolicyOutput, error) {
	stack := middleware.NewStack("PutPermissionPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutPermissionPolicyMiddlewares(stack)
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
	addOpPutPermissionPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutPermissionPolicy(options.Region), middleware.Before)
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
			OperationName: "PutPermissionPolicy",
			Err:           err,
		}
	}
	out := result.(*PutPermissionPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPermissionPolicyInput struct {

	// The policy to attach to the specified RuleGroup.
	//
	// This member is required.
	Policy *string

	// The Amazon Resource Name (ARN) of the RuleGroup to which you want to attach the
	// policy.
	//
	// This member is required.
	ResourceArn *string
}

type PutPermissionPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutPermissionPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutPermissionPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPermissionPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutPermissionPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "PutPermissionPolicy",
	}
}
