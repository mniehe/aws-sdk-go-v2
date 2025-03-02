// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Sets the amount of time in days that a given AppInstance retains data. This API
// is is no longer supported and will not be updated. We recommend using the latest
// version, PutAppInstanceRetentionSettings (https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_identity-chime_PutAppInstanceRetentionSettings.html)
// , in the Amazon Chime SDK. Using the latest version requires migrating to a
// dedicated namespace. For more information, refer to Migrating from the Amazon
// Chime namespace (https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html)
// in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by PutAppInstanceRetentionSettings in the Amazon Chime SDK
// Identity Namespace
func (c *Client) PutAppInstanceRetentionSettings(ctx context.Context, params *PutAppInstanceRetentionSettingsInput, optFns ...func(*Options)) (*PutAppInstanceRetentionSettingsOutput, error) {
	if params == nil {
		params = &PutAppInstanceRetentionSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAppInstanceRetentionSettings", params, optFns, c.addOperationPutAppInstanceRetentionSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAppInstanceRetentionSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAppInstanceRetentionSettingsInput struct {

	// The ARN of the AppInstance .
	//
	// This member is required.
	AppInstanceArn *string

	// The time in days to retain data. Data type: number.
	//
	// This member is required.
	AppInstanceRetentionSettings *types.AppInstanceRetentionSettings

	noSmithyDocumentSerde
}

type PutAppInstanceRetentionSettingsOutput struct {

	// The time in days to retain data. Data type: number.
	AppInstanceRetentionSettings *types.AppInstanceRetentionSettings

	// The time at which the API deletes data.
	InitiateDeletionTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAppInstanceRetentionSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutAppInstanceRetentionSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutAppInstanceRetentionSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAppInstanceRetentionSettings"); err != nil {
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
	if err = addEndpointPrefix_opPutAppInstanceRetentionSettingsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutAppInstanceRetentionSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAppInstanceRetentionSettings(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opPutAppInstanceRetentionSettingsMiddleware struct {
}

func (*endpointPrefix_opPutAppInstanceRetentionSettingsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutAppInstanceRetentionSettingsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "identity-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opPutAppInstanceRetentionSettingsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opPutAppInstanceRetentionSettingsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opPutAppInstanceRetentionSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAppInstanceRetentionSettings",
	}
}
