// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the schedule state of the specified crawler to NOT_SCHEDULED, but does not
// stop the crawler if it is already running.
func (c *Client) StopCrawlerSchedule(ctx context.Context, params *StopCrawlerScheduleInput, optFns ...func(*Options)) (*StopCrawlerScheduleOutput, error) {
	stack := middleware.NewStack("StopCrawlerSchedule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopCrawlerScheduleMiddlewares(stack)
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
	addOpStopCrawlerScheduleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopCrawlerSchedule(options.Region), middleware.Before)
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
			OperationName: "StopCrawlerSchedule",
			Err:           err,
		}
	}
	out := result.(*StopCrawlerScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopCrawlerScheduleInput struct {

	// Name of the crawler whose schedule state to set.
	//
	// This member is required.
	CrawlerName *string
}

type StopCrawlerScheduleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopCrawlerScheduleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopCrawlerSchedule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopCrawlerSchedule{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopCrawlerSchedule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StopCrawlerSchedule",
	}
}
