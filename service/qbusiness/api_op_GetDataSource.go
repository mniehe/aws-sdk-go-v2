// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/qbusiness/document"
	"github.com/mniehe/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about an existing Amazon Q data source connector.
func (c *Client) GetDataSource(ctx context.Context, params *GetDataSourceInput, optFns ...func(*Options)) (*GetDataSourceOutput, error) {
	if params == nil {
		params = &GetDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataSource", params, optFns, c.addOperationGetDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataSourceInput struct {

	// The identifier of the Amazon Q application.
	//
	// This member is required.
	ApplicationId *string

	// The identifier of the data source connector.
	//
	// This member is required.
	DataSourceId *string

	// The identfier of the index used with the data source connector.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type GetDataSourceOutput struct {

	// The identifier of the Amazon Q application.
	ApplicationId *string

	// The details of how the data source connector is configured.
	Configuration document.Interface

	// The Unix timestamp when the data source connector was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the data source.
	DataSourceArn *string

	// The identifier of the data source connector.
	DataSourceId *string

	// The description for the data source connector.
	Description *string

	// The name for the data source connector.
	DisplayName *string

	// Provides the configuration information for altering document metadata and
	// content during the document ingestion process. For more information, see Custom
	// document enrichment (https://docs.aws.amazon.com/amazonq/latest/business-use-dg/custom-document-enrichment.html)
	// .
	DocumentEnrichmentConfiguration *types.DocumentEnrichmentConfiguration

	// When the Status field value is FAILED , the ErrorMessage field contains a
	// description of the error that caused the data source connector to fail.
	Error *types.ErrorDetail

	// The identifier of the index linked to the data source connector.
	IndexId *string

	// The Amazon Resource Name (ARN) of the role with permission to access the data
	// source and required resources.
	RoleArn *string

	// The current status of the data source connector. When the Status field value is
	// FAILED , the ErrorMessage field contains a description of the error that caused
	// the data source connector to fail.
	Status types.DataSourceStatus

	// The schedule for Amazon Q to update the index.
	SyncSchedule *string

	// The type of the data source connector. For example, S3 .
	Type *string

	// The Unix timestamp when the data source connector was last updated.
	UpdatedAt *time.Time

	// Configuration information for an Amazon VPC (Virtual Private Cloud) to connect
	// to your data source.
	VpcConfiguration *types.DataSourceVpcConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDataSource"); err != nil {
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
	if err = addOpGetDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDataSource",
	}
}
