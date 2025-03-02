// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an auto scaling configuration to be the default. The existing default
// auto scaling configuration will be set to non-default automatically.
func (c *Client) UpdateDefaultAutoScalingConfiguration(ctx context.Context, params *UpdateDefaultAutoScalingConfigurationInput, optFns ...func(*Options)) (*UpdateDefaultAutoScalingConfigurationOutput, error) {
	if params == nil {
		params = &UpdateDefaultAutoScalingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDefaultAutoScalingConfiguration", params, optFns, c.addOperationUpdateDefaultAutoScalingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDefaultAutoScalingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDefaultAutoScalingConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the App Runner auto scaling configuration
	// that you want to set as the default. The ARN can be a full auto scaling
	// configuration ARN, or a partial ARN ending with either .../name  or
	// .../name/revision . If a revision isn't specified, the latest active revision
	// is set as the default.
	//
	// This member is required.
	AutoScalingConfigurationArn *string

	noSmithyDocumentSerde
}

type UpdateDefaultAutoScalingConfigurationOutput struct {

	// A description of the App Runner auto scaling configuration that was set as
	// default.
	//
	// This member is required.
	AutoScalingConfiguration *types.AutoScalingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDefaultAutoScalingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateDefaultAutoScalingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateDefaultAutoScalingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDefaultAutoScalingConfiguration"); err != nil {
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
	if err = addOpUpdateDefaultAutoScalingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDefaultAutoScalingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDefaultAutoScalingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDefaultAutoScalingConfiguration",
	}
}
