// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the specified rule. A disabled rule won't match any events, and won't
// self-trigger if it has a schedule expression.  <p>When you disable a rule,
// incoming events might continue to match to the disabled rule. Allow a short
// period of time for changes to take effect.</p>
func (c *Client) DisableRule(ctx context.Context, params *DisableRuleInput, optFns ...func(*Options)) (*DisableRuleOutput, error) {
	stack := middleware.NewStack("DisableRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisableRuleMiddlewares(stack)
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
	addOpDisableRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableRule(options.Region), middleware.Before)
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
			OperationName: "DisableRule",
			Err:           err,
		}
	}
	out := result.(*DisableRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableRuleInput struct {

	// The name of the rule.
	//
	// This member is required.
	Name *string

	// The event bus associated with the rule. If you omit this, the default event bus
	// is used.
	EventBusName *string
}

type DisableRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisableRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisableRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DisableRule",
	}
}
