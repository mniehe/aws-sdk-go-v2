// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about an MFA device for a specified user.
func (c *Client) GetMFADevice(ctx context.Context, params *GetMFADeviceInput, optFns ...func(*Options)) (*GetMFADeviceOutput, error) {
	if params == nil {
		params = &GetMFADeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMFADevice", params, optFns, c.addOperationGetMFADeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMFADeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMFADeviceInput struct {

	// Serial number that uniquely identifies the MFA device. For this API, we only
	// accept FIDO security key ARNs (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html)
	// .
	//
	// This member is required.
	SerialNumber *string

	// The friendly name identifying the user.
	UserName *string

	noSmithyDocumentSerde
}

type GetMFADeviceOutput struct {

	// Serial number that uniquely identifies the MFA device. For this API, we only
	// accept FIDO security key ARNs (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html)
	// .
	//
	// This member is required.
	SerialNumber *string

	// The certifications of a specified user's MFA device. We currently provide
	// FIPS-140-2, FIPS-140-3, and FIDO certification levels obtained from FIDO
	// Alliance Metadata Service (MDS) (https://fidoalliance.org/metadata/) .
	Certifications map[string]string

	// The date that a specified user's MFA device was first enabled.
	EnableDate *time.Time

	// The friendly name identifying the user.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMFADeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetMFADevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetMFADevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMFADevice"); err != nil {
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
	if err = addOpGetMFADeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMFADevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMFADevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMFADevice",
	}
}
