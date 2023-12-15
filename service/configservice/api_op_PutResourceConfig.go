// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Records the configuration state for the resource provided in the request. The
// configuration state of a resource is represented in Config as Configuration
// Items. Once this API records the configuration item, you can retrieve the list
// of configuration items for the custom resource type using existing Config APIs.
// The custom resource type must be registered with CloudFormation. This API
// accepts the configuration item registered with CloudFormation. When you call
// this API, Config only stores configuration state of the resource provided in the
// request. This API does not change or remediate the configuration of the
// resource. Write-only schema properites are not recorded as part of the published
// configuration item.
func (c *Client) PutResourceConfig(ctx context.Context, params *PutResourceConfigInput, optFns ...func(*Options)) (*PutResourceConfigOutput, error) {
	if params == nil {
		params = &PutResourceConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResourceConfig", params, optFns, c.addOperationPutResourceConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourceConfigInput struct {

	// The configuration object of the resource in valid JSON format. It must match
	// the schema registered with CloudFormation. The configuration JSON must not
	// exceed 64 KB.
	//
	// This member is required.
	Configuration *string

	// Unique identifier of the resource.
	//
	// This member is required.
	ResourceId *string

	// The type of the resource. The custom resource type must be registered with
	// CloudFormation. You cannot use the organization names “amzn”, “amazon”, “alexa”,
	// “custom” with custom resource types. It is the first part of the ResourceType up
	// to the first ::.
	//
	// This member is required.
	ResourceType *string

	// Version of the schema registered for the ResourceType in CloudFormation.
	//
	// This member is required.
	SchemaVersionId *string

	// Name of the resource.
	ResourceName *string

	// Tags associated with the resource. This field is not to be confused with the
	// Amazon Web Services-wide tag feature for Amazon Web Services resources. Tags for
	// PutResourceConfig are tags that you supply for the configuration items of your
	// custom resources.
	Tags map[string]string

	noSmithyDocumentSerde
}

type PutResourceConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutResourceConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutResourceConfig"); err != nil {
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
	if err = addOpPutResourceConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourceConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutResourceConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutResourceConfig",
	}
}
