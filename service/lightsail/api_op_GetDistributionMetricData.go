// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the data points of a specific metric for an Amazon Lightsail content
// delivery network (CDN) distribution. Metrics report the utilization of your
// resources, and the error counts generated by them. Monitor and collect metric
// data regularly to maintain the reliability, availability, and performance of
// your resources.
func (c *Client) GetDistributionMetricData(ctx context.Context, params *GetDistributionMetricDataInput, optFns ...func(*Options)) (*GetDistributionMetricDataOutput, error) {
	if params == nil {
		params = &GetDistributionMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDistributionMetricData", params, optFns, c.addOperationGetDistributionMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDistributionMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDistributionMetricDataInput struct {

	// The name of the distribution for which to get metric data. Use the
	// GetDistributions action to get a list of distribution names that you can specify.
	//
	// This member is required.
	DistributionName *string

	// The end of the time interval for which to get metric data. Constraints:
	//   - Specified in Coordinated Universal Time (UTC).
	//   - Specified in the Unix time format. For example, if you wish to use an end
	//   time of October 1, 2018, at 9 PM UTC, specify 1538427600 as the end time.
	// You can convert a human-friendly time to Unix time format using a converter
	// like Epoch converter (https://www.epochconverter.com/) .
	//
	// This member is required.
	EndTime *time.Time

	// The metric for which you want to return information. Valid distribution metric
	// names are listed below, along with the most useful statistics to include in
	// your request, and the published unit value.
	//   - Requests - The total number of viewer requests received by your Lightsail
	//   distribution, for all HTTP methods, and for both HTTP and HTTPS requests.
	//   Statistics : The most useful statistic is Sum . Unit : The published unit is
	//   None .
	//   - BytesDownloaded - The number of bytes downloaded by viewers for GET, HEAD,
	//   and OPTIONS requests. Statistics : The most useful statistic is Sum . Unit :
	//   The published unit is None .
	//   - BytesUploaded - The number of bytes uploaded to your origin by your
	//   Lightsail distribution, using POST and PUT requests. Statistics : The most
	//   useful statistic is Sum . Unit : The published unit is None .
	//   - TotalErrorRate - The percentage of all viewer requests for which the
	//   response's HTTP status code was 4xx or 5xx. Statistics : The most useful
	//   statistic is Average . Unit : The published unit is Percent .
	//   - 4xxErrorRate - The percentage of all viewer requests for which the
	//   response's HTTP status cod was 4xx. In these cases, the client or client viewer
	//   may have made an error. For example, a status code of 404 (Not Found) means that
	//   the client requested an object that could not be found. Statistics : The most
	//   useful statistic is Average . Unit : The published unit is Percent .
	//   - 5xxErrorRate - The percentage of all viewer requests for which the
	//   response's HTTP status code was 5xx. In these cases, the origin server did not
	//   satisfy the requests. For example, a status code of 503 (Service Unavailable)
	//   means that the origin server is currently unavailable. Statistics : The most
	//   useful statistic is Average . Unit : The published unit is Percent .
	//
	// This member is required.
	MetricName types.DistributionMetricName

	// The granularity, in seconds, for the metric data points that will be returned.
	//
	// This member is required.
	Period *int32

	// The start of the time interval for which to get metric data. Constraints:
	//   - Specified in Coordinated Universal Time (UTC).
	//   - Specified in the Unix time format. For example, if you wish to use a start
	//   time of October 1, 2018, at 8 PM UTC, specify 1538424000 as the start time.
	// You can convert a human-friendly time to Unix time format using a converter
	// like Epoch converter (https://www.epochconverter.com/) .
	//
	// This member is required.
	StartTime *time.Time

	// The statistic for the metric. The following statistics are available:
	//   - Minimum - The lowest value observed during the specified period. Use this
	//   value to determine low volumes of activity for your application.
	//   - Maximum - The highest value observed during the specified period. Use this
	//   value to determine high volumes of activity for your application.
	//   - Sum - All values submitted for the matching metric added together. You can
	//   use this statistic to determine the total volume of a metric.
	//   - Average - The value of Sum / SampleCount during the specified period. By
	//   comparing this statistic with the Minimum and Maximum values, you can determine
	//   the full scope of a metric and how close the average use is to the Minimum and
	//   Maximum values. This comparison helps you to know when to increase or decrease
	//   your resources.
	//   - SampleCount - The count, or number, of data points used for the statistical
	//   calculation.
	//
	// This member is required.
	Statistics []types.MetricStatistic

	// The unit for the metric data request. Valid units depend on the metric data
	// being requested. For the valid units with each available metric, see the
	// metricName parameter.
	//
	// This member is required.
	Unit types.MetricUnit

	noSmithyDocumentSerde
}

type GetDistributionMetricDataOutput struct {

	// An array of objects that describe the metric data returned.
	MetricData []types.MetricDatapoint

	// The name of the metric returned.
	MetricName types.DistributionMetricName

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDistributionMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDistributionMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDistributionMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDistributionMetricData"); err != nil {
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
	if err = addOpGetDistributionMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDistributionMetricData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDistributionMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDistributionMetricData",
	}
}
