// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all LaunchConfigurations available, filtered by Source Server IDs.
func (c *Client) GetLaunchConfiguration(ctx context.Context, params *GetLaunchConfigurationInput, optFns ...func(*Options)) (*GetLaunchConfigurationOutput, error) {
	if params == nil {
		params = &GetLaunchConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLaunchConfiguration", params, optFns, c.addOperationGetLaunchConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLaunchConfigurationInput struct {

	// Request to get Launch Configuration information by Source Server ID.
	//
	// This member is required.
	SourceServerID *string

	noSmithyDocumentSerde
}

type GetLaunchConfigurationOutput struct {

	// Launch configuration boot mode.
	BootMode types.BootMode

	// Copy Private IP during Launch Configuration.
	CopyPrivateIp *bool

	// Copy Tags during Launch Configuration.
	CopyTags *bool

	// Launch configuration EC2 Launch template ID.
	Ec2LaunchTemplateID *string

	// Launch disposition for launch configuration.
	LaunchDisposition types.LaunchDisposition

	// Launch configuration OS licensing.
	Licensing *types.Licensing

	// Launch configuration name.
	Name *string

	// Launch configuration Source Server ID.
	SourceServerID *string

	// Launch configuration Target instance type right sizing method.
	TargetInstanceTypeRightSizingMethod types.TargetInstanceTypeRightSizingMethod

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLaunchConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetLaunchConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLaunchConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetLaunchConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "GetLaunchConfiguration",
	}
}
