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
	"time"
)

// Provides details about how an DataSync transfer location for Microsoft Azure
// Blob Storage is configured.
func (c *Client) DescribeLocationAzureBlob(ctx context.Context, params *DescribeLocationAzureBlobInput, optFns ...func(*Options)) (*DescribeLocationAzureBlobOutput, error) {
	if params == nil {
		params = &DescribeLocationAzureBlobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationAzureBlob", params, optFns, c.addOperationDescribeLocationAzureBlobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationAzureBlobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLocationAzureBlobInput struct {

	// Specifies the Amazon Resource Name (ARN) of your Azure Blob Storage transfer
	// location.
	//
	// This member is required.
	LocationArn *string

	noSmithyDocumentSerde
}

type DescribeLocationAzureBlobOutput struct {

	// The access tier that you want your objects or files transferred into. This only
	// applies when using the location as a transfer destination. For more information,
	// see Access tiers (https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access-tiers)
	// .
	AccessTier types.AzureAccessTier

	// The ARNs of the DataSync agents that can connect with your Azure Blob Storage
	// container.
	AgentArns []string

	// The authentication method DataSync uses to access your Azure Blob Storage.
	// DataSync can access blob storage using a shared access signature (SAS).
	AuthenticationType types.AzureBlobAuthenticationType

	// The type of blob that you want your objects or files to be when transferring
	// them into Azure Blob Storage. Currently, DataSync only supports moving data into
	// Azure Blob Storage as block blobs. For more information on blob types, see the
	// Azure Blob Storage documentation (https://learn.microsoft.com/en-us/rest/api/storageservices/understanding-block-blobs--append-blobs--and-page-blobs)
	// .
	BlobType types.AzureBlobType

	// The time that your Azure Blob Storage transfer location was created.
	CreationTime *time.Time

	// The ARN of your Azure Blob Storage transfer location.
	LocationArn *string

	// The URL of the Azure Blob Storage container involved in your transfer.
	LocationUri *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocationAzureBlobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationAzureBlob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationAzureBlob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeLocationAzureBlob"); err != nil {
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
	if err = addOpDescribeLocationAzureBlobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationAzureBlob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocationAzureBlob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeLocationAzureBlob",
	}
}
