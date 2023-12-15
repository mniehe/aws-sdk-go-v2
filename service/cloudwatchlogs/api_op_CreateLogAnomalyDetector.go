// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an anomaly detector that regularly scans one or more log groups and
// look for patterns and anomalies in the logs. An anomaly detector can help
// surface issues by automatically discovering anomalies in your log event traffic.
// An anomaly detector uses machine learning algorithms to scan log events and find
// patterns. A pattern is a shared text structure that recurs among your log
// fields. Patterns provide a useful tool for analyzing large sets of logs because
// a large number of log events can often be compressed into a few patterns. The
// anomaly detector uses pattern recognition to find anomalies , which are unusual
// log events. It uses the evaluationFrequency to compare current log events and
// patterns with trained baselines. Fields within a pattern are called tokens.
// Fields that vary within a pattern, such as a request ID or timestamp, are
// referred to as dynamic tokens and represented by <> . The following is an
// example of a pattern: [INFO] Request time: < > ms This pattern represents log
// events like [INFO] Request time: 327 ms and other similar log events that
// differ only by the number, in this csse 327. When the pattern is displayed, the
// different numbers are replaced by <*> Any parts of log events that are masked
// as sensitive data are not scanned for anomalies. For more information about
// masking sensitive data, see Help protect sensitive log data with masking (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html)
// .
func (c *Client) CreateLogAnomalyDetector(ctx context.Context, params *CreateLogAnomalyDetectorInput, optFns ...func(*Options)) (*CreateLogAnomalyDetectorOutput, error) {
	if params == nil {
		params = &CreateLogAnomalyDetectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLogAnomalyDetector", params, optFns, c.addOperationCreateLogAnomalyDetectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLogAnomalyDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLogAnomalyDetectorInput struct {

	// An array containing the ARN of the log group that this anomaly detector will
	// watch. You can specify only one log group ARN.
	//
	// This member is required.
	LogGroupArnList []string

	// The number of days to have visibility on an anomaly. After this time period has
	// elapsed for an anomaly, it will be automatically baselined and the anomaly
	// detector will treat new occurrences of a similar anomaly as normal. Therefore,
	// if you do not correct the cause of an anomaly during the time period specified
	// in anomalyVisibilityTime , it will be considered normal going forward and will
	// not be detected as an anomaly.
	AnomalyVisibilityTime *int64

	// A name for this anomaly detector.
	DetectorName *string

	// Specifies how often the anomaly detector is to run and look for anomalies. Set
	// this value according to the frequency that the log group receives new logs. For
	// example, if the log group receives new log events every 10 minutes, then 15
	// minutes might be a good setting for evaluationFrequency .
	EvaluationFrequency types.EvaluationFrequency

	// You can use this parameter to limit the anomaly detection model to examine only
	// log events that match the pattern you specify here. For more information, see
	// Filter and Pattern Syntax (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html)
	// .
	FilterPattern *string

	// Optionally assigns a KMS key to secure this anomaly detector and its findings.
	// If a key is assigned, the anomalies found and the model used by this detector
	// are encrypted at rest with the key. If a key is assigned to an anomaly detector,
	// a user must have permissions for both this key and for the anomaly detector to
	// retrieve information about the anomalies that it finds. For more information
	// about using a KMS key and to see the required IAM policy, see Use a KMS key
	// with an anomaly detector (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/LogsAnomalyDetection-KMS.html)
	// .
	KmsKeyId *string

	// An optional list of key-value pairs to associate with the resource. For more
	// information about tagging, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateLogAnomalyDetectorOutput struct {

	// The ARN of the log anomaly detector that you just created.
	AnomalyDetectorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLogAnomalyDetectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLogAnomalyDetector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLogAnomalyDetector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLogAnomalyDetector"); err != nil {
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
	if err = addOpCreateLogAnomalyDetectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLogAnomalyDetector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLogAnomalyDetector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLogAnomalyDetector",
	}
}
