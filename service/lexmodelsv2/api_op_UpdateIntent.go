// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the settings for an intent.
func (c *Client) UpdateIntent(ctx context.Context, params *UpdateIntentInput, optFns ...func(*Options)) (*UpdateIntentOutput, error) {
	if params == nil {
		params = &UpdateIntentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIntent", params, optFns, c.addOperationUpdateIntentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIntentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIntentInput struct {

	// The identifier of the bot that contains the intent.
	//
	// This member is required.
	BotId *string

	// The version of the bot that contains the intent. Must be DRAFT .
	//
	// This member is required.
	BotVersion *string

	// The unique identifier of the intent to update.
	//
	// This member is required.
	IntentId *string

	// The new name for the intent.
	//
	// This member is required.
	IntentName *string

	// The identifier of the language and locale where this intent is used. The string
	// must match one of the supported locales. For more information, see Supported
	// languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html) .
	//
	// This member is required.
	LocaleId *string

	// The new description of the intent.
	Description *string

	// The new Lambda function to use between each turn of the conversation with the
	// bot.
	DialogCodeHook *types.DialogCodeHookSettings

	// The new Lambda function to call when all of the intents required slots are
	// provided and the intent is ready for fulfillment.
	FulfillmentCodeHook *types.FulfillmentCodeHookSettings

	// Configuration settings for a response sent to the user before Amazon Lex starts
	// eliciting slots.
	InitialResponseSetting *types.InitialResponseSetting

	// A new list of contexts that must be active in order for Amazon Lex to consider
	// the intent.
	InputContexts []types.InputContext

	// The new response that Amazon Lex sends the user when the intent is closed.
	IntentClosingSetting *types.IntentClosingSetting

	// New prompts that Amazon Lex sends to the user to confirm the completion of an
	// intent.
	IntentConfirmationSetting *types.IntentConfirmationSetting

	// New configuration settings for connecting to an Amazon Kendra index.
	KendraConfiguration *types.KendraConfiguration

	// A new list of contexts that Amazon Lex activates when the intent is fulfilled.
	OutputContexts []types.OutputContext

	// The signature of the new built-in intent to use as the parent of this intent.
	ParentIntentSignature *string

	// New utterances used to invoke the intent.
	SampleUtterances []types.SampleUtterance

	// A new list of slots and their priorities that are contained by the intent.
	SlotPriorities []types.SlotPriority

	noSmithyDocumentSerde
}

type UpdateIntentOutput struct {

	// The identifier of the bot that contains the intent.
	BotId *string

	// The version of the bot that contains the intent. Will always be DRAFT .
	BotVersion *string

	// A timestamp of when the intent was created.
	CreationDateTime *time.Time

	// The updated description of the intent.
	Description *string

	// The updated Lambda function called during each turn of the conversation with
	// the user.
	DialogCodeHook *types.DialogCodeHookSettings

	// The updated Lambda function called when the intent is ready for fulfillment.
	FulfillmentCodeHook *types.FulfillmentCodeHookSettings

	// Configuration settings for a response sent to the user before Amazon Lex starts
	// eliciting slots.
	InitialResponseSetting *types.InitialResponseSetting

	// The updated list of contexts that must be active for the intent to be
	// considered by Amazon Lex.
	InputContexts []types.InputContext

	// The updated response that Amazon Lex sends the user when the intent is closed.
	IntentClosingSetting *types.IntentClosingSetting

	// The updated prompts that Amazon Lex sends to the user to confirm the completion
	// of an intent.
	IntentConfirmationSetting *types.IntentConfirmationSetting

	// The identifier of the intent that was updated.
	IntentId *string

	// The updated name of the intent.
	IntentName *string

	// The updated configuration for connecting to an Amazon Kendra index with the
	// AMAZON.KendraSearchIntent intent.
	KendraConfiguration *types.KendraConfiguration

	// A timestamp of the last time that the intent was modified.
	LastUpdatedDateTime *time.Time

	// The updated language and locale of the intent.
	LocaleId *string

	// The updated list of contexts that Amazon Lex activates when the intent is
	// fulfilled.
	OutputContexts []types.OutputContext

	// The updated built-in intent that is the parent of this intent.
	ParentIntentSignature *string

	// The updated list of sample utterances for the intent.
	SampleUtterances []types.SampleUtterance

	// The updated list of slots and their priorities that are elicited from the user
	// for the intent.
	SlotPriorities []types.SlotPriority

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIntentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIntent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIntent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateIntent"); err != nil {
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
	if err = addOpUpdateIntentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIntent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIntent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateIntent",
	}
}
