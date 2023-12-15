// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns metadata about a query, including query run time in milliseconds,
// number of events scanned and matched, and query status. If the query results
// were delivered to an S3 bucket, the response also provides the S3 URI and the
// delivery status. You must specify either a QueryID or a QueryAlias . Specifying
// the QueryAlias parameter returns information about the last query run for the
// alias.
func (c *Client) DescribeQuery(ctx context.Context, params *DescribeQueryInput, optFns ...func(*Options)) (*DescribeQueryOutput, error) {
	if params == nil {
		params = &DescribeQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeQuery", params, optFns, c.addOperationDescribeQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeQueryInput struct {

	// The ARN (or the ID suffix of the ARN) of an event data store on which the
	// specified query was run.
	//
	// Deprecated: EventDataStore is no longer required by DescribeQueryRequest
	EventDataStore *string

	// The alias that identifies a query template.
	QueryAlias *string

	// The query ID.
	QueryId *string

	noSmithyDocumentSerde
}

type DescribeQueryOutput struct {

	// The URI for the S3 bucket where CloudTrail delivered query results, if
	// applicable.
	DeliveryS3Uri *string

	// The delivery status.
	DeliveryStatus types.DeliveryStatus

	// The error message returned if a query failed.
	ErrorMessage *string

	// The ID of the query.
	QueryId *string

	// Metadata about a query, including the number of events that were matched, the
	// total number of events scanned, the query run time in milliseconds, and the
	// query's creation time.
	QueryStatistics *types.QueryStatisticsForDescribeQuery

	// The status of a query. Values for QueryStatus include QUEUED , RUNNING ,
	// FINISHED , FAILED , TIMED_OUT , or CANCELLED
	QueryStatus types.QueryStatus

	// The SQL code of a query.
	QueryString *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeQuery{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeQuery"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeQuery",
	}
}
