// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// ServiceSetting is an account-level setting for an AWS service. This setting
// defines how a user interacts with or uses a service or a feature of a service.
// For example, if an AWS service charges money to the account based on feature or
// service usage, then the AWS service team might create a default setting of
// "false". This means the user can't use this feature unless they change the
// setting to "true" and intentionally opt in for a paid feature. Services map a
// SettingId object to a setting value. AWS services teams define the default value
// for a SettingId. You can't create a new SettingId, but you can overwrite the
// default value if you have the ssm:UpdateServiceSetting permission for the
// setting. Use the GetServiceSetting () API action to view the current value. Or,
// use the ResetServiceSetting () to change the value back to the original value
// defined by the AWS service team.  <p>Update the service setting for the account.
// </p>
func (c *Client) UpdateServiceSetting(ctx context.Context, params *UpdateServiceSettingInput, optFns ...func(*Options)) (*UpdateServiceSettingOutput, error) {
	stack := middleware.NewStack("UpdateServiceSetting", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateServiceSettingMiddlewares(stack)
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
	addOpUpdateServiceSettingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServiceSetting(options.Region), middleware.Before)
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
			OperationName: "UpdateServiceSetting",
			Err:           err,
		}
	}
	out := result.(*UpdateServiceSettingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body of the UpdateServiceSetting API action.
type UpdateServiceSettingInput struct {

	// The Amazon Resource Name (ARN) of the service setting to reset. For example,
	// arn:aws:ssm:us-east-1:111122223333:servicesetting/ssm/parameter-store/high-throughput-enabled.
	// The setting ID can be one of the following.
	//
	//     *
	// /ssm/parameter-store/default-parameter-tier
	//
	//     *
	// /ssm/parameter-store/high-throughput-enabled
	//
	//     *
	// /ssm/managed-instance/activation-tier
	//
	// This member is required.
	SettingId *string

	// The new value to specify for the service setting. For the
	// /ssm/parameter-store/default-parameter-tier setting ID, the setting value can be
	// one of the following.
	//
	//     * Standard
	//
	//     * Advanced
	//
	//     *
	// Intelligent-Tiering
	//
	// For the /ssm/parameter-store/high-throughput-enabled, and
	// /ssm/managed-instance/activation-tier setting IDs, the setting value can be true
	// or false.
	//
	// This member is required.
	SettingValue *string
}

// The result body of the UpdateServiceSetting API action.
type UpdateServiceSettingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateServiceSettingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateServiceSetting{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateServiceSetting{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateServiceSetting(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdateServiceSetting",
	}
}
