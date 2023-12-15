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

// Starts a request for the descriptive bot builder to generate a bot locale
// configuration based on the prompt you provide it. After you make this call, use
// the DescribeBotResourceGeneration operation to check on the status of the
// generation and for the generatedBotLocaleUrl when the generation is complete.
// Use that value to retrieve the Amazon S3 object containing the bot locale
// configuration. You can then modify and import this configuration.
func (c *Client) StartBotResourceGeneration(ctx context.Context, params *StartBotResourceGenerationInput, optFns ...func(*Options)) (*StartBotResourceGenerationOutput, error) {
	if params == nil {
		params = &StartBotResourceGenerationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartBotResourceGeneration", params, optFns, c.addOperationStartBotResourceGenerationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartBotResourceGenerationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBotResourceGenerationInput struct {

	// The unique identifier of the bot for which to generate intents and slot types.
	//
	// This member is required.
	BotId *string

	// The version of the bot for which to generate intents and slot types.
	//
	// This member is required.
	BotVersion *string

	// The prompt to generate intents and slot types for the bot locale. Your
	// description should be both detailed and precise to help generate appropriate and
	// sufficient intents for your bot. Include a list of actions to improve the intent
	// creation process.
	//
	// This member is required.
	GenerationInputPrompt *string

	// The locale of the bot for which to generate intents and slot types.
	//
	// This member is required.
	LocaleId *string

	noSmithyDocumentSerde
}

type StartBotResourceGenerationOutput struct {

	// The unique identifier of the bot for which the generation request was made.
	BotId *string

	// The version of the bot for which the generation request was made.
	BotVersion *string

	// The date and time at which the generation request was made.
	CreationDateTime *time.Time

	// The unique identifier of the generation request.
	GenerationId *string

	// The prompt that was used generate intents and slot types for the bot locale.
	GenerationInputPrompt *string

	// The status of the generation request.
	GenerationStatus types.GenerationStatus

	// The locale of the bot for which the generation request was made.
	LocaleId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartBotResourceGenerationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartBotResourceGeneration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartBotResourceGeneration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartBotResourceGeneration"); err != nil {
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
	if err = addOpStartBotResourceGenerationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartBotResourceGeneration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartBotResourceGeneration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartBotResourceGeneration",
	}
}
