// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a unique symmetric data key for use outside of KMS. This operation
// returns a data key that is encrypted under a symmetric encryption KMS key that
// you specify. The bytes in the key are random; they are not related to the caller
// or to the KMS key. GenerateDataKeyWithoutPlaintext is identical to the
// GenerateDataKey operation except that it does not return a plaintext copy of the
// data key. This operation is useful for systems that need to encrypt data at some
// point, but not immediately. When you need to encrypt the data, you call the
// Decrypt operation on the encrypted copy of the key. It's also useful in
// distributed systems with different levels of trust. For example, you might store
// encrypted data in containers. One component of your system creates new
// containers and stores an encrypted data key with each container. Then, a
// different component puts the data into the containers. That component first
// decrypts the data key, uses the plaintext data key to encrypt data, puts the
// encrypted data into the container, and then destroys the plaintext data key. In
// this system, the component that creates the containers never sees the plaintext
// data key. To request an asymmetric data key pair, use the GenerateDataKeyPair
// or GenerateDataKeyPairWithoutPlaintext operations. To generate a data key, you
// must specify the symmetric encryption KMS key that is used to encrypt the data
// key. You cannot use an asymmetric KMS key or a key in a custom key store to
// generate a data key. To get the type of your KMS key, use the DescribeKey
// operation. You must also specify the length of the data key. Use either the
// KeySpec or NumberOfBytes parameters (but not both). For 128-bit and 256-bit
// data keys, use the KeySpec parameter. To generate an SM4 data key (China
// Regions only), specify a KeySpec value of AES_128 or NumberOfBytes value of 16 .
// The symmetric encryption key used in China Regions to encrypt your data key is
// an SM4 encryption key. If the operation succeeds, you will find the encrypted
// copy of the data key in the CiphertextBlob field. You can use an optional
// encryption context to add additional security to the encryption operation. If
// you specify an EncryptionContext , you must specify the same encryption context
// (a case-sensitive exact match) when decrypting the encrypted data key.
// Otherwise, the request to decrypt fails with an InvalidCiphertextException . For
// more information, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
// in the Key Management Service Developer Guide. The KMS key that you use for this
// operation must be in a compatible key state. For details, see Key states of KMS
// keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in
// the Key Management Service Developer Guide. Cross-account use: Yes. To perform
// this operation with a KMS key in a different Amazon Web Services account,
// specify the key ARN or alias ARN in the value of the KeyId parameter. Required
// permissions: kms:GenerateDataKeyWithoutPlaintext (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations:
//   - Decrypt
//   - Encrypt
//   - GenerateDataKey
//   - GenerateDataKeyPair
//   - GenerateDataKeyPairWithoutPlaintext
func (c *Client) GenerateDataKeyWithoutPlaintext(ctx context.Context, params *GenerateDataKeyWithoutPlaintextInput, optFns ...func(*Options)) (*GenerateDataKeyWithoutPlaintextOutput, error) {
	if params == nil {
		params = &GenerateDataKeyWithoutPlaintextInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateDataKeyWithoutPlaintext", params, optFns, c.addOperationGenerateDataKeyWithoutPlaintextMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateDataKeyWithoutPlaintextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateDataKeyWithoutPlaintextInput struct {

	// Specifies the symmetric encryption KMS key that encrypts the data key. You
	// cannot specify an asymmetric KMS key or a KMS key in a custom key store. To get
	// the type and origin of your KMS key, use the DescribeKey operation. To specify
	// a KMS key, use its key ID, key ARN, alias name, or alias ARN. When using an
	// alias name, prefix it with "alias/" . To specify a KMS key in a different Amazon
	// Web Services account, you must use the key ARN or alias ARN. For example:
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Alias name: alias/ExampleAlias
	//   - Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey . To
	// get the alias name and alias ARN, use ListAliases .
	//
	// This member is required.
	KeyId *string

	// Checks if your request will succeed. DryRun is an optional parameter. To learn
	// more about how to use this parameter, see Testing your KMS API calls (https://docs.aws.amazon.com/kms/latest/developerguide/programming-dryrun.html)
	// in the Key Management Service Developer Guide.
	DryRun *bool

	// Specifies the encryption context that will be used when encrypting the data
	// key. Do not include confidential or sensitive information in this field. This
	// field may be displayed in plaintext in CloudTrail logs and other output. An
	// encryption context is a collection of non-secret key-value pairs that represent
	// additional authenticated data. When you use an encryption context to encrypt
	// data, you must specify the same (an exact case-sensitive match) encryption
	// context to decrypt the data. An encryption context is supported only on
	// operations with symmetric encryption KMS keys. On operations with symmetric
	// encryption KMS keys, an encryption context is optional, but it is strongly
	// recommended. For more information, see Encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the Key Management Service Developer Guide.
	EncryptionContext map[string]string

	// A list of grant tokens. Use a grant token when your permission to call this
	// operation comes from a new grant that has not yet achieved eventual consistency.
	// For more information, see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)
	// and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)
	// in the Key Management Service Developer Guide.
	GrantTokens []string

	// The length of the data key. Use AES_128 to generate a 128-bit symmetric key, or
	// AES_256 to generate a 256-bit symmetric key.
	KeySpec types.DataKeySpec

	// The length of the data key in bytes. For example, use the value 64 to generate
	// a 512-bit data key (64 bytes is 512 bits). For common key lengths (128-bit and
	// 256-bit symmetric keys), we recommend that you use the KeySpec field instead of
	// this one.
	NumberOfBytes *int32

	noSmithyDocumentSerde
}

type GenerateDataKeyWithoutPlaintextOutput struct {

	// The encrypted data key. When you use the HTTP API or the Amazon Web Services
	// CLI, the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	CiphertextBlob []byte

	// The Amazon Resource Name ( key ARN (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN)
	// ) of the KMS key that encrypted the data key.
	KeyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateDataKeyWithoutPlaintextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateDataKeyWithoutPlaintext{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateDataKeyWithoutPlaintext{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GenerateDataKeyWithoutPlaintext"); err != nil {
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
	if err = addOpGenerateDataKeyWithoutPlaintextValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateDataKeyWithoutPlaintext(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateDataKeyWithoutPlaintext(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GenerateDataKeyWithoutPlaintext",
	}
}
