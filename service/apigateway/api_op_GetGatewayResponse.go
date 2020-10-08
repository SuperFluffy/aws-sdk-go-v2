// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a GatewayResponse () of a specified response type on the given RestApi ().
func (c *Client) GetGatewayResponse(ctx context.Context, params *GetGatewayResponseInput, optFns ...func(*Options)) (*GetGatewayResponseOutput, error) {
	stack := middleware.NewStack("GetGatewayResponse", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetGatewayResponseMiddlewares(stack)
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
	addOpGetGatewayResponseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetGatewayResponse(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "GetGatewayResponse",
			Err:           err,
		}
	}
	out := result.(*GetGatewayResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets a GatewayResponse () of a specified response type on the given RestApi ().
type GetGatewayResponseInput struct {

	// [Required] The response type of the associated GatewayResponse (). Valid values
	// are
	//
	//     * ACCESS_DENIED
	//
	//     * API_CONFIGURATION_ERROR
	//
	//     *
	// AUTHORIZER_FAILURE
	//
	//     * AUTHORIZER_CONFIGURATION_ERROR
	//
	//     *
	// BAD_REQUEST_PARAMETERS
	//
	//     * BAD_REQUEST_BODY
	//
	//     * DEFAULT_4XX
	//
	//     *
	// DEFAULT_5XX
	//
	//     * EXPIRED_TOKEN
	//
	//     * INVALID_SIGNATURE
	//
	//     *
	// INTEGRATION_FAILURE
	//
	//     * INTEGRATION_TIMEOUT
	//
	//     * INVALID_API_KEY
	//
	//     *
	// MISSING_AUTHENTICATION_TOKEN
	//
	//     * QUOTA_EXCEEDED
	//
	//     * REQUEST_TOO_LARGE
	//
	//
	// * RESOURCE_NOT_FOUND
	//
	//     * THROTTLED
	//
	//     * UNAUTHORIZED
	//
	//     *
	// UNSUPPORTED_MEDIA_TYPE
	//
	// This member is required.
	ResponseType types.GatewayResponseType

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// A gateway response of a given response type and status code, with optional
// response parameters and mapping templates. For more information about valid
// gateway response types, see Gateway Response Types Supported by API Gateway
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/supported-gateway-response-types.html)
// Example:
// Get a Gateway Response of a given response type
//
// Request
//
// This example shows how
// to get a gateway response of the MISSING_AUTHENTICATION_TOKEN type. GET
// /restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN HTTP/1.1
// Host: beta-apigateway.us-east-1.amazonaws.com Content-Type: application/json
// X-Amz-Date: 20170503T202516Z Authorization: AWS4-HMAC-SHA256
// Credential={access-key-id}/20170503/us-east-1/apigateway/aws4_request,
// SignedHeaders=content-type;host;x-amz-date,
// Signature=1b52460e3159c1a26cff29093855d50ea141c1c5b937528fecaf60f51129697a
// Cache-Control: no-cache Postman-Token: 3b2a1ce9-c848-2e26-2e2f-9c2caefbed45  The
// response type is specified as a URL path.
// Response
//
// The successful operation
// returns the 200 OK status code and a payload similar to the following: {
// "_links": { "curies": { "href":
// "http://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-gatewayresponse-{rel}.html",
// "name": "gatewayresponse", "templated": true }, "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN" },
// "gatewayresponse:delete": { "href":
// "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN" } },
// "defaultResponse": false, "responseParameters": {
// "gatewayresponse.header.x-request-path": "method.request.path.petId",
// "gatewayresponse.header.Access-Control-Allow-Origin": "'a.b.c'",
// "gatewayresponse.header.x-request-query": "method.request.querystring.q",
// "gatewayresponse.header.x-request-header": "method.request.header.Accept" },
// "responseTemplates": { "application/json": "{\n \"message\":
// $context.error.messageString,\n \"type\": \"$context.error.responseType\",\n
// \"stage\": \"$context.stage\",\n \"resourcePath\": \"$context.resourcePath\",\n
// \"stageVariables.a\": \"$stageVariables.a\",\n \"statusCode\": \"'404'\"\n}" },
// "responseType": "MISSING_AUTHENTICATION_TOKEN", "statusCode": "404" }
//     </div>
// <div class="seeAlso"> <a
// href="https://docs.aws.amazon.com/apigateway/latest/developerguide/customize-gateway-responses.html">Customize
// Gateway Responses</a> </div>
type GetGatewayResponseOutput struct {

	// A Boolean flag to indicate whether this GatewayResponse () is the default
	// gateway response (true) or not (false). A default gateway response is one
	// generated by API Gateway without any customization by an API developer.
	DefaultResponse *bool

	// Response parameters (paths, query strings and headers) of the GatewayResponse ()
	// as a string-to-string map of key-value pairs.
	ResponseParameters map[string]*string

	// Response templates of the GatewayResponse () as a string-to-string map of
	// key-value pairs.
	ResponseTemplates map[string]*string

	// The response type of the associated GatewayResponse (). Valid values are
	//
	//     *
	// ACCESS_DENIED
	//
	//     * API_CONFIGURATION_ERROR
	//
	//     * AUTHORIZER_FAILURE
	//
	//     *
	// AUTHORIZER_CONFIGURATION_ERROR
	//
	//     * BAD_REQUEST_PARAMETERS
	//
	//     *
	// BAD_REQUEST_BODY
	//
	//     * DEFAULT_4XX
	//
	//     * DEFAULT_5XX
	//
	//     * EXPIRED_TOKEN
	//
	//
	// * INVALID_SIGNATURE
	//
	//     * INTEGRATION_FAILURE
	//
	//     * INTEGRATION_TIMEOUT
	//
	//     *
	// INVALID_API_KEY
	//
	//     * MISSING_AUTHENTICATION_TOKEN
	//
	//     * QUOTA_EXCEEDED
	//
	//     *
	// REQUEST_TOO_LARGE
	//
	//     * RESOURCE_NOT_FOUND
	//
	//     * THROTTLED
	//
	//     *
	// UNAUTHORIZED
	//
	//     * UNSUPPORTED_MEDIA_TYPE
	ResponseType types.GatewayResponseType

	// The HTTP status code for this GatewayResponse ().
	StatusCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetGatewayResponseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetGatewayResponse{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGatewayResponse{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetGatewayResponse(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetGatewayResponse",
	}
}
