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

// Verifies a digital signature that was generated by the Sign operation.
// Verification confirms that an authorized user signed the message with the
// specified KMS key and signing algorithm, and the message hasn't changed since it
// was signed. If the signature is verified, the value of the SignatureValid field
// in the response is True . If the signature verification fails, the Verify
// operation fails with an KMSInvalidSignatureException exception. A digital
// signature is generated by using the private key in an asymmetric KMS key. The
// signature is verified by using the public key in the same asymmetric KMS key.
// For information about asymmetric KMS keys, see Asymmetric KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
// in the Key Management Service Developer Guide. To use the Verify operation,
// specify the same asymmetric KMS key, message, and signing algorithm that were
// used to produce the signature. The message type does not need to be the same as
// the one used for signing, but it must indicate whether the value of the Message
// parameter should be hashed as part of the verification process. You can also
// verify the digital signature by using the public key of the KMS key outside of
// KMS. Use the GetPublicKey operation to download the public key in the
// asymmetric KMS key and then use the public key to verify the signature outside
// of KMS. The advantage of using the Verify operation is that it is performed
// within KMS. As a result, it's easy to call, the operation is performed within
// the FIPS boundary, it is logged in CloudTrail, and you can use key policy and
// IAM policy to determine who is authorized to use the KMS key to verify
// signatures. To verify a signature outside of KMS with an SM2 public key (China
// Regions only), you must specify the distinguishing ID. By default, KMS uses
// 1234567812345678 as the distinguishing ID. For more information, see Offline
// verification with SM2 key pairs (https://docs.aws.amazon.com/kms/latest/developerguide/asymmetric-key-specs.html#key-spec-sm-offline-verification)
// . The KMS key that you use for this operation must be in a compatible key state.
// For details, see Key states of KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the Key Management Service Developer Guide. Cross-account use: Yes. To
// perform this operation with a KMS key in a different Amazon Web Services
// account, specify the key ARN or alias ARN in the value of the KeyId parameter.
// Required permissions: kms:Verify (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations: Sign
func (c *Client) Verify(ctx context.Context, params *VerifyInput, optFns ...func(*Options)) (*VerifyOutput, error) {
	if params == nil {
		params = &VerifyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Verify", params, optFns, c.addOperationVerifyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type VerifyInput struct {

	// Identifies the asymmetric KMS key that will be used to verify the signature.
	// This must be the same KMS key that was used to generate the signature. If you
	// specify a different KMS key, the signature verification fails. To specify a KMS
	// key, use its key ID, key ARN, alias name, or alias ARN. When using an alias
	// name, prefix it with "alias/" . To specify a KMS key in a different Amazon Web
	// Services account, you must use the key ARN or alias ARN. For example:
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

	// Specifies the message that was signed. You can submit a raw message of up to
	// 4096 bytes, or a hash digest of the message. If you submit a digest, use the
	// MessageType parameter with a value of DIGEST . If the message specified here is
	// different from the message that was signed, the signature verification fails. A
	// message and its hash digest are considered to be the same message.
	//
	// This member is required.
	Message []byte

	// The signature that the Sign operation generated.
	//
	// This member is required.
	Signature []byte

	// The signing algorithm that was used to sign the message. If you submit a
	// different algorithm, the signature verification fails.
	//
	// This member is required.
	SigningAlgorithm types.SigningAlgorithmSpec

	// Checks if your request will succeed. DryRun is an optional parameter. To learn
	// more about how to use this parameter, see Testing your KMS API calls (https://docs.aws.amazon.com/kms/latest/developerguide/programming-dryrun.html)
	// in the Key Management Service Developer Guide.
	DryRun *bool

	// A list of grant tokens. Use a grant token when your permission to call this
	// operation comes from a new grant that has not yet achieved eventual consistency.
	// For more information, see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)
	// and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)
	// in the Key Management Service Developer Guide.
	GrantTokens []string

	// Tells KMS whether the value of the Message parameter should be hashed as part
	// of the signing algorithm. Use RAW for unhashed messages; use DIGEST for message
	// digests, which are already hashed. When the value of MessageType is RAW , KMS
	// uses the standard signing algorithm, which begins with a hash function. When the
	// value is DIGEST , KMS skips the hashing step in the signing algorithm. Use the
	// DIGEST value only when the value of the Message parameter is a message digest.
	// If you use the DIGEST value with an unhashed message, the security of the
	// verification operation can be compromised. When the value of MessageType is
	// DIGEST , the length of the Message value must match the length of hashed
	// messages for the specified signing algorithm. You can submit a message digest
	// and omit the MessageType or specify RAW so the digest is hashed again while
	// signing. However, if the signed message is hashed once while signing, but twice
	// while verifying, verification fails, even when the message hasn't changed. The
	// hashing algorithm in that Verify uses is based on the SigningAlgorithm value.
	//   - Signing algorithms that end in SHA_256 use the SHA_256 hashing algorithm.
	//   - Signing algorithms that end in SHA_384 use the SHA_384 hashing algorithm.
	//   - Signing algorithms that end in SHA_512 use the SHA_512 hashing algorithm.
	//   - SM2DSA uses the SM3 hashing algorithm. For details, see Offline
	//   verification with SM2 key pairs (https://docs.aws.amazon.com/kms/latest/developerguide/asymmetric-key-specs.html#key-spec-sm-offline-verification)
	//   .
	MessageType types.MessageType

	noSmithyDocumentSerde
}

type VerifyOutput struct {

	// The Amazon Resource Name ( key ARN (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN)
	// ) of the asymmetric KMS key that was used to verify the signature.
	KeyId *string

	// A Boolean value that indicates whether the signature was verified. A value of
	// True indicates that the Signature was produced by signing the Message with the
	// specified KeyID and SigningAlgorithm. If the signature is not verified, the
	// Verify operation fails with a KMSInvalidSignatureException exception.
	SignatureValid bool

	// The signing algorithm that was used to verify the signature.
	SigningAlgorithm types.SigningAlgorithmSpec

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVerifyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpVerify{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpVerify{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "Verify"); err != nil {
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
	if err = addOpVerifyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVerify(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVerify(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "Verify",
	}
}
