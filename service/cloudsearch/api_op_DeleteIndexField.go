// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes an IndexField () from the search domain. For more information, see
// Configuring Index Fields
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-index-fields.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DeleteIndexField(ctx context.Context, params *DeleteIndexFieldInput, optFns ...func(*Options)) (*DeleteIndexFieldOutput, error) {
	stack := middleware.NewStack("DeleteIndexField", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteIndexFieldMiddlewares(stack)
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
	addOpDeleteIndexFieldValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIndexField(options.Region), middleware.Before)
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
			OperationName: "DeleteIndexField",
			Err:           err,
		}
	}
	out := result.(*DeleteIndexFieldOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DeleteIndexField () operation. Specifies the
// name of the domain you want to update and the name of the index field you want
// to delete.
type DeleteIndexFieldInput struct {

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string

	// The name of the index field your want to remove from the domain's indexing
	// options.
	//
	// This member is required.
	IndexFieldName *string
}

// The result of a DeleteIndexField () request.
type DeleteIndexFieldOutput struct {

	// The status of the index field being deleted.
	//
	// This member is required.
	IndexField *types.IndexFieldStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteIndexFieldMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteIndexField{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteIndexField{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteIndexField(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DeleteIndexField",
	}
}
