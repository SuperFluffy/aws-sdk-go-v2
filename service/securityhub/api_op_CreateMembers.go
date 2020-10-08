// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a member association in Security Hub between the specified accounts and
// the account used to make the request, which is the master account. To
// successfully create a member, you must use this action from an account that
// already has Security Hub enabled. To enable Security Hub, you can use the
// EnableSecurityHub () operation. After you use CreateMembers to create member
// account associations in Security Hub, you must use the InviteMembers ()
// operation to invite the accounts to enable Security Hub and become member
// accounts in Security Hub. If the account owner accepts the invitation, the
// account becomes a member account in Security Hub. A permissions policy is added
// that permits the master account to view the findings generated in the member
// account. When Security Hub is enabled in the invited account, findings start to
// be sent to both the member and master accounts. To remove the association
// between the master and member accounts, use the DisassociateFromMasterAccount ()
// or DisassociateMembers () operation.
func (c *Client) CreateMembers(ctx context.Context, params *CreateMembersInput, optFns ...func(*Options)) (*CreateMembersOutput, error) {
	stack := middleware.NewStack("CreateMembers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateMembersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMembers(options.Region), middleware.Before)
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
			OperationName: "CreateMembers",
			Err:           err,
		}
	}
	out := result.(*CreateMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMembersInput struct {

	// The list of accounts to associate with the Security Hub master account. For each
	// account, the list includes the account ID and the email address.
	AccountDetails []*types.AccountDetails
}

type CreateMembersOutput struct {

	// The list of AWS accounts that were not processed. For each account, the list
	// includes the account ID and the email address.
	UnprocessedAccounts []*types.Result

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateMembersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateMembers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMembers{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateMembers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "CreateMembers",
	}
}
