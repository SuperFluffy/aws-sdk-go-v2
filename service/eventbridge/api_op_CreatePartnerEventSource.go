// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Called by an SaaS partner to create a partner event source. This operation is
// not used by AWS customers. Each partner event source can be used by one AWS
// account to create a matching partner event bus in that AWS account. A SaaS
// partner must create one partner event source for each AWS account that wants to
// receive those event types. A partner event source creates events based on
// resources within the SaaS partner's service or application. An AWS account that
// creates a partner event bus that matches the partner event source can use that
// event bus to receive events from the partner, and then process them using AWS
// Events rules and targets. Partner event source names follow this format:
// partner_name/event_namespace/event_name  partner_name is determined during
// partner registration and identifies the partner to AWS customers.
// event_namespace is determined by the partner and is a way for the partner to
// categorize their events. event_name is determined by the partner, and should
// uniquely identify an event-generating resource within the partner system. The
// combination of event_namespace and event_name should help AWS customers decide
// whether to create an event bus to receive these events.
func (c *Client) CreatePartnerEventSource(ctx context.Context, params *CreatePartnerEventSourceInput, optFns ...func(*Options)) (*CreatePartnerEventSourceOutput, error) {
	stack := middleware.NewStack("CreatePartnerEventSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreatePartnerEventSourceMiddlewares(stack)
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
	addOpCreatePartnerEventSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePartnerEventSource(options.Region), middleware.Before)
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
			OperationName: "CreatePartnerEventSource",
			Err:           err,
		}
	}
	out := result.(*CreatePartnerEventSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePartnerEventSourceInput struct {

	// The AWS account ID that is permitted to create a matching partner event bus for
	// this partner event source.
	//
	// This member is required.
	Account *string

	// The name of the partner event source. This name must be unique and must be in
	// the format  partner_name/event_namespace/event_name . The AWS account that wants
	// to use this partner event source must create a partner event bus with a name
	// that matches the name of the partner event source.
	//
	// This member is required.
	Name *string
}

type CreatePartnerEventSourceOutput struct {

	// The ARN of the partner event source.
	EventSourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreatePartnerEventSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePartnerEventSource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePartnerEventSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePartnerEventSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "CreatePartnerEventSource",
	}
}
