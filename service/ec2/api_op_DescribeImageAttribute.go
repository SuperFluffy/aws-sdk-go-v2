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

// Describes the specified attribute of the specified AMI. You can specify only one
// attribute at a time.
func (c *Client) DescribeImageAttribute(ctx context.Context, params *DescribeImageAttributeInput, optFns ...func(*Options)) (*DescribeImageAttributeOutput, error) {
	stack := middleware.NewStack("DescribeImageAttribute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeImageAttributeMiddlewares(stack)
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
	addOpDescribeImageAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImageAttribute(options.Region), middleware.Before)
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
			OperationName: "DescribeImageAttribute",
			Err:           err,
		}
	}
	out := result.(*DescribeImageAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeImageAttribute.
type DescribeImageAttributeInput struct {

	// The AMI attribute. Note: Depending on your account privileges, the
	// blockDeviceMapping attribute may return a Client.AuthFailure error. If this
	// happens, use DescribeImages () to get information about the block device mapping
	// for the AMI.
	//
	// This member is required.
	Attribute types.ImageAttributeName

	// The ID of the AMI.
	//
	// This member is required.
	ImageId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

// Describes an image attribute.
type DescribeImageAttributeOutput struct {

	// The block device mapping entries.
	BlockDeviceMappings []*types.BlockDeviceMapping

	// A description for the AMI.
	Description *types.AttributeValue

	// The ID of the AMI.
	ImageId *string

	// The kernel ID.
	KernelId *types.AttributeValue

	// The launch permissions.
	LaunchPermissions []*types.LaunchPermission

	// The product codes.
	ProductCodes []*types.ProductCode

	// The RAM disk ID.
	RamdiskId *types.AttributeValue

	// Indicates whether enhanced networking with the Intel 82599 Virtual Function
	// interface is enabled.
	SriovNetSupport *types.AttributeValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeImageAttributeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeImageAttribute{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeImageAttribute{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeImageAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeImageAttribute",
	}
}
