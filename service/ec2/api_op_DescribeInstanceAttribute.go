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

// Describes the specified attribute of the specified instance. You can specify
// only one attribute at a time. Valid attribute values are: instanceType | kernel
// | ramdisk | userData | disableApiTermination | instanceInitiatedShutdownBehavior
// | rootDeviceName | blockDeviceMapping | productCodes | sourceDestCheck |
// groupSet | ebsOptimized | sriovNetSupport
func (c *Client) DescribeInstanceAttribute(ctx context.Context, params *DescribeInstanceAttributeInput, optFns ...func(*Options)) (*DescribeInstanceAttributeOutput, error) {
	stack := middleware.NewStack("DescribeInstanceAttribute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeInstanceAttributeMiddlewares(stack)
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
	addOpDescribeInstanceAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceAttribute(options.Region), middleware.Before)
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
			OperationName: "DescribeInstanceAttribute",
			Err:           err,
		}
	}
	out := result.(*DescribeInstanceAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceAttributeInput struct {

	// The instance attribute. Note: The enaSupport attribute is not supported at this
	// time.
	//
	// This member is required.
	Attribute types.InstanceAttributeName

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

// Describes an instance attribute.
type DescribeInstanceAttributeOutput struct {

	// The block device mapping of the instance.
	BlockDeviceMappings []*types.InstanceBlockDeviceMapping

	// If the value is true, you can't terminate the instance through the Amazon EC2
	// console, CLI, or API; otherwise, you can.
	DisableApiTermination *types.AttributeBooleanValue

	// Indicates whether the instance is optimized for Amazon EBS I/O.
	EbsOptimized *types.AttributeBooleanValue

	// Indicates whether enhanced networking with ENA is enabled.
	EnaSupport *types.AttributeBooleanValue

	// The security groups associated with the instance.
	Groups []*types.GroupIdentifier

	// The ID of the instance.
	InstanceId *string

	// Indicates whether an instance stops or terminates when you initiate shutdown
	// from the instance (using the operating system command for system shutdown).
	InstanceInitiatedShutdownBehavior *types.AttributeValue

	// The instance type.
	InstanceType *types.AttributeValue

	// The kernel ID.
	KernelId *types.AttributeValue

	// A list of product codes.
	ProductCodes []*types.ProductCode

	// The RAM disk ID.
	RamdiskId *types.AttributeValue

	// The device name of the root device volume (for example, /dev/sda1).
	RootDeviceName *types.AttributeValue

	// Indicates whether source/destination checking is enabled. A value of true means
	// that checking is enabled, and false means that checking is disabled. This value
	// must be false for a NAT instance to perform NAT.
	SourceDestCheck *types.AttributeBooleanValue

	// Indicates whether enhanced networking with the Intel 82599 Virtual Function
	// interface is enabled.
	SriovNetSupport *types.AttributeValue

	// The user data.
	UserData *types.AttributeValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeInstanceAttributeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstanceAttribute{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstanceAttribute{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInstanceAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeInstanceAttribute",
	}
}
