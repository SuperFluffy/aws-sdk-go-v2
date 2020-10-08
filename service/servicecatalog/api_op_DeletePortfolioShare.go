// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops sharing the specified portfolio with the specified account or organization
// node. Shares to an organization node can only be deleted by the master account
// of an organization or by a delegated administrator. Note that if a delegated
// admin is de-registered, portfolio shares created from that account are removed.
func (c *Client) DeletePortfolioShare(ctx context.Context, params *DeletePortfolioShareInput, optFns ...func(*Options)) (*DeletePortfolioShareOutput, error) {
	stack := middleware.NewStack("DeletePortfolioShare", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeletePortfolioShareMiddlewares(stack)
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
	addOpDeletePortfolioShareValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePortfolioShare(options.Region), middleware.Before)
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
			OperationName: "DeletePortfolioShare",
			Err:           err,
		}
	}
	out := result.(*DeletePortfolioShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePortfolioShareInput struct {

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string

	// The AWS account ID.
	AccountId *string

	// The organization node to whom you are going to stop sharing.
	OrganizationNode *types.OrganizationNode
}

type DeletePortfolioShareOutput struct {

	// The portfolio share unique identifier. This will only be returned if delete is
	// made to an organization node.
	PortfolioShareToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeletePortfolioShareMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePortfolioShare{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePortfolioShare{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePortfolioShare(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DeletePortfolioShare",
	}
}
