// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified interconnect. Intended for use by AWS Direct Connect
// Partners only.
func (c *Client) DeleteInterconnect(ctx context.Context, params *DeleteInterconnectInput, optFns ...func(*Options)) (*DeleteInterconnectOutput, error) {
	stack := middleware.NewStack("DeleteInterconnect", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteInterconnectMiddlewares(stack)
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
	addOpDeleteInterconnectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInterconnect(options.Region), middleware.Before)
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
			OperationName: "DeleteInterconnect",
			Err:           err,
		}
	}
	out := result.(*DeleteInterconnectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteInterconnectInput struct {

	// The ID of the interconnect.
	//
	// This member is required.
	InterconnectId *string
}

type DeleteInterconnectOutput struct {

	// The state of the interconnect. The following are the possible values:
	//
	//     *
	// requested: The initial state of an interconnect. The interconnect stays in the
	// requested state until the Letter of Authorization (LOA) is sent to the
	// customer.
	//
	//     * pending: The interconnect is approved, and is being
	// initialized.
	//
	//     * available: The network link is up, and the interconnect is
	// ready for use.
	//
	//     * down: The network link is down.
	//
	//     * deleting: The
	// interconnect is being deleted.
	//
	//     * deleted: The interconnect is deleted.
	//
	//
	// * unknown: The state of the interconnect is not available.
	InterconnectState types.InterconnectState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteInterconnectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteInterconnect{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteInterconnect{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteInterconnect(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DeleteInterconnect",
	}
}
