// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the IP address filters associated with your AWS account in the current AWS
// Region. For information about managing IP address filters, see the Amazon SES
// Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-ip-filters.html).
// You can execute this operation no more than once per second.
func (c *Client) ListReceiptFilters(ctx context.Context, params *ListReceiptFiltersInput, optFns ...func(*Options)) (*ListReceiptFiltersOutput, error) {
	stack := middleware.NewStack("ListReceiptFilters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListReceiptFiltersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListReceiptFilters(options.Region), middleware.Before)
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
			OperationName: "ListReceiptFilters",
			Err:           err,
		}
	}
	out := result.(*ListReceiptFiltersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list the IP address filters that exist under your AWS
// account. You use IP address filters when you receive email with Amazon SES. For
// more information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type ListReceiptFiltersInput struct {
}

// A list of IP address filters that exist under your AWS account.
type ListReceiptFiltersOutput struct {

	// A list of IP address filter data structures, which each consist of a name, an IP
	// address range, and whether to allow or block mail from it.
	Filters []*types.ReceiptFilter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListReceiptFiltersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListReceiptFilters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListReceiptFilters{}, middleware.After)
}

func newServiceMetadataMiddleware_opListReceiptFilters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListReceiptFilters",
	}
}
