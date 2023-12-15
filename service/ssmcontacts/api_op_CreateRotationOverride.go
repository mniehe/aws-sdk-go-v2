// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an override for a rotation in an on-call schedule.
func (c *Client) CreateRotationOverride(ctx context.Context, params *CreateRotationOverrideInput, optFns ...func(*Options)) (*CreateRotationOverrideOutput, error) {
	if params == nil {
		params = &CreateRotationOverrideInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRotationOverride", params, optFns, c.addOperationCreateRotationOverrideMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRotationOverrideOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRotationOverrideInput struct {

	// The date and time when the override ends.
	//
	// This member is required.
	EndTime *time.Time

	// The Amazon Resource Names (ARNs) of the contacts to replace those in the
	// current on-call rotation with. If you want to include any current team members
	// in the override shift, you must include their ARNs in the new contact ID list.
	//
	// This member is required.
	NewContactIds []string

	// The Amazon Resource Name (ARN) of the rotation to create an override for.
	//
	// This member is required.
	RotationId *string

	// The date and time when the override goes into effect.
	//
	// This member is required.
	StartTime *time.Time

	// A token that ensures that the operation is called only once with the specified
	// details.
	IdempotencyToken *string

	noSmithyDocumentSerde
}

type CreateRotationOverrideOutput struct {

	// The Amazon Resource Name (ARN) of the created rotation override.
	//
	// This member is required.
	RotationOverrideId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRotationOverrideMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRotationOverride{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRotationOverride{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRotationOverride"); err != nil {
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
	if err = addOpCreateRotationOverrideValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRotationOverride(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRotationOverride(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRotationOverride",
	}
}
