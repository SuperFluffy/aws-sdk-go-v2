// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates the role from a group.
func (c *Client) DisassociateRoleFromGroup(ctx context.Context, params *DisassociateRoleFromGroupInput, optFns ...func(*Options)) (*DisassociateRoleFromGroupOutput, error) {
	stack := middleware.NewStack("DisassociateRoleFromGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisassociateRoleFromGroupMiddlewares(stack)
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
	addOpDisassociateRoleFromGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateRoleFromGroup(options.Region), middleware.Before)
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
			OperationName: "DisassociateRoleFromGroup",
			Err:           err,
		}
	}
	out := result.(*DisassociateRoleFromGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateRoleFromGroupInput struct {

	// The ID of the Greengrass group.
	//
	// This member is required.
	GroupId *string
}

type DisassociateRoleFromGroupOutput struct {

	// The time, in milliseconds since the epoch, when the role was disassociated from
	// the group.
	DisassociatedAt *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisassociateRoleFromGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateRoleFromGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateRoleFromGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateRoleFromGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "DisassociateRoleFromGroup",
	}
}
