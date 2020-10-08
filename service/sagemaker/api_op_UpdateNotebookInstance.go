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
)

// Updates a notebook instance. NotebookInstance updates include upgrading or
// downgrading the ML compute instance used for your notebook instance to
// accommodate changes in your workload requirements.
func (c *Client) UpdateNotebookInstance(ctx context.Context, params *UpdateNotebookInstanceInput, optFns ...func(*Options)) (*UpdateNotebookInstanceOutput, error) {
	stack := middleware.NewStack("UpdateNotebookInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateNotebookInstanceMiddlewares(stack)
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
	addOpUpdateNotebookInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNotebookInstance(options.Region), middleware.Before)
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
			OperationName: "UpdateNotebookInstance",
			Err:           err,
		}
	}
	out := result.(*UpdateNotebookInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNotebookInstanceInput struct {

	// The name of the notebook instance to update.
	//
	// This member is required.
	NotebookInstanceName *string

	// A list of the Elastic Inference (EI) instance types to associate with this
	// notebook instance. Currently only one EI instance type can be associated with a
	// notebook instance. For more information, see Using Elastic Inference in Amazon
	// SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html).
	AcceleratorTypes []types.NotebookInstanceAcceleratorType

	// An array of up to three Git repositories to associate with the notebook
	// instance. These can be either the names of Git repositories stored as resources
	// in your account, or the URL of Git repositories in AWS CodeCommit
	// (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any
	// other Git repository. These repositories are cloned at the same level as the
	// default repository of your notebook instance. For more information, see
	// Associating Git Repositories with Amazon SageMaker Notebook Instances
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	AdditionalCodeRepositories []*string

	// The Git repository to associate with the notebook instance as its default code
	// repository. This can be either the name of a Git repository stored as a resource
	// in your account, or the URL of a Git repository in AWS CodeCommit
	// (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any
	// other Git repository. When you open a notebook instance, it opens in the
	// directory that contains this repository. For more information, see Associating
	// Git Repositories with Amazon SageMaker Notebook Instances
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	DefaultCodeRepository *string

	// A list of the Elastic Inference (EI) instance types to remove from this notebook
	// instance. This operation is idempotent. If you specify an accelerator type that
	// is not associated with the notebook instance when you call this method, it does
	// not throw an error.
	DisassociateAcceleratorTypes *bool

	// A list of names or URLs of the default Git repositories to remove from this
	// notebook instance. This operation is idempotent. If you specify a Git repository
	// that is not associated with the notebook instance when you call this method, it
	// does not throw an error.
	DisassociateAdditionalCodeRepositories *bool

	// The name or URL of the default Git repository to remove from this notebook
	// instance. This operation is idempotent. If you specify a Git repository that is
	// not associated with the notebook instance when you call this method, it does not
	// throw an error.
	DisassociateDefaultCodeRepository *bool

	// Set to true to remove the notebook instance lifecycle configuration currently
	// associated with the notebook instance. This operation is idempotent. If you
	// specify a lifecycle configuration that is not associated with the notebook
	// instance when you call this method, it does not throw an error.
	DisassociateLifecycleConfig *bool

	// The Amazon ML compute instance type.
	InstanceType types.InstanceType

	// The name of a lifecycle configuration to associate with the notebook instance.
	// For information about lifestyle configurations, see Step 2.1: (Optional)
	// Customize a Notebook Instance
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html).
	LifecycleConfigName *string

	// The Amazon Resource Name (ARN) of the IAM role that Amazon SageMaker can assume
	// to access the notebook instance. For more information, see Amazon SageMaker
	// Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html). To
	// be able to pass this role to Amazon SageMaker, the caller of this API must have
	// the iam:PassRole permission.
	RoleArn *string

	// Whether root access is enabled or disabled for users of the notebook instance.
	// The default value is Enabled. If you set this to Disabled, users don't have root
	// access on the notebook instance, but lifecycle configuration scripts still run
	// with root permissions.
	RootAccess types.RootAccess

	// The size, in GB, of the ML storage volume to attach to the notebook instance.
	// The default value is 5 GB. ML storage volumes are encrypted, so Amazon SageMaker
	// can't determine the amount of available free space on the volume. Because of
	// this, you can increase the volume size when you update a notebook instance, but
	// you can't decrease the volume size. If you want to decrease the size of the ML
	// storage volume in use, create a new notebook instance with the desired size.
	VolumeSizeInGB *int32
}

type UpdateNotebookInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateNotebookInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNotebookInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNotebookInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateNotebookInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateNotebookInstance",
	}
}
