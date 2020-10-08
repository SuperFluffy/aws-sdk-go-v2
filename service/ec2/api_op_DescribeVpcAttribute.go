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

// Describes the specified attribute of the specified VPC. You can specify only one
// attribute at a time.
func (c *Client) DescribeVpcAttribute(ctx context.Context, params *DescribeVpcAttributeInput, optFns ...func(*Options)) (*DescribeVpcAttributeOutput, error) {
	stack := middleware.NewStack("DescribeVpcAttribute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeVpcAttributeMiddlewares(stack)
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
	addOpDescribeVpcAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcAttribute(options.Region), middleware.Before)
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
			OperationName: "DescribeVpcAttribute",
			Err:           err,
		}
	}
	out := result.(*DescribeVpcAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcAttributeInput struct {

	// The VPC attribute.
	//
	// This member is required.
	Attribute types.VpcAttributeName

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DescribeVpcAttributeOutput struct {

	// Indicates whether the instances launched in the VPC get DNS hostnames. If this
	// attribute is true, instances in the VPC get DNS hostnames; otherwise, they do
	// not.
	EnableDnsHostnames *types.AttributeBooleanValue

	// Indicates whether DNS resolution is enabled for the VPC. If this attribute is
	// true, the Amazon DNS server resolves DNS hostnames for your instances to their
	// corresponding IP addresses; otherwise, it does not.
	EnableDnsSupport *types.AttributeBooleanValue

	// The ID of the VPC.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeVpcAttributeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcAttribute{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcAttribute{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVpcAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcAttribute",
	}
}
