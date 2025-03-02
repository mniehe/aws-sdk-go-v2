// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the properties associated with an InferRxNorm job. Use this operation to
// get the status of an inference job.
func (c *Client) DescribeRxNormInferenceJob(ctx context.Context, params *DescribeRxNormInferenceJobInput, optFns ...func(*Options)) (*DescribeRxNormInferenceJobOutput, error) {
	if params == nil {
		params = &DescribeRxNormInferenceJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRxNormInferenceJob", params, optFns, c.addOperationDescribeRxNormInferenceJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRxNormInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRxNormInferenceJobInput struct {

	// The identifier that Amazon Comprehend Medical generated for the job. The
	// StartRxNormInferenceJob operation returns this identifier in its response.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type DescribeRxNormInferenceJobOutput struct {

	// An object that contains the properties associated with a detection job.
	ComprehendMedicalAsyncJobProperties *types.ComprehendMedicalAsyncJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRxNormInferenceJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRxNormInferenceJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRxNormInferenceJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRxNormInferenceJob"); err != nil {
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
	if err = addOpDescribeRxNormInferenceJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRxNormInferenceJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRxNormInferenceJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRxNormInferenceJob",
	}
}
