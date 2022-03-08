// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Amazon Kendra index.
func (c *Client) UpdateIndex(ctx context.Context, params *UpdateIndexInput, optFns ...func(*Options)) (*UpdateIndexOutput, error) {
	if params == nil {
		params = &UpdateIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIndex", params, optFns, c.addOperationUpdateIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIndexInput struct {

	// The identifier of the index to update.
	//
	// This member is required.
	Id *string

	// Sets the number of additional storage and query capacity units that should be
	// used by the index. You can change the capacity of the index up to 5 times per
	// day. If you are using extra storage units, you can't reduce the storage capacity
	// below that required to meet the storage needs for your index.
	CapacityUnits *types.CapacityUnitsConfiguration

	// A new description for the index.
	Description *string

	// The document metadata you want to update.
	DocumentMetadataConfigurationUpdates []types.DocumentMetadataConfiguration

	// The name of the index to update.
	Name *string

	// A new IAM role that gives Amazon Kendra permission to access your Amazon
	// CloudWatch logs.
	RoleArn *string

	// The user context policy.
	UserContextPolicy types.UserContextPolicy

	// Enables fetching access levels of groups and users from an Amazon Web Services
	// Single Sign On identity source. To configure this, see
	// UserGroupResolutionConfiguration
	// (https://docs.aws.amazon.com/kendra/latest/dg/API_UserGroupResolutionConfiguration.html).
	UserGroupResolutionConfiguration *types.UserGroupResolutionConfiguration

	// The user token configuration.
	UserTokenConfigurations []types.UserTokenConfiguration

	noSmithyDocumentSerde
}

type UpdateIndexOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateIndex{}, middleware.After)
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
	if err = addOpUpdateIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIndex(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "UpdateIndex",
	}
}
