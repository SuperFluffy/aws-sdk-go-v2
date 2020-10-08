// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays the details of the dashboard that you specify. To copy an existing
// dashboard, use GetDashboard, and then use the data returned within DashboardBody
// as the template for the new dashboard when you call PutDashboard to create the
// copy.
func (c *Client) GetDashboard(ctx context.Context, params *GetDashboardInput, optFns ...func(*Options)) (*GetDashboardOutput, error) {
	stack := middleware.NewStack("GetDashboard", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetDashboardMiddlewares(stack)
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
	addOpGetDashboardValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDashboard(options.Region), middleware.Before)
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
			OperationName: "GetDashboard",
			Err:           err,
		}
	}
	out := result.(*GetDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDashboardInput struct {

	// The name of the dashboard to be described.
	//
	// This member is required.
	DashboardName *string
}

type GetDashboardOutput struct {

	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string

	// The detailed information about the dashboard, including what widgets are
	// included and their location on the dashboard. For more information about the
	// DashboardBody syntax, see Dashboard Body Structure and Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
	DashboardBody *string

	// The name of the dashboard.
	DashboardName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetDashboardMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetDashboard{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetDashboard{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDashboard(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "GetDashboard",
	}
}
