// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a specified job and whether that job has been received
// by the job worker. Used for custom actions only.
func (c *Client) AcknowledgeJob(ctx context.Context, params *AcknowledgeJobInput, optFns ...func(*Options)) (*AcknowledgeJobOutput, error) {
	stack := middleware.NewStack("AcknowledgeJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAcknowledgeJobMiddlewares(stack)
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
	addOpAcknowledgeJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcknowledgeJob(options.Region), middleware.Before)
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
			OperationName: "AcknowledgeJob",
			Err:           err,
		}
	}
	out := result.(*AcknowledgeJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an AcknowledgeJob action.
type AcknowledgeJobInput struct {

	// The unique system-generated ID of the job for which you want to confirm receipt.
	//
	// This member is required.
	JobId *string

	// A system-generated random number that AWS CodePipeline uses to ensure that the
	// job is being worked on by only one job worker. Get this number from the response
	// of the PollForJobs () request that returned this job.
	//
	// This member is required.
	Nonce *string
}

// Represents the output of an AcknowledgeJob action.
type AcknowledgeJobOutput struct {

	// Whether the job worker has received the specified job.
	Status types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAcknowledgeJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAcknowledgeJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAcknowledgeJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opAcknowledgeJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "AcknowledgeJob",
	}
}
