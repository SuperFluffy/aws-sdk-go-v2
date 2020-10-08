// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the properties of an endpoint in an Amazon Aurora DB cluster. This
// action only applies to Aurora DB clusters.
func (c *Client) ModifyDBClusterEndpoint(ctx context.Context, params *ModifyDBClusterEndpointInput, optFns ...func(*Options)) (*ModifyDBClusterEndpointOutput, error) {
	stack := middleware.NewStack("ModifyDBClusterEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyDBClusterEndpointMiddlewares(stack)
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
	addOpModifyDBClusterEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBClusterEndpoint(options.Region), middleware.Before)
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
			OperationName: "ModifyDBClusterEndpoint",
			Err:           err,
		}
	}
	out := result.(*ModifyDBClusterEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBClusterEndpointInput struct {

	// The identifier of the endpoint to modify. This parameter is stored as a
	// lowercase string.
	//
	// This member is required.
	DBClusterEndpointIdentifier *string

	// The type of the endpoint. One of: READER, WRITER, ANY.
	EndpointType *string

	// List of DB instance identifiers that aren't part of the custom endpoint group.
	// All other eligible instances are reachable through the custom endpoint. Only
	// relevant if the list of static members is empty.
	ExcludedMembers []*string

	// List of DB instance identifiers that are part of the custom endpoint group.
	StaticMembers []*string
}

// This data type represents the information you need to connect to an Amazon
// Aurora DB cluster. This data type is used as a response element in the following
// actions:
//
//     * CreateDBClusterEndpoint
//
//     * DescribeDBClusterEndpoints
//
//     *
// ModifyDBClusterEndpoint
//
//     * DeleteDBClusterEndpoint
//
// For the data structure
// that represents Amazon RDS DB instance endpoints, see Endpoint.
type ModifyDBClusterEndpointOutput struct {

	// The type associated with a custom endpoint. One of: READER, WRITER, ANY.
	CustomEndpointType *string

	// The Amazon Resource Name (ARN) for the endpoint.
	DBClusterEndpointArn *string

	// The identifier associated with the endpoint. This parameter is stored as a
	// lowercase string.
	DBClusterEndpointIdentifier *string

	// A unique system-generated identifier for an endpoint. It remains the same for
	// the whole life of the endpoint.
	DBClusterEndpointResourceIdentifier *string

	// The DB cluster identifier of the DB cluster associated with the endpoint. This
	// parameter is stored as a lowercase string.
	DBClusterIdentifier *string

	// The DNS address of the endpoint.
	Endpoint *string

	// The type of the endpoint. One of: READER, WRITER, CUSTOM.
	EndpointType *string

	// List of DB instance identifiers that aren't part of the custom endpoint group.
	// All other eligible instances are reachable through the custom endpoint. Only
	// relevant if the list of static members is empty.
	ExcludedMembers []*string

	// List of DB instance identifiers that are part of the custom endpoint group.
	StaticMembers []*string

	// The current status of the endpoint. One of: creating, available, deleting,
	// modifying.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyDBClusterEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBClusterEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBClusterEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyDBClusterEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBClusterEndpoint",
	}
}
