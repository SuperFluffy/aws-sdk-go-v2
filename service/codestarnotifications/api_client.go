// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	cryptorand "crypto/rand"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/awslabs/smithy-go/middleware"
	smithyrand "github.com/awslabs/smithy-go/rand"
	"net/http"
	"time"
)

const ServiceID = "codestar notifications"
const ServiceAPIVersion = "2019-10-15"

// This AWS CodeStar Notifications API Reference provides descriptions and usage
// examples of the operations and data types for the AWS CodeStar Notifications
// API. You can use the AWS CodeStar Notifications API to work with the following
// objects:  <p>Notification rules, by calling the following: </p> <ul> <li> <p>
// <a>CreateNotificationRule</a>, which creates a notification rule for a resource
// in your account. </p> </li> <li> <p> <a>DeleteNotificationRule</a>, which
// deletes a notification rule. </p> </li> <li> <p>
// <a>DescribeNotificationRule</a>, which provides information about a notification
// rule. </p> </li> <li> <p> <a>ListNotificationRules</a>, which lists the
// notification rules associated with your account. </p> </li> <li> <p>
// <a>UpdateNotificationRule</a>, which changes the name, events, or targets
// associated with a notification rule. </p> </li> <li> <p> <a>Subscribe</a>, which
// subscribes a target to a notification rule. </p> </li> <li> <p>
// <a>Unsubscribe</a>, which removes a target from a notification rule. </p> </li>
// </ul> <p>Targets, by calling the following: </p> <ul> <li> <p>
// <a>DeleteTarget</a>, which removes a notification rule target (SNS topic) from a
// notification rule. </p> </li> <li> <p> <a>ListTargets</a>, which lists the
// targets associated with a notification rule. </p> </li> </ul> <p>Events, by
// calling the following: </p> <ul> <li> <p> <a>ListEventTypes</a>, which lists the
// event types you can include in a notification rule. </p> </li> </ul> <p>Tags, by
// calling the following: </p> <ul> <li> <p> <a>ListTagsForResource</a>, which
// lists the tags already associated with a notification rule in your account. </p>
// </li> <li> <p> <a>TagResource</a>, which associates a tag you provide with a
// notification rule in your account. </p> </li> <li> <p> <a>UntagResource</a>,
// which removes a tag from a notification rule in your account. </p> </li> </ul>
// <p> For information about how to use AWS CodeStar Notifications, see link in the
// CodeStarNotifications User Guide. </p>
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	resolveIdempotencyTokenProvider(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// Provides idempotency tokens values that will be automatically populated into
	// idempotent API operations.
	IdempotencyTokenProvider IdempotencyTokenProvider

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
		APIOptions:  cfg.APIOptions,
	}
	resolveAWSEndpointResolver(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil {
		return
	}
	o.EndpointResolver = WithEndpointResolver(cfg.EndpointResolver, NewDefaultEndpointResolver())
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("codestarnotifications")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner()
}

func resolveIdempotencyTokenProvider(o *Options) {
	if o.IdempotencyTokenProvider != nil {
		return
	}
	o.IdempotencyTokenProvider = smithyrand.NewUUIDIdempotencyToken(cryptorand.Reader)
}

func addRetryMiddlewares(stack *middleware.Stack, o Options) error {
	mo := retry.AddRetryMiddlewaresOptions{
		Retryer: o.Retryer,
	}
	return retry.AddRetryMiddlewares(stack, mo)
}

// IdempotencyTokenProvider interface for providing idempotency token
type IdempotencyTokenProvider interface {
	GetIdempotencyToken() (string, error)
}

func addRequestIDRetrieverMiddleware(stack *middleware.Stack) {
	awsmiddleware.AddRequestIDRetrieverMiddleware(stack)
}

func addResponseErrorMiddleware(stack *middleware.Stack) {
	awshttp.AddResponseErrorMiddleware(stack)
}
