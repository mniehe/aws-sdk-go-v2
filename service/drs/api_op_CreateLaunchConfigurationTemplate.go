// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Launch Configuration Template.
func (c *Client) CreateLaunchConfigurationTemplate(ctx context.Context, params *CreateLaunchConfigurationTemplateInput, optFns ...func(*Options)) (*CreateLaunchConfigurationTemplateOutput, error) {
	if params == nil {
		params = &CreateLaunchConfigurationTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLaunchConfigurationTemplate", params, optFns, c.addOperationCreateLaunchConfigurationTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLaunchConfigurationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchConfigurationTemplateInput struct {

	// Copy private IP.
	CopyPrivateIp *bool

	// Copy tags.
	CopyTags *bool

	// S3 bucket ARN to export Source Network templates.
	ExportBucketArn *string

	// Launch disposition.
	LaunchDisposition types.LaunchDisposition

	// DRS will set the 'launch into instance ID' of any source server when performing
	// a drill, recovery or failback to the previous region or availability zone, using
	// the instance ID of the source instance.
	LaunchIntoSourceInstance *bool

	// Licensing.
	Licensing *types.Licensing

	// Whether we want to activate post-launch actions.
	PostLaunchEnabled *bool

	// Request to associate tags during creation of a Launch Configuration Template.
	Tags map[string]string

	// Target instance type right-sizing method.
	TargetInstanceTypeRightSizingMethod types.TargetInstanceTypeRightSizingMethod

	noSmithyDocumentSerde
}

type CreateLaunchConfigurationTemplateOutput struct {

	// Created Launch Configuration Template.
	LaunchConfigurationTemplate *types.LaunchConfigurationTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLaunchConfigurationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLaunchConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLaunchConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLaunchConfigurationTemplate"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchConfigurationTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLaunchConfigurationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLaunchConfigurationTemplate",
	}
}
