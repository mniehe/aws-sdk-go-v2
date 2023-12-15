// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves (queries) pre-aggregated data for a standard run execution metric
// that applies to a journey activity.
func (c *Client) GetJourneyRunExecutionActivityMetrics(ctx context.Context, params *GetJourneyRunExecutionActivityMetricsInput, optFns ...func(*Options)) (*GetJourneyRunExecutionActivityMetricsOutput, error) {
	if params == nil {
		params = &GetJourneyRunExecutionActivityMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetJourneyRunExecutionActivityMetrics", params, optFns, c.addOperationGetJourneyRunExecutionActivityMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetJourneyRunExecutionActivityMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJourneyRunExecutionActivityMetricsInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier for the journey activity.
	//
	// This member is required.
	JourneyActivityId *string

	// The unique identifier for the journey.
	//
	// This member is required.
	JourneyId *string

	// The unique identifier for the journey run.
	//
	// This member is required.
	RunId *string

	// The  string that specifies which page of results to return in a paginated
	// response. This parameter is not supported for application, campaign, and journey
	// metrics.
	NextToken *string

	// The maximum number of items to include in each page of a paginated response.
	// This parameter is not supported for application, campaign, and journey metrics.
	PageSize *string

	noSmithyDocumentSerde
}

type GetJourneyRunExecutionActivityMetricsOutput struct {

	// Provides the results of a query that retrieved the data for a standard
	// execution metric that applies to a journey activity for a particular journey
	// run, and provides information about that query.
	//
	// This member is required.
	JourneyRunExecutionActivityMetricsResponse *types.JourneyRunExecutionActivityMetricsResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetJourneyRunExecutionActivityMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetJourneyRunExecutionActivityMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetJourneyRunExecutionActivityMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetJourneyRunExecutionActivityMetrics"); err != nil {
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
	if err = addOpGetJourneyRunExecutionActivityMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetJourneyRunExecutionActivityMetrics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetJourneyRunExecutionActivityMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetJourneyRunExecutionActivityMetrics",
	}
}
