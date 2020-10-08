// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets all of the variables or the specific variable. This is a paginated API.
// Providing null maxSizePerPage results in retrieving maximum of 100 records per
// page. If you provide maxSizePerPage the value must be between 50 and 100. To get
// the next page result, a provide a pagination token from GetVariablesResult as
// part of your request. Null pagination token fetches the records from the
// beginning.
func (c *Client) GetVariables(ctx context.Context, params *GetVariablesInput, optFns ...func(*Options)) (*GetVariablesOutput, error) {
	stack := middleware.NewStack("GetVariables", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetVariablesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetVariables(options.Region), middleware.Before)
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
			OperationName: "GetVariables",
			Err:           err,
		}
	}
	out := result.(*GetVariablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVariablesInput struct {

	// The max size per page determined for the get variable request.
	MaxResults *int32

	// The name of the variable.
	Name *string

	// The next page token of the get variable request.
	NextToken *string
}

type GetVariablesOutput struct {

	// The next page token to be used in subsequent requests.
	NextToken *string

	// The names of the variables returned.
	Variables []*types.Variable

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetVariablesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetVariables{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetVariables{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetVariables(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetVariables",
	}
}
