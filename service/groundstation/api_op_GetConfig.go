// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns Config information. Only one Config response can be returned.
func (c *Client) GetConfig(ctx context.Context, params *GetConfigInput, optFns ...func(*Options)) (*GetConfigOutput, error) {
	stack := middleware.NewStack("GetConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetConfigMiddlewares(stack)
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
	addOpGetConfigValidationMiddleware(stack)
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
			OperationName: "GetConfig",
			Err:           err,
		}
	}
	out := result.(*GetConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type GetConfigInput struct {

	// UUID of a Config.
	//
	// This member is required.
	ConfigId *string

	// Type of a Config.
	//
	// This member is required.
	ConfigType types.ConfigCapabilityType
}

//
type GetConfigOutput struct {

	// ARN of a Config
	//
	// This member is required.
	ConfigArn *string

	// Data elements in a Config.
	//
	// This member is required.
	ConfigData types.ConfigTypeData

	// UUID of a Config.
	//
	// This member is required.
	ConfigId *string

	// Name of a Config.
	//
	// This member is required.
	Name *string

	// Type of a Config.
	ConfigType types.ConfigCapabilityType

	// Tags assigned to a Config.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfig{}, middleware.After)
}
