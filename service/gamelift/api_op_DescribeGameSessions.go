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

// Retrieves a set of one or more game sessions. Request a specific game session or
// request all game sessions on a fleet. Alternatively, use SearchGameSessions ()
// to request a set of active game sessions that are filtered by certain criteria.
// To retrieve protection policy settings for game sessions, use
// DescribeGameSessionDetails (). To get game sessions, specify one of the
// following: game session ID, fleet ID, or alias ID. You can filter this request
// by game session status. Use the pagination parameters to retrieve results as a
// set of sequential pages. If successful, a GameSession () object is returned for
// each game session matching the request. Available in Amazon GameLift Local.
//
//
// * CreateGameSession ()
//
//     * DescribeGameSessions ()
//
//     *
// DescribeGameSessionDetails ()
//
//     * SearchGameSessions ()
//
//     *
// UpdateGameSession ()
//
//     * GetGameSessionLogUrl ()
//
//     * Game session
// placements
//
//         * StartGameSessionPlacement ()
//
//         *
// DescribeGameSessionPlacement ()
//
//         * StopGameSessionPlacement ()
func (c *Client) DescribeGameSessions(ctx context.Context, params *DescribeGameSessionsInput, optFns ...func(*Options)) (*DescribeGameSessionsOutput, error) {
	stack := middleware.NewStack("DescribeGameSessions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeGameSessionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGameSessions(options.Region), middleware.Before)
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
			OperationName: "DescribeGameSessions",
			Err:           err,
		}
	}
	out := result.(*DescribeGameSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DescribeGameSessionsInput struct {

	// A unique identifier for an alias associated with the fleet to retrieve all game
	// sessions for. You can use either the alias ID or ARN value.
	AliasId *string

	// A unique identifier for a fleet to retrieve all game sessions for. You can use
	// either the fleet ID or ARN value.
	FleetId *string

	// A unique identifier for the game session to retrieve.
	GameSessionId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this action. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// Game session status to filter results on. Possible game session statuses include
	// ACTIVE, TERMINATED, ACTIVATING, and TERMINATING (the last two are transitory).
	StatusFilter *string
}

// Represents the returned data in response to a request action.
type DescribeGameSessionsOutput struct {

	// A collection of objects containing game session properties for each session
	// matching the request.
	GameSessions []*types.GameSession

	// Token that indicates where to resume retrieving results on the next call to this
	// action. If no token is returned, these results represent the end of the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeGameSessionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGameSessions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGameSessions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeGameSessions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeGameSessions",
	}
}
