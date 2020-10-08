// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the programs that currently exist for a specific multiplex.
func (c *Client) ListMultiplexPrograms(ctx context.Context, params *ListMultiplexProgramsInput, optFns ...func(*Options)) (*ListMultiplexProgramsOutput, error) {
	stack := middleware.NewStack("ListMultiplexPrograms", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListMultiplexProgramsMiddlewares(stack)
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
	addOpListMultiplexProgramsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMultiplexPrograms(options.Region), middleware.Before)
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
			OperationName: "ListMultiplexPrograms",
			Err:           err,
		}
	}
	out := result.(*ListMultiplexProgramsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListMultiplexProgramsRequest
type ListMultiplexProgramsInput struct {

	// The ID of the multiplex that the programs belong to.
	//
	// This member is required.
	MultiplexId *string

	// The maximum number of items to return.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string
}

// Placeholder documentation for ListMultiplexProgramsResponse
type ListMultiplexProgramsOutput struct {

	// List of multiplex programs.
	MultiplexPrograms []*types.MultiplexProgramSummary

	// Token for the next ListMultiplexProgram request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListMultiplexProgramsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListMultiplexPrograms{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListMultiplexPrograms{}, middleware.After)
}

func newServiceMetadataMiddleware_opListMultiplexPrograms(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "ListMultiplexPrograms",
	}
}
