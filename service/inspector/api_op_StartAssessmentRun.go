// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts the assessment run specified by the ARN of the assessment template. For
// this API to function properly, you must not exceed the limit of running up to
// 500 concurrent agents per AWS account.
func (c *Client) StartAssessmentRun(ctx context.Context, params *StartAssessmentRunInput, optFns ...func(*Options)) (*StartAssessmentRunOutput, error) {
	stack := middleware.NewStack("StartAssessmentRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartAssessmentRunMiddlewares(stack)
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
	addOpStartAssessmentRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartAssessmentRun(options.Region), middleware.Before)
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
			OperationName: "StartAssessmentRun",
			Err:           err,
		}
	}
	out := result.(*StartAssessmentRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAssessmentRunInput struct {

	// The ARN of the assessment template of the assessment run that you want to start.
	//
	// This member is required.
	AssessmentTemplateArn *string

	// You can specify the name for the assessment run. The name must be unique for the
	// assessment template whose ARN is used to start the assessment run.
	AssessmentRunName *string
}

type StartAssessmentRunOutput struct {

	// The ARN of the assessment run that has been started.
	//
	// This member is required.
	AssessmentRunArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartAssessmentRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartAssessmentRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartAssessmentRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartAssessmentRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "StartAssessmentRun",
	}
}
