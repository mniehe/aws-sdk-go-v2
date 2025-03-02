// Code generated by smithy-go-codegen DO NOT EDIT.

package controltower

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/controltower/document"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API call updates the landing zone. It starts an asynchronous operation
// that updates the landing zone based on the new landing zone version, or on the
// changed parameters specified in the updated manifest file.
func (c *Client) UpdateLandingZone(ctx context.Context, params *UpdateLandingZoneInput, optFns ...func(*Options)) (*UpdateLandingZoneOutput, error) {
	if params == nil {
		params = &UpdateLandingZoneInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLandingZone", params, optFns, c.addOperationUpdateLandingZoneMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLandingZoneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLandingZoneInput struct {

	// The unique identifier of the landing zone.
	//
	// This member is required.
	LandingZoneIdentifier *string

	// The manifest JSON file is a text file that describes your Amazon Web Services
	// resources. For examples, review Launch your landing zone (https://docs.aws.amazon.com/controltower/latest/userguide/lz-api-launch)
	// .
	//
	// This member is required.
	Manifest document.Interface

	// The landing zone version, for example, 3.2.
	//
	// This member is required.
	Version *string

	noSmithyDocumentSerde
}

type UpdateLandingZoneOutput struct {

	// A unique identifier assigned to a UpdateLandingZone operation. You can use this
	// identifier as an input of GetLandingZoneOperation to check the operation's
	// status.
	//
	// This member is required.
	OperationIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLandingZoneMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLandingZone{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLandingZone{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLandingZone"); err != nil {
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
	if err = addOpUpdateLandingZoneValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLandingZone(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLandingZone(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLandingZone",
	}
}
