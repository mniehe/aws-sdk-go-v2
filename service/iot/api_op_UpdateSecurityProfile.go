// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a Device Defender security profile. Requires permission to access the
// UpdateSecurityProfile (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) UpdateSecurityProfile(ctx context.Context, params *UpdateSecurityProfileInput, optFns ...func(*Options)) (*UpdateSecurityProfileOutput, error) {
	if params == nil {
		params = &UpdateSecurityProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSecurityProfile", params, optFns, c.addOperationUpdateSecurityProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSecurityProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSecurityProfileInput struct {

	// The name of the security profile you want to update.
	//
	// This member is required.
	SecurityProfileName *string

	// Please use UpdateSecurityProfileRequest$additionalMetricsToRetainV2 instead. A
	// list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors , but it is also retained for any
	// metric specified here. Can be used with custom metrics; cannot be used with
	// dimensions.
	//
	// Deprecated: Use additionalMetricsToRetainV2.
	AdditionalMetricsToRetain []string

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here. Can be used with custom metrics; cannot be used with
	// dimensions.
	AdditionalMetricsToRetainV2 []types.MetricToRetain

	// Where the alerts are sent. (Alerts are always sent to the console.)
	AlertTargets map[string]types.AlertTarget

	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors []types.Behavior

	// If true, delete all additionalMetricsToRetain defined for this security
	// profile. If any additionalMetricsToRetain are defined in the current
	// invocation, an exception occurs.
	DeleteAdditionalMetricsToRetain bool

	// If true, delete all alertTargets defined for this security profile. If any
	// alertTargets are defined in the current invocation, an exception occurs.
	DeleteAlertTargets bool

	// If true, delete all behaviors defined for this security profile. If any
	// behaviors are defined in the current invocation, an exception occurs.
	DeleteBehaviors bool

	// Set the value as true to delete metrics export related configurations.
	DeleteMetricsExportConfig bool

	// The expected version of the security profile. A new version is generated
	// whenever the security profile is updated. If you specify a value that is
	// different from the actual version, a VersionConflictException is thrown.
	ExpectedVersion *int64

	// Specifies the MQTT topic and role ARN required for metric export.
	MetricsExportConfig *types.MetricsExportConfig

	// A description of the security profile.
	SecurityProfileDescription *string

	noSmithyDocumentSerde
}

type UpdateSecurityProfileOutput struct {

	// Please use UpdateSecurityProfileResponse$additionalMetricsToRetainV2 instead. A
	// list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the security profile's behaviors , but it is also
	// retained for any metric specified here.
	//
	// Deprecated: Use additionalMetricsToRetainV2.
	AdditionalMetricsToRetain []string

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here. Can be used with custom metrics; cannot be used with
	// dimensions.
	AdditionalMetricsToRetainV2 []types.MetricToRetain

	// Where the alerts are sent. (Alerts are always sent to the console.)
	AlertTargets map[string]types.AlertTarget

	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors []types.Behavior

	// The time the security profile was created.
	CreationDate *time.Time

	// The time the security profile was last modified.
	LastModifiedDate *time.Time

	// Specifies the MQTT topic and role ARN required for metric export.
	MetricsExportConfig *types.MetricsExportConfig

	// The ARN of the security profile that was updated.
	SecurityProfileArn *string

	// The description of the security profile.
	SecurityProfileDescription *string

	// The name of the security profile that was updated.
	SecurityProfileName *string

	// The updated version of the security profile.
	Version int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSecurityProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSecurityProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSecurityProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSecurityProfile"); err != nil {
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
	if err = addOpUpdateSecurityProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSecurityProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSecurityProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSecurityProfile",
	}
}
