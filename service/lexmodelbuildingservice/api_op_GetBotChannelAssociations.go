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

// Returns a list of all of the channels associated with the specified bot. The
// GetBotChannelAssociations operation requires permissions for the
// lex:GetBotChannelAssociations action.
func (c *Client) GetBotChannelAssociations(ctx context.Context, params *GetBotChannelAssociationsInput, optFns ...func(*Options)) (*GetBotChannelAssociationsOutput, error) {
	stack := middleware.NewStack("GetBotChannelAssociations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetBotChannelAssociationsMiddlewares(stack)
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
	addOpGetBotChannelAssociationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBotChannelAssociations(options.Region), middleware.Before)
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
			OperationName: "GetBotChannelAssociations",
			Err:           err,
		}
	}
	out := result.(*GetBotChannelAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotChannelAssociationsInput struct {

	// An alias pointing to the specific version of the Amazon Lex bot to which this
	// association is being made.
	//
	// This member is required.
	BotAlias *string

	// The name of the Amazon Lex bot in the association.
	//
	// This member is required.
	BotName *string

	// The maximum number of associations to return in the response. The default is 50.
	MaxResults *int32

	// Substring to match in channel association names. An association will be returned
	// if any part of its name matches the substring. For example, "xyz" matches both
	// "xyzabc" and "abcxyz." To return all bot channel associations, use a hyphen
	// ("-") as the nameContains parameter.
	NameContains *string

	// A pagination token for fetching the next page of associations. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of associations, specify the pagination token
	// in the next request.
	NextToken *string
}

type GetBotChannelAssociationsOutput struct {

	// An array of objects, one for each association, that provides information about
	// the Amazon Lex bot and its association with the channel.
	BotChannelAssociations []*types.BotChannelAssociation

	// A pagination token that fetches the next page of associations. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of associations, specify the pagination token
	// in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetBotChannelAssociationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetBotChannelAssociations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBotChannelAssociations{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBotChannelAssociations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBotChannelAssociations",
	}
}
