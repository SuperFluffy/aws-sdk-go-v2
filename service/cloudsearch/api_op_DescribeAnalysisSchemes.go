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

// Gets the analysis schemes configured for a domain. An analysis scheme defines
// language-specific text processing options for a text field. Can be limited to
// specific analysis schemes by name. By default, shows all analysis schemes and
// includes any pending changes to the configuration. Set the Deployed option to
// true to show the active configuration and exclude pending changes. For more
// information, see Configuring Analysis Schemes
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-analysis-schemes.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DescribeAnalysisSchemes(ctx context.Context, params *DescribeAnalysisSchemesInput, optFns ...func(*Options)) (*DescribeAnalysisSchemesOutput, error) {
	stack := middleware.NewStack("DescribeAnalysisSchemes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeAnalysisSchemesMiddlewares(stack)
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
	addOpDescribeAnalysisSchemesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAnalysisSchemes(options.Region), middleware.Before)
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
			OperationName: "DescribeAnalysisSchemes",
			Err:           err,
		}
	}
	out := result.(*DescribeAnalysisSchemesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeAnalysisSchemes () operation.
// Specifies the name of the domain you want to describe. To limit the response to
// particular analysis schemes, specify the names of the analysis schemes you want
// to describe. To show the active configuration and exclude any pending changes,
// set the Deployed option to true.
type DescribeAnalysisSchemesInput struct {

	// The name of the domain you want to describe.
	//
	// This member is required.
	DomainName *string

	// The analysis schemes you want to describe.
	AnalysisSchemeNames []*string

	// Whether to display the deployed configuration (true) or include any pending
	// changes (false). Defaults to false.
	Deployed *bool
}

// The result of a DescribeAnalysisSchemes request. Contains the analysis schemes
// configured for the domain specified in the request.
type DescribeAnalysisSchemesOutput struct {

	// The analysis scheme descriptions.
	//
	// This member is required.
	AnalysisSchemes []*types.AnalysisSchemeStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeAnalysisSchemesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAnalysisSchemes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAnalysisSchemes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAnalysisSchemes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DescribeAnalysisSchemes",
	}
}
