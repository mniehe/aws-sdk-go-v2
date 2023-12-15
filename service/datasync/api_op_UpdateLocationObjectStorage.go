// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates some parameters of an existing object storage location that DataSync
// accesses for a transfer. For information about creating a self-managed object
// storage location, see Creating a location for object storage (https://docs.aws.amazon.com/datasync/latest/userguide/create-object-location.html)
// .
func (c *Client) UpdateLocationObjectStorage(ctx context.Context, params *UpdateLocationObjectStorageInput, optFns ...func(*Options)) (*UpdateLocationObjectStorageOutput, error) {
	if params == nil {
		params = &UpdateLocationObjectStorageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationObjectStorage", params, optFns, c.addOperationUpdateLocationObjectStorageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationObjectStorageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationObjectStorageInput struct {

	// Specifies the ARN of the object storage system location that you're updating.
	//
	// This member is required.
	LocationArn *string

	// Specifies the access key (for example, a user name) if credentials are required
	// to authenticate with the object storage server.
	AccessKey *string

	// Specifies the Amazon Resource Names (ARNs) of the DataSync agents that can
	// securely connect with your location.
	AgentArns []string

	// Specifies the secret key (for example, a password) if credentials are required
	// to authenticate with the object storage server.
	SecretKey *string

	// Specifies a certificate to authenticate with an object storage system that uses
	// a private or self-signed certificate authority (CA). You must specify a
	// Base64-encoded .pem file (for example,
	// file:///home/user/.ssh/storage_sys_certificate.pem ). The certificate can be up
	// to 32768 bytes (before Base64 encoding). To use this parameter, configure
	// ServerProtocol to HTTPS . Updating the certificate doesn't interfere with tasks
	// that you have in progress.
	ServerCertificate []byte

	// Specifies the port that your object storage server accepts inbound network
	// traffic on (for example, port 443).
	ServerPort *int32

	// Specifies the protocol that your object storage server uses to communicate.
	ServerProtocol types.ObjectStorageServerProtocol

	// Specifies the object prefix for your object storage server. If this is a source
	// location, DataSync only copies objects with this prefix. If this is a
	// destination location, DataSync writes all objects with this prefix.
	Subdirectory *string

	noSmithyDocumentSerde
}

type UpdateLocationObjectStorageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationObjectStorageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationObjectStorage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationObjectStorage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLocationObjectStorage"); err != nil {
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
	if err = addOpUpdateLocationObjectStorageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationObjectStorage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLocationObjectStorage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLocationObjectStorage",
	}
}
