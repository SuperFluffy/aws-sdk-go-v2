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

// Describes the specified conversion tasks or all your conversion tasks. For more
// information, see the VM Import/Export User Guide
// (https://docs.aws.amazon.com/vm-import/latest/userguide/). For information about
// the import manifest referenced by this API action, see VM Import Manifest
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/manifest.html).
func (c *Client) DescribeConversionTasks(ctx context.Context, params *DescribeConversionTasksInput, optFns ...func(*Options)) (*DescribeConversionTasksOutput, error) {
	stack := middleware.NewStack("DescribeConversionTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeConversionTasksMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConversionTasks(options.Region), middleware.Before)
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
			OperationName: "DescribeConversionTasks",
			Err:           err,
		}
	}
	out := result.(*DescribeConversionTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConversionTasksInput struct {

	// The conversion task IDs.
	ConversionTaskIds []*string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DescribeConversionTasksOutput struct {

	// Information about the conversion tasks.
	ConversionTasks []*types.ConversionTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeConversionTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeConversionTasks{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeConversionTasks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeConversionTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeConversionTasks",
	}
}
