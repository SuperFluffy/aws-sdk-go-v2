// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a log pattern to a LogPatternSet.
func (c *Client) UpdateLogPattern(ctx context.Context, params *UpdateLogPatternInput, optFns ...func(*Options)) (*UpdateLogPatternOutput, error) {
	stack := middleware.NewStack("UpdateLogPattern", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateLogPatternMiddlewares(stack)
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
	addOpUpdateLogPatternValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLogPattern(options.Region), middleware.Before)
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
			OperationName: "UpdateLogPattern",
			Err:           err,
		}
	}
	out := result.(*UpdateLogPatternOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLogPatternInput struct {

	// The name of the log pattern.
	//
	// This member is required.
	PatternName *string

	// The name of the log pattern set.
	//
	// This member is required.
	PatternSetName *string

	// The name of the resource group.
	//
	// This member is required.
	ResourceGroupName *string

	// The log pattern.
	Pattern *string

	// Rank of the log pattern.
	Rank *int32
}

type UpdateLogPatternOutput struct {

	// The successfully created log pattern.
	LogPattern *types.LogPattern

	// The name of the resource group.
	ResourceGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateLogPatternMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLogPattern{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLogPattern{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateLogPattern(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "UpdateLogPattern",
	}
}
