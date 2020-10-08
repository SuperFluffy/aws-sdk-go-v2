// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a set of DKIM tokens for a domain identity. When you execute the
// VerifyDomainDkim operation, the domain that you specify is added to the list of
// identities that are associated with your account. This is true even if you
// haven't already associated the domain with your account by using the
// VerifyDomainIdentity operation. However, you can't send email from the domain
// until you either successfully verify it
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-domains.html) or
// you successfully set up DKIM for it
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html). You use
// the tokens that are generated by this operation to create CNAME records. When
// Amazon SES detects that you've added these records to the DNS configuration for
// a domain, you can start sending email from that domain. You can start sending
// email even if you haven't added the TXT record provided by the
// VerifyDomainIdentity operation to the DNS configuration for your domain. All
// email that you send from the domain is authenticated using DKIM. To create the
// CNAME records for DKIM authentication, use the following values:
//
//     * Name:
// token._domainkey.example.com
//
//     * Type: CNAME
//
//     * Value:
// token.dkim.amazonses.com
//
// In the preceding example, replace token with one of
// the tokens that are generated when you execute this operation. Replace
// example.com with your domain. Repeat this process for each token that's
// generated by this operation. You can execute this operation no more than once
// per second.
func (c *Client) VerifyDomainDkim(ctx context.Context, params *VerifyDomainDkimInput, optFns ...func(*Options)) (*VerifyDomainDkimOutput, error) {
	stack := middleware.NewStack("VerifyDomainDkim", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpVerifyDomainDkimMiddlewares(stack)
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
	addOpVerifyDomainDkimValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyDomainDkim(options.Region), middleware.Before)
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
			OperationName: "VerifyDomainDkim",
			Err:           err,
		}
	}
	out := result.(*VerifyDomainDkimOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to generate the CNAME records needed to set up Easy DKIM
// with Amazon SES. For more information about setting up Easy DKIM, see the Amazon
// SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).
type VerifyDomainDkimInput struct {

	// The name of the domain to be verified for Easy DKIM signing.
	//
	// This member is required.
	Domain *string
}

// Returns CNAME records that you must publish to the DNS server of your domain to
// set up Easy DKIM with Amazon SES.
type VerifyDomainDkimOutput struct {

	// A set of character strings that represent the domain's identity. If the identity
	// is an email address, the tokens represent the domain of that address. Using
	// these tokens, you need to create DNS CNAME records that point to DKIM public
	// keys that are hosted by Amazon SES. Amazon Web Services eventually detects that
	// you've updated your DNS records. This detection process might take up to 72
	// hours. After successful detection, Amazon SES is able to DKIM-sign email
	// originating from that domain. (This only applies to domain identities, not email
	// address identities.) For more information about creating DNS records using DKIM
	// tokens, see the Amazon SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).
	//
	// This member is required.
	DkimTokens []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpVerifyDomainDkimMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpVerifyDomainDkim{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpVerifyDomainDkim{}, middleware.After)
}

func newServiceMetadataMiddleware_opVerifyDomainDkim(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "VerifyDomainDkim",
	}
}
