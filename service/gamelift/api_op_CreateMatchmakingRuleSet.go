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

// Creates a new rule set for FlexMatch matchmaking. A rule set describes the type
// of match to create, such as the number and size of teams. It also sets the
// parameters for acceptable player matches, such as minimum skill level or
// character type. A rule set is used by a MatchmakingConfiguration (). To create a
// matchmaking rule set, provide unique rule set name and the rule set body in JSON
// format. Rule sets must be defined in the same Region as the matchmaking
// configuration they are used with. Since matchmaking rule sets cannot be edited,
// it is a good idea to check the rule set syntax using ValidateMatchmakingRuleSet
// () before creating a new rule set. Learn more
//
//     * Build a Rule Set
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-rulesets.html)
//
//
// * Design a Matchmaker
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-configuration.html)
//
//
// * Matchmaking with FlexMatch
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-intro.html)
//
// Related
// operations
//
//     * CreateMatchmakingConfiguration ()
//
//     *
// DescribeMatchmakingConfigurations ()
//
//     * UpdateMatchmakingConfiguration ()
//
//
// * DeleteMatchmakingConfiguration ()
//
//     * CreateMatchmakingRuleSet ()
//
//     *
// DescribeMatchmakingRuleSets ()
//
//     * ValidateMatchmakingRuleSet ()
//
//     *
// DeleteMatchmakingRuleSet ()
func (c *Client) CreateMatchmakingRuleSet(ctx context.Context, params *CreateMatchmakingRuleSetInput, optFns ...func(*Options)) (*CreateMatchmakingRuleSetOutput, error) {
	stack := middleware.NewStack("CreateMatchmakingRuleSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateMatchmakingRuleSetMiddlewares(stack)
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
	addOpCreateMatchmakingRuleSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMatchmakingRuleSet(options.Region), middleware.Before)
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
			OperationName: "CreateMatchmakingRuleSet",
			Err:           err,
		}
	}
	out := result.(*CreateMatchmakingRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type CreateMatchmakingRuleSetInput struct {

	// A unique identifier for a matchmaking rule set. A matchmaking configuration
	// identifies the rule set it uses by this name value. Note that the rule set name
	// is different from the optional name field in the rule set body.
	//
	// This member is required.
	Name *string

	// A collection of matchmaking rules, formatted as a JSON string. Comments are not
	// allowed in JSON, but most elements support a description field.
	//
	// This member is required.
	RuleSetBody *string

	// A list of labels to assign to the new matchmaking rule set resource. Tags are
	// developer-defined key-value pairs. Tagging AWS resources are useful for resource
	// management, access management and cost allocation. For more information, see
	// Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the AWS
	// General Reference. Once the resource is created, you can use TagResource (),
	// UntagResource (), and ListTagsForResource () to add, remove, and view tags. The
	// maximum tag limit may be lower than stated. See the AWS General Reference for
	// actual tagging limits.
	Tags []*types.Tag
}

// Represents the returned data in response to a request action.
type CreateMatchmakingRuleSetOutput struct {

	// The newly created matchmaking rule set.
	//
	// This member is required.
	RuleSet *types.MatchmakingRuleSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateMatchmakingRuleSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMatchmakingRuleSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMatchmakingRuleSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateMatchmakingRuleSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateMatchmakingRuleSet",
	}
}
