// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an Amazon SageMaker experiment. All trials associated with the
// experiment must be deleted first. Use the ListTrials () API to get a list of the
// trials associated with the experiment.
func (c *Client) DeleteExperiment(ctx context.Context, params *DeleteExperimentInput, optFns ...func(*Options)) (*DeleteExperimentOutput, error) {
	stack := middleware.NewStack("DeleteExperiment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteExperimentMiddlewares(stack)
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
	addOpDeleteExperimentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteExperiment(options.Region), middleware.Before)
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
			OperationName: "DeleteExperiment",
			Err:           err,
		}
	}
	out := result.(*DeleteExperimentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteExperimentInput struct {

	// The name of the experiment to delete.
	//
	// This member is required.
	ExperimentName *string
}

type DeleteExperimentOutput struct {

	// The Amazon Resource Name (ARN) of the experiment that is being deleted.
	ExperimentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteExperimentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteExperiment{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteExperiment{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteExperiment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DeleteExperiment",
	}
}
