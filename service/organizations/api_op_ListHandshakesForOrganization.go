// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the handshakes that are associated with the organization that the
// requesting user is part of. The ListHandshakesForOrganization operation returns
// a list of handshake structures. Each structure contains details and status about
// a handshake. Handshakes that are ACCEPTED, DECLINED, or CANCELED appear in the
// results of this API for only 30 days after changing to that state. After that,
// they're deleted and no longer accessible. Always check the NextToken response
// parameter for a null value when calling a List* operation. These operations can
// occasionally return an empty set of results even when there are more results
// available. The NextToken response parameter value is null only when there are no
// more results to display. This operation can be called only from the
// organization's master account or by a member account that is a delegated
// administrator for an AWS service.
func (c *Client) ListHandshakesForOrganization(ctx context.Context, params *ListHandshakesForOrganizationInput, optFns ...func(*Options)) (*ListHandshakesForOrganizationOutput, error) {
	stack := middleware.NewStack("ListHandshakesForOrganization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListHandshakesForOrganizationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListHandshakesForOrganization(options.Region), middleware.Before)
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
			OperationName: "ListHandshakesForOrganization",
			Err:           err,
		}
	}
	out := result.(*ListHandshakesForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHandshakesForOrganizationInput struct {

	// A filter of the handshakes that you want included in the response. The default
	// is all types. Use the ActionType element to limit the output to only a specified
	// type, such as INVITE, ENABLE-ALL-FEATURES, or APPROVE-ALL-FEATURES.
	// Alternatively, for the ENABLE-ALL-FEATURES handshake that generates a separate
	// child handshake for each member account, you can specify the ParentHandshakeId
	// to see only the handshakes that were generated by that parent request.
	Filter *types.HandshakeFilter

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string
}

type ListHandshakesForOrganizationOutput struct {

	// A list of Handshake () objects with details about each of the handshakes that
	// are associated with an organization.
	Handshakes []*types.Handshake

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListHandshakesForOrganizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListHandshakesForOrganization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHandshakesForOrganization{}, middleware.After)
}

func newServiceMetadataMiddleware_opListHandshakesForOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListHandshakesForOrganization",
	}
}
