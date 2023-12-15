// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used to enable or disable feedback forwarding for an identity. This setting
// determines what happens when an identity is used to send an email that results
// in a bounce or complaint event. When you enable feedback forwarding, Amazon
// Pinpoint sends you email notifications when bounce or complaint events occur.
// Amazon Pinpoint sends this notification to the address that you specified in the
// Return-Path header of the original email. When you disable feedback forwarding,
// Amazon Pinpoint sends notifications through other mechanisms, such as by
// notifying an Amazon SNS topic. You're required to have a method of tracking
// bounces and complaints. If you haven't set up another mechanism for receiving
// bounce or complaint notifications, Amazon Pinpoint sends an email notification
// when these events occur (even if this setting is disabled).
func (c *Client) PutEmailIdentityFeedbackAttributes(ctx context.Context, params *PutEmailIdentityFeedbackAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityFeedbackAttributesOutput, error) {
	if params == nil {
		params = &PutEmailIdentityFeedbackAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEmailIdentityFeedbackAttributes", params, optFns, c.addOperationPutEmailIdentityFeedbackAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEmailIdentityFeedbackAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to set the attributes that control how bounce and complaint events
// are processed.
type PutEmailIdentityFeedbackAttributesInput struct {

	// The email identity that you want to configure bounce and complaint feedback
	// forwarding for.
	//
	// This member is required.
	EmailIdentity *string

	// Sets the feedback forwarding configuration for the identity. If the value is
	// true , Amazon Pinpoint sends you email notifications when bounce or complaint
	// events occur. Amazon Pinpoint sends this notification to the address that you
	// specified in the Return-Path header of the original email. When you set this
	// value to false , Amazon Pinpoint sends notifications through other mechanisms,
	// such as by notifying an Amazon SNS topic or another event destination. You're
	// required to have a method of tracking bounces and complaints. If you haven't set
	// up another mechanism for receiving bounce or complaint notifications, Amazon
	// Pinpoint sends an email notification when these events occur (even if this
	// setting is disabled).
	EmailForwardingEnabled bool

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutEmailIdentityFeedbackAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEmailIdentityFeedbackAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutEmailIdentityFeedbackAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEmailIdentityFeedbackAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutEmailIdentityFeedbackAttributes"); err != nil {
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
	if err = addOpPutEmailIdentityFeedbackAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEmailIdentityFeedbackAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEmailIdentityFeedbackAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutEmailIdentityFeedbackAttributes",
	}
}
