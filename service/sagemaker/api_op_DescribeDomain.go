// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The description of the domain.
func (c *Client) DescribeDomain(ctx context.Context, params *DescribeDomainInput, optFns ...func(*Options)) (*DescribeDomainOutput, error) {
	stack := middleware.NewStack("DescribeDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDomainMiddlewares(stack)
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
	addOpDescribeDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDomain(options.Region), middleware.Before)
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
			OperationName: "DescribeDomain",
			Err:           err,
		}
	}
	out := result.(*DescribeDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDomainInput struct {

	// The domain ID.
	//
	// This member is required.
	DomainId *string
}

type DescribeDomainOutput struct {

	// The domain's authentication mode.
	AuthMode types.AuthMode

	// The creation time.
	CreationTime *time.Time

	// Settings which are applied to all UserProfile in this domain, if settings are
	// not explicitly specified in a given UserProfile.
	DefaultUserSettings *types.UserSettings

	// The domain's Amazon Resource Name (ARN).
	DomainArn *string

	// The domain ID.
	DomainId *string

	// The domain name.
	DomainName *string

	// The failure reason.
	FailureReason *string

	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEfsFileSystemId *string

	// The AWS Key Management Service encryption key ID.
	HomeEfsFileSystemKmsKeyId *string

	// The last modified time.
	LastModifiedTime *time.Time

	// The SSO managed application instance ID.
	SingleSignOnManagedApplicationInstanceId *string

	// The status.
	Status types.DomainStatus

	// Security setting to limit to a set of subnets.
	SubnetIds []*string

	// The domain's URL.
	Url *string

	// The ID of the Amazon Virtual Private Cloud.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDomain{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeDomain",
	}
}
