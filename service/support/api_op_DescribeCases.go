// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of cases that you specify by passing one or more case IDs. You
// can use the afterTime and beforeTime parameters to filter the cases by date. You
// can set values for the includeResolvedCases and includeCommunications parameters
// to specify how much information to return. The response returns the following in
// JSON format:
//
//     * One or more CaseDetails
// (https://docs.aws.amazon.com/awssupport/latest/APIReference/API_CaseDetails.html)
// data types.
//
//     * One or more nextToken values, which specify where to paginate
// the returned records represented by the CaseDetails objects.
//
// Case data is
// available for 12 months after creation. If a case was created more than 12
// months ago, a request might return an error.
//
//     * You must have a Business or
// Enterprise support plan to use the AWS Support API.
//
//     * If you call the AWS
// Support API from an account that does not have a Business or Enterprise support
// plan, the SubscriptionRequiredException error message appears. For information
// about changing your support plan, see AWS Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) DescribeCases(ctx context.Context, params *DescribeCasesInput, optFns ...func(*Options)) (*DescribeCasesOutput, error) {
	stack := middleware.NewStack("DescribeCases", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeCasesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCases(options.Region), middleware.Before)
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
			OperationName: "DescribeCases",
			Err:           err,
		}
	}
	out := result.(*DescribeCasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCasesInput struct {

	// The start date for a filtered date search on support case communications. Case
	// communications are available for 12 months after creation.
	AfterTime *string

	// The end date for a filtered date search on support case communications. Case
	// communications are available for 12 months after creation.
	BeforeTime *string

	// A list of ID numbers of the support cases you want returned. The maximum number
	// of cases is 100.
	CaseIdList []*string

	// The ID displayed for a case in the AWS Support Center user interface.
	DisplayId *string

	// Specifies whether to include communications in the DescribeCases response. By
	// default, communications are incuded.
	IncludeCommunications *bool

	// Specifies whether to include resolved support cases in the DescribeCases
	// response. By default, resolved cases aren't included.
	IncludeResolvedCases *bool

	// The ISO 639-1 code for the language in which AWS provides support. AWS Support
	// currently supports English ("en") and Japanese ("ja"). Language parameters must
	// be passed explicitly for operations that take them.
	Language *string

	// The maximum number of results to return before paginating.
	MaxResults *int32

	// A resumption point for pagination.
	NextToken *string
}

// Returns an array of CaseDetails
// (https://docs.aws.amazon.com/awssupport/latest/APIReference/API_CaseDetails.html)
// objects and a nextToken that defines a point for pagination in the result set.
type DescribeCasesOutput struct {

	// The details for the cases that match the request.
	Cases []*types.CaseDetails

	// A resumption point for pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeCasesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCases{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCases{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCases(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeCases",
	}
}
