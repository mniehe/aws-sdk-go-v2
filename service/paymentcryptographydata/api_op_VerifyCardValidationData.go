// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptographydata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/paymentcryptographydata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Verifies card-related validation data using algorithms such as Card
// Verification Values (CVV/CVV2), Dynamic Card Verification Values (dCVV/dCVV2)
// and Card Security Codes (CSC). For more information, see Verify card data (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/verify-card-data.html)
// in the Amazon Web Services Payment Cryptography User Guide. This operation
// validates the CVV or CSC codes that is printed on a payment credit or debit card
// during card payment transaction. The input values are typically provided as part
// of an inbound transaction to an issuer or supporting platform partner. Amazon
// Web Services Payment Cryptography uses CVV or CSC, PAN (Primary Account Number)
// and expiration date of the card to check its validity during transaction
// processing. In this operation, the CVK (Card Verification Key) encryption key
// for use with card data verification is same as the one in used for
// GenerateCardValidationData . For information about valid keys for this
// operation, see Understanding key attributes (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html)
// and Key types for specific data operations (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html)
// in the Amazon Web Services Payment Cryptography User Guide. Cross-account use:
// This operation can't be used across different Amazon Web Services accounts.
// Related operations:
//   - GenerateCardValidationData
//   - VerifyAuthRequestCryptogram
//   - VerifyPinData
func (c *Client) VerifyCardValidationData(ctx context.Context, params *VerifyCardValidationDataInput, optFns ...func(*Options)) (*VerifyCardValidationDataOutput, error) {
	if params == nil {
		params = &VerifyCardValidationDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VerifyCardValidationData", params, optFns, c.addOperationVerifyCardValidationDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifyCardValidationDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type VerifyCardValidationDataInput struct {

	// The keyARN of the CVK encryption key that Amazon Web Services Payment
	// Cryptography uses to verify card data.
	//
	// This member is required.
	KeyIdentifier *string

	// The Primary Account Number (PAN), a unique identifier for a payment credit or
	// debit card that associates the card with a specific account holder.
	//
	// This member is required.
	PrimaryAccountNumber *string

	// The CVV or CSC value for use for card data verification within Amazon Web
	// Services Payment Cryptography.
	//
	// This member is required.
	ValidationData *string

	// The algorithm to use for verification of card data within Amazon Web Services
	// Payment Cryptography.
	//
	// This member is required.
	VerificationAttributes types.CardVerificationAttributes

	noSmithyDocumentSerde
}

type VerifyCardValidationDataOutput struct {

	// The keyARN of the CVK encryption key that Amazon Web Services Payment
	// Cryptography uses to verify CVV or CSC.
	//
	// This member is required.
	KeyArn *string

	// The key check value (KCV) of the encryption key. The KCV is used to check if
	// all parties holding a given key have the same key or to detect that a key has
	// changed. Amazon Web Services Payment Cryptography calculates the KCV by using
	// standard algorithms, typically by encrypting 8 or 16 bytes or "00" or "01" and
	// then truncating the result to the first 3 bytes, or 6 hex digits, of the
	// resulting cryptogram.
	//
	// This member is required.
	KeyCheckValue *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVerifyCardValidationDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpVerifyCardValidationData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpVerifyCardValidationData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "VerifyCardValidationData"); err != nil {
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
	if err = addOpVerifyCardValidationDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyCardValidationData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVerifyCardValidationData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "VerifyCardValidationData",
	}
}
