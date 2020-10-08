// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a particular detector version.
func (c *Client) GetDetectorVersion(ctx context.Context, params *GetDetectorVersionInput, optFns ...func(*Options)) (*GetDetectorVersionOutput, error) {
	stack := middleware.NewStack("GetDetectorVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDetectorVersionMiddlewares(stack)
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
	addOpGetDetectorVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDetectorVersion(options.Region), middleware.Before)
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
			OperationName: "GetDetectorVersion",
			Err:           err,
		}
	}
	out := result.(*GetDetectorVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDetectorVersionInput struct {

	// The detector ID.
	//
	// This member is required.
	DetectorId *string

	// The detector version ID.
	//
	// This member is required.
	DetectorVersionId *string
}

type GetDetectorVersionOutput struct {

	// The detector version ARN.
	Arn *string

	// The timestamp when the detector version was created.
	CreatedTime *string

	// The detector version description.
	Description *string

	// The detector ID.
	DetectorId *string

	// The detector version ID.
	DetectorVersionId *string

	// The Amazon SageMaker model endpoints included in the detector version.
	ExternalModelEndpoints []*string

	// The timestamp when the detector version was last updated.
	LastUpdatedTime *string

	// The model versions included in the detector version.
	ModelVersions []*types.ModelVersion

	// The execution mode of the rule in the dectector FIRST_MATCHED indicates that
	// Amazon Fraud Detector evaluates rules sequentially, first to last, stopping at
	// the first matched rule. Amazon Fraud dectector then provides the outcomes for
	// that single rule. ALL_MATCHED indicates that Amazon Fraud Detector evaluates all
	// rules and returns the outcomes for all matched rules. You can define and edit
	// the rule mode at the detector version level, when it is in draft status.
	RuleExecutionMode types.RuleExecutionMode

	// The rules included in the detector version.
	Rules []*types.Rule

	// The status of the detector version.
	Status types.DetectorVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDetectorVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDetectorVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDetectorVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDetectorVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetDetectorVersion",
	}
}
