// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This action is part of Amazon GameLift FleetIQ with game server groups, which is
// in preview release and is subject to change. Creates a GameLift FleetIQ game
// server group to manage a collection of EC2 instances for game hosting. In
// addition to creating the game server group, this action also creates an Auto
// Scaling group in your AWS account and establishes a link between the two groups.
// You have full control over configuration of the Auto Scaling group, but GameLift
// FleetIQ routinely certain Auto Scaling group properties in order to optimize the
// group's instances for low-cost game hosting. You can view the status of your
// game server groups in the GameLift Console. Game server group metrics and events
// are emitted to Amazon CloudWatch. Prior creating a new game server group, you
// must set up the following:
//
//     * An EC2 launch template. The template provides
// configuration settings for a set of EC2 instances and includes the game server
// build that you want to deploy and run on each instance. For more information on
// creating a launch template, see  Launching an Instance from a Launch Template
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html)
// in the Amazon EC2 User Guide.
//
//     * An IAM role. The role sets up limited
// access to your AWS account, allowing GameLift FleetIQ to create and manage the
// EC2 Auto Scaling group, get instance data, and emit metrics and events to
// CloudWatch. For more information on setting up an IAM permissions policy with
// principal access for GameLift, see  Specifying a Principal in a Policy
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-bucket-user-policy-specifying-principal-intro.html)
// in the Amazon S3 Developer Guide.
//
// To create a new game server group, provide a
// name and specify the IAM role and EC2 launch template. You also need to provide
// a list of instance types to be used in the group and set initial maximum and
// minimum limits on the group's instance count. You can optionally set an
// autoscaling policy with target tracking based on a GameLift FleetIQ metric. Once
// the game server group and corresponding Auto Scaling group are created, you have
// full access to change the Auto Scaling group's configuration as needed. Keep in
// mind, however, that some properties are periodically updated by GameLift FleetIQ
// as it balances the group's instances based on availability and cost. Learn more
// GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-intro.html)Updating
// a GameLift FleetIQ-Linked Auto Scaling Group
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-asgroups.html)
// Related operations
//
//     * CreateGameServerGroup ()
//
//     * ListGameServerGroups
// ()
//
//     * DescribeGameServerGroup ()
//
//     * UpdateGameServerGroup ()
//
//     *
// DeleteGameServerGroup ()
//
//     * ResumeGameServerGroup ()
//
//     *
// SuspendGameServerGroup ()
func (c *Client) CreateGameServerGroup(ctx context.Context, params *CreateGameServerGroupInput, optFns ...func(*Options)) (*CreateGameServerGroupOutput, error) {
	stack := middleware.NewStack("CreateGameServerGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateGameServerGroupMiddlewares(stack)
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
	addOpCreateGameServerGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGameServerGroup(options.Region), middleware.Before)
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
			OperationName: "CreateGameServerGroup",
			Err:           err,
		}
	}
	out := result.(*CreateGameServerGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGameServerGroupInput struct {

	// An identifier for the new game server group. This value is used to generate
	// unique ARN identifiers for the EC2 Auto Scaling group and the GameLift FleetIQ
	// game server group. The name must be unique per Region per AWS account.
	//
	// This member is required.
	GameServerGroupName *string

	// A set of EC2 instance types to use when creating instances in the group. The
	// instance definitions must specify at least two different instance types that are
	// supported by GameLift FleetIQ. For more information on instance types, see EC2
	// Instance Types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the
	// Amazon EC2 User Guide.
	//
	// This member is required.
	InstanceDefinitions []*types.InstanceDefinition

	// The EC2 launch template that contains configuration settings and game server
	// code to be deployed to all instances in the game server group. You can specify
	// the template using either the template name or ID. For help with creating a
	// launch template, see Creating a Launch Template for an Auto Scaling Group
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	//
	// This member is required.
	LaunchTemplate *types.LaunchTemplateSpecification

	// The maximum number of instances allowed in the EC2 Auto Scaling group. During
	// autoscaling events, GameLift FleetIQ and EC2 do not scale up the group above
	// this maximum.
	//
	// This member is required.
	MaxSize *int32

	// The minimum number of instances allowed in the EC2 Auto Scaling group. During
	// autoscaling events, GameLift FleetIQ and EC2 do not scale down the group below
	// this minimum. In production, this value should be set to at least 1.
	//
	// This member is required.
	MinSize *int32

	// The Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) for an IAM
	// role that allows Amazon GameLift to access your EC2 Auto Scaling groups. The
	// submitted role is validated to ensure that it contains the necessary permissions
	// for game server groups.
	//
	// This member is required.
	RoleArn *string

	// Configuration settings to define a scaling policy for the Auto Scaling group
	// that is optimized for game hosting. The scaling policy uses the metric
	// "PercentUtilizedGameServers" to maintain a buffer of idle game servers that can
	// immediately accommodate new games and players. Once the game server and Auto
	// Scaling groups are created, you can update the scaling policy settings directly
	// in Auto Scaling Groups.
	AutoScalingPolicy *types.GameServerGroupAutoScalingPolicy

	// The fallback balancing method to use for the game server group when Spot
	// instances in a Region become unavailable or are not viable for game hosting.
	// Once triggered, this method remains active until Spot instances can once again
	// be used. Method options include:
	//
	//     * SPOT_ONLY -- If Spot instances are
	// unavailable, the game server group provides no hosting capacity. No new
	// instances are started, and the existing nonviable Spot instances are terminated
	// (once current gameplay ends) and not replaced.
	//
	//     * SPOT_PREFERRED -- If Spot
	// instances are unavailable, the game server group continues to provide hosting
	// capacity by using On-Demand instances. Existing nonviable Spot instances are
	// terminated (once current gameplay ends) and replaced with new On-Demand
	// instances.
	BalancingStrategy types.BalancingStrategy

	// A flag that indicates whether instances in the game server group are protected
	// from early termination. Unprotected instances that have active game servers
	// running may by terminated during a scale-down event, causing players to be
	// dropped from the game. Protected instances cannot be terminated while there are
	// active game servers running. An exception to this is Spot Instances, which may
	// be terminated by AWS regardless of protection status. This property is set to
	// NO_PROTECTION by default.
	GameServerProtectionPolicy types.GameServerProtectionPolicy

	// A list of labels to assign to the new game server group resource. Tags are
	// developer-defined key-value pairs. Tagging AWS resources are useful for resource
	// management, access management, and cost allocation. For more information, see
	// Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the AWS
	// General Reference. Once the resource is created, you can use TagResource (),
	// UntagResource (), and ListTagsForResource () to add, remove, and view tags. The
	// maximum tag limit may be lower than stated. See the AWS General Reference for
	// actual tagging limits.
	Tags []*types.Tag

	// A list of virtual private cloud (VPC) subnets to use with instances in the game
	// server group. By default, all GameLift FleetIQ-supported availability zones are
	// used; this parameter allows you to specify VPCs that you've set up.
	VpcSubnets []*string
}

type CreateGameServerGroupOutput struct {

	// The newly created game server group object, including the new ARN value for the
	// GameLift FleetIQ game server group and the object's status. The EC2 Auto Scaling
	// group ARN is initially null, since the group has not yet been created. This
	// value is added once the game server group status reaches ACTIVE.
	GameServerGroup *types.GameServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateGameServerGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateGameServerGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateGameServerGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateGameServerGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateGameServerGroup",
	}
}
