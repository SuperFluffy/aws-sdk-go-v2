// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the IPSet specified by the IPSet ID.
func (c *Client) UpdateIPSet(ctx context.Context, params *UpdateIPSetInput, optFns ...func(*Options)) (*UpdateIPSetOutput, error) {
	stack := middleware.NewStack("UpdateIPSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateIPSetMiddlewares(stack)
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
	addOpUpdateIPSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIPSet(options.Region), middleware.Before)
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
			OperationName: "UpdateIPSet",
			Err:           err,
		}
	}
	out := result.(*UpdateIPSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIPSetInput struct {

	// The detectorID that specifies the GuardDuty service whose IPSet you want to
	// update.
	//
	// This member is required.
	DetectorId *string

	// The unique ID that specifies the IPSet that you want to update.
	//
	// This member is required.
	IpSetId *string

	// The updated Boolean value that specifies whether the IPSet is active or not.
	Activate *bool

	// The updated URI of the file that contains the IPSet. For example:
	// https://s3.us-west-2.amazonaws.com/my-bucket/my-object-key.
	Location *string

	// The unique ID that specifies the IPSet that you want to update.
	Name *string
}

type UpdateIPSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateIPSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIPSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIPSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateIPSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "UpdateIPSet",
	}
}
