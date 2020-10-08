// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
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
// global use. Inserts or deletes XssMatchTuple () objects (filters) in an
// XssMatchSet (). For each XssMatchTuple object, you specify the following
// values:
//
//     * Action: Whether to insert the object into or delete the object
// from the array. To change an XssMatchTuple, you delete the existing object and
// add a new one.
//
//     * FieldToMatch: The part of web requests that you want AWS
// WAF to inspect and, if you want AWS WAF to inspect a header or custom query
// parameter, the name of the header or parameter.
//
//     * TextTransformation: Which
// text transformation, if any, to perform on the web request before inspecting the
// request for cross-site scripting attacks. You can only specify a single type of
// TextTransformation.
//
// You use XssMatchSet objects to specify which CloudFront
// requests that you want to allow, block, or count. For example, if you're
// receiving requests that contain cross-site scripting attacks in the request body
// and you want to block the requests, you can create an XssMatchSet with the
// applicable settings, and then configure AWS WAF to block the requests. To create
// and configure an XssMatchSet, perform the following steps:
//
//     * Submit a
// CreateXssMatchSet () request.
//
//     * Use GetChangeToken () to get the change
// token that you provide in the ChangeToken parameter of an UpdateIPSet ()
// request.
//
//     * Submit an UpdateXssMatchSet request to specify the parts of web
// requests that you want AWS WAF to inspect for cross-site scripting attacks.
//
// For
// more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) UpdateXssMatchSet(ctx context.Context, params *UpdateXssMatchSetInput, optFns ...func(*Options)) (*UpdateXssMatchSetOutput, error) {
	stack := middleware.NewStack("UpdateXssMatchSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateXssMatchSetMiddlewares(stack)
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
	addOpUpdateXssMatchSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateXssMatchSet(options.Region), middleware.Before)
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
			OperationName: "UpdateXssMatchSet",
			Err:           err,
		}
	}
	out := result.(*UpdateXssMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update an XssMatchSet ().
type UpdateXssMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken ().
	//
	// This member is required.
	ChangeToken *string

	// An array of XssMatchSetUpdate objects that you want to insert into or delete
	// from an XssMatchSet (). For more information, see the applicable data types:
	//
	//
	// * XssMatchSetUpdate (): Contains Action and XssMatchTuple
	//
	//     * XssMatchTuple
	// (): Contains FieldToMatch and TextTransformation
	//
	//     * FieldToMatch ():
	// Contains Data and Type
	//
	// This member is required.
	Updates []*types.XssMatchSetUpdate

	// The XssMatchSetId of the XssMatchSet that you want to update. XssMatchSetId is
	// returned by CreateXssMatchSet () and by ListXssMatchSets ().
	//
	// This member is required.
	XssMatchSetId *string
}

// The response to an UpdateXssMatchSets () request.
type UpdateXssMatchSetOutput struct {

	// The ChangeToken that you used to submit the UpdateXssMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus ().
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateXssMatchSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateXssMatchSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateXssMatchSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateXssMatchSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "UpdateXssMatchSet",
	}
}
