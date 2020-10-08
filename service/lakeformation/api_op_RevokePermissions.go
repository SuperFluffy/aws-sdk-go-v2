// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Revokes permissions to the principal to access metadata in the Data Catalog and
// data organized in underlying data storage such as Amazon S3.
func (c *Client) RevokePermissions(ctx context.Context, params *RevokePermissionsInput, optFns ...func(*Options)) (*RevokePermissionsOutput, error) {
	stack := middleware.NewStack("RevokePermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRevokePermissionsMiddlewares(stack)
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
	addOpRevokePermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRevokePermissions(options.Region), middleware.Before)
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
			OperationName: "RevokePermissions",
			Err:           err,
		}
	}
	out := result.(*RevokePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokePermissionsInput struct {

	// The permissions revoked to the principal on the resource. For information about
	// permissions, see Security and Access Control to Metadata and Data
	// (https://docs-aws.amazon.com/lake-formation/latest/dg/security-data-access.html).
	//
	// This member is required.
	Permissions []types.Permission

	// The principal to be revoked permissions on the resource.
	//
	// This member is required.
	Principal *types.DataLakePrincipal

	// The resource to which permissions are to be revoked.
	//
	// This member is required.
	Resource *types.Resource

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string

	// Indicates a list of permissions for which to revoke the grant option allowing
	// the principal to pass permissions to other principals.
	PermissionsWithGrantOption []types.Permission
}

type RevokePermissionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRevokePermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRevokePermissions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRevokePermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opRevokePermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "RevokePermissions",
	}
}
