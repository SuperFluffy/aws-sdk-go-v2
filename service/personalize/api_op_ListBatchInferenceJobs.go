// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of the batch inference jobs that have been performed off of a
// solution version.
func (c *Client) ListBatchInferenceJobs(ctx context.Context, params *ListBatchInferenceJobsInput, optFns ...func(*Options)) (*ListBatchInferenceJobsOutput, error) {
	stack := middleware.NewStack("ListBatchInferenceJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListBatchInferenceJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBatchInferenceJobs(options.Region), middleware.Before)
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
			OperationName: "ListBatchInferenceJobs",
			Err:           err,
		}
	}
	out := result.(*ListBatchInferenceJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBatchInferenceJobsInput struct {

	// The maximum number of batch inference job results to return in each page. The
	// default value is 100.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The Amazon Resource Name (ARN) of the solution version from which the batch
	// inference jobs were created.
	SolutionVersionArn *string
}

type ListBatchInferenceJobsOutput struct {

	// A list containing information on each job that is returned.
	BatchInferenceJobs []*types.BatchInferenceJobSummary

	// The token to use to retreive the next page of results. The value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListBatchInferenceJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListBatchInferenceJobs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBatchInferenceJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListBatchInferenceJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListBatchInferenceJobs",
	}
}
