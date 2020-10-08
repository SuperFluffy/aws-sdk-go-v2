// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates the associated AWS Key Management Service (AWS KMS) customer
// master key (CMK) from the specified log group. After the AWS KMS CMK is
// disassociated from the log group, AWS CloudWatch Logs stops encrypting newly
// ingested data for the log group. All previously ingested data remains encrypted,
// and AWS CloudWatch Logs requires permissions for the CMK whenever the encrypted
// data is requested. Note that it can take up to 5 minutes for this operation to
// take effect.
func (c *Client) DisassociateKmsKey(ctx context.Context, params *DisassociateKmsKeyInput, optFns ...func(*Options)) (*DisassociateKmsKeyOutput, error) {
	stack := middleware.NewStack("DisassociateKmsKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisassociateKmsKeyMiddlewares(stack)
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
	addOpDisassociateKmsKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateKmsKey(options.Region), middleware.Before)
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
			OperationName: "DisassociateKmsKey",
			Err:           err,
		}
	}
	out := result.(*DisassociateKmsKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateKmsKeyInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string
}

type DisassociateKmsKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisassociateKmsKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateKmsKey{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateKmsKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateKmsKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DisassociateKmsKey",
	}
}
