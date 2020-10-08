// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the email templates present in your Amazon SES account in the current AWS
// Region.  <p>You can execute this operation no more than once per second.</p>
func (c *Client) ListEmailTemplates(ctx context.Context, params *ListEmailTemplatesInput, optFns ...func(*Options)) (*ListEmailTemplatesOutput, error) {
	stack := middleware.NewStack("ListEmailTemplates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListEmailTemplatesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEmailTemplates(options.Region), middleware.Before)
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
			OperationName: "ListEmailTemplates",
			Err:           err,
		}
	}
	out := result.(*ListEmailTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list the email templates present in your Amazon SES
// account in the current AWS Region. For more information, see the Amazon SES
// Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html).
type ListEmailTemplatesInput struct {

	// A token returned from a previous call to ListEmailTemplates to indicate the
	// position in the list of email templates.
	NextToken *string

	// The number of results to show in a single call to ListEmailTemplates. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results. The value you specify has to be at least 1, and can be no
	// more than 10.
	PageSize *int32
}

// The following elements are returned by the service.
type ListEmailTemplatesOutput struct {

	// A token indicating that there are additional email templates available to be
	// listed. Pass this token to a subsequent ListEmailTemplates call to retrieve the
	// next 10 email templates.
	NextToken *string

	// An array the contains the name and creation time stamp for each template in your
	// Amazon SES account.
	TemplatesMetadata []*types.EmailTemplateMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListEmailTemplatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListEmailTemplates{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListEmailTemplates{}, middleware.After)
}

func newServiceMetadataMiddleware_opListEmailTemplates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListEmailTemplates",
	}
}
