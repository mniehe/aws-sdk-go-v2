// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates the specified phone numbers from the specified Amazon Chime Voice
// Connector. This API is is no longer supported and will not be updated. We
// recommend using the latest version, DisassociatePhoneNumbersFromVoiceConnector (https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_voice-chime_DisassociatePhoneNumbersFromVoiceConnector.html)
// , in the Amazon Chime SDK. Using the latest version requires migrating to a
// dedicated namespace. For more information, refer to Migrating from the Amazon
// Chime namespace (https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html)
// in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by DisassociatePhoneNumbersFromVoiceConnector in the
// Amazon Chime SDK Voice Namespace
func (c *Client) DisassociatePhoneNumbersFromVoiceConnector(ctx context.Context, params *DisassociatePhoneNumbersFromVoiceConnectorInput, optFns ...func(*Options)) (*DisassociatePhoneNumbersFromVoiceConnectorOutput, error) {
	if params == nil {
		params = &DisassociatePhoneNumbersFromVoiceConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociatePhoneNumbersFromVoiceConnector", params, optFns, c.addOperationDisassociatePhoneNumbersFromVoiceConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociatePhoneNumbersFromVoiceConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociatePhoneNumbersFromVoiceConnectorInput struct {

	// List of phone numbers, in E.164 format.
	//
	// This member is required.
	E164PhoneNumbers []string

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	noSmithyDocumentSerde
}

type DisassociatePhoneNumbersFromVoiceConnectorOutput struct {

	// If the action fails for one or more of the phone numbers in the request, a list
	// of the phone numbers is returned, along with error codes and error messages.
	PhoneNumberErrors []types.PhoneNumberError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociatePhoneNumbersFromVoiceConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociatePhoneNumbersFromVoiceConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociatePhoneNumbersFromVoiceConnector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociatePhoneNumbersFromVoiceConnector"); err != nil {
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
	if err = addOpDisassociatePhoneNumbersFromVoiceConnectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociatePhoneNumbersFromVoiceConnector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociatePhoneNumbersFromVoiceConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociatePhoneNumbersFromVoiceConnector",
	}
}
