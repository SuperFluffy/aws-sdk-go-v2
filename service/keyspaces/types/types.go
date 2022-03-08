// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Amazon Keyspaces has two read/write capacity modes for processing reads and
// writes on your tables:
//
// * On-demand (default)
//
// * Provisioned
//
// The read/write
// capacity mode that you choose controls how you are charged for read and write
// throughput and how table throughput capacity is managed. For more information,
// see Read/write capacity modes
// (https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html)
// in the Amazon Keyspaces Developer Guide.
type CapacitySpecification struct {

	// The read/write throughput capacity mode for a table. The options are:
	//
	// *
	// throughputMode:PAY_PER_REQUEST and
	//
	// * throughputMode:PROVISIONED. The
	// provisioned capacity mode requires readCapacityUnits and writeCapacityUnits as
	// inputs.
	//
	// The default is throughput_mode:PAY_PER_REQUEST. For more information,
	// see Read/write capacity modes
	// (https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html)
	// in the Amazon Keyspaces Developer Guide.
	//
	// This member is required.
	ThroughputMode ThroughputMode

	// The throughput capacity specified for read operations defined in read capacity
	// units(RCUs).
	ReadCapacityUnits *int64

	// The throughput capacity specified for write operations defined in write capacity
	// units(WCUs).
	WriteCapacityUnits *int64

	noSmithyDocumentSerde
}

// The read/write throughput capacity mode for a table. The options are:
//
// *
// throughputMode:PAY_PER_REQUEST and
//
// * throughputMode:PROVISIONED.
//
// For more
// information, see Read/write capacity modes
// (https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html)
// in the Amazon Keyspaces Developer Guide.
type CapacitySpecificationSummary struct {

	// The read/write throughput capacity mode for a table. The options are:
	//
	// *
	// throughputMode:PAY_PER_REQUEST and
	//
	// * throughputMode:PROVISIONED. The
	// provisioned capacity mode requires readCapacityUnits and writeCapacityUnits as
	// inputs.
	//
	// The default is throughput_mode:PAY_PER_REQUEST. For more information,
	// see Read/write capacity modes
	// (https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html)
	// in the Amazon Keyspaces Developer Guide.
	//
	// This member is required.
	ThroughputMode ThroughputMode

	// The timestamp of the last operation that changed the provisioned throughput
	// capacity of a table.
	LastUpdateToPayPerRequestTimestamp *time.Time

	// The throughput capacity specified for read operations defined in read capacity
	// units(RCUs).
	ReadCapacityUnits *int64

	// The throughput capacity specified for write operations defined in write capacity
	// units(WCUs).
	WriteCapacityUnits *int64

	noSmithyDocumentSerde
}

// The optional clustering column portion of your primary key determines how the
// data is clustered and sorted within each partition.
type ClusteringKey struct {

	// The name(s) of the clustering column(s).
	//
	// This member is required.
	Name *string

	// Sets the ascendant (ASC) or descendant (DESC) order modifier.
	//
	// This member is required.
	OrderBy SortOrder

	noSmithyDocumentSerde
}

// The names and data types of regular columns.
type ColumnDefinition struct {

	// The name of the column.
	//
	// This member is required.
	Name *string

	// The data type of the column. For a list of available data types, see Data types
	// (https://docs.aws.amazon.com/keyspaces/latest/devguide/cql.elements.html#cql.data-types)
	// in the Amazon Keyspaces Developer Guide.
	//
	// This member is required.
	Type *string

	noSmithyDocumentSerde
}

// An optional comment that describes the table.
type Comment struct {

	// An optional description of the table.
	//
	// This member is required.
	Message *string

	noSmithyDocumentSerde
}

// Amazon Keyspaces encrypts and decrypts the table data at rest transparently and
// integrates with Key Management Service for storing and managing the encryption
// key. You can choose one of the following KMS keys (KMS keys):
//
// * Amazon Web
// Services owned key - This is the default encryption type. The key is owned by
// Amazon Keyspaces (no additional charge).
//
// * Customer managed key - This key is
// stored in your account and is created, owned, and managed by you. You have full
// control over the customer managed key (KMS charges apply).
//
// For more information
// about encryption at rest in Amazon Keyspaces, see Encryption at rest
// (https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html) in
// the Amazon Keyspaces Developer Guide. For more information about KMS, see KMS
// management service concepts
// (https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html) in
// the Key Management Service Developer Guide.
type EncryptionSpecification struct {

	// The encryption option specified for the table. You can choose one of the
	// following KMS keys (KMS keys):
	//
	// * type:AWS_OWNED_KMS_KEY - This key is owned by
	// Amazon Keyspaces.
	//
	// * type:CUSTOMER_MANAGED_KMS_KEY - This key is stored in your
	// account and is created, owned, and managed by you. This option requires the
	// kms_key_identifier of the KMS key in Amazon Resource Name (ARN) format as
	// input.
	//
	// The default is type:AWS_OWNED_KMS_KEY. For more information, see
	// Encryption at rest
	// (https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html) in
	// the Amazon Keyspaces Developer Guide.
	//
	// This member is required.
	Type EncryptionType

	// The Amazon Resource Name (ARN) of the customer managed KMS key, for example
	// kms_key_identifier:ARN.
	KmsKeyIdentifier *string

	noSmithyDocumentSerde
}

