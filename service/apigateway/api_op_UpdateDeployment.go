// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Changes information about a Deployment () resource.
func (c *Client) UpdateDeployment(ctx context.Context, params *UpdateDeploymentInput, optFns ...func(*Options)) (*UpdateDeploymentOutput, error) {
	stack := middleware.NewStack("UpdateDeployment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDeploymentMiddlewares(stack)
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
	addOpUpdateDeploymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDeployment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "UpdateDeployment",
			Err:           err,
		}
	}
	out := result.(*UpdateDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to change information about a Deployment () resource.
type UpdateDeploymentInput struct {

	// The replacement identifier for the Deployment () resource to change information
	// about.
	//
	// This member is required.
	DeploymentId *string

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	Name *string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// An immutable representation of a RestApi () resource that can be called by users
// using Stages (). A deployment must be associated with a Stage () for it to be
// callable over the Internet. To create a deployment, call POST on the Deployments
// () resource of a RestApi (). To view, update, or delete a deployment, call GET,
// PATCH, or DELETE on the specified deployment resource
// (/restapis/{restapi_id}/deployments/{deployment_id}). RestApi (), Deployments
// (), Stage (), AWS CLI
// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-deployment.html),
// AWS SDKs (https://aws.amazon.com/tools/)
type UpdateDeploymentOutput struct {

	// A summary of the RestApi () at the date and time that the deployment resource
	// was created.
	ApiSummary map[string]map[string]*types.MethodSnapshot

	// The date and time that the deployment resource was created.
	CreatedDate *time.Time

	// The description for the deployment resource.
	Description *string

	// The identifier for the deployment resource.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDeploymentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDeployment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDeployment{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDeployment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateDeployment",
	}
}
