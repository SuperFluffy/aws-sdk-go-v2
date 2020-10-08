// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves core properties, including configuration, status, and metadata, for a
// fleet. To get attributes for one or more fleets, provide a list of fleet IDs or
// fleet ARNs. To get attributes for all fleets, do not specify a fleet identifier.
// When requesting attributes for multiple fleets, use the pagination parameters to
// retrieve results as a set of sequential pages. If successful, a FleetAttributes
// () object is returned for each fleet requested, unless the fleet identifier is
// not found. Some API actions may limit the number of fleet IDs allowed in one
// request. If a request exceeds this limit, the request fails and the error
// message includes the maximum allowed number. Learn more Setting up GameLift
// Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
//     * CreateFleet ()
//
//     * ListFleets ()
//
//     * DeleteFleet
// ()
//
//     * Describe fleets:
//
//         * DescribeFleetAttributes ()
//
//         *
// DescribeFleetCapacity ()
//
//         * DescribeFleetPortSettings ()
//
//         *
// DescribeFleetUtilization ()
//
//         * DescribeRuntimeConfiguration ()
//
//
// * DescribeEC2InstanceLimits ()
//
//         * DescribeFleetEvents ()
//
//     *
// UpdateFleetAttributes ()
//
//     * StartFleetActions () or StopFleetActions ()
func (c *Client) DescribeFleetAttributes(ctx context.Context, params *DescribeFleetAttributesInput, optFns ...func(*Options)) (*DescribeFleetAttributesOutput, error) {
	stack := middleware.NewStack("DescribeFleetAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeFleetAttributesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetAttributes(options.Region), middleware.Before)
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
			OperationName: "DescribeFleetAttributes",
			Err:           err,
		}
	}
	out := result.(*DescribeFleetAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DescribeFleetAttributesInput struct {

	// A list of unique fleet identifiers to retrieve attributes for. You can use
	// either the fleet ID or ARN value. To retrieve attributes for all current fleets,
	// do not include this parameter. If the list of fleet identifiers includes fleets
	// that don't currently exist, the request succeeds but no attributes for that
	// fleet are returned.
	FleetIds []*string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is ignored when the
	// request specifies one or a list of fleet IDs.
	Limit *int32

	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this action. To start at the
	// beginning of the result set, do not specify a value. This parameter is ignored
	// when the request specifies one or a list of fleet IDs.
	NextToken *string
}

// Represents the returned data in response to a request action.
type DescribeFleetAttributesOutput struct {

	// A collection of objects containing attribute metadata for each requested fleet
	// ID. Attribute objects are returned only for fleets that currently exist.
	FleetAttributes []*types.FleetAttributes

	// Token that indicates where to resume retrieving results on the next call to this
	// action. If no token is returned, these results represent the end of the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeFleetAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeFleetAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeFleetAttributes",
	}
}
