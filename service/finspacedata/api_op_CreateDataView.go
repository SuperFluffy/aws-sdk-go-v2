// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Dataview for a Dataset.
func (c *Client) CreateDataView(ctx context.Context, params *CreateDataViewInput, optFns ...func(*Options)) (*CreateDataViewOutput, error) {
	if params == nil {
		params = &CreateDataViewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataView", params, optFns, c.addOperationCreateDataViewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataViewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request for creating a data view.
type CreateDataViewInput struct {

	// The unique Dataset identifier that is used to create a Dataview.
	//
	// This member is required.
	DatasetId *string

	// Options that define the destination type for the Dataview.
	//
	// This member is required.
	DestinationTypeParams *types.DataViewDestinationTypeParams

	// Beginning time to use for the Dataview. The value is determined as epoch time in
	// milliseconds. For example, the value for Monday, November 1, 2021 12:00:00 PM
	// UTC is specified as 1635768000000.
	AsOfTimestamp int64

	// Flag to indicate Dataview should be updated automatically.
	AutoUpdate bool

	// A token that ensures idempotency. This token expires in 10 minutes.
	ClientToken *string

	// Ordered set of column names used to partition data.
	PartitionColumns []string

	// Columns to be used for sorting the data.
	SortColumns []string

	noSmithyDocumentSerde
}

// Response for creating a data view.
type CreateDataViewOutput struct {

	// The unique identifier for the created Dataview.
	DataViewId *string

	// The unique identifier of the Dataset used for the Dataview.
	DatasetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataViewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataView{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataView{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateDataViewMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDataViewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataView(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDataView struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDataView) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDataView) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDataViewInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDataViewInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDataViewMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDataView{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDataView(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "CreateDataView",
	}
}
