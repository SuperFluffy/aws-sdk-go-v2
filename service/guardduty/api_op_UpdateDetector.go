// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the Amazon GuardDuty detector specified by the detectorId.
func (c *Client) UpdateDetector(ctx context.Context, params *UpdateDetectorInput, optFns ...func(*Options)) (*UpdateDetectorOutput, error) {
	stack := middleware.NewStack("UpdateDetector", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDetectorMiddlewares(stack)
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
	addOpUpdateDetectorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDetector(options.Region), middleware.Before)
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
			OperationName: "UpdateDetector",
			Err:           err,
		}
	}
	out := result.(*UpdateDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDetectorInput struct {

	// The unique ID of the detector to update.
	//
	// This member is required.
	DetectorId *string

	// An object that describes which data sources will be updated.
	DataSources *types.DataSourceConfigurations

	// Specifies whether the detector is enabled or not enabled.
	Enable *bool

	// An enum value that specifies how frequently findings are exported, such as to
	// CloudWatch Events.
	FindingPublishingFrequency types.FindingPublishingFrequency
}

type UpdateDetectorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDetectorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDetector{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDetector{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDetector(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "UpdateDetector",
	}
}
