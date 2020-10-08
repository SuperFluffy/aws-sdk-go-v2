// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Returns the GeoMatchSet () that is specified by GeoMatchSetId.
func (c *Client) GetGeoMatchSet(ctx context.Context, params *GetGeoMatchSetInput, optFns ...func(*Options)) (*GetGeoMatchSetOutput, error) {
	stack := middleware.NewStack("GetGeoMatchSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetGeoMatchSetMiddlewares(stack)
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
	addOpGetGeoMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetGeoMatchSet(options.Region), middleware.Before)
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
			OperationName: "GetGeoMatchSet",
			Err:           err,
		}
	}
	out := result.(*GetGeoMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGeoMatchSetInput struct {

	// The GeoMatchSetId of the GeoMatchSet () that you want to get. GeoMatchSetId is
	// returned by CreateGeoMatchSet () and by ListGeoMatchSets ().
	//
	// This member is required.
	GeoMatchSetId *string
}

type GetGeoMatchSetOutput struct {

	// Information about the GeoMatchSet () that you specified in the GetGeoMatchSet
	// request. This includes the Type, which for a GeoMatchContraint is always
	// Country, as well as the Value, which is the identifier for a specific country.
	GeoMatchSet *types.GeoMatchSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetGeoMatchSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetGeoMatchSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetGeoMatchSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetGeoMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetGeoMatchSet",
	}
}
