// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of aliases for a specified Amazon Lex bot. This operation
// requires permissions for the lex:GetBotAliases action.
func (c *Client) GetBotAliases(ctx context.Context, params *GetBotAliasesInput, optFns ...func(*Options)) (*GetBotAliasesOutput, error) {
	stack := middleware.NewStack("GetBotAliases", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetBotAliasesMiddlewares(stack)
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
	addOpGetBotAliasesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBotAliases(options.Region), middleware.Before)
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
			OperationName: "GetBotAliases",
			Err:           err,
		}
	}
	out := result.(*GetBotAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotAliasesInput struct {

	// The name of the bot.
	//
	// This member is required.
	BotName *string

	// The maximum number of aliases to return in the response. The default is 50. .
	MaxResults *int32

	// Substring to match in bot alias names. An alias will be returned if any part of
	// its name matches the substring. For example, "xyz" matches both "xyzabc" and
	// "abcxyz."
	NameContains *string

	// A pagination token for fetching the next page of aliases. If the response to
	// this call is truncated, Amazon Lex returns a pagination token in the response.
	// To fetch the next page of aliases, specify the pagination token in the next
	// request.
	NextToken *string
}

type GetBotAliasesOutput struct {

	// An array of BotAliasMetadata objects, each describing a bot alias.
	BotAliases []*types.BotAliasMetadata

	// A pagination token for fetching next page of aliases. If the response to this
	// call is truncated, Amazon Lex returns a pagination token in the response. To
	// fetch the next page of aliases, specify the pagination token in the next
	// request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetBotAliasesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetBotAliases{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBotAliases{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBotAliases(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBotAliases",
	}
}
