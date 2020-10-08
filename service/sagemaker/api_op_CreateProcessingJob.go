// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a processing job.
func (c *Client) CreateProcessingJob(ctx context.Context, params *CreateProcessingJobInput, optFns ...func(*Options)) (*CreateProcessingJobOutput, error) {
	stack := middleware.NewStack("CreateProcessingJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateProcessingJobMiddlewares(stack)
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
	addOpCreateProcessingJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProcessingJob(options.Region), middleware.Before)
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
			OperationName: "CreateProcessingJob",
			Err:           err,
		}
	}
	out := result.(*CreateProcessingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProcessingJobInput struct {

	// Configures the processing job to run a specified Docker container image.
	//
	// This member is required.
	AppSpecification *types.AppSpecification

	// The name of the processing job. The name must be unique within an AWS Region in
	// the AWS account.
	//
	// This member is required.
	ProcessingJobName *string

	// Identifies the resources, ML compute instances, and ML storage volumes to deploy
	// for a processing job. In distributed training, you specify more than one
	// instance.
	//
	// This member is required.
	ProcessingResources *types.ProcessingResources

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform tasks on your behalf.
	//
	// This member is required.
	RoleArn *string

	// Sets the environment variables in the Docker container.
	Environment map[string]*string

	// Associates a SageMaker job as a trial component with an experiment and trial.
	// Specified when you call the following APIs:
	//
	//     * CreateProcessingJob ()
	//
	//     *
	// CreateTrainingJob ()
	//
	//     * CreateTransformJob ()
	ExperimentConfig *types.ExperimentConfig

	// Networking options for a processing job.
	NetworkConfig *types.NetworkConfig

	// For each input, data is downloaded from S3 into the processing container before
	// the processing job begins running if "S3InputMode" is set to File.
	ProcessingInputs []*types.ProcessingInput

	// Output configuration for the processing job.
	ProcessingOutputConfig *types.ProcessingOutputConfig

	// The time limit for how long the processing job is allowed to run.
	StoppingCondition *types.ProcessingStoppingCondition

	// (Optional) An array of key-value pairs. For more information, see Using Cost
	// Allocation Tags
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL)
	// in the AWS Billing and Cost Management User Guide.
	Tags []*types.Tag
}

type CreateProcessingJobOutput struct {

	// The Amazon Resource Name (ARN) of the processing job.
	//
	// This member is required.
	ProcessingJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateProcessingJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProcessingJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProcessingJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateProcessingJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateProcessingJob",
	}
}
