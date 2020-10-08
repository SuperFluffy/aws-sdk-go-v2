// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a CloudWatch Logs Insights query that is in progress. If the query has
// already ended, the operation returns an error indicating that the specified
// query is not running.
func (c *Client) StopQuery(ctx context.Context, params *StopQueryInput, optFns ...func(*Options)) (*StopQueryOutput, error) {
	stack := middleware.NewStack("StopQuery", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopQueryMiddlewares(stack)
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
	addOpStopQueryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopQuery(options.Region), middleware.Before)
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
			OperationName: "StopQuery",
			Err:           err,
		}
	}
	out := result.(*StopQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopQueryInput struct {

	// The ID number of the query to stop. If necessary, you can use DescribeQueries to
	// find this ID number.
	//
	// This member is required.
	QueryId *string
}

type StopQueryOutput struct {

	// This is true if the query was stopped by the StopQuery operation.
	Success *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopQueryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopQuery{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopQuery{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopQuery(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "StopQuery",
	}
}
