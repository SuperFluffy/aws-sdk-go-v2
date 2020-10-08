// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// DeleteReportGroup: Deletes a report group. Before you delete a report group, you
// must delete its reports. Use ListReportsForReportGroup
// (https://docs.aws.amazon.com/codebuild/latest/APIReference/API_ListReportsForReportGroup.html)
// to get the reports in a report group. Use DeleteReport
// (https://docs.aws.amazon.com/codebuild/latest/APIReference/API_DeleteReport.html)
// to delete the reports. If you call DeleteReportGroup for a report group that
// contains one or more reports, an exception is thrown.
func (c *Client) DeleteReportGroup(ctx context.Context, params *DeleteReportGroupInput, optFns ...func(*Options)) (*DeleteReportGroupOutput, error) {
	stack := middleware.NewStack("DeleteReportGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteReportGroupMiddlewares(stack)
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
	addOpDeleteReportGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReportGroup(options.Region), middleware.Before)
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
			OperationName: "DeleteReportGroup",
			Err:           err,
		}
	}
	out := result.(*DeleteReportGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteReportGroupInput struct {

	// The ARN of the report group to delete.
	//
	// This member is required.
	Arn *string
}

type DeleteReportGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteReportGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteReportGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteReportGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteReportGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "DeleteReportGroup",
	}
}
