// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds tags to a resource group with the specified ARN. Existing tags on a
// resource group are not changed if they are not specified in the request
// parameters. Do not store personally identifiable information (PII) or other
// confidential or sensitive information in tags. We use tags to provide you with
// billing and administration services. Tags are not intended to be used for
// private or sensitive data.
func (c *Client) Tag(ctx context.Context, params *TagInput, optFns ...func(*Options)) (*TagOutput, error) {
	stack := middleware.NewStack("Tag", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpTagMiddlewares(stack)
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
	addOpTagValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTag(options.Region), middleware.Before)
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
			OperationName: "Tag",
			Err:           err,
		}
	}
	out := result.(*TagOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagInput struct {

	// The ARN of the resource group to which to add tags.
	//
	// This member is required.
	Arn *string

	// The tags to add to the specified resource group. A tag is a string-to-string map
	// of key-value pairs.
	//
	// This member is required.
	Tags map[string]*string
}

type TagOutput struct {

	// The ARN of the tagged resource.
	Arn *string

	// The tags that have been added to the specified resource group.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpTagMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpTag{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpTag{}, middleware.After)
}

func newServiceMetadataMiddleware_opTag(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "Tag",
	}
}
