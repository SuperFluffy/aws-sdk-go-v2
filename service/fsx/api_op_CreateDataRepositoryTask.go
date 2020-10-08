// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Amazon FSx for Lustre data repository task. You use data repository
// tasks to perform bulk operations between your Amazon FSx file system and its
// linked data repository. An example of a data repository task is exporting any
// data and metadata changes, including POSIX metadata, to files, directories, and
// symbolic links (symlinks) from your FSx file system to its linked data
// repository. A CreateDataRepositoryTask operation will fail if a data repository
// is not linked to the FSx file system. To learn more about data repository tasks,
// see Using Data Repository Tasks
// (https://docs.aws.amazon.com/fsx/latest/LustreGuide/data-repository-tasks.html).
// To learn more about linking a data repository to your file system, see Setting
// the Export Prefix
// (https://docs.aws.amazon.com/fsx/latest/LustreGuide/export-data-repository.html#export-prefix).
func (c *Client) CreateDataRepositoryTask(ctx context.Context, params *CreateDataRepositoryTaskInput, optFns ...func(*Options)) (*CreateDataRepositoryTaskOutput, error) {
	stack := middleware.NewStack("CreateDataRepositoryTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDataRepositoryTaskMiddlewares(stack)
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
	addIdempotencyToken_opCreateDataRepositoryTaskMiddleware(stack, options)
	addOpCreateDataRepositoryTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataRepositoryTask(options.Region), middleware.Before)
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
			OperationName: "CreateDataRepositoryTask",
			Err:           err,
		}
	}
	out := result.(*CreateDataRepositoryTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataRepositoryTaskInput struct {

	// The globally unique ID of the file system, assigned by Amazon FSx.
	//
	// This member is required.
	FileSystemId *string

	// Defines whether or not Amazon FSx provides a CompletionReport once the task has
	// completed. A CompletionReport provides a detailed report on the files that
	// Amazon FSx processed that meet the criteria specified by the Scope parameter.
	// For more information, see Working with Task Completion Reports
	// (https://docs.aws.amazon.com/fsx/latest/LustreGuide/task-completion-report.html).
	//
	// This member is required.
	Report *types.CompletionReport

	// Specifies the type of data repository task to create.
	//
	// This member is required.
	Type types.DataRepositoryTaskType

	// (Optional) An idempotency token for resource creation, in a string of up to 64
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the AWS Command Line Interface (AWS CLI) or an AWS SDK.
	ClientRequestToken *string

	// (Optional) The path or paths on the Amazon FSx file system to use when the data
	// repository task is processed. The default path is the file system root
	// directory. The paths you provide need to be relative to the mount point of the
	// file system. If the mount point is /mnt/fsx and /mnt/fsx/path1 is a directory or
	// file on the file system you want to export, then the path to provide is path1.
	// If a path that you provide isn't valid, the task fails.
	Paths []*string

	// A list of Tag values, with a maximum of 50 elements.
	Tags []*types.Tag
}

type CreateDataRepositoryTaskOutput struct {

	// The description of the data repository task that you just created.
	DataRepositoryTask *types.DataRepositoryTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDataRepositoryTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataRepositoryTask{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataRepositoryTask{}, middleware.After)
}

type idempotencyToken_initializeOpCreateDataRepositoryTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDataRepositoryTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDataRepositoryTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDataRepositoryTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDataRepositoryTaskInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDataRepositoryTaskMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateDataRepositoryTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDataRepositoryTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "CreateDataRepositoryTask",
	}
}
