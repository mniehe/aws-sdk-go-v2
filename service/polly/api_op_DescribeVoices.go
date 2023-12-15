// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/polly/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of voices that are available for use when requesting speech
// synthesis. Each voice speaks a specified language, is either male or female, and
// is identified by an ID, which is the ASCII version of the voice name. When
// synthesizing speech ( SynthesizeSpeech ), you provide the voice ID for the
// voice you want from the list of voices returned by DescribeVoices . For example,
// you want your news reader application to read news in a specific language, but
// giving a user the option to choose the voice. Using the DescribeVoices
// operation you can provide the user with a list of available voices to select
// from. You can optionally specify a language code to filter the available voices.
// For example, if you specify en-US , the operation returns a list of all
// available US English voices. This operation requires permissions to perform the
// polly:DescribeVoices action.
func (c *Client) DescribeVoices(ctx context.Context, params *DescribeVoicesInput, optFns ...func(*Options)) (*DescribeVoicesOutput, error) {
	if params == nil {
		params = &DescribeVoicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVoices", params, optFns, c.addOperationDescribeVoicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVoicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVoicesInput struct {

	// Specifies the engine ( standard , neural or long-form ) used by Amazon Polly
	// when processing input text for speech synthesis.
	Engine types.Engine

	// Boolean value indicating whether to return any bilingual voices that use the
	// specified language as an additional language. For instance, if you request all
	// languages that use US English (es-US), and there is an Italian voice that speaks
	// both Italian (it-IT) and US English, that voice will be included if you specify
	// yes but not if you specify no .
	IncludeAdditionalLanguageCodes bool

	// The language identification tag (ISO 639 code for the language name-ISO 3166
	// country code) for filtering the list of voices returned. If you don't specify
	// this optional parameter, all available voices are returned.
	LanguageCode types.LanguageCode

	// An opaque pagination token returned from the previous DescribeVoices operation.
	// If present, this indicates where to continue the listing.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeVoicesOutput struct {

	// The pagination token to use in the next request to continue the listing of
	// voices. NextToken is returned only if the response is truncated.
	NextToken *string

	// A list of voices with their properties.
	Voices []types.Voice

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVoicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeVoices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeVoices{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeVoices"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVoices(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVoices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeVoices",
	}
}
