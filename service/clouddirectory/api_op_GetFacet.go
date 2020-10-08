// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets details of the Facet (), such as facet name, attributes, Rule ()s, or
// ObjectType. You can call this on all kinds of schema facets -- published,
// development, or applied.
func (c *Client) GetFacet(ctx context.Context, params *GetFacetInput, optFns ...func(*Options)) (*GetFacetOutput, error) {
	stack := middleware.NewStack("GetFacet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetFacetMiddlewares(stack)
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
	addOpGetFacetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFacet(options.Region), middleware.Before)
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
			OperationName: "GetFacet",
			Err:           err,
		}
	}
	out := result.(*GetFacetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFacetInput struct {

	// The name of the facet to retrieve.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) that is associated with the Facet (). For more
	// information, see arns ().
	//
	// This member is required.
	SchemaArn *string
}

type GetFacetOutput struct {

	// The Facet () structure that is associated with the facet.
	Facet *types.Facet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetFacetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetFacet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFacet{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFacet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "GetFacet",
	}
}
