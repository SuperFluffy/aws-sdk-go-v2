// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new custom vocabulary that you can use to change how Amazon Transcribe
// Medical transcribes your audio file.
func (c *Client) CreateMedicalVocabulary(ctx context.Context, params *CreateMedicalVocabularyInput, optFns ...func(*Options)) (*CreateMedicalVocabularyOutput, error) {
	stack := middleware.NewStack("CreateMedicalVocabulary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateMedicalVocabularyMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateMedicalVocabularyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMedicalVocabulary(options.Region), middleware.Before)
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
			OperationName: "CreateMedicalVocabulary",
			Err:           err,
		}
	}
	out := result.(*CreateMedicalVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMedicalVocabularyInput struct {
	// The Amazon S3 location of the text file you use to define your custom
	// vocabulary. The URI must be in the same AWS region as the API endpoint you're
	// calling. Enter information about your VocabularyFileUri in the following format:
	// https://s3..amazonaws.com///  This is an example of a vocabulary file uri
	// location in Amazon S3:
	// https://s3.us-east-1.amazonaws.com/AWSDOC-EXAMPLE-BUCKET/vocab.txt For more
	// information about S3 object names, see Object Keys
	// (http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#object-keys)
	// in the Amazon S3 Developer Guide. For more information about custom
	// vocabularies, see Medical Custom Vocabularies
	// (http://docs.aws.amazon.com/transcribe/latest/dg/how-it-works.html#how-vocabulary-med).
	VocabularyFileUri *string
	// The language code used for the entries within your custom vocabulary. The
	// language code of your custom vocabulary must match the language code of your
	// transcription job. US English (en-US) is the only language code available for
	// Amazon Transcribe Medical.
	LanguageCode types.LanguageCode
	// The name of the custom vocabulary. This case-sensitive name must be unique
	// within an AWS account. If you try to create a vocabulary with the same name as a
	// previous vocabulary you will receive a ConflictException error.
	VocabularyName *string
}

type CreateMedicalVocabularyOutput struct {
	// The language code you chose to describe the entries in your custom vocabulary.
	// US English (en-US) is the only valid language code for Amazon Transcribe
	// Medical.
	LanguageCode types.LanguageCode
	// The name of the vocabulary. The name must be unique within an AWS account. It is
	// also case-sensitive.
	VocabularyName *string
	// If the VocabularyState field is FAILED, this field contains information about
	// why the job failed.
	FailureReason *string
	// The date and time you created the vocabulary.
	LastModifiedTime *time.Time
	// The processing state of your custom vocabulary in Amazon Transcribe Medical. If
	// the state is READY you can use the vocabulary in a StartMedicalTranscriptionJob
	// request.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateMedicalVocabularyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMedicalVocabulary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMedicalVocabulary{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateMedicalVocabulary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "CreateMedicalVocabulary",
	}
}