// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Advanced event selectors let you create fine-grained selectors for the following
// CloudTrail event record ﬁelds. They help you control costs by logging only those
// events that are important to you. For more information about advanced event
// selectors, see Logging data events for trails
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html)
// in the CloudTrail User Guide.
//
// * readOnly
//
// * eventSource
//
// * eventName
//
// *
// eventCategory
//
// * resources.type
//
// * resources.ARN
//
// You cannot apply both event
// selectors and advanced event selectors to a trail.
type AdvancedEventSelector struct {

	// Contains all selector statements in an advanced event selector.
	//
	// This member is required.
	FieldSelectors []AdvancedFieldSelector

	// An optional, descriptive name for an advanced event selector, such as "Log data
	// events for only two S3 buckets".
	Name *string

	noSmithyDocumentSerde
}

// A single selector statement in an advanced event selector.
type AdvancedFieldSelector struct {

	// A field in an event record on which to filter events to be logged. Supported
	// fields include readOnly, eventCategory, eventSource (for management events),
	// eventName, resources.type, and resources.ARN.
	//
	// * readOnly - Optional. Can be set
	// to Equals a value of true or false. If you do not add this field, CloudTrail
	// logs both both read and write events. A value of true logs only read events. A
	// value of false logs only write events.
	//
	// * eventSource - For filtering management
	// events only. This can be set only to NotEqualskms.amazonaws.com.
	//
	// * eventName -
	// Can use any operator. You can use it to ﬁlter in or ﬁlter out any data event
	// logged to CloudTrail, such as PutBucket or GetSnapshotBlock. You can have
	// multiple values for this ﬁeld, separated by commas.
	//
	// * eventCategory - This is
	// required. It must be set to Equals, and the value must be Management or Data.
	//
	// *
	// resources.type - This ﬁeld is required. resources.type can only use the Equals
	// operator, and the value can be one of the following:
	//
	// * AWS::S3::Object
	//
	// *
	// AWS::Lambda::Function
	//
	// * AWS::DynamoDB::Table
	//
	// * AWS::S3Outposts::Object
	//
	// *
	// AWS::ManagedBlockchain::Node
	//
	// * AWS::S3ObjectLambda::AccessPoint
	//
	// *
	// AWS::EC2::Snapshot
	//
	// * AWS::S3::AccessPoint
	//
	// * AWS::DynamoDB::Stream
	//
	// *
	// AWS::Glue::Table
	//
	// You can have only one resources.type ﬁeld per selector. To log
	// data events on more than one resource type, add another selector.
	//
	// *
	// resources.ARN - You can use any operator with resources.ARN, but if you use
	// Equals or NotEquals, the value must exactly match the ARN of a valid resource of
	// the type you've speciﬁed in the template as the value of resources.type. For
	// example, if resources.type equals AWS::S3::Object, the ARN must be in one of the
	// following formats. To log all data events for all objects in a specific S3
	// bucket, use the StartsWith operator, and include only the bucket ARN as the
	// matching value. The trailing slash is intentional; do not exclude it. Replace
	// the text between less than and greater than symbols (<>) with resource-specific
	// information.
	//
	// * arn::s3:::/
	//
	// * arn::s3::://
	//
	// When resources.type equals
	// AWS::S3::AccessPoint, and the operator is set to Equals or NotEquals, the ARN
	// must be in one of the following formats. To log events on all objects in an S3
	// access point, we recommend that you use only the access point ARN, don’t include
	// the object path, and use the StartsWith or NotStartsWith operators.
	//
	// *
	// arn::s3:::accesspoint/
	//
	// * arn::s3:::accesspoint//object/
	//
	// When resources.type
	// equals AWS::Lambda::Function, and the operator is set to Equals or NotEquals,
	// the ARN must be in the following format:
	//
	// * arn::lambda:::function:
	//
	// When
	// resources.type equals AWS::DynamoDB::Table, and the operator is set to Equals or
	// NotEquals, the ARN must be in the following format:
	//
	// *
	// arn::dynamodb:::table/
	//
	// When resources.type equals AWS::S3Outposts::Object, and
	// the operator is set to Equals or NotEquals, the ARN must be in the following
	// format:
	//
	// * arn::s3-outposts:::
	//
	// When resources.type equals
	// AWS::ManagedBlockchain::Node, and the operator is set to Equals or NotEquals,
	// the ARN must be in the following format:
	//
	// *
	// arn::managedblockchain:::nodes/
	//
	// When resources.type equals
	// AWS::S3ObjectLambda::AccessPoint, and the operator is set to Equals or
	// NotEquals, the ARN must be in the following format:
	//
	// *
	// arn::s3-object-lambda:::accesspoint/
	//
	// When resources.type equals
	// AWS::EC2::Snapshot, and the operator is set to Equals or NotEquals, the ARN must
	// be in the following format:
	//
	// * arn::ec2:::snapshot/
	//
	// When resources.type equals
	// AWS::DynamoDB::Stream, and the operator is set to Equals or NotEquals, the ARN
	// must be in the following format:
	//
	// * arn::dynamodb:::table//stream/
	//
	// When
	// resources.type equals AWS::Glue::Table, and the operator is set to Equals or
	// NotEquals, the ARN must be in the following format:
	//
	// * arn::glue:::table//
	//
	// This member is required.
	Field *string

	// An operator that includes events that match the last few characters of the event
	// record field specified as the value of Field.
	EndsWith []string

	// An operator that includes events that match the exact value of the event record
	// field specified as the value of Field. This is the only valid operator that you
	// can use with the readOnly, eventCategory, and resources.type fields.
	Equals []string

	// An operator that excludes events that match the last few characters of the event
	// record field specified as the value of Field.
	NotEndsWith []string

	// An operator that excludes events that match the exact value of the event record
	// field specified as the value of Field.
	NotEquals []string

	// An operator that excludes events that match the first few characters of the
	// event record field specified as the value of Field.
	NotStartsWith []string

	// An operator that includes events that match the first few characters of the
	// event record field specified as the value of Field.
	StartsWith []string

	noSmithyDocumentSerde
}

