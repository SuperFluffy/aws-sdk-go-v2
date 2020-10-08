// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about Aurora global database clusters. This API supports
// pagination. For more information on Amazon Aurora, see  What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
func (c *Client) DescribeGlobalClusters(ctx context.Context, params *DescribeGlobalClustersInput, optFns ...func(*Options)) (*DescribeGlobalClustersOutput, error) {
	stack := middleware.NewStack("DescribeGlobalClusters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeGlobalClustersMiddlewares(stack)
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
	addOpDescribeGlobalClustersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGlobalClusters(options.Region), middleware.Before)
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
			OperationName: "DescribeGlobalClusters",
			Err:           err,
		}
	}
	out := result.(*DescribeGlobalClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGlobalClustersInput struct {

	// A filter that specifies one or more global DB clusters to describe. Supported
	// filters:
	//
	//     * db-cluster-id - Accepts DB cluster identifiers and DB cluster
	// Amazon Resource Names (ARNs). The results list will only include information
	// about the DB clusters identified by these ARNs.
	Filters []*types.Filter

	// The user-supplied DB cluster identifier. If this parameter is specified,
	// information from only the specific DB cluster is returned. This parameter isn't
	// case-sensitive. Constraints:
	//
	//     * If supplied, must match an existing
	// DBClusterIdentifier.
	GlobalClusterIdentifier *string

	// An optional pagination token provided by a previous DescribeGlobalClusters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

type DescribeGlobalClustersOutput struct {

	// The list of global clusters returned by this request.
	GlobalClusters []*types.GlobalCluster

	// An optional pagination token provided by a previous DescribeGlobalClusters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeGlobalClustersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeGlobalClusters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeGlobalClusters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeGlobalClusters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeGlobalClusters",
	}
}
