// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the record of a single premigration assessment run. This operation
// removes all metadata that AWS DMS maintains about this assessment run. However,
// the operation leaves untouched all information about this assessment run that is
// stored in your Amazon S3 bucket.
func (c *Client) DeleteReplicationTaskAssessmentRun(ctx context.Context, params *DeleteReplicationTaskAssessmentRunInput, optFns ...func(*Options)) (*DeleteReplicationTaskAssessmentRunOutput, error) {
	stack := middleware.NewStack("DeleteReplicationTaskAssessmentRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteReplicationTaskAssessmentRunMiddlewares(stack)
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
	addOpDeleteReplicationTaskAssessmentRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReplicationTaskAssessmentRun(options.Region), middleware.Before)
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
			OperationName: "DeleteReplicationTaskAssessmentRun",
			Err:           err,
		}
	}
	out := result.(*DeleteReplicationTaskAssessmentRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteReplicationTaskAssessmentRunInput struct {

	// Amazon Resource Name (ARN) of the premigration assessment run to be deleted.
	//
	// This member is required.
	ReplicationTaskAssessmentRunArn *string
}

//
type DeleteReplicationTaskAssessmentRunOutput struct {

	// The ReplicationTaskAssessmentRun object for the deleted assessment run.
	ReplicationTaskAssessmentRun *types.ReplicationTaskAssessmentRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteReplicationTaskAssessmentRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteReplicationTaskAssessmentRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteReplicationTaskAssessmentRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteReplicationTaskAssessmentRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DeleteReplicationTaskAssessmentRun",
	}
}
