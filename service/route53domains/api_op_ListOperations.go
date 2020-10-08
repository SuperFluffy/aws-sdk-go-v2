// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about all of the operations that return an operation ID and
// that have ever been performed on domains that were registered by the current
// account.
func (c *Client) ListOperations(ctx context.Context, params *ListOperationsInput, optFns ...func(*Options)) (*ListOperationsOutput, error) {
	stack := middleware.NewStack("ListOperations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListOperationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListOperations(options.Region), middleware.Before)
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
			OperationName: "ListOperations",
			Err:           err,
		}
	}
	out := result.(*ListOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListOperations request includes the following elements.
type ListOperationsInput struct {

	// For an initial request for a list of operations, omit this element. If the
	// number of operations that are not yet complete is greater than the value that
	// you specified for MaxItems, you can use Marker to return additional operations.
	// Get the value of NextPageMarker from the previous response, and submit another
	// request that includes the value of NextPageMarker in the Marker element.
	Marker *string

	// Number of domains to be returned. Default: 20
	MaxItems *int32

	// An optional parameter that lets you get information about all the operations
	// that you submitted after a specified date and time. Specify the date and time in
	// Unix time format and Coordinated Universal time (UTC).
	SubmittedSince *time.Time
}

// The ListOperations response includes the following elements.
type ListOperationsOutput struct {

	// Lists summaries of the operations.
	//
	// This member is required.
	Operations []*types.OperationSummary

	// If there are more operations than you specified for MaxItems in the request,
	// submit another request and include the value of NextPageMarker in the value of
	// Marker.
	NextPageMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListOperationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListOperations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOperations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListOperations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "ListOperations",
	}
}
