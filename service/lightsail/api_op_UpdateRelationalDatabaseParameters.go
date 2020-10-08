// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows the update of one or more parameters of a database in Amazon Lightsail.
// Parameter updates don't cause outages; therefore, their application is not
// subject to the preferred maintenance window. However, there are two ways in
// which parameter updates are applied: dynamic or pending-reboot. Parameters
// marked with a dynamic apply type are applied immediately. Parameters marked with
// a pending-reboot apply type are applied only after the database is rebooted
// using the reboot relational database operation. The update relational database
// parameters operation supports tag-based access control via resource tags applied
// to the resource identified by relationalDatabaseName. For more information, see
// the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) UpdateRelationalDatabaseParameters(ctx context.Context, params *UpdateRelationalDatabaseParametersInput, optFns ...func(*Options)) (*UpdateRelationalDatabaseParametersOutput, error) {
	stack := middleware.NewStack("UpdateRelationalDatabaseParameters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateRelationalDatabaseParametersMiddlewares(stack)
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
	addOpUpdateRelationalDatabaseParametersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRelationalDatabaseParameters(options.Region), middleware.Before)
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
			OperationName: "UpdateRelationalDatabaseParameters",
			Err:           err,
		}
	}
	out := result.(*UpdateRelationalDatabaseParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRelationalDatabaseParametersInput struct {

	// The database parameters to update.
	//
	// This member is required.
	Parameters []*types.RelationalDatabaseParameter

	// The name of your database for which to update parameters.
	//
	// This member is required.
	RelationalDatabaseName *string
}

type UpdateRelationalDatabaseParametersOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateRelationalDatabaseParametersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRelationalDatabaseParameters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRelationalDatabaseParameters{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRelationalDatabaseParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "UpdateRelationalDatabaseParameters",
	}
}
