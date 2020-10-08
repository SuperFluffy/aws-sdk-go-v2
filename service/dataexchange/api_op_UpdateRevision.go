// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This operation updates a revision.
func (c *Client) UpdateRevision(ctx context.Context, params *UpdateRevisionInput, optFns ...func(*Options)) (*UpdateRevisionOutput, error) {
	stack := middleware.NewStack("UpdateRevision", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateRevisionMiddlewares(stack)
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
	addOpUpdateRevisionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRevision(options.Region), middleware.Before)
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
			OperationName: "UpdateRevision",
			Err:           err,
		}
	}
	out := result.(*UpdateRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body for UpdateRevision.
type UpdateRevisionInput struct {

	// The unique identifier for a data set.
	//
	// This member is required.
	DataSetId *string

	// The unique identifier for a revision.
	//
	// This member is required.
	RevisionId *string

	// An optional comment about the revision.
	Comment *string

	// Finalizing a revision tells AWS Data Exchange that your changes to the assets in
	// the revision are complete. After it's in this read-only state, you can publish
	// the revision to your products.
	Finalized *bool
}

type UpdateRevisionOutput struct {

	// The ARN for the revision.
	Arn *string

	// An optional comment about the revision.
	Comment *string

	// The date and time that the revision was created, in ISO 8601 format.
	CreatedAt *time.Time

	// The unique identifier for the data set associated with this revision.
	DataSetId *string

	// To publish a revision to a data set in a product, the revision must first be
	// finalized. Finalizing a revision tells AWS Data Exchange that changes to the
	// assets in the revision are complete. After it's in this read-only state, you can
	// publish the revision to your products. Finalized revisions can be published
	// through the AWS Data Exchange console or the AWS Marketplace Catalog API, using
	// the StartChangeSet AWS Marketplace Catalog API action. When using the API,
	// revisions are uniquely identified by their ARN.
	Finalized *bool

	// The unique identifier for the revision.
	Id *string

	// The revision ID of the owned revision corresponding to the entitled revision
	// being viewed. This parameter is returned when a revision owner is viewing the
	// entitled copy of its owned revision.
	SourceId *string

	// The date and time that the revision was last updated, in ISO 8601 format.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateRevisionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRevision{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRevision{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRevision(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "UpdateRevision",
	}
}
