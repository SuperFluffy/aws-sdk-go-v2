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
// global use. Creates a ByteMatchSet. You then use UpdateByteMatchSet () to
// identify the part of a web request that you want AWS WAF to inspect, such as the
// values of the User-Agent header or the query string. For example, you can create
// a ByteMatchSet that matches any requests with User-Agent headers that contain
// the string BadBot. You can then configure AWS WAF to reject those requests. To
// create and configure a ByteMatchSet, perform the following steps:
//
//     * Use
// GetChangeToken () to get the change token that you provide in the ChangeToken
// parameter of a CreateByteMatchSet request.
//
//     * Submit a CreateByteMatchSet
// request.
//
//     * Use GetChangeToken to get the change token that you provide in
// the ChangeToken parameter of an UpdateByteMatchSet request.
//
//     * Submit an
// UpdateByteMatchSet () request to specify the part of the request that you want
// AWS WAF to inspect (for example, the header or the URI) and the value that you
// want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF
// API to allow or block HTTP requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateByteMatchSet(ctx context.Context, params *CreateByteMatchSetInput, optFns ...func(*Options)) (*CreateByteMatchSetOutput, error) {
	stack := middleware.NewStack("CreateByteMatchSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateByteMatchSetMiddlewares(stack)
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
	addOpCreateByteMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateByteMatchSet(options.Region), middleware.Before)
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
			OperationName: "CreateByteMatchSet",
			Err:           err,
		}
	}
	out := result.(*CreateByteMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateByteMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken ().
	//
	// This member is required.
	ChangeToken *string

	// A friendly name or description of the ByteMatchSet (). You can't change Name
	// after you create a ByteMatchSet.
	//
	// This member is required.
	Name *string
}

type CreateByteMatchSetOutput struct {

	// A ByteMatchSet () that contains no ByteMatchTuple objects.
	ByteMatchSet *types.ByteMatchSet

	// The ChangeToken that you used to submit the CreateByteMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateByteMatchSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateByteMatchSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateByteMatchSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateByteMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "CreateByteMatchSet",
	}
}
