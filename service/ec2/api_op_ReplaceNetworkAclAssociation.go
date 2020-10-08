// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes which network ACL a subnet is associated with. By default when you
// create a subnet, it's automatically associated with the default network ACL. For
// more information, see Network ACLs
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html) in the Amazon
// Virtual Private Cloud User Guide. This is an idempotent operation.
func (c *Client) ReplaceNetworkAclAssociation(ctx context.Context, params *ReplaceNetworkAclAssociationInput, optFns ...func(*Options)) (*ReplaceNetworkAclAssociationOutput, error) {
	stack := middleware.NewStack("ReplaceNetworkAclAssociation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpReplaceNetworkAclAssociationMiddlewares(stack)
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
	addOpReplaceNetworkAclAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReplaceNetworkAclAssociation(options.Region), middleware.Before)
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
			OperationName: "ReplaceNetworkAclAssociation",
			Err:           err,
		}
	}
	out := result.(*ReplaceNetworkAclAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplaceNetworkAclAssociationInput struct {

	// The ID of the current association between the original network ACL and the
	// subnet.
	//
	// This member is required.
	AssociationId *string

	// The ID of the new network ACL to associate with the subnet.
	//
	// This member is required.
	NetworkAclId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ReplaceNetworkAclAssociationOutput struct {

	// The ID of the new association.
	NewAssociationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpReplaceNetworkAclAssociationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpReplaceNetworkAclAssociation{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpReplaceNetworkAclAssociation{}, middleware.After)
}

func newServiceMetadataMiddleware_opReplaceNetworkAclAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReplaceNetworkAclAssociation",
	}
}
