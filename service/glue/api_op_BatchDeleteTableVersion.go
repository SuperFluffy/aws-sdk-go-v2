// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a specified batch of versions of a table.
func (c *Client) BatchDeleteTableVersion(ctx context.Context, params *BatchDeleteTableVersionInput, optFns ...func(*Options)) (*BatchDeleteTableVersionOutput, error) {
	stack := middleware.NewStack("BatchDeleteTableVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchDeleteTableVersionMiddlewares(stack)
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
	addOpBatchDeleteTableVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteTableVersion(options.Region), middleware.Before)
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
			OperationName: "BatchDeleteTableVersion",
			Err:           err,
		}
	}
	out := result.(*BatchDeleteTableVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteTableVersionInput struct {

	// The database in the catalog in which the table resides. For Hive compatibility,
	// this name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// The name of the table. For Hive compatibility, this name is entirely lowercase.
	//
	// This member is required.
	TableName *string

	// A list of the IDs of versions to be deleted. A VersionId is a string
	// representation of an integer. Each version is incremented by 1.
	//
	// This member is required.
	VersionIds []*string

	// The ID of the Data Catalog where the tables reside. If none is provided, the AWS
	// account ID is used by default.
	CatalogId *string
}

type BatchDeleteTableVersionOutput struct {

	// A list of errors encountered while trying to delete the specified table
	// versions.
	Errors []*types.TableVersionError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchDeleteTableVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeleteTableVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeleteTableVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchDeleteTableVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "BatchDeleteTableVersion",
	}
}
