// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates one or more partitions in a batch operation.
func (c *Client) BatchUpdatePartition(ctx context.Context, params *BatchUpdatePartitionInput, optFns ...func(*Options)) (*BatchUpdatePartitionOutput, error) {
	if params == nil {
		params = &BatchUpdatePartitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdatePartition", params, optFns, c.addOperationBatchUpdatePartitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdatePartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdatePartitionInput struct {

	// The name of the metadata database in which the partition is to be updated.
	//
	// This member is required.
	DatabaseName *string

	// A list of up to 100 BatchUpdatePartitionRequestEntry objects to update.
	//
	// This member is required.
	Entries []types.BatchUpdatePartitionRequestEntry

	// The name of the metadata table in which the partition is to be updated.
	//
	// This member is required.
	TableName *string

	// The ID of the catalog in which the partition is to be updated. Currently, this
	// should be the Amazon Web Services account ID.
	CatalogId *string

	noSmithyDocumentSerde
}

type BatchUpdatePartitionOutput struct {

	// The errors encountered when trying to update the requested partitions. A list
	// of BatchUpdatePartitionFailureEntry objects.
	Errors []types.BatchUpdatePartitionFailureEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdatePartitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchUpdatePartition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchUpdatePartition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchUpdatePartition"); err != nil {
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
	if err = addOpBatchUpdatePartitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdatePartition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdatePartition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchUpdatePartition",
	}
}
