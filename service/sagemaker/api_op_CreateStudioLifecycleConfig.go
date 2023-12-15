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
)

// Creates a new Amazon SageMaker Studio Lifecycle Configuration.
func (c *Client) CreateStudioLifecycleConfig(ctx context.Context, params *CreateStudioLifecycleConfigInput, optFns ...func(*Options)) (*CreateStudioLifecycleConfigOutput, error) {
	if params == nil {
		params = &CreateStudioLifecycleConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStudioLifecycleConfig", params, optFns, c.addOperationCreateStudioLifecycleConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStudioLifecycleConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStudioLifecycleConfigInput struct {

	// The App type that the Lifecycle Configuration is attached to.
	//
	// This member is required.
	StudioLifecycleConfigAppType types.StudioLifecycleConfigAppType

	// The content of your Amazon SageMaker Studio Lifecycle Configuration script.
	// This content must be base64 encoded.
	//
	// This member is required.
	StudioLifecycleConfigContent *string

	// The name of the Amazon SageMaker Studio Lifecycle Configuration to create.
	//
	// This member is required.
	StudioLifecycleConfigName *string

	// Tags to be associated with the Lifecycle Configuration. Each tag consists of a
	// key and an optional value. Tag keys must be unique per resource. Tags are
	// searchable using the Search API.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateStudioLifecycleConfigOutput struct {

	// The ARN of your created Lifecycle Configuration.
	StudioLifecycleConfigArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStudioLifecycleConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStudioLifecycleConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStudioLifecycleConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateStudioLifecycleConfig"); err != nil {
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
	if err = addOpCreateStudioLifecycleConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStudioLifecycleConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStudioLifecycleConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateStudioLifecycleConfig",
	}
}
