// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a configuration for DNS query logging. After you create a query logging
// configuration, Amazon Route 53 begins to publish log data to an Amazon
// CloudWatch Logs log group. DNS query logs contain information about the queries
// that Route 53 receives for a specified public hosted zone, such as the
// following:
//
//     * Route 53 edge location that responded to the DNS query
//
//     *
// Domain or subdomain that was requested
//
//     * DNS record type, such as A or
// AAAA
//
//     * DNS response code, such as NoError or ServFail
//
//     <dl> <dt>Log
// Group and Resource Policy</dt> <dd> <p>Before you create a query logging
// configuration, perform the following operations.</p> <note> <p>If you create a
// query logging configuration using the Route 53 console, Route 53 performs these
// operations automatically.</p> </note> <ol> <li> <p>Create a CloudWatch Logs log
// group, and make note of the ARN, which you specify when you create a query
// logging configuration. Note the following:</p> <ul> <li> <p>You must create the
// log group in the us-east-1 region.</p> </li> <li> <p>You must use the same AWS
// account to create the log group and the hosted zone that you want to configure
// query logging for.</p> </li> <li> <p>When you create log groups for query
// logging, we recommend that you use a consistent prefix, for example:</p> <p>
// <code>/aws/route53/<i>hosted zone name</i> </code> </p> <p>In the next step,
// you'll create a resource policy, which controls access to one or more log groups
// and the associated AWS resources, such as Route 53 hosted zones. There's a limit
// on the number of resource policies that you can create, so we recommend that you
// use a consistent prefix so you can use the same resource policy for all the log
// groups that you create for query logging.</p> </li> </ul> </li> <li> <p>Create a
// CloudWatch Logs resource policy, and give it the permissions that Route 53 needs
// to create log streams and to send query logs to log streams. For the value of
// <code>Resource</code>, specify the ARN for the log group that you created in the
// previous step. To use the same resource policy for all the CloudWatch Logs log
// groups that you created for query logging configurations, replace the hosted
// zone name with <code>*</code>, for example:</p> <p>
// <code>arn:aws:logs:us-east-1:123412341234:log-group:/aws/route53/*</code> </p>
// <note> <p>You can't use the CloudWatch console to create or edit a resource
// policy. You must use the CloudWatch API, one of the AWS SDKs, or the AWS
// CLI.</p> </note> </li> </ol> </dd> <dt>Log Streams and Edge Locations</dt> <dd>
// <p>When Route 53 finishes creating the configuration for DNS query logging, it
// does the following:</p> <ul> <li> <p>Creates a log stream for an edge location
// the first time that the edge location responds to DNS queries for the specified
// hosted zone. That log stream is used to log all queries that Route 53 responds
// to for that edge location.</p> </li> <li> <p>Begins to send query logs to the
// applicable log stream.</p> </li> </ul> <p>The name of each log stream is in the
// following format:</p> <p> <code> <i>hosted zone ID</i>/<i>edge location code</i>
// </code> </p> <p>The edge location code is a three-letter code and an arbitrarily
// assigned number, for example, DFW3. The three-letter code typically corresponds
// with the International Air Transport Association airport code for an airport
// near the edge location. (These abbreviations might change in the future.) For a
// list of edge locations, see "The Route 53 Global Network" on the <a
// href="http://aws.amazon.com/route53/details/">Route 53 Product Details</a>
// page.</p> </dd> <dt>Queries That Are Logged</dt> <dd> <p>Query logs contain only
// the queries that DNS resolvers forward to Route 53. If a DNS resolver has
// already cached the response to a query (such as the IP address for a load
// balancer for example.com), the resolver will continue to return the cached
// response. It doesn't forward another query to Route 53 until the TTL for the
// corresponding resource record set expires. Depending on how many DNS queries are
// submitted for a resource record set, and depending on the TTL for that resource
// record set, query logs might contain information about only one query out of
// every several thousand queries that are submitted to DNS. For more information
// about how DNS works, see <a
// href="https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/welcome-dns-service.html">Routing
// Internet Traffic to Your Website or Web Application</a> in the <i>Amazon Route
// 53 Developer Guide</i>.</p> </dd> <dt>Log File Format</dt> <dd> <p>For a list of
// the values in each query log and the format of each value, see <a
// href="https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html">Logging
// DNS Queries</a> in the <i>Amazon Route 53 Developer Guide</i>.</p> </dd>
// <dt>Pricing</dt> <dd> <p>For information about charges for query logs, see <a
// href="http://aws.amazon.com/cloudwatch/pricing/">Amazon CloudWatch
// Pricing</a>.</p> </dd> <dt>How to Stop Logging</dt> <dd> <p>If you want Route 53
// to stop sending query logs to CloudWatch Logs, delete the query logging
// configuration. For more information, see <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_DeleteQueryLoggingConfig.html">DeleteQueryLoggingConfig</a>.</p>
// </dd> </dl>
func (c *Client) CreateQueryLoggingConfig(ctx context.Context, params *CreateQueryLoggingConfigInput, optFns ...func(*Options)) (*CreateQueryLoggingConfigOutput, error) {
	stack := middleware.NewStack("CreateQueryLoggingConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpCreateQueryLoggingConfigMiddlewares(stack)
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
	addOpCreateQueryLoggingConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateQueryLoggingConfig(options.Region), middleware.Before)
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
			OperationName: "CreateQueryLoggingConfig",
			Err:           err,
		}
	}
	out := result.(*CreateQueryLoggingConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateQueryLoggingConfigInput struct {

	// The Amazon Resource Name (ARN) for the log group that you want to Amazon Route
	// 53 to send query logs to. This is the format of the ARN:
	// <p>arn:aws:logs:<i>region</i>:<i>account-id</i>:log-group:<i>log_group_name</i>
	// </p> <p>To get the ARN for a log group, you can use the CloudWatch console, the
	// <a
	// href="https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeLogGroups.html">DescribeLogGroups</a>
	// API action, the <a
	// href="https://docs.aws.amazon.com/cli/latest/reference/logs/describe-log-groups.html">describe-log-groups</a>
	// command, or the applicable command in one of the AWS SDKs.</p>
	//
	// This member is required.
	CloudWatchLogsLogGroupArn *string

	// The ID of the hosted zone that you want to log queries for. You can log queries
	// only for public hosted zones.
	//
	// This member is required.
	HostedZoneId *string
}

type CreateQueryLoggingConfigOutput struct {

	// The unique URL representing the new query logging configuration.
	//
	// This member is required.
	Location *string

	// A complex type that contains the ID for a query logging configuration, the ID of
	// the hosted zone that you want to log queries for, and the ARN for the log group
	// that you want Amazon Route 53 to send query logs to.
	//
	// This member is required.
	QueryLoggingConfig *types.QueryLoggingConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpCreateQueryLoggingConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpCreateQueryLoggingConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpCreateQueryLoggingConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateQueryLoggingConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "CreateQueryLoggingConfig",
	}
}
