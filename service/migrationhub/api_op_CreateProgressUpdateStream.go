// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a progress update stream which is an AWS resource used for access
// control as well as a namespace for migration task names that is implicitly
// linked to your AWS account. It must uniquely identify the migration tool as it
// is used for all updates made by the tool; however, it does not need to be unique
// for each AWS account because it is scoped to the AWS account.
func (c *Client) CreateProgressUpdateStream(ctx context.Context, params *CreateProgressUpdateStreamInput, optFns ...func(*Options)) (*CreateProgressUpdateStreamOutput, error) {
	stack := middleware.NewStack("CreateProgressUpdateStream", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateProgressUpdateStreamMiddlewares(stack)
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
	addOpCreateProgressUpdateStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProgressUpdateStream(options.Region), middleware.Before)
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
			OperationName: "CreateProgressUpdateStream",
			Err:           err,
		}
	}
	out := result.(*CreateProgressUpdateStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProgressUpdateStreamInput struct {

	// The name of the ProgressUpdateStream. Do not store personal data in this field.
	//
	// This member is required.
	ProgressUpdateStreamName *string

	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun *bool
}

type CreateProgressUpdateStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateProgressUpdateStreamMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProgressUpdateStream{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProgressUpdateStream{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateProgressUpdateStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "CreateProgressUpdateStream",
	}
}
