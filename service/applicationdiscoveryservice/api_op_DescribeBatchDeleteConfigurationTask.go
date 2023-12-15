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

// Takes a unique deletion task identifier as input and returns metadata about a
// configuration deletion task.
func (c *Client) DescribeBatchDeleteConfigurationTask(ctx context.Context, params *DescribeBatchDeleteConfigurationTaskInput, optFns ...func(*Options)) (*DescribeBatchDeleteConfigurationTaskOutput, error) {
	if params == nil {
		params = &DescribeBatchDeleteConfigurationTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBatchDeleteConfigurationTask", params, optFns, c.addOperationDescribeBatchDeleteConfigurationTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBatchDeleteConfigurationTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBatchDeleteConfigurationTaskInput struct {

	// The ID of the task to delete.
	//
	// This member is required.
	TaskId *string

	noSmithyDocumentSerde
}

type DescribeBatchDeleteConfigurationTaskOutput struct {

	// The BatchDeleteConfigurationTask that represents the deletion task being
	// executed.
	Task *types.BatchDeleteConfigurationTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBatchDeleteConfigurationTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBatchDeleteConfigurationTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBatchDeleteConfigurationTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBatchDeleteConfigurationTask"); err != nil {
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
	if err = addOpDescribeBatchDeleteConfigurationTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBatchDeleteConfigurationTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBatchDeleteConfigurationTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBatchDeleteConfigurationTask",
	}
}
