// Code generated by smithy-go-codegen DO NOT EDIT.

package rum

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/rum/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a destination to receive extended metrics from CloudWatch
// RUM. You can send extended metrics to CloudWatch or to a CloudWatch Evidently
// experiment. For more information about extended metrics, see
// BatchCreateRumMetricDefinitions (https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_BatchCreateRumMetricDefinitions.html)
// .
func (c *Client) PutRumMetricsDestination(ctx context.Context, params *PutRumMetricsDestinationInput, optFns ...func(*Options)) (*PutRumMetricsDestinationOutput, error) {
	if params == nil {
		params = &PutRumMetricsDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRumMetricsDestination", params, optFns, c.addOperationPutRumMetricsDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRumMetricsDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRumMetricsDestinationInput struct {

	// The name of the CloudWatch RUM app monitor that will send the metrics.
	//
	// This member is required.
	AppMonitorName *string

	// Defines the destination to send the metrics to. Valid values are CloudWatch and
	// Evidently . If you specify Evidently , you must also specify the ARN of the
	// CloudWatchEvidently experiment that is to be the destination and an IAM role
	// that has permission to write to the experiment.
	//
	// This member is required.
	Destination types.MetricDestination

	// Use this parameter only if Destination is Evidently . This parameter specifies
	// the ARN of the Evidently experiment that will receive the extended metrics.
	DestinationArn *string

	// This parameter is required if Destination is Evidently . If Destination is
	// CloudWatch , do not use this parameter. This parameter specifies the ARN of an
	// IAM role that RUM will assume to write to the Evidently experiment that you are
	// sending metrics to. This role must have permission to write to that experiment.
	IamRoleArn *string

	noSmithyDocumentSerde
}

type PutRumMetricsDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRumMetricsDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutRumMetricsDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutRumMetricsDestination{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRumMetricsDestination"); err != nil {
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
	if err = addOpPutRumMetricsDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRumMetricsDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRumMetricsDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRumMetricsDestination",
	}
}
