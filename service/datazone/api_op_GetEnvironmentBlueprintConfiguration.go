// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the blueprint configuration in Amazon DataZone.
func (c *Client) GetEnvironmentBlueprintConfiguration(ctx context.Context, params *GetEnvironmentBlueprintConfigurationInput, optFns ...func(*Options)) (*GetEnvironmentBlueprintConfigurationOutput, error) {
	if params == nil {
		params = &GetEnvironmentBlueprintConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnvironmentBlueprintConfiguration", params, optFns, c.addOperationGetEnvironmentBlueprintConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnvironmentBlueprintConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnvironmentBlueprintConfigurationInput struct {

	// The ID of the Amazon DataZone domain where this blueprint exists.
	//
	// This member is required.
	DomainIdentifier *string

	// He ID of the blueprint.
	//
	// This member is required.
	EnvironmentBlueprintIdentifier *string

	noSmithyDocumentSerde
}

type GetEnvironmentBlueprintConfigurationOutput struct {

	// The ID of the Amazon DataZone domain where this blueprint exists.
	//
	// This member is required.
	DomainId *string

	// The ID of the blueprint.
	//
	// This member is required.
	EnvironmentBlueprintId *string

	// The timestamp of when this blueprint was created.
	CreatedAt *time.Time

	// The Amazon Web Services regions in which this blueprint is enabled.
	EnabledRegions []string

	// The ARN of the manage access role with which this blueprint is created.
	ManageAccessRoleArn *string

	// The ARN of the provisioning role with which this blueprint is created.
	ProvisioningRoleArn *string

	// The regional parameters of the blueprint.
	RegionalParameters map[string]map[string]string

	// The timestamp of when this blueprint was upated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEnvironmentBlueprintConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEnvironmentBlueprintConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEnvironmentBlueprintConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEnvironmentBlueprintConfiguration"); err != nil {
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
	if err = addOpGetEnvironmentBlueprintConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnvironmentBlueprintConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEnvironmentBlueprintConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEnvironmentBlueprintConfiguration",
	}
}
