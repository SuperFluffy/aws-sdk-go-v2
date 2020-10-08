// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete a deployment strategy. Deleting a deployment strategy does not delete a
// configuration from a host.
func (c *Client) DeleteDeploymentStrategy(ctx context.Context, params *DeleteDeploymentStrategyInput, optFns ...func(*Options)) (*DeleteDeploymentStrategyOutput, error) {
	stack := middleware.NewStack("DeleteDeploymentStrategy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteDeploymentStrategyMiddlewares(stack)
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
	addOpDeleteDeploymentStrategyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDeploymentStrategy(options.Region), middleware.Before)
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
			OperationName: "DeleteDeploymentStrategy",
			Err:           err,
		}
	}
	out := result.(*DeleteDeploymentStrategyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDeploymentStrategyInput struct {

	// The ID of the deployment strategy you want to delete.
	//
	// This member is required.
	DeploymentStrategyId *string
}

type DeleteDeploymentStrategyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteDeploymentStrategyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDeploymentStrategy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDeploymentStrategy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDeploymentStrategy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "DeleteDeploymentStrategy",
	}
}
