// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Replaces an IAM instance profile for the specified running instance. You can use
// this action to change the IAM instance profile that's associated with an
// instance without having to disassociate the existing IAM instance profile first.
// Use DescribeIamInstanceProfileAssociations () to get the association ID.
func (c *Client) ReplaceIamInstanceProfileAssociation(ctx context.Context, params *ReplaceIamInstanceProfileAssociationInput, optFns ...func(*Options)) (*ReplaceIamInstanceProfileAssociationOutput, error) {
	stack := middleware.NewStack("ReplaceIamInstanceProfileAssociation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpReplaceIamInstanceProfileAssociationMiddlewares(stack)
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
	addOpReplaceIamInstanceProfileAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opReplaceIamInstanceProfileAssociation(options.Region), middleware.Before)
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
			OperationName: "ReplaceIamInstanceProfileAssociation",
			Err:           err,
		}
	}
	out := result.(*ReplaceIamInstanceProfileAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplaceIamInstanceProfileAssociationInput struct {

	// The ID of the existing IAM instance profile association.
	//
	// This member is required.
	AssociationId *string

	// The IAM instance profile.
	//
	// This member is required.
	IamInstanceProfile *types.IamInstanceProfileSpecification
}

type ReplaceIamInstanceProfileAssociationOutput struct {

	// Information about the IAM instance profile association.
	IamInstanceProfileAssociation *types.IamInstanceProfileAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpReplaceIamInstanceProfileAssociationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpReplaceIamInstanceProfileAssociation{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpReplaceIamInstanceProfileAssociation{}, middleware.After)
}

func newServiceMetadataMiddleware_opReplaceIamInstanceProfileAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReplaceIamInstanceProfileAssociation",
	}
}
