// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes automated backups based on the source instance's DbiResourceId value or
// the restorable instance's resource ID.
func (c *Client) DeleteDBInstanceAutomatedBackup(ctx context.Context, params *DeleteDBInstanceAutomatedBackupInput, optFns ...func(*Options)) (*DeleteDBInstanceAutomatedBackupOutput, error) {
	stack := middleware.NewStack("DeleteDBInstanceAutomatedBackup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteDBInstanceAutomatedBackupMiddlewares(stack)
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
	addOpDeleteDBInstanceAutomatedBackupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBInstanceAutomatedBackup(options.Region), middleware.Before)
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
			OperationName: "DeleteDBInstanceAutomatedBackup",
			Err:           err,
		}
	}
	out := result.(*DeleteDBInstanceAutomatedBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Parameter input for the DeleteDBInstanceAutomatedBackup operation.
type DeleteDBInstanceAutomatedBackupInput struct {

	// The identifier for the source DB instance, which can't be changed and which is
	// unique to an AWS Region.
	//
	// This member is required.
	DbiResourceId *string
}

type DeleteDBInstanceAutomatedBackupOutput struct {

	// An automated backup of a DB instance. It it consists of system backups,
	// transaction logs, and the database instance properties that existed at the time
	// you deleted the source instance.
	DBInstanceAutomatedBackup *types.DBInstanceAutomatedBackup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteDBInstanceAutomatedBackupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBInstanceAutomatedBackup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBInstanceAutomatedBackup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDBInstanceAutomatedBackup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBInstanceAutomatedBackup",
	}
}
