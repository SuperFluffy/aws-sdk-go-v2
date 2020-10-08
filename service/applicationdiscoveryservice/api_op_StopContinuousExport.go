// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Stop the continuous flow of agent's discovered data into Amazon Athena.
func (c *Client) StopContinuousExport(ctx context.Context, params *StopContinuousExportInput, optFns ...func(*Options)) (*StopContinuousExportOutput, error) {
	stack := middleware.NewStack("StopContinuousExport", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopContinuousExportMiddlewares(stack)
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
	addOpStopContinuousExportValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopContinuousExport(options.Region), middleware.Before)
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
			OperationName: "StopContinuousExport",
			Err:           err,
		}
	}
	out := result.(*StopContinuousExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopContinuousExportInput struct {

	// The unique ID assigned to this export.
	//
	// This member is required.
	ExportId *string
}

type StopContinuousExportOutput struct {

	// Timestamp that represents when this continuous export started collecting data.
	StartTime *time.Time

	// Timestamp that represents when this continuous export was stopped.
	StopTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopContinuousExportMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopContinuousExport{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopContinuousExport{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopContinuousExport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "StopContinuousExport",
	}
}
