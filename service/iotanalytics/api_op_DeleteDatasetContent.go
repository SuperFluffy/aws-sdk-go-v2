// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the content of the specified data set.
func (c *Client) DeleteDatasetContent(ctx context.Context, params *DeleteDatasetContentInput, optFns ...func(*Options)) (*DeleteDatasetContentOutput, error) {
	stack := middleware.NewStack("DeleteDatasetContent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteDatasetContentMiddlewares(stack)
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
	addOpDeleteDatasetContentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDatasetContent(options.Region), middleware.Before)
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
			OperationName: "DeleteDatasetContent",
			Err:           err,
		}
	}
	out := result.(*DeleteDatasetContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDatasetContentInput struct {

	// The name of the data set whose content is deleted.
	//
	// This member is required.
	DatasetName *string

	// The version of the data set whose content is deleted. You can also use the
	// strings "$LATEST" or "$LATEST_SUCCEEDED" to delete the latest or latest
	// successfully completed data set. If not specified, "$LATEST_SUCCEEDED" is the
	// default.
	VersionId *string
}

type DeleteDatasetContentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteDatasetContentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDatasetContent{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDatasetContent{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDatasetContent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "DeleteDatasetContent",
	}
}
