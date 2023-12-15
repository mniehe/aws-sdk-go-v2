// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This API operation contains deprecated parameters. Register your account again
// without the Timestream resources parameter so that Amazon Web Services IoT
// FleetWise can remove the Timestream metadata stored. You should then pass the
// data destination into the CreateCampaign (https://docs.aws.amazon.com/iot-fleetwise/latest/APIReference/API_CreateCampaign.html)
// API operation. You must delete any existing campaigns that include an empty data
// destination before you register your account again. For more information, see
// the DeleteCampaign (https://docs.aws.amazon.com/iot-fleetwise/latest/APIReference/API_DeleteCampaign.html)
// API operation. If you want to delete the Timestream inline policy from the
// service-linked role, such as to mitigate an overly permissive policy, you must
// first delete any existing campaigns. Then delete the service-linked role and
// register your account again to enable CloudWatch metrics. For more information,
// see DeleteServiceLinkedRole (https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteServiceLinkedRole.html)
// in the Identity and Access Management API Reference. Registers your Amazon Web
// Services account, IAM, and Amazon Timestream resources so Amazon Web Services
// IoT FleetWise can transfer your vehicle data to the Amazon Web Services Cloud.
// For more information, including step-by-step procedures, see Setting up Amazon
// Web Services IoT FleetWise (https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/setting-up.html)
// . An Amazon Web Services account is not the same thing as a "user." An Amazon
// Web Services user (https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction_identity-management.html#intro-identity-users)
// is an identity that you create using Identity and Access Management (IAM) and
// takes the form of either an IAM user (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users.html)
// or an IAM role, both with credentials (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)
// . A single Amazon Web Services account can, and typically does, contain many
// users and roles.
func (c *Client) RegisterAccount(ctx context.Context, params *RegisterAccountInput, optFns ...func(*Options)) (*RegisterAccountOutput, error) {
	if params == nil {
		params = &RegisterAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterAccount", params, optFns, c.addOperationRegisterAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterAccountInput struct {

	// The IAM resource that allows Amazon Web Services IoT FleetWise to send data to
	// Amazon Timestream.
	//
	// Deprecated: iamResources is no longer used or needed as input
	IamResources *types.IamResources

	// The registered Amazon Timestream resources that Amazon Web Services IoT
	// FleetWise edge agent software can transfer your vehicle data to.
	//
	// Deprecated: Amazon Timestream metadata is now passed in the CreateCampaign API.
	TimestreamResources *types.TimestreamResources

	noSmithyDocumentSerde
}

type RegisterAccountOutput struct {

	// The time the account was registered, in seconds since epoch (January 1, 1970 at
	// midnight UTC time).
	//
	// This member is required.
	CreationTime *time.Time

	// The registered IAM resource that allows Amazon Web Services IoT FleetWise to
	// send data to Amazon Timestream.
	//
	// This member is required.
	IamResources *types.IamResources

	// The time this registration was last updated, in seconds since epoch (January 1,
	// 1970 at midnight UTC time).
	//
	// This member is required.
	LastModificationTime *time.Time

	// The status of registering your Amazon Web Services account, IAM role, and
	// Timestream resources.
	//
	// This member is required.
	RegisterAccountStatus types.RegistrationStatus

	// The registered Amazon Timestream resources that Amazon Web Services IoT
	// FleetWise edge agent software can transfer your vehicle data to.
	TimestreamResources *types.TimestreamResources

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRegisterAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRegisterAccount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterAccount"); err != nil {
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
	if err = addOpRegisterAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterAccount",
	}
}
