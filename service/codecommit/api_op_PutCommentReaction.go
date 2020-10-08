// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or updates a reaction to a specified comment for the user whose identity is
// used to make the request. You can only add or update a reaction for yourself.
// You cannot add, modify, or delete a reaction for another user.
func (c *Client) PutCommentReaction(ctx context.Context, params *PutCommentReactionInput, optFns ...func(*Options)) (*PutCommentReactionOutput, error) {
	stack := middleware.NewStack("PutCommentReaction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutCommentReactionMiddlewares(stack)
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
	addOpPutCommentReactionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutCommentReaction(options.Region), middleware.Before)
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
			OperationName: "PutCommentReaction",
			Err:           err,
		}
	}
	out := result.(*PutCommentReactionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutCommentReactionInput struct {

	// The ID of the comment to which you want to add or update a reaction.
	//
	// This member is required.
	CommentId *string

	// The emoji reaction you want to add or update. To remove a reaction, provide a
	// value of blank or null. You can also provide the value of none. For information
	// about emoji reaction values supported in AWS CodeCommit, see the AWS CodeCommit
	// User Guide
	// (https://docs.aws.amazon.com/codecommit/latest/userguide/how-to-commit-comment.html#emoji-reaction-table).
	//
	// This member is required.
	ReactionValue *string
}

type PutCommentReactionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutCommentReactionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutCommentReaction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutCommentReaction{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutCommentReaction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "PutCommentReaction",
	}
}
