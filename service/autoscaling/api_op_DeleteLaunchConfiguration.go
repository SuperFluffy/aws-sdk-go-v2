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

// Deletes the specified launch configuration. The launch configuration must not be
// attached to an Auto Scaling group. When this call completes, the launch
// configuration is no longer available for use.
func (c *Client) DeleteLaunchConfiguration(ctx context.Context, params *DeleteLaunchConfigurationInput, optFns ...func(*Options)) (*DeleteLaunchConfigurationOutput, error) {
	stack := middleware.NewStack("DeleteLaunchConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteLaunchConfigurationMiddlewares(stack)
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
	addOpDeleteLaunchConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLaunchConfiguration(options.Region), middleware.Before)
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
			OperationName: "DeleteLaunchConfiguration",
			Err:           err,
		}
	}
	out := result.(*DeleteLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLaunchConfigurationInput struct {

	// The name of the launch configuration.
	//
	// This member is required.
	LaunchConfigurationName *string
}

type DeleteLaunchConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteLaunchConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteLaunchConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteLaunchConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteLaunchConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DeleteLaunchConfiguration",
	}
}
