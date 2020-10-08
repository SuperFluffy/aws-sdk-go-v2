// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detaches a specified Elastic Load Balancing instance from its layer. Required
// Permissions: To use this action, an IAM user must have a Manage permissions
// level for the stack, or an attached policy that explicitly grants permissions.
// For more information on user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DetachElasticLoadBalancer(ctx context.Context, params *DetachElasticLoadBalancerInput, optFns ...func(*Options)) (*DetachElasticLoadBalancerOutput, error) {
	stack := middleware.NewStack("DetachElasticLoadBalancer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDetachElasticLoadBalancerMiddlewares(stack)
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
	addOpDetachElasticLoadBalancerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachElasticLoadBalancer(options.Region), middleware.Before)
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
			OperationName: "DetachElasticLoadBalancer",
			Err:           err,
		}
	}
	out := result.(*DetachElasticLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachElasticLoadBalancerInput struct {

	// The Elastic Load Balancing instance's name.
	//
	// This member is required.
	ElasticLoadBalancerName *string

	// The ID of the layer that the Elastic Load Balancing instance is attached to.
	//
	// This member is required.
	LayerId *string
}

type DetachElasticLoadBalancerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDetachElasticLoadBalancerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDetachElasticLoadBalancer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetachElasticLoadBalancer{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetachElasticLoadBalancer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DetachElasticLoadBalancer",
	}
}
