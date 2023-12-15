// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptography

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/paymentcryptography/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the key material and metadata associated with Amazon Web Services
// Payment Cryptography key. Key deletion is irreversible. After a key is deleted,
// you can't perform cryptographic operations using the key. For example, you can't
// decrypt data that was encrypted by a deleted Amazon Web Services Payment
// Cryptography key, and the data may become unrecoverable. Because key deletion is
// destructive, Amazon Web Services Payment Cryptography has a safety mechanism to
// prevent accidental deletion of a key. When you call this operation, Amazon Web
// Services Payment Cryptography disables the specified key but doesn't delete it
// until after a waiting period set using DeleteKeyInDays . The default waiting
// period is 7 days. During the waiting period, the KeyState is DELETE_PENDING .
// After the key is deleted, the KeyState is DELETE_COMPLETE . You should delete a
// key only when you are sure that you don't need to use it anymore and no other
// parties are utilizing this key. If you aren't sure, consider deactivating it
// instead by calling StopKeyUsage . Cross-account use: This operation can't be
// used across different Amazon Web Services accounts. Related operations:
//   - RestoreKey
//   - StartKeyUsage
//   - StopKeyUsage
func (c *Client) DeleteKey(ctx context.Context, params *DeleteKeyInput, optFns ...func(*Options)) (*DeleteKeyOutput, error) {
	if params == nil {
		params = &DeleteKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteKey", params, optFns, c.addOperationDeleteKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteKeyInput struct {

	// The KeyARN of the key that is scheduled for deletion.
	//
	// This member is required.
	KeyIdentifier *string

	// The waiting period for key deletion. The default value is seven days.
	DeleteKeyInDays *int32

	noSmithyDocumentSerde
}

type DeleteKeyOutput struct {

	// The KeyARN of the key that is scheduled for deletion.
	//
	// This member is required.
	Key *types.Key

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteKey"); err != nil {
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
	if err = addOpDeleteKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteKey",
	}
}
