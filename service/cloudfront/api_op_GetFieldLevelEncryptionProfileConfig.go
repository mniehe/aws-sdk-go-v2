// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the field-level encryption profile configuration information.
func (c *Client) GetFieldLevelEncryptionProfileConfig(ctx context.Context, params *GetFieldLevelEncryptionProfileConfigInput, optFns ...func(*Options)) (*GetFieldLevelEncryptionProfileConfigOutput, error) {
	if params == nil {
		params = &GetFieldLevelEncryptionProfileConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFieldLevelEncryptionProfileConfig", params, optFns, c.addOperationGetFieldLevelEncryptionProfileConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFieldLevelEncryptionProfileConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFieldLevelEncryptionProfileConfigInput struct {

	// Get the ID for the field-level encryption profile configuration information.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetFieldLevelEncryptionProfileConfigOutput struct {

	// The current version of the field-level encryption profile configuration result.
	// For example: E2QWRUHAPOMQZL .
	ETag *string

	// Return the field-level encryption profile configuration information.
	FieldLevelEncryptionProfileConfig *types.FieldLevelEncryptionProfileConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFieldLevelEncryptionProfileConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetFieldLevelEncryptionProfileConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetFieldLevelEncryptionProfileConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFieldLevelEncryptionProfileConfig"); err != nil {
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
	if err = addOpGetFieldLevelEncryptionProfileConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFieldLevelEncryptionProfileConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFieldLevelEncryptionProfileConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFieldLevelEncryptionProfileConfig",
	}
}
