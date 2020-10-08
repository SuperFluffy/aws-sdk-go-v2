// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches one or more target groups to the specified Auto Scaling group. To
// describe the target groups for an Auto Scaling group, call the
// DescribeLoadBalancerTargetGroups () API. To detach the target group from the
// Auto Scaling group, call the DetachLoadBalancerTargetGroups () API. With
// Application Load Balancers and Network Load Balancers, instances are registered
// as targets with a target group. With Classic Load Balancers, instances are
// registered with the load balancer. For more information, see Attaching a Load
// Balancer to Your Auto Scaling Group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/attach-load-balancer-asg.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) AttachLoadBalancerTargetGroups(ctx context.Context, params *AttachLoadBalancerTargetGroupsInput, optFns ...func(*Options)) (*AttachLoadBalancerTargetGroupsOutput, error) {
	stack := middleware.NewStack("AttachLoadBalancerTargetGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpAttachLoadBalancerTargetGroupsMiddlewares(stack)
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
	addOpAttachLoadBalancerTargetGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachLoadBalancerTargetGroups(options.Region), middleware.Before)
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
			OperationName: "AttachLoadBalancerTargetGroups",
			Err:           err,
		}
	}
	out := result.(*AttachLoadBalancerTargetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachLoadBalancerTargetGroupsInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The Amazon Resource Names (ARN) of the target groups. You can specify up to 10
	// target groups.
	//
	// This member is required.
	TargetGroupARNs []*string
}

type AttachLoadBalancerTargetGroupsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpAttachLoadBalancerTargetGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpAttachLoadBalancerTargetGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpAttachLoadBalancerTargetGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachLoadBalancerTargetGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "AttachLoadBalancerTargetGroups",
	}
}
