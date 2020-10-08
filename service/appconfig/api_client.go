// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/awslabs/smithy-go/middleware"
	"net/http"
	"time"
)

const ServiceID = "AppConfig"
const ServiceAPIVersion = "2019-10-09"

// AWS AppConfig Use AWS AppConfig, a capability of AWS Systems Manager, to create,
// manage, and quickly deploy application configurations. AppConfig supports
// controlled deployments to applications of any size and includes built-in
// validation checks and monitoring. You can use AppConfig with applications hosted
// on Amazon EC2 instances, AWS Lambda, containers, mobile applications, or IoT
// devices.  <p>To prevent errors when deploying application configurations,
// especially for production systems where a simple typo could cause an unexpected
// outage, AppConfig includes validators. A validator provides a syntactic or
// semantic check to ensure that the configuration you want to deploy works as
// intended. To validate your application configuration data, you provide a schema
// or a Lambda function that runs against the configuration. The configuration
// deployment or update can only proceed when the configuration data is valid.</p>
// <p>During a configuration deployment, AppConfig monitors the application to
// ensure that the deployment is successful. If the system encounters an error,
// AppConfig rolls back the change to minimize impact for your application users.
// You can configure a deployment strategy for each application or environment that
// includes deployment criteria, including velocity, bake time, and alarms to
// monitor. Similar to error monitoring, if a deployment triggers an alarm,
// AppConfig automatically rolls back to the previous version. </p> <p>AppConfig
// supports multiple use cases. Here are some examples.</p> <ul> <li> <p>
// <b>Application tuning</b>: Use AppConfig to carefully introduce changes to your
// application that can only be tested with production traffic.</p> </li> <li> <p>
// <b>Feature toggle</b>: Use AppConfig to turn on new features that require a
// timely deployment, such as a product launch or announcement. </p> </li> <li> <p>
// <b>Allow list</b>: Use AppConfig to allow premium subscribers to access paid
// content. </p> </li> <li> <p> <b>Operational issues</b>: Use AppConfig to reduce
// stress on your application when a dependency or other external factor impacts
// the system.</p> </li> </ul> <p>This reference is intended to be used with the <a
// href="http://docs.aws.amazon.com/systems-manager/latest/userguide/appconfig.html">AWS
// AppConfig User Guide</a>.</p>
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
	awsmiddleware.AddUserAgentKey("appconfig")(stack)
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

func addRetryMiddlewares(stack *middleware.Stack, o Options) error {
	mo := retry.AddRetryMiddlewaresOptions{
		Retryer: o.Retryer,
	}
	return retry.AddRetryMiddlewares(stack, mo)
}

func addRequestIDRetrieverMiddleware(stack *middleware.Stack) {
	awsmiddleware.AddRequestIDRetrieverMiddleware(stack)
}

func addResponseErrorMiddleware(stack *middleware.Stack) {
	awshttp.AddResponseErrorMiddleware(stack)
}
