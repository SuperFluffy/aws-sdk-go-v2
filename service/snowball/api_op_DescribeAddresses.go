// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a specified number of ADDRESS objects. Calling this API in one of the US
// regions will return addresses from the list of all addresses associated with
// this account in all US regions.
func (c *Client) DescribeAddresses(ctx context.Context, params *DescribeAddressesInput, optFns ...func(*Options)) (*DescribeAddressesOutput, error) {
	stack := middleware.NewStack("DescribeAddresses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAddressesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAddresses(options.Region), middleware.Before)
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
			OperationName: "DescribeAddresses",
			Err:           err,
		}
	}
	out := result.(*DescribeAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAddressesInput struct {

	// The number of ADDRESS objects to return.
	MaxResults *int32

	// HTTP requests are stateless. To identify what object comes "next" in the list of
	// ADDRESS objects, you have the option of specifying a value for NextToken as the
	// starting point for your list of returned addresses.
	NextToken *string
}

type DescribeAddressesOutput struct {

	// The Snowball shipping addresses that were created for this account.
	Addresses []*types.Address

	// HTTP requests are stateless. If you use the automatically generated NextToken
	// value in your next DescribeAddresses call, your list of returned addresses will
	// start from this point in the array.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAddressesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAddresses{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAddresses{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAddresses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "DescribeAddresses",
	}
}
