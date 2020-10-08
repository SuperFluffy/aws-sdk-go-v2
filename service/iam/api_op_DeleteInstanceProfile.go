// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified instance profile. The instance profile must not have an
// associated role. Make sure that you do not have any Amazon EC2 instances running
// with the instance profile you are about to delete. Deleting a role or instance
// profile that is associated with a running instance will break any applications
// running on the instance. For more information about instance profiles, go to
// About Instance Profiles
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/AboutInstanceProfiles.html).
func (c *Client) DeleteInstanceProfile(ctx context.Context, params *DeleteInstanceProfileInput, optFns ...func(*Options)) (*DeleteInstanceProfileOutput, error) {
	stack := middleware.NewStack("DeleteInstanceProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteInstanceProfileMiddlewares(stack)
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
	addOpDeleteInstanceProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInstanceProfile(options.Region), middleware.Before)
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
			OperationName: "DeleteInstanceProfile",
			Err:           err,
		}
	}
	out := result.(*DeleteInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteInstanceProfileInput struct {

	// The name of the instance profile to delete. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex)) a string of characters
	// consisting of upper and lowercase alphanumeric characters with no spaces. You
	// can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	InstanceProfileName *string
}

type DeleteInstanceProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteInstanceProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteInstanceProfile{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteInstanceProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteInstanceProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteInstanceProfile",
	}
}