// The Amazon S3 buckets, Lambda functions, or Amazon DynamoDB tables that you
// specify in your event selectors for your trail to log data events. Data events
// provide information about the resource operations performed on or within a
// resource itself. These are also known as data plane operations. You can specify
// up to 250 data resources for a trail. The total number of allowed data resources
// is 250. This number can be distributed between 1 and 5 event selectors, but the
// total cannot exceed 250 across all selectors. If you are using advanced event
// selectors, the maximum total number of values for all conditions, across all
// advanced event selectors for the trail, is 500. The following example
// demonstrates how logging works when you configure logging of all data events for
// an S3 bucket named bucket-1. In this example, the CloudTrail user specified an
// empty prefix, and the option to log both Read and Write data events.
//
// * A user
// uploads an image file to bucket-1.
//
// * The PutObject API operation is an Amazon
// S3 object-level API. It is recorded as a data event in CloudTrail. Because the
// CloudTrail user specified an S3 bucket with an empty prefix, events that occur
// on any object in that bucket are logged. The trail processes and logs the
// event.
//
// * A user uploads an object to an Amazon S3 bucket named
// arn:aws:s3:::bucket-2.
//
// * The PutObject API operation occurred for an object in
// an S3 bucket that the CloudTrail user didn't specify for the trail. The trail
// doesn’t log the event.
//
// The following example demonstrates how logging works
// when you configure logging of Lambda data events for a Lambda function named
// MyLambdaFunction, but not for all Lambda functions.
//
// * A user runs a script that
// includes a call to the MyLambdaFunction function and the MyOtherLambdaFunction
// function.
//
// * The Invoke API operation on MyLambdaFunction is an Lambda API. It
// is recorded as a data event in CloudTrail. Because the CloudTrail user specified
// logging data events for MyLambdaFunction, any invocations of that function are
// logged. The trail processes and logs the event.
//
// * The Invoke API operation on
// MyOtherLambdaFunction is an Lambda API. Because the CloudTrail user did not
// specify logging data events for all Lambda functions, the Invoke operation for
// MyOtherLambdaFunction does not match the function specified for the trail. The
// trail doesn’t log the event.
type DataResource struct {

	// The resource type in which you want to log data events. You can specify the
	// following basic event selector resource types:
	//
	// * AWS::S3::Object
	//
	// *
	// AWS::Lambda::Function
	//
	// * AWS::DynamoDB::Table
	//
	// The following resource types are
	// also availble through advanced event selectors. Basic event selector resource
	// types are valid in advanced event selectors, but advanced event selector
	// resource types are not valid in basic event selectors. For more information, see
	// AdvancedFieldSelector$Field.
	//
	// * AWS::S3Outposts::Object
	//
	// *
	// AWS::ManagedBlockchain::Node
	//
	// * AWS::S3ObjectLambda::AccessPoint
	//
	// *
	// AWS::EC2::Snapshot
	//
	// * AWS::S3::AccessPoint
	//
	// * AWS::DynamoDB::Stream
	//
	// *
	// AWS::Glue::Table
	Type *string

	// An array of Amazon Resource Name (ARN) strings or partial ARN strings for the
	// specified objects.
	//
	// * To log data events for all objects in all S3 buckets in
	// your Amazon Web Services account, specify the prefix as arn:aws:s3:::. This also
	// enables logging of data event activity performed by any user or role in your
	// Amazon Web Services account, even if that activity is performed on a bucket that
	// belongs to another Amazon Web Services account.
	//
	// * To log data events for all
	// objects in an S3 bucket, specify the bucket and an empty object prefix such as
	// arn:aws:s3:::bucket-1/. The trail logs data events for all objects in this S3
	// bucket.
	//
	// * To log data events for specific objects, specify the S3 bucket and
	// object prefix such as arn:aws:s3:::bucket-1/example-images. The trail logs data
	// events for objects in this S3 bucket that match the prefix.
	//
	// * To log data
	// events for all Lambda functions in your Amazon Web Services account, specify the
	// prefix as arn:aws:lambda. This also enables logging of Invoke activity performed
	// by any user or role in your Amazon Web Services account, even if that activity
	// is performed on a function that belongs to another Amazon Web Services
	// account.
	//
	// * To log data events for a specific Lambda function, specify the
	// function ARN. Lambda function ARNs are exact. For example, if you specify a
	// function ARN arn:aws:lambda:us-west-2:111111111111:function:helloworld, data
	// events will only be logged for
	// arn:aws:lambda:us-west-2:111111111111:function:helloworld. They will not be
	// logged for arn:aws:lambda:us-west-2:111111111111:function:helloworld2.
	//
	// * To log
	// data events for all DynamoDB tables in your Amazon Web Services account, specify
	// the prefix as arn:aws:dynamodb.
	Values []string

	noSmithyDocumentSerde
}

