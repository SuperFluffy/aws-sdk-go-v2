// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides metrics on the accuracy of the models that were trained by the
// CreatePredictor () operation. Use metrics to see how well the model performed
// and to decide whether to use the predictor to generate a forecast. For more
// information, see metrics (). This operation generates metrics for each backtest
// window that was evaluated. The number of backtest windows
// (NumberOfBacktestWindows) is specified using the EvaluationParameters () object,
// which is optionally included in the CreatePredictor request. If
// NumberOfBacktestWindows isn't specified, the number defaults to one. The
// parameters of the filling method determine which items contribute to the
// metrics. If you want all items to contribute, specify zero. If you want only
// those items that have complete data in the range being evaluated to contribute,
// specify nan. For more information, see FeaturizationMethod ().  <note> <p>Before
// you can get accuracy metrics, the <code>Status</code> of the predictor must be
// <code>ACTIVE</code>, signifying that training has completed. To get the status,
// use the <a>DescribePredictor</a> operation.</p> </note>
func (c *Client) GetAccuracyMetrics(ctx context.Context, params *GetAccuracyMetricsInput, optFns ...func(*Options)) (*GetAccuracyMetricsOutput, error) {
	stack := middleware.NewStack("GetAccuracyMetrics", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetAccuracyMetricsMiddlewares(stack)
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
	addOpGetAccuracyMetricsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccuracyMetrics(options.Region), middleware.Before)
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
			OperationName: "GetAccuracyMetrics",
			Err:           err,
		}
	}
	out := result.(*GetAccuracyMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccuracyMetricsInput struct {

	// The Amazon Resource Name (ARN) of the predictor to get metrics for.
	//
	// This member is required.
	PredictorArn *string
}

type GetAccuracyMetricsOutput struct {

	// An array of results from evaluating the predictor.
	PredictorEvaluationResults []*types.EvaluationResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetAccuracyMetricsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetAccuracyMetrics{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAccuracyMetrics{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAccuracyMetrics(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "GetAccuracyMetrics",
	}
}
