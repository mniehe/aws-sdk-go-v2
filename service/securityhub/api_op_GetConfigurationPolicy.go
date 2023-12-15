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
	"time"
)

// Provides information about a configuration policy. Only the Security Hub
// delegated administrator can invoke this operation from the home Region.
func (c *Client) GetConfigurationPolicy(ctx context.Context, params *GetConfigurationPolicyInput, optFns ...func(*Options)) (*GetConfigurationPolicyOutput, error) {
	if params == nil {
		params = &GetConfigurationPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConfigurationPolicy", params, optFns, c.addOperationGetConfigurationPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConfigurationPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConfigurationPolicyInput struct {

	// The Amazon Resource Name (ARN) or universally unique identifier (UUID) of the
	// configuration policy.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetConfigurationPolicyOutput struct {

	// The ARN of the configuration policy.
	Arn *string

	// An object that defines how Security Hub is configured. It includes whether
	// Security Hub is enabled or disabled, a list of enabled security standards, a
	// list of enabled or disabled security controls, and a list of custom parameter
	// values for specified controls. If the policy includes a list of security
	// controls that are enabled, Security Hub disables all other controls (including
	// newly released controls). If the policy includes a list of security controls
	// that are disabled, Security Hub enables all other controls (including newly
	// released controls).
	ConfigurationPolicy types.Policy

	// The date and time, in UTC and ISO 8601 format, that the configuration policy
	// was created.
	CreatedAt *time.Time

	// The description of the configuration policy.
	Description *string

	// The UUID of the configuration policy.
	Id *string

	// The name of the configuration policy.
	Name *string

	// The date and time, in UTC and ISO 8601 format, that the configuration policy
	// was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConfigurationPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConfigurationPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfigurationPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetConfigurationPolicy"); err != nil {
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
	if err = addOpGetConfigurationPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfigurationPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConfigurationPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetConfigurationPolicy",
	}
}