// Contains information about an event that was returned by a lookup request. The
// result includes a representation of a CloudTrail event.
type Event struct {

	// The Amazon Web Services access key ID that was used to sign the request. If the
	// request was made with temporary security credentials, this is the access key ID
	// of the temporary credentials.
	AccessKeyId *string

	// A JSON string that contains a representation of the event returned.
	CloudTrailEvent *string

	// The CloudTrail ID of the event returned.
	EventId *string

	// The name of the event returned.
	EventName *string

	// The Amazon Web Services service to which the request was made.
	EventSource *string

	// The date and time of the event returned.
	EventTime *time.Time

	// Information about whether the event is a write event or a read event.
	ReadOnly *string

	// A list of resources referenced by the event returned.
	Resources []Resource

	// A user name or role name of the requester that called the API in the event
	// returned.
	Username *string

	noSmithyDocumentSerde
}

// A storage lake of event data against which you can run complex SQL-based
// queries. An event data store can include events that you have logged on your
// account from the last 90 to 2555 days (about three months to up to seven years).
// To select events for an event data store, use advanced event selectors
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced).
type EventDataStore struct {

	// The advanced event selectors that were used to select events for the data store.
	AdvancedEventSelectors []AdvancedEventSelector

	// The timestamp of the event data store's creation.
	CreatedTimestamp *time.Time

	// The ARN of the event data store.
	EventDataStoreArn *string

	// Indicates whether the event data store includes events from all regions, or only
	// from the region in which it was created.
	MultiRegionEnabled *bool

	// The name of the event data store.
	Name *string

	// Indicates that an event data store is collecting logged events for an
	// organization.
	OrganizationEnabled *bool

	// The retention period, in days.
	RetentionPeriod *int32

	// The status of an event data store. Values are ENABLED and PENDING_DELETION.
	Status EventDataStoreStatus

	// Indicates whether the event data store is protected from termination.
	TerminationProtectionEnabled *bool

	// The timestamp showing when an event data store was updated, if applicable.
	// UpdatedTimestamp is always either the same or newer than the time shown in
	// CreatedTimestamp.
	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

// Use event selectors to further specify the management and data event settings
// for your trail. By default, trails created without specific event selectors will
// be configured to log all read and write management events, and no data events.
// When an event occurs in your account, CloudTrail evaluates the event selector
// for all trails. For each trail, if the event matches any event selector, the
// trail processes and logs the event. If the event doesn't match any event
// selector, the trail doesn't log the event. You can configure up to five event
// selectors for a trail. You cannot apply both event selectors and advanced event
// selectors to a trail.
type EventSelector struct {

	// CloudTrail supports data event logging for Amazon S3 objects, Lambda functions,
	// and Amazon DynamoDB tables with basic event selectors. You can specify up to 250
	// resources for an individual event selector, but the total number of data
	// resources cannot exceed 250 across all event selectors in a trail. This limit
	// does not apply if you configure resource logging for all data events. For more
	// information, see Data Events
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-and-data-events-with-cloudtrail.html#logging-data-events)
	// and Limits in CloudTrail
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html)
	// in the CloudTrail User Guide.
	DataResources []DataResource

	// An optional list of service event sources from which you do not want management
	// events to be logged on your trail. In this release, the list can be empty
	// (disables the filter), or it can filter out Key Management Service or Amazon RDS
	// Data API events by containing kms.amazonaws.com or rdsdata.amazonaws.com. By
	// default, ExcludeManagementEventSources is empty, and KMS and Amazon RDS Data API
	// events are logged to your trail. You can exclude management event sources only
	// in regions that support the event source.
	ExcludeManagementEventSources []string

	// Specify if you want your event selector to include management events for your
	// trail. For more information, see Management Events
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-and-data-events-with-cloudtrail.html#logging-management-events)
	// in the CloudTrail User Guide. By default, the value is true. The first copy of
	// management events is free. You are charged for additional copies of management
	// events that you are logging on any subsequent trail in the same region. For more
	// information about CloudTrail pricing, see CloudTrail Pricing
	// (http://aws.amazon.com/cloudtrail/pricing/).
	IncludeManagementEvents *bool

	// Specify if you want your trail to log read-only events, write-only events, or
	// all. For example, the EC2 GetConsoleOutput is a read-only API operation and
	// RunInstances is a write-only API operation. By default, the value is All.
	ReadWriteType ReadWriteType

	noSmithyDocumentSerde
}

