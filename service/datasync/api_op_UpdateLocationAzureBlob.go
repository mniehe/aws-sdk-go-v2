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

// Modifies some configurations of the Microsoft Azure Blob Storage transfer
// location that you're using with DataSync.
func (c *Client) UpdateLocationAzureBlob(ctx context.Context, params *UpdateLocationAzureBlobInput, optFns ...func(*Options)) (*UpdateLocationAzureBlobOutput, error) {
	if params == nil {
		params = &UpdateLocationAzureBlobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationAzureBlob", params, optFns, c.addOperationUpdateLocationAzureBlobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationAzureBlobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationAzureBlobInput struct {

	// Specifies the ARN of the Azure Blob Storage transfer location that you're
	// updating.
	//
	// This member is required.
	LocationArn *string

	// Specifies the access tier that you want your objects or files transferred into.
	// This only applies when using the location as a transfer destination. For more
	// information, see Access tiers (https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access-tiers)
	// .
	AccessTier types.AzureAccessTier

	// Specifies the Amazon Resource Name (ARN) of the DataSync agent that can connect
	// with your Azure Blob Storage container. You can specify more than one agent. For
	// more information, see Using multiple agents for your transfer (https://docs.aws.amazon.com/datasync/latest/userguide/multiple-agents.html)
	// .
	AgentArns []string

	// Specifies the authentication method DataSync uses to access your Azure Blob
	// Storage. DataSync can access blob storage using a shared access signature (SAS).
	AuthenticationType types.AzureBlobAuthenticationType

	// Specifies the type of blob that you want your objects or files to be when
	// transferring them into Azure Blob Storage. Currently, DataSync only supports
	// moving data into Azure Blob Storage as block blobs. For more information on blob
	// types, see the Azure Blob Storage documentation (https://learn.microsoft.com/en-us/rest/api/storageservices/understanding-block-blobs--append-blobs--and-page-blobs)
	// .
	BlobType types.AzureBlobType

	// Specifies the SAS configuration that allows DataSync to access your Azure Blob
	// Storage.
	SasConfiguration *types.AzureBlobSasConfiguration

	// Specifies path segments if you want to limit your transfer to a virtual
	// directory in your container (for example, /my/images ).
	Subdirectory *string

	noSmithyDocumentSerde
}

type UpdateLocationAzureBlobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationAzureBlobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationAzureBlob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationAzureBlob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLocationAzureBlob"); err != nil {
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
	if err = addOpUpdateLocationAzureBlobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationAzureBlob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLocationAzureBlob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLocationAzureBlob",
	}
}
