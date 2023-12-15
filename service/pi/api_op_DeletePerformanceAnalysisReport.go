// Code generated by smithy-go-codegen DO NOT EDIT.

package pi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/pi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a performance analysis report.
func (c *Client) DeletePerformanceAnalysisReport(ctx context.Context, params *DeletePerformanceAnalysisReportInput, optFns ...func(*Options)) (*DeletePerformanceAnalysisReportOutput, error) {
	if params == nil {
		params = &DeletePerformanceAnalysisReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePerformanceAnalysisReport", params, optFns, c.addOperationDeletePerformanceAnalysisReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePerformanceAnalysisReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePerformanceAnalysisReportInput struct {

	// The unique identifier of the analysis report for deletion.
	//
	// This member is required.
	AnalysisReportId *string

	// An immutable identifier for a data source that is unique for an Amazon Web
	// Services Region. Performance Insights gathers metrics from this data source. In
	// the console, the identifier is shown as ResourceID. When you call
	// DescribeDBInstances , the identifier is returned as DbiResourceId . To use a DB
	// instance as a data source, specify its DbiResourceId value. For example,
	// specify db-ABCDEFGHIJKLMNOPQRSTU1VW2X .
	//
	// This member is required.
	Identifier *string

	// The Amazon Web Services service for which Performance Insights will return
	// metrics. Valid value is RDS .
	//
	// This member is required.
	ServiceType types.ServiceType

	noSmithyDocumentSerde
}

type DeletePerformanceAnalysisReportOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePerformanceAnalysisReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePerformanceAnalysisReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePerformanceAnalysisReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeletePerformanceAnalysisReport"); err != nil {
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
	if err = addOpDeletePerformanceAnalysisReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePerformanceAnalysisReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePerformanceAnalysisReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeletePerformanceAnalysisReport",
	}
}
