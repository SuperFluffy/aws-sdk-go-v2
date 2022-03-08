// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A temporary Amazon S3 location, where you can copy your files from a source
// location to stage or use as a scratch space in FinSpace notebook.
func (c *Client) GetWorkingLocation(ctx context.Context, params *GetWorkingLocationInput, optFns ...func(*Options)) (*GetWorkingLocationOutput, error) {
	if params == nil {
		params = &GetWorkingLocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkingLocation", params, optFns, c.addOperationGetWorkingLocationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkingLocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkingLocationInput struct {

	// Specify the type of the working location.
	//
	// * SAGEMAKER – Use the Amazon S3
	// location as a temporary location to store data content when working with
	// FinSpace Notebooks that run on SageMaker studio.
	//
	// * INGESTION – Use the Amazon
	// S3 location as a staging location to copy your data content and then use the
	// location with the Changeset creation operation.
	LocationType types.LocationType

	noSmithyDocumentSerde
}

type GetWorkingLocationOutput struct {

	// Returns the Amazon S3 bucket name for the working location.
	S3Bucket *string

	// Returns the Amazon S3 Path for the working location.
	S3Path *string

	// Returns the Amazon S3 URI for the working location.
	S3Uri *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkingLocationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWorkingLocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWorkingLocation{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkingLocation(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetWorkingLocation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "GetWorkingLocation",
	}
}
