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

// Describes the ClassicLink status of one or more VPCs.
func (c *Client) DescribeVpcClassicLink(ctx context.Context, params *DescribeVpcClassicLinkInput, optFns ...func(*Options)) (*DescribeVpcClassicLinkOutput, error) {
	stack := middleware.NewStack("DescribeVpcClassicLink", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeVpcClassicLinkMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcClassicLink(options.Region), middleware.Before)
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
			OperationName: "DescribeVpcClassicLink",
			Err:           err,
		}
	}
	out := result.(*DescribeVpcClassicLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcClassicLinkInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * is-classic-link-enabled - Whether the VPC is enabled
	// for ClassicLink (true | false).
	//
	//     * tag: - The key/value combination of a tag
	// assigned to the resource. Use the tag key in the filter name and the tag value
	// as the filter value. For example, to find all resources that have a tag with the
	// key Owner and the value TeamA, specify tag:Owner for the filter name and TeamA
	// for the filter value.
	//
	//     * tag-key - The key of a tag assigned to the
	// resource. Use this filter to find all resources assigned a tag with a specific
	// key, regardless of the tag value.
	Filters []*types.Filter

	// One or more VPCs for which you want to describe the ClassicLink status.
	VpcIds []*string
}

type DescribeVpcClassicLinkOutput struct {

	// The ClassicLink status of one or more VPCs.
	Vpcs []*types.VpcClassicLink

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeVpcClassicLinkMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcClassicLink{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcClassicLink{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVpcClassicLink(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcClassicLink",
	}
}
