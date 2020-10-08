// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about a deployment configuration.
func (c *Client) GetDeploymentConfig(ctx context.Context, params *GetDeploymentConfigInput, optFns ...func(*Options)) (*GetDeploymentConfigOutput, error) {
	stack := middleware.NewStack("GetDeploymentConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDeploymentConfigMiddlewares(stack)
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
	addOpGetDeploymentConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeploymentConfig(options.Region), middleware.Before)
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
			OperationName: "GetDeploymentConfig",
			Err:           err,
		}
	}
	out := result.(*GetDeploymentConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetDeploymentConfig operation.
type GetDeploymentConfigInput struct {

	// The name of a deployment configuration associated with the IAM user or AWS
	// account.
	//
	// This member is required.
	DeploymentConfigName *string
}

// Represents the output of a GetDeploymentConfig operation.
type GetDeploymentConfigOutput struct {

	// Information about the deployment configuration.
	DeploymentConfigInfo *types.DeploymentConfigInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDeploymentConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDeploymentConfig{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDeploymentConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDeploymentConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "GetDeploymentConfig",
	}
}
