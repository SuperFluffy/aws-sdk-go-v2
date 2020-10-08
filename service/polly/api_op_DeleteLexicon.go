// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified pronunciation lexicon stored in an AWS Region. A lexicon
// which has been deleted is not available for speech synthesis, nor is it possible
// to retrieve it using either the GetLexicon or ListLexicon APIs. For more
// information, see Managing Lexicons
// (https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
func (c *Client) DeleteLexicon(ctx context.Context, params *DeleteLexiconInput, optFns ...func(*Options)) (*DeleteLexiconOutput, error) {
	stack := middleware.NewStack("DeleteLexicon", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteLexiconMiddlewares(stack)
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
	addOpDeleteLexiconValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLexicon(options.Region), middleware.Before)
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
			OperationName: "DeleteLexicon",
			Err:           err,
		}
	}
	out := result.(*DeleteLexiconOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLexiconInput struct {

	// The name of the lexicon to delete. Must be an existing lexicon in the region.
	//
	// This member is required.
	Name *string
}

type DeleteLexiconOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteLexiconMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteLexicon{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteLexicon{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteLexicon(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "DeleteLexicon",
	}
}