// Represents the properties of a keyspace.
type KeyspaceSummary struct {

	// The name of the keyspace.
	//
	// This member is required.
	KeyspaceName *string

	// The unique identifier of the keyspace in the format of an Amazon Resource Name
	// (ARN).
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

// The partition key portion of the primary key is required and determines how
// Amazon Keyspaces stores the data. The partition key can be a single column, or
// it can be a compound value composed of two or more columns.
type PartitionKey struct {

	// The name(s) of the partition key column(s).
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Point-in-time recovery (PITR) helps protect your Amazon Keyspaces tables from
// accidental write or delete operations by providing you continuous backups of
// your table data. For more information, see Point-in-time recovery
// (https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html)
// in the Amazon Keyspaces Developer Guide.
type PointInTimeRecovery struct {

	// The options are:
	//
	// * ENABLED
	//
	// * DISABLED
	//
	// This member is required.
	Status PointInTimeRecoveryStatus

	noSmithyDocumentSerde
}

// The point-in-time recovery status of the specified table.
type PointInTimeRecoverySummary struct {

	// Shows if point-in-time recovery is enabled or disabled for the specified table.
	//
	// This member is required.
	Status PointInTimeRecoveryStatus

	// Specifies the earliest possible restore point of the table in ISO 8601 format.
	EarliestRestorableTimestamp *time.Time

	noSmithyDocumentSerde
}

// Describes the schema of the table.
type SchemaDefinition struct {

	// The regular columns of the table.
	//
	// This member is required.
	AllColumns []ColumnDefinition

	// The columns that are part of the partition key of the table .
	//
	// This member is required.
	PartitionKeys []PartitionKey

	// The columns that are part of the clustering key of the table.
	ClusteringKeys []ClusteringKey

	// The columns that have been defined as STATIC. Static columns store values that
	// are shared by all rows in the same partition.
	StaticColumns []StaticColumn

	noSmithyDocumentSerde
}

// The static columns of the table. Static columns store values that are shared by
// all rows in the same partition.
type StaticColumn struct {

	// The name of the static column.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Returns the name of the specified table, the keyspace it is stored in, and the
// unique identifier in the format of an Amazon Resource Name (ARN).
type TableSummary struct {

	// The name of the keyspace that the table is stored in.
	//
	// This member is required.
	KeyspaceName *string

	// The unique identifier of the table in the format of an Amazon Resource Name
	// (ARN).
	//
	// This member is required.
	ResourceArn *string

	// The name of the table.
	//
	// This member is required.
	TableName *string

	noSmithyDocumentSerde
}

// Describes a tag. A tag is a key-value pair. You can add up to 50 tags to a
// single Amazon Keyspaces resource. Amazon Web Services-assigned tag names and
// values are automatically assigned the aws: prefix, which the user cannot assign.
// Amazon Web Services-assigned tag names do not count towards the tag limit of 50.
// User-assigned tag names have the prefix user: in the Cost Allocation Report. You
// cannot backdate the application of a tag. For more information, see Adding tags
// and labels to Amazon Keyspaces resources
// (https://docs.aws.amazon.com/keyspaces/latest/devguide/tagging-keyspaces.html)
// in the Amazon Keyspaces Developer Guide.
type Tag struct {

	// The key of the tag. Tag keys are case sensitive. Each Amazon Keyspaces resource
	// can only have up to one tag with the same key. If you try to add an existing tag
	// (same key), the existing tag value will be updated to the new value.
	//
	// This member is required.
	Key *string

	// The value of the tag. Tag values are case-sensitive and can be null.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Enable custom Time to Live (TTL) settings for rows and columns without setting a
// TTL default for the specified table. For more information, see Enabling TTL on
// tables
// (https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_enabling)
// in the Amazon Keyspaces Developer Guide.
type TimeToLive struct {

	// Shows how to enable custom Time to Live (TTL) settings for the specified table.
	//
	// This member is required.
	Status TimeToLiveStatus

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
