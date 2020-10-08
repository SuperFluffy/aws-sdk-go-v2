// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes all packages available to Amazon ES. Includes options for filtering,
// limiting the number of results, and pagination.
func (c *Client) DescribePackages(ctx context.Context, params *DescribePackagesInput, optFns ...func(*Options)) (*DescribePackagesOutput, error) {
	stack := middleware.NewStack("DescribePackages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribePackagesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePackages(options.Region), middleware.Before)
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
			OperationName: "DescribePackages",
			Err:           err,
		}
	}
	out := result.(*DescribePackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to DescribePackage () operation.
type DescribePackagesInput struct {

	// Only returns packages that match the DescribePackagesFilterList values.
	Filters []*types.DescribePackagesFilter

	// Limits results to a maximum number of packages.
	MaxResults *int32

	// Used for pagination. Only necessary if a previous API call includes a non-null
	// NextToken value. If provided, returns results for the next page.
	NextToken *string
}

// Container for response returned by DescribePackages () operation.
type DescribePackagesOutput struct {
	NextToken *string

	// List of PackageDetails objects.
	PackageDetailsList []*types.PackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribePackagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribePackages{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePackages{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePackages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribePackages",
	}
}
