// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified change set. Deleting change sets ensures that no one
// executes the wrong change set. If the call successfully completes, AWS
// CloudFormation successfully deleted the change set.
func (c *Client) DeleteChangeSet(ctx context.Context, params *DeleteChangeSetInput, optFns ...func(*Options)) (*DeleteChangeSetOutput, error) {
	stack := middleware.NewStack("DeleteChangeSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteChangeSetMiddlewares(stack)
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
	addOpDeleteChangeSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteChangeSet(options.Region), middleware.Before)
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
			OperationName: "DeleteChangeSet",
			Err:           err,
		}
	}
	out := result.(*DeleteChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DeleteChangeSet () action.
type DeleteChangeSetInput struct {

	// The name or Amazon Resource Name (ARN) of the change set that you want to
	// delete.
	//
	// This member is required.
	ChangeSetName *string

	// If you specified the name of a change set to delete, specify the stack name or
	// ID (ARN) that is associated with it.
	StackName *string
}

// The output for the DeleteChangeSet () action.
type DeleteChangeSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteChangeSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteChangeSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteChangeSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteChangeSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DeleteChangeSet",
	}
}