// A JSON string that contains a list of insight types that are logged on a trail.
type InsightSelector struct {

	// The type of insights to log on a trail. ApiCallRateInsight and
	// ApiErrorRateInsight are valid insight types.
	InsightType InsightType

	noSmithyDocumentSerde
}

// Specifies an attribute and value that filter the events returned.
type LookupAttribute struct {

	// Specifies an attribute on which to filter the events returned.
	//
	// This member is required.
	AttributeKey LookupAttributeKey

	// Specifies a value for the specified AttributeKey.
	//
	// This member is required.
	AttributeValue *string

	noSmithyDocumentSerde
}

// Contains information about a returned public key.
type PublicKey struct {

	// The fingerprint of the public key.
	Fingerprint *string

	// The ending time of validity of the public key.
	ValidityEndTime *time.Time

	// The starting time of validity of the public key.
	ValidityStartTime *time.Time

	// The DER encoded public key value in PKCS#1 format.
	Value []byte

	noSmithyDocumentSerde
}

// A SQL string of criteria about events that you want to collect in an event data
// store.
type Query struct {

	// The creation time of a query.
	CreationTime *time.Time

	// The ID of a query.
	QueryId *string

	// The status of the query. This can be QUEUED, RUNNING, FINISHED, FAILED,
	// TIMED_OUT, or CANCELLED.
	QueryStatus QueryStatus

	noSmithyDocumentSerde
}

// Metadata about a query, such as the number of results.
type QueryStatistics struct {

	// The total bytes that the query scanned in the event data store. This value
	// matches the number of bytes for which your account is billed for the query,
	// unless the query is still running.
	BytesScanned *int64

	// The number of results returned.
	ResultsCount *int32

	// The total number of results returned by a query.
	TotalResultsCount *int32

	noSmithyDocumentSerde
}

