// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// A remediation exception is when a specific resource is no longer considered for
// auto-remediation. This API adds a new exception or updates an exisiting
// exception for a specific resource with a specific AWS Config rule. AWS Config
// generates a remediation exception when a problem occurs executing a remediation
// action to a specific resource. Remediation exceptions blocks auto-remediation
// until the exception is cleared.
func (c *Client) PutRemediationExceptions(ctx context.Context, params *PutRemediationExceptionsInput, optFns ...func(*Options)) (*PutRemediationExceptionsOutput, error) {
	stack := middleware.NewStack("PutRemediationExceptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutRemediationExceptionsMiddlewares(stack)
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
	addOpPutRemediationExceptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutRemediationExceptions(options.Region), middleware.Before)
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
			OperationName: "PutRemediationExceptions",
			Err:           err,
		}
	}
	out := result.(*PutRemediationExceptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRemediationExceptionsInput struct {

	// The name of the AWS Config rule for which you want to create remediation
	// exception.
	//
	// This member is required.
	ConfigRuleName *string

	// An exception list of resource exception keys to be processed with the current
	// request. AWS Config adds exception for each resource key. For example, AWS
	// Config adds 3 exceptions for 3 resource keys.
	//
	// This member is required.
	ResourceKeys []*types.RemediationExceptionResourceKey

	// The exception is automatically deleted after the expiration date.
	ExpirationTime *time.Time

	// The message contains an explanation of the exception.
	Message *string
}

type PutRemediationExceptionsOutput struct {

	// Returns a list of failed remediation exceptions batch objects. Each object in
	// the batch consists of a list of failed items and failure messages.
	FailedBatches []*types.FailedRemediationExceptionBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutRemediationExceptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutRemediationExceptions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRemediationExceptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutRemediationExceptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutRemediationExceptions",
	}
}
