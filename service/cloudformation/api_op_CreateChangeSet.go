// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a list of changes that will be applied to a stack so that you can review
// the changes before executing them. You can create a change set for a stack that
// doesn't exist or an existing stack. If you create a change set for a stack that
// doesn't exist, the change set shows all of the resources that AWS CloudFormation
// will create. If you create a change set for an existing stack, AWS
// CloudFormation compares the stack's information with the information that you
// submit in the change set and lists the differences. Use change sets to
// understand which resources AWS CloudFormation will create or change, and how it
// will change resources in an existing stack, before you create or update a stack.
// To create a change set for a stack that doesn't exist, for the ChangeSetType
// parameter, specify CREATE. To create a change set for an existing stack, specify
// UPDATE for the ChangeSetType parameter. To create a change set for an import
// operation, specify IMPORT for the ChangeSetType parameter. After the
// CreateChangeSet call successfully completes, AWS CloudFormation starts creating
// the change set. To check the status of the change set or to review it, use the
// DescribeChangeSet () action. When you are satisfied with the changes the change
// set will make, execute the change set by using the ExecuteChangeSet () action.
// AWS CloudFormation doesn't make changes until you execute the change set.
func (c *Client) CreateChangeSet(ctx context.Context, params *CreateChangeSetInput, optFns ...func(*Options)) (*CreateChangeSetOutput, error) {
	stack := middleware.NewStack("CreateChangeSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateChangeSetMiddlewares(stack)
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
	addOpCreateChangeSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChangeSet(options.Region), middleware.Before)
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
			OperationName: "CreateChangeSet",
			Err:           err,
		}
	}
	out := result.(*CreateChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CreateChangeSet () action.
type CreateChangeSetInput struct {

	// The name of the change set. The name must be unique among all change sets that
	// are associated with the specified stack. A change set name can contain only
	// alphanumeric, case sensitive characters and hyphens. It must start with an
	// alphabetic character and cannot exceed 128 characters.
	//
	// This member is required.
	ChangeSetName *string

	// The name or the unique ID of the stack for which you are creating a change set.
	// AWS CloudFormation generates the change set by comparing this stack's
	// information with the information that you submit, such as a modified template or
	// different parameter input values.
	//
	// This member is required.
	StackName *string

	// In some cases, you must explicitly acknowledge that your stack template contains
	// certain capabilities in order for AWS CloudFormation to create the stack.
	//
	//     *
	// CAPABILITY_IAM and CAPABILITY_NAMED_IAM Some stack templates might include
	// resources that can affect permissions in your AWS account; for example, by
	// creating new AWS Identity and Access Management (IAM) users. For those stacks,
	// you must explicitly acknowledge this by specifying one of these capabilities.
	// The following IAM resources require you to specify either the CAPABILITY_IAM or
	// CAPABILITY_NAMED_IAM capability.
	//
	//         * If you have IAM resources, you can
	// specify either capability.
	//
	//         * If you have IAM resources with custom
	// names, you must specify CAPABILITY_NAMED_IAM.
	//
	//         * If you don't specify
	// either of these capabilities, AWS CloudFormation returns an
	// InsufficientCapabilities error.
	//
	//     If your stack template contains these
	// resources, we recommend that you review all permissions associated with them and
	// edit their permissions if necessary.
	//
	//         * AWS::IAM::AccessKey
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html)
	//
	//
	// * AWS::IAM::Group
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html)
	//
	//
	// * AWS::IAM::InstanceProfile
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html)
	//
	//
	// * AWS::IAM::Policy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html)
	//
	//
	// * AWS::IAM::Role
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html)
	//
	//
	// * AWS::IAM::User
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html)
	//
	//
	// * AWS::IAM::UserToGroupAddition
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html)
	//
	//
	// For more information, see Acknowledging IAM Resources in AWS CloudFormation
	// Templates
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities).
	//
	//
	// * CAPABILITY_AUTO_EXPAND Some template contain macros. Macros perform custom
	// processing on templates; this can include simple actions like find-and-replace
	// operations, all the way to extensive transformations of entire templates.
	// Because of this, users typically create a change set from the processed
	// template, so that they can review the changes resulting from the macros before
	// actually creating the stack. If your stack template contains one or more macros,
	// and you choose to create a stack directly from the processed template, without
	// first reviewing the resulting changes in a change set, you must acknowledge this
	// capability. This includes the AWS::Include
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/create-reusable-transform-function-snippets-and-add-to-your-template-with-aws-include-transform.html)
	// and AWS::Serverless
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-aws-serverless.html)
	// transforms, which are macros hosted by AWS CloudFormation. This capacity does
	// not apply to creating change sets, and specifying it when creating change sets
	// has no effect. Also, change sets do not currently support nested stacks. If you
	// want to create a stack from a stack template that contains macros and nested
	// stacks, you must create or update the stack directly from the template using the
	// CreateStack () or UpdateStack () action, and specifying this capability. For
	// more information on macros, see Using AWS CloudFormation Macros to Perform
	// Custom Processing on Templates
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html).
	Capabilities []types.Capability

	// The type of change set operation. To create a change set for a new stack,
	// specify CREATE. To create a change set for an existing stack, specify UPDATE. To
	// create a change set for an import operation, specify IMPORT. If you create a
	// change set for a new stack, AWS Cloudformation creates a stack with a unique
	// stack ID, but no template or resources. The stack will be in the
	// REVIEW_IN_PROGRESS
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-describing-stacks.html#d0e11995)
	// state until you execute the change set. By default, AWS CloudFormation specifies
	// UPDATE. You can't use the UPDATE type to create a change set for a new stack or
	// the CREATE type to create a change set for an existing stack.
	ChangeSetType types.ChangeSetType

	// A unique identifier for this CreateChangeSet request. Specify this token if you
	// plan to retry requests so that AWS CloudFormation knows that you're not
	// attempting to create another change set with the same name. You might retry
	// CreateChangeSet requests to ensure that AWS CloudFormation successfully received
	// them.
	ClientToken *string

	// A description to help you identify this change set.
	Description *string

	// The Amazon Resource Names (ARNs) of Amazon Simple Notification Service (Amazon
	// SNS) topics that AWS CloudFormation associates with the stack. To remove all
	// associated notification topics, specify an empty list.
	NotificationARNs []*string

	// A list of Parameter structures that specify input parameters for the change set.
	// For more information, see the Parameter () data type.
	Parameters []*types.Parameter

	// The template resource types that you have permissions to work with if you
	// execute this change set, such as AWS::EC2::Instance, AWS::EC2::*, or
	// Custom::MyCustomInstance. If the list of resource types doesn't include a
	// resource type that you're updating, the stack update fails. By default, AWS
	// CloudFormation grants permissions to all resource types. AWS Identity and Access
	// Management (IAM) uses this parameter for condition keys in IAM policies for AWS
	// CloudFormation. For more information, see Controlling Access with AWS Identity
	// and Access Management
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html)
	// in the AWS CloudFormation User Guide.
	ResourceTypes []*string

	// The resources to import into your stack.
	ResourcesToImport []*types.ResourceToImport

	// The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM)
	// role that AWS CloudFormation assumes when executing the change set. AWS
	// CloudFormation uses the role's credentials to make calls on your behalf. AWS
	// CloudFormation uses this role for all future operations on the stack. As long as
	// users have permission to operate on the stack, AWS CloudFormation uses this role
	// even if the users don't have permission to pass it. Ensure that the role grants
	// least privilege. If you don't specify a value, AWS CloudFormation uses the role
	// that was previously associated with the stack. If no role is available, AWS
	// CloudFormation uses a temporary session that is generated from your user
	// credentials.
	RoleARN *string

	// The rollback triggers for AWS CloudFormation to monitor during stack creation
	// and updating operations, and for the specified monitoring period afterwards.
	RollbackConfiguration *types.RollbackConfiguration

	// Key-value pairs to associate with this stack. AWS CloudFormation also propagates
	// these tags to resources in the stack. You can specify a maximum of 50 tags.
	Tags []*types.Tag

	// A structure that contains the body of the revised template, with a minimum
	// length of 1 byte and a maximum length of 51,200 bytes. AWS CloudFormation
	// generates the change set by comparing this template with the template of the
	// stack that you specified. Conditional: You must specify only TemplateBody or
	// TemplateURL.
	TemplateBody *string

	// The location of the file that contains the revised template. The URL must point
	// to a template (max size: 460,800 bytes) that is located in an S3 bucket. AWS
	// CloudFormation generates the change set by comparing this template with the
	// stack that you specified. Conditional: You must specify only TemplateBody or
	// TemplateURL.
	TemplateURL *string

	// Whether to reuse the template that is associated with the stack to create the
	// change set.
	UsePreviousTemplate *bool
}

// The output for the CreateChangeSet () action.
type CreateChangeSetOutput struct {

	// The Amazon Resource Name (ARN) of the change set.
	Id *string

	// The unique ID of the stack.
	StackId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateChangeSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateChangeSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateChangeSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateChangeSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "CreateChangeSet",
	}
}
