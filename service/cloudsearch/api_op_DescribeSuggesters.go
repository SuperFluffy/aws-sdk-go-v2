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

// Gets the suggesters configured for a domain. A suggester enables you to display
// possible matches before users finish typing their queries. Can be limited to
// specific suggesters by name. By default, shows all suggesters and includes any
// pending changes to the configuration. Set the Deployed option to true to show
// the active configuration and exclude pending changes. For more information, see
// Getting Search Suggestions
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DescribeSuggesters(ctx context.Context, params *DescribeSuggestersInput, optFns ...func(*Options)) (*DescribeSuggestersOutput, error) {
	stack := middleware.NewStack("DescribeSuggesters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeSuggestersMiddlewares(stack)
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
	addOpDescribeSuggestersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSuggesters(options.Region), middleware.Before)
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
			OperationName: "DescribeSuggesters",
			Err:           err,
		}
	}
	out := result.(*DescribeSuggestersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeSuggester () operation. Specifies
// the name of the domain you want to describe. To restrict the response to
// particular suggesters, specify the names of the suggesters you want to describe.
// To show the active configuration and exclude any pending changes, set the
// Deployed option to true.
type DescribeSuggestersInput struct {

	// The name of the domain you want to describe.
	//
	// This member is required.
	DomainName *string

	// Whether to display the deployed configuration (true) or include any pending
	// changes (false). Defaults to false.
	Deployed *bool

	// The suggesters you want to describe.
	SuggesterNames []*string
}

// The result of a DescribeSuggesters request.
type DescribeSuggestersOutput struct {

	// The suggesters configured for the domain specified in the request.
	//
	// This member is required.
	Suggesters []*types.SuggesterStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeSuggestersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeSuggesters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeSuggesters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSuggesters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DescribeSuggesters",
	}
}
