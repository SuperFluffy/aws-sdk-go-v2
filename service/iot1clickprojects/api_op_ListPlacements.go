// Code generated by smithy-go-codegen DO NOT EDIT.

package iot1clickprojects

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the placement(s) of a project.
func (c *Client) ListPlacements(ctx context.Context, params *ListPlacementsInput, optFns ...func(*Options)) (*ListPlacementsOutput, error) {
	stack := middleware.NewStack("ListPlacements", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPlacementsMiddlewares(stack)
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
	addOpListPlacementsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPlacements(options.Region), middleware.Before)
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
			OperationName: "ListPlacements",
			Err:           err,
		}
	}
	out := result.(*ListPlacementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPlacementsInput struct {

	// The project containing the placements to be listed.
	//
	// This member is required.
	ProjectName *string

	// The maximum number of results to return per request. If not set, a default value
	// of 100 is used.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string
}

type ListPlacementsOutput struct {

	// An object listing the requested placements.
	//
	// This member is required.
	Placements []*types.PlacementSummary

	// The token used to retrieve the next set of results - will be effectively empty
	// if there are no further results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPlacementsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPlacements{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPlacements{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPlacements(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot1click",
		OperationName: "ListPlacements",
	}
}
