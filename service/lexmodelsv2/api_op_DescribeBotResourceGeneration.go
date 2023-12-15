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

// Returns information about a request to generate a bot through natural language
// description, made through the StartBotResource API. Use the
// generatedBotLocaleUrl to retrieve the Amazon S3 object containing the bot locale
// configuration. You can then modify and import this configuration.
func (c *Client) DescribeBotResourceGeneration(ctx context.Context, params *DescribeBotResourceGenerationInput, optFns ...func(*Options)) (*DescribeBotResourceGenerationOutput, error) {
	if params == nil {
		params = &DescribeBotResourceGenerationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBotResourceGeneration", params, optFns, c.addOperationDescribeBotResourceGenerationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBotResourceGenerationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBotResourceGenerationInput struct {

	// The unique identifier of the bot for which to return the generation details.
	//
	// This member is required.
	BotId *string

	// The version of the bot for which to return the generation details.
	//
	// This member is required.
	BotVersion *string

	// The unique identifier of the generation request for which to return the
	// generation details.
	//
	// This member is required.
	GenerationId *string

	// The locale of the bot for which to return the generation details.
	//
	// This member is required.
	LocaleId *string

	noSmithyDocumentSerde
}

type DescribeBotResourceGenerationOutput struct {

	// The unique identifier of the bot for which the generation request was made.
	BotId *string

	// The version of the bot for which the generation request was made.
	BotVersion *string

	// The date and time at which the item was generated.
	CreationDateTime *time.Time

	// A list of reasons why the generation of bot resources through natural language
	// description failed.
	FailureReasons []string

	// The Amazon S3 location of the generated bot locale configuration.
	GeneratedBotLocaleUrl *string

	// The generation ID for which to return the generation details.
	GenerationId *string

	// The prompt used in the generation request.
	GenerationInputPrompt *string

	// The status of the generation request.
	GenerationStatus types.GenerationStatus

	// The date and time at which the generated item was updated.
	LastUpdatedDateTime *time.Time

	// The locale of the bot for which the generation request was made.
	LocaleId *string

	// The ARN of the model used to generate the bot resources.
	ModelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBotResourceGenerationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBotResourceGeneration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBotResourceGeneration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBotResourceGeneration"); err != nil {
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
	if err = addOpDescribeBotResourceGenerationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBotResourceGeneration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBotResourceGeneration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBotResourceGeneration",
	}
}
