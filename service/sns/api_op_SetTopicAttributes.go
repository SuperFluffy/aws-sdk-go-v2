// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows a topic owner to set an attribute of the topic to a new value.
func (c *Client) SetTopicAttributes(ctx context.Context, params *SetTopicAttributesInput, optFns ...func(*Options)) (*SetTopicAttributesOutput, error) {
	stack := middleware.NewStack("SetTopicAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetTopicAttributesMiddlewares(stack)
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
	addOpSetTopicAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetTopicAttributes(options.Region), middleware.Before)
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
			OperationName: "SetTopicAttributes",
			Err:           err,
		}
	}
	out := result.(*SetTopicAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for SetTopicAttributes action.
type SetTopicAttributesInput struct {

	// A map of attributes with their corresponding values. The following lists the
	// names, descriptions, and values of the special request parameters that the
	// SetTopicAttributes action uses:
	//
	//     * DeliveryPolicy – The policy that defines
	// how Amazon SNS retries failed deliveries to HTTP/S endpoints.
	//
	//     * DisplayName
	// – The display name to use for a topic with SMS subscriptions.
	//
	//     * Policy –
	// The policy that defines who can access your topic. By default, only the topic
	// owner can publish or subscribe to the topic.
	//
	//     <p>The following attribute
	// applies only to <a
	// href="https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html">server-side-encryption</a>:</p>
	// <ul> <li> <p> <code>KmsMasterKeyId</code> – The ID of an AWS-managed customer
	// master key (CMK) for Amazon SNS or a custom CMK. For more information, see <a
	// href="https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms">Key
	// Terms</a>. For more examples, see <a
	// href="https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters">KeyId</a>
	// in the <i>AWS Key Management Service API Reference</i>. </p> </li> </ul> <p>The
	// following attribute applies only to FIFO topics:</p> <ul> <li> <p>
	// <code>ContentBasedDeduplication</code> – Enables content-based deduplication.
	// Amazon SNS uses a SHA-256 hash to generate the
	// <code>MessageDeduplicationId</code> using the body of the message (but not the
	// attributes of the message). </p> </li> <li> <p> When
	// <code>ContentBasedDeduplication</code> is in effect, messages with identical
	// content sent within the deduplication interval are treated as duplicates and
	// only one copy of the message is delivered. </p> </li> <li> <p> If the queue has
	// <code>ContentBasedDeduplication</code> set, your
	// <code>MessageDeduplicationId</code> overrides the generated one. </p> </li>
	// </ul>
	//
	// This member is required.
	AttributeName *string

	// The ARN of the topic to modify.
	//
	// This member is required.
	TopicArn *string

	// The new value for the attribute.
	AttributeValue *string
}

type SetTopicAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetTopicAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetTopicAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetTopicAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetTopicAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "SetTopicAttributes",
	}
}
