// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the details for the specified service quota. This operation provides a
// different Value than the GetAWSDefaultServiceQuota operation. This operation
// returns the applied value for each quota. GetAWSDefaultServiceQuota returns the
// default AWS value for each quota.
func (c *Client) GetServiceQuota(ctx context.Context, params *GetServiceQuotaInput, optFns ...func(*Options)) (*GetServiceQuotaOutput, error) {
	stack := middleware.NewStack("GetServiceQuota", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetServiceQuotaMiddlewares(stack)
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
	addOpGetServiceQuotaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceQuota(options.Region), middleware.Before)
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
			OperationName: "GetServiceQuota",
			Err:           err,
		}
	}
	out := result.(*GetServiceQuotaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceQuotaInput struct {

	// Identifies the service quota you want to select.
	//
	// This member is required.
	QuotaCode *string

	// Specifies the service that you want to use.
	//
	// This member is required.
	ServiceCode *string
}

type GetServiceQuotaOutput struct {

	// Returns the ServiceQuota () object which contains all values for a quota.
	Quota *types.ServiceQuota

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetServiceQuotaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetServiceQuota{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetServiceQuota{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetServiceQuota(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "GetServiceQuota",
	}
}
