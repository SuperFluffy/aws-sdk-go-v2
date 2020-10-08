// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Applies a pending maintenance action to a resource (for example, to a
// replication instance).
func (c *Client) ApplyPendingMaintenanceAction(ctx context.Context, params *ApplyPendingMaintenanceActionInput, optFns ...func(*Options)) (*ApplyPendingMaintenanceActionOutput, error) {
	stack := middleware.NewStack("ApplyPendingMaintenanceAction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpApplyPendingMaintenanceActionMiddlewares(stack)
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
	addOpApplyPendingMaintenanceActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opApplyPendingMaintenanceAction(options.Region), middleware.Before)
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
			OperationName: "ApplyPendingMaintenanceAction",
			Err:           err,
		}
	}
	out := result.(*ApplyPendingMaintenanceActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ApplyPendingMaintenanceActionInput struct {

	// The pending maintenance action to apply to this resource.
	//
	// This member is required.
	ApplyAction *string

	// A value that specifies the type of opt-in request, or undoes an opt-in request.
	// You can't undo an opt-in request of type immediate. Valid values:
	//
	//     *
	// immediate - Apply the maintenance action immediately.
	//
	//     * next-maintenance -
	// Apply the maintenance action during the next maintenance window for the
	// resource.
	//
	//     * undo-opt-in - Cancel any existing next-maintenance opt-in
	// requests.
	//
	// This member is required.
	OptInType *string

	// The Amazon Resource Name (ARN) of the AWS DMS resource that the pending
	// maintenance action applies to.
	//
	// This member is required.
	ReplicationInstanceArn *string
}

//
type ApplyPendingMaintenanceActionOutput struct {

	// The AWS DMS resource that the pending maintenance action will be applied to.
	ResourcePendingMaintenanceActions *types.ResourcePendingMaintenanceActions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpApplyPendingMaintenanceActionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpApplyPendingMaintenanceAction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpApplyPendingMaintenanceAction{}, middleware.After)
}

func newServiceMetadataMiddleware_opApplyPendingMaintenanceAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "ApplyPendingMaintenanceAction",
	}
}
