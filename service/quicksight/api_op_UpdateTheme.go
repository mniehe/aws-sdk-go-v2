// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a theme.
func (c *Client) UpdateTheme(ctx context.Context, params *UpdateThemeInput, optFns ...func(*Options)) (*UpdateThemeOutput, error) {
	if params == nil {
		params = &UpdateThemeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTheme", params, optFns, c.addOperationUpdateThemeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateThemeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateThemeInput struct {

	// The ID of the Amazon Web Services account that contains the theme that you're
	// updating.
	//
	// This member is required.
	AwsAccountId *string

	// The theme ID, defined by Amazon QuickSight, that a custom theme inherits from.
	// All themes initially inherit from a default Amazon QuickSight theme.
	//
	// This member is required.
	BaseThemeId *string

	// The ID for the theme.
	//
	// This member is required.
	ThemeId *string

	// The theme configuration, which contains the theme display properties.
	Configuration *types.ThemeConfiguration

	// The name for the theme.
	Name *string

	// A description of the theme version that you're updating Every time that you
	// call UpdateTheme , you create a new version of the theme. Each version of the
	// theme maintains a description of the version in VersionDescription .
	VersionDescription *string

	noSmithyDocumentSerde
}

type UpdateThemeOutput struct {

	// The Amazon Resource Name (ARN) for the theme.
	Arn *string

	// The creation status of the theme.
	CreationStatus types.ResourceStatus

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The ID for the theme.
	ThemeId *string

	// The Amazon Resource Name (ARN) for the new version of the theme.
	VersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateThemeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTheme{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTheme{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTheme"); err != nil {
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
	if err = addOpUpdateThemeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTheme(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTheme(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTheme",
	}
}
