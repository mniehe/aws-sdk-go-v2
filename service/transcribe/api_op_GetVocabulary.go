// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about the specified custom vocabulary. To view the status
// of the specified custom vocabulary, check the VocabularyState field. If the
// status is READY , your custom vocabulary is available to use. If the status is
// FAILED , FailureReason provides details on why your custom vocabulary failed.
// To get a list of your custom vocabularies, use the operation.
func (c *Client) GetVocabulary(ctx context.Context, params *GetVocabularyInput, optFns ...func(*Options)) (*GetVocabularyOutput, error) {
	if params == nil {
		params = &GetVocabularyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVocabulary", params, optFns, c.addOperationGetVocabularyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVocabularyInput struct {

	// The name of the custom vocabulary you want information about. Custom vocabulary
	// names are case sensitive.
	//
	// This member is required.
	VocabularyName *string

	noSmithyDocumentSerde
}

type GetVocabularyOutput struct {

	// The Amazon S3 location where the custom vocabulary is stored; use this URI to
	// view or download the custom vocabulary.
	DownloadUri *string

	// If VocabularyState is FAILED , FailureReason contains information about why the
	// custom vocabulary request failed. See also: Common Errors (https://docs.aws.amazon.com/transcribe/latest/APIReference/CommonErrors.html)
	// .
	FailureReason *string

	// The language code you selected for your custom vocabulary.
	LanguageCode types.LanguageCode

	// The date and time the specified custom vocabulary was last modified. Timestamps
	// are in the format YYYY-MM-DD'T'HH:MM:SS.SSSSSS-UTC . For example,
	// 2022-05-04T12:32:58.761000-07:00 represents 12:32 PM UTC-7 on May 4, 2022.
	LastModifiedTime *time.Time

	// The name of the custom vocabulary you requested information about.
	VocabularyName *string

	// The processing state of your custom vocabulary. If the state is READY , you can
	// use the custom vocabulary in a StartTranscriptionJob request.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetVocabularyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetVocabulary"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetVocabularyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVocabulary(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetVocabulary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetVocabulary",
	}
}
