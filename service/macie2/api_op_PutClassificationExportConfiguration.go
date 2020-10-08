// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates the configuration settings for storing data classification
// results.
func (c *Client) PutClassificationExportConfiguration(ctx context.Context, params *PutClassificationExportConfigurationInput, optFns ...func(*Options)) (*PutClassificationExportConfigurationOutput, error) {
	stack := middleware.NewStack("PutClassificationExportConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutClassificationExportConfigurationMiddlewares(stack)
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
	addOpPutClassificationExportConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutClassificationExportConfiguration(options.Region), middleware.Before)
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
			OperationName: "PutClassificationExportConfiguration",
			Err:           err,
		}
	}
	out := result.(*PutClassificationExportConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutClassificationExportConfigurationInput struct {

	// The location to store data classification results in, and the encryption
	// settings to use when storing results in that location.
	//
	// This member is required.
	Configuration *types.ClassificationExportConfiguration
}

type PutClassificationExportConfigurationOutput struct {

	// The location where the data classification results are stored, and the
	// encryption settings that are used when storing results in that location.
	Configuration *types.ClassificationExportConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutClassificationExportConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutClassificationExportConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutClassificationExportConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutClassificationExportConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "PutClassificationExportConfiguration",
	}
}
