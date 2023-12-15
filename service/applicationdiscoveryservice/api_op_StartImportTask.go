// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an import task, which allows you to import details of your on-premises
// environment directly into Amazon Web Services Migration Hub without having to
// use the Amazon Web Services Application Discovery Service (Application Discovery
// Service) tools such as the Amazon Web Services Application Discovery Service
// Agentless Collector or Application Discovery Agent. This gives you the option to
// perform migration assessment and planning directly from your imported data,
// including the ability to group your devices as applications and track their
// migration status. To start an import request, do this:
//   - Download the specially formatted comma separated value (CSV) import
//     template, which you can find here:
//     https://s3.us-west-2.amazonaws.com/templates-7cffcf56-bd96-4b1c-b45b-a5b42f282e46/import_template.csv (https://s3.us-west-2.amazonaws.com/templates-7cffcf56-bd96-4b1c-b45b-a5b42f282e46/import_template.csv)
//     .
//   - Fill out the template with your server and application data.
//   - Upload your import file to an Amazon S3 bucket, and make a note of it's
//     Object URL. Your import file must be in the CSV format.
//   - Use the console or the StartImportTask command with the Amazon Web Services
//     CLI or one of the Amazon Web Services SDKs to import the records from your file.
//
// For more information, including step-by-step procedures, see Migration Hub
// Import (https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-import.html)
// in the Amazon Web Services Application Discovery Service User Guide. There are
// limits to the number of import tasks you can create (and delete) in an Amazon
// Web Services account. For more information, see Amazon Web Services Application
// Discovery Service Limits (https://docs.aws.amazon.com/application-discovery/latest/userguide/ads_service_limits.html)
// in the Amazon Web Services Application Discovery Service User Guide.
func (c *Client) StartImportTask(ctx context.Context, params *StartImportTaskInput, optFns ...func(*Options)) (*StartImportTaskOutput, error) {
	if params == nil {
		params = &StartImportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartImportTask", params, optFns, c.addOperationStartImportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartImportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartImportTaskInput struct {

	// The URL for your import file that you've uploaded to Amazon S3. If you're using
	// the Amazon Web Services CLI, this URL is structured as follows:
	// s3://BucketName/ImportFileName.CSV
	//
	// This member is required.
	ImportUrl *string

	// A descriptive name for this request. You can use this name to filter future
	// requests related to this import task, such as identifying applications and
	// servers that were included in this import task. We recommend that you use a
	// meaningful name for each import task.
	//
	// This member is required.
	Name *string

	// Optional. A unique token that you can provide to prevent the same import
	// request from occurring more than once. If you don't provide a token, a token is
	// automatically generated. Sending more than one StartImportTask request with the
	// same client request token will return information about the original import task
	// with that client request token.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type StartImportTaskOutput struct {

	// An array of information related to the import task request including status
	// information, times, IDs, the Amazon S3 Object URL for the import file, and more.
	Task *types.ImportTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartImportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartImportTask"); err != nil {
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
	if err = addIdempotencyToken_opStartImportTaskMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartImportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartImportTask(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartImportTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartImportTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartImportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartImportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartImportTaskInput ")
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
func addIdempotencyToken_opStartImportTaskMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartImportTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartImportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartImportTask",
	}
}
