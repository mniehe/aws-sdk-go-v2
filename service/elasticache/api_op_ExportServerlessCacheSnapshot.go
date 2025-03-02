// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides the functionality to export the serverless cache snapshot data to
// Amazon S3. Available for Redis only.
func (c *Client) ExportServerlessCacheSnapshot(ctx context.Context, params *ExportServerlessCacheSnapshotInput, optFns ...func(*Options)) (*ExportServerlessCacheSnapshotOutput, error) {
	if params == nil {
		params = &ExportServerlessCacheSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportServerlessCacheSnapshot", params, optFns, c.addOperationExportServerlessCacheSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportServerlessCacheSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportServerlessCacheSnapshotInput struct {

	// Name of the Amazon S3 bucket to export the snapshot to. The Amazon S3 bucket
	// must also be in same region as the snapshot. Available for Redis only.
	//
	// This member is required.
	S3BucketName *string

	// The identifier of the serverless cache snapshot to be exported to S3. Available
	// for Redis only.
	//
	// This member is required.
	ServerlessCacheSnapshotName *string

	noSmithyDocumentSerde
}

type ExportServerlessCacheSnapshotOutput struct {

	// The state of a serverless cache at a specific point in time, to the
	// millisecond. Available for Redis only.
	ServerlessCacheSnapshot *types.ServerlessCacheSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportServerlessCacheSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpExportServerlessCacheSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpExportServerlessCacheSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExportServerlessCacheSnapshot"); err != nil {
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
	if err = addOpExportServerlessCacheSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportServerlessCacheSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportServerlessCacheSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExportServerlessCacheSnapshot",
	}
}
