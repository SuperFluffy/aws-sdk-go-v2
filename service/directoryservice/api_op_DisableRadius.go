// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables multi-factor authentication (MFA) with the Remote Authentication Dial
// In User Service (RADIUS) server for an AD Connector or Microsoft AD directory.
func (c *Client) DisableRadius(ctx context.Context, params *DisableRadiusInput, optFns ...func(*Options)) (*DisableRadiusOutput, error) {
	stack := middleware.NewStack("DisableRadius", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisableRadiusMiddlewares(stack)
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
	addOpDisableRadiusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableRadius(options.Region), middleware.Before)
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
			OperationName: "DisableRadius",
			Err:           err,
		}
	}
	out := result.(*DisableRadiusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DisableRadius () operation.
type DisableRadiusInput struct {

	// The identifier of the directory for which to disable MFA.
	//
	// This member is required.
	DirectoryId *string
}

// Contains the results of the DisableRadius () operation.
type DisableRadiusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisableRadiusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisableRadius{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableRadius{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableRadius(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DisableRadius",
	}
}
