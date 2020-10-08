// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Exports routes from the specified transit gateway route table to the specified
// S3 bucket. By default, all routes are exported. Alternatively, you can filter by
// CIDR range. The routes are saved to the specified bucket in a JSON file. For
// more information, see Export Route Tables to Amazon S3
// (https://docs.aws.amazon.com/vpc/latest/tgw/tgw-route-tables.html#tgw-export-route-tables)
// in Transit Gateways.
func (c *Client) ExportTransitGatewayRoutes(ctx context.Context, params *ExportTransitGatewayRoutesInput, optFns ...func(*Options)) (*ExportTransitGatewayRoutesOutput, error) {
	stack := middleware.NewStack("ExportTransitGatewayRoutes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpExportTransitGatewayRoutesMiddlewares(stack)
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
	addOpExportTransitGatewayRoutesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExportTransitGatewayRoutes(options.Region), middleware.Before)
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
			OperationName: "ExportTransitGatewayRoutes",
			Err:           err,
		}
	}
	out := result.(*ExportTransitGatewayRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportTransitGatewayRoutesInput struct {

	// The name of the S3 bucket.
	//
	// This member is required.
	S3Bucket *string

	// The ID of the route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	//     *
	// attachment.transit-gateway-attachment-id - The id of the transit gateway
	// attachment.
	//
	//     * attachment.resource-id - The resource id of the transit
	// gateway attachment.
	//
	//     * route-search.exact-match - The exact match of the
	// specified filter.
	//
	//     * route-search.longest-prefix-match - The longest prefix
	// that matches the route.
	//
	//     * route-search.subnet-of-match - The routes with a
	// subnet that match the specified CIDR filter.
	//
	//     *
	// route-search.supernet-of-match - The routes with a CIDR that encompass the CIDR
	// filter. For example, if you have 10.0.1.0/29 and 10.0.1.0/31 routes in your
	// route table and you specify supernet-of-match as 10.0.1.0/30, then the result
	// returns 10.0.1.0/29.
	//
	//     * state - The state of the attachment (available |
	// deleted | deleting | failed | modifying | pendingAcceptance | pending |
	// rollingBack | rejected | rejecting).
	//
	//     *
	// transit-gateway-route-destination-cidr-block - The CIDR range.
	//
	//     * type - The
	// type of route (active | blackhole).
	Filters []*types.Filter
}

type ExportTransitGatewayRoutesOutput struct {

	// The URL of the exported file in Amazon S3. For example,
	// s3://bucket_name/VPCTransitGateway/TransitGatewayRouteTables/file_name.
	S3Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpExportTransitGatewayRoutesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpExportTransitGatewayRoutes{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpExportTransitGatewayRoutes{}, middleware.After)
}

func newServiceMetadataMiddleware_opExportTransitGatewayRoutes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ExportTransitGatewayRoutes",
	}
}