// Gets metadata about a query, including the number of events that were matched,
// the total number of events scanned, the query run time in milliseconds, and the
// query's creation time.
type QueryStatisticsForDescribeQuery struct {

	// The total bytes that the query scanned in the event data store. This value
	// matches the number of bytes for which your account is billed for the query,
	// unless the query is still running.
	BytesScanned *int64

	// The creation time of the query.
	CreationTime *time.Time

	// The number of events that matched a query.
	EventsMatched *int64

	// The number of events that the query scanned in the event data store.
	EventsScanned *int64

	// The query's run time, in milliseconds.
	ExecutionTimeInMillis *int32

	noSmithyDocumentSerde
}

// Specifies the type and name of a resource referenced by an event.
type Resource struct {

	// The name of the resource referenced by the event returned. These are
	// user-created names whose values will depend on the environment. For example, the
	// resource name might be "auto-scaling-test-group" for an Auto Scaling Group or
	// "i-1234567" for an EC2 Instance.
	ResourceName *string

	// The type of a resource referenced by the event returned. When the resource type
	// cannot be determined, null is returned. Some examples of resource types are:
	// Instance for EC2, Trail for CloudTrail, DBInstance for Amazon RDS, and AccessKey
	// for IAM. To learn more about how to look up and filter events by the resource
	// types supported for a service, see Filtering CloudTrail Events
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/view-cloudtrail-events-console.html#filtering-cloudtrail-events).
	ResourceType *string

	noSmithyDocumentSerde
}

// A resource tag.
type ResourceTag struct {

	// Specifies the ARN of the resource.
	ResourceId *string

	// A list of tags.
	TagsList []Tag

	noSmithyDocumentSerde
}

// A custom key-value pair associated with a resource such as a CloudTrail trail.
type Tag struct {

	// The key in a key-value pair. The key must be must be no longer than 128 Unicode
	// characters. The key must be unique for the resource to which it applies.
	//
	// This member is required.
	Key *string

	// The value in a key-value pair of a tag. The value must be no longer than 256
	// Unicode characters.
	Value *string

	noSmithyDocumentSerde
}

// The settings for a trail.
type Trail struct {

	// Specifies an Amazon Resource Name (ARN), a unique identifier that represents the
	// log group to which CloudTrail logs will be delivered.
	CloudWatchLogsLogGroupArn *string

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to a
	// user's log group.
	CloudWatchLogsRoleArn *string

	// Specifies if the trail has custom event selectors.
	HasCustomEventSelectors *bool

	// Specifies whether a trail has insight types specified in an InsightSelector
	// list.
	HasInsightSelectors *bool

	// The region in which the trail was created.
	HomeRegion *string

	// Set to True to include Amazon Web Services API calls from Amazon Web Services
	// global services such as IAM. Otherwise, False.
	IncludeGlobalServiceEvents *bool

	// Specifies whether the trail exists only in one region or exists in all regions.
	IsMultiRegionTrail *bool

	// Specifies whether the trail is an organization trail.
	IsOrganizationTrail *bool

	// Specifies the KMS key ID that encrypts the logs delivered by CloudTrail. The
	// value is a fully specified ARN to a KMS key in the following format.
	// arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	KmsKeyId *string

	// Specifies whether log file validation is enabled.
	LogFileValidationEnabled *bool

	// Name of the trail set by calling CreateTrail. The maximum length is 128
	// characters.
	Name *string

	// Name of the Amazon S3 bucket into which CloudTrail delivers your trail files.
	// See Amazon S3 Bucket Naming Requirements
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html).
	S3BucketName *string

	// Specifies the Amazon S3 key prefix that comes after the name of the bucket you
	// have designated for log file delivery. For more information, see Finding Your
	// CloudTrail Log Files
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html).
	// The maximum length is 200 characters.
	S3KeyPrefix *string

	// Specifies the ARN of the Amazon SNS topic that CloudTrail uses to send
	// notifications when log files are delivered. The following is the format of a
	// topic ARN. arn:aws:sns:us-east-2:123456789012:MyTopic
	SnsTopicARN *string

	// This field is no longer in use. Use SnsTopicARN.
	//
	// Deprecated: This member has been deprecated.
	SnsTopicName *string

	// Specifies the ARN of the trail. The following is the format of a trail ARN.
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	TrailARN *string

	noSmithyDocumentSerde
}

// Information about a CloudTrail trail, including the trail's name, home region,
// and Amazon Resource Name (ARN).
type TrailInfo struct {

	// The Amazon Web Services Region in which a trail was created.
	HomeRegion *string

	// The name of a trail.
	Name *string

	// The ARN of a trail.
	TrailARN *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
