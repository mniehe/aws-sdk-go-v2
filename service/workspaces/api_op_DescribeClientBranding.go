// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified client branding. Client branding allows you to
// customize the log in page of various device types for your users. You can add
// your company logo, the support email address, support link, link to reset
// password, and a custom message for users trying to sign in. Only device types
// that have branding information configured will be shown in the response.
func (c *Client) DescribeClientBranding(ctx context.Context, params *DescribeClientBrandingInput, optFns ...func(*Options)) (*DescribeClientBrandingOutput, error) {
	if params == nil {
		params = &DescribeClientBrandingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClientBranding", params, optFns, c.addOperationDescribeClientBrandingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClientBrandingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClientBrandingInput struct {

	// The directory identifier of the WorkSpace for which you want to view client
	// branding information.
	//
	// This member is required.
	ResourceId *string

	noSmithyDocumentSerde
}

type DescribeClientBrandingOutput struct {

	// The branding information for Android devices.
	DeviceTypeAndroid *types.DefaultClientBrandingAttributes

	// The branding information for iOS devices.
	DeviceTypeIos *types.IosClientBrandingAttributes

	// The branding information for Linux devices.
	DeviceTypeLinux *types.DefaultClientBrandingAttributes

	// The branding information for macOS devices.
	DeviceTypeOsx *types.DefaultClientBrandingAttributes

	// The branding information for Web access.
	DeviceTypeWeb *types.DefaultClientBrandingAttributes

	// The branding information for Windows devices.
	DeviceTypeWindows *types.DefaultClientBrandingAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClientBrandingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeClientBranding{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeClientBranding{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeClientBranding"); err != nil {
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
	if err = addOpDescribeClientBrandingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClientBranding(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeClientBranding(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeClientBranding",
	}
}
