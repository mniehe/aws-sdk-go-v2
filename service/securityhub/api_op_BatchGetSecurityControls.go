// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides details about a batch of security controls for the current Amazon Web
// Services account and Amazon Web Services Region.
func (c *Client) BatchGetSecurityControls(ctx context.Context, params *BatchGetSecurityControlsInput, optFns ...func(*Options)) (*BatchGetSecurityControlsOutput, error) {
	if params == nil {
		params = &BatchGetSecurityControlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetSecurityControls", params, optFns, c.addOperationBatchGetSecurityControlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetSecurityControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetSecurityControlsInput struct {

	// A list of security controls (identified with SecurityControlId ,
	// SecurityControlArn , or a mix of both parameters). The security control ID or
	// Amazon Resource Name (ARN) is the same across standards.
	//
	// This member is required.
	SecurityControlIds []string

	noSmithyDocumentSerde
}

type BatchGetSecurityControlsOutput struct {

	// An array that returns the identifier, Amazon Resource Name (ARN), and other
	// details about a security control. The same information is returned whether the
	// request includes SecurityControlId or SecurityControlArn .
	//
	// This member is required.
	SecurityControls []types.SecurityControl

	// A security control (identified with SecurityControlId , SecurityControlArn , or
	// a mix of both parameters) for which details cannot be returned.
	UnprocessedIds []types.UnprocessedSecurityControl

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetSecurityControlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetSecurityControls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetSecurityControls{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetSecurityControls"); err != nil {
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
	if err = addOpBatchGetSecurityControlsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetSecurityControls(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetSecurityControls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetSecurityControls",
	}
}
