// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes an Amazon SageMaker Model Card export job.
func (c *Client) DescribeModelCardExportJob(ctx context.Context, params *DescribeModelCardExportJobInput, optFns ...func(*Options)) (*DescribeModelCardExportJobOutput, error) {
	if params == nil {
		params = &DescribeModelCardExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeModelCardExportJob", params, optFns, c.addOperationDescribeModelCardExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeModelCardExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeModelCardExportJobInput struct {

	// The Amazon Resource Name (ARN) of the model card export job to describe.
	//
	// This member is required.
	ModelCardExportJobArn *string

	noSmithyDocumentSerde
}

type DescribeModelCardExportJobOutput struct {

	// The date and time that the model export job was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The date and time that the model export job was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The Amazon Resource Name (ARN) of the model card export job.
	//
	// This member is required.
	ModelCardExportJobArn *string

	// The name of the model card export job to describe.
	//
	// This member is required.
	ModelCardExportJobName *string

	// The name or Amazon Resource Name (ARN) of the model card that the model export
	// job exports.
	//
	// This member is required.
	ModelCardName *string

	// The version of the model card that the model export job exports.
	//
	// This member is required.
	ModelCardVersion *int32

	// The export output details for the model card.
	//
	// This member is required.
	OutputConfig *types.ModelCardExportOutputConfig

	// The completion status of the model card export job.
	//   - InProgress : The model card export job is in progress.
	//   - Completed : The model card export job is complete.
	//   - Failed : The model card export job failed. To see the reason for the
	//   failure, see the FailureReason field in the response to a
	//   DescribeModelCardExportJob call.
	//
	// This member is required.
	Status types.ModelCardExportJobStatus

	// The exported model card artifacts.
	ExportArtifacts *types.ModelCardExportArtifacts

	// The failure reason if the model export job fails.
	FailureReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeModelCardExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeModelCardExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeModelCardExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeModelCardExportJob"); err != nil {
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
	if err = addOpDescribeModelCardExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeModelCardExportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeModelCardExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeModelCardExportJob",
	}
}
