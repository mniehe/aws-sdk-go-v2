// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes the specified Kinesis data stream. This API has been revised. It's
// highly recommended that you use the DescribeStreamSummary API to get a
// summarized description of the specified Kinesis data stream and the ListShards
// API to list the shards in a specified data stream and obtain information about
// each shard. When invoking this API, you must use either the StreamARN or the
// StreamName parameter, or both. It is recommended that you use the StreamARN
// input parameter when you invoke this API. The information returned includes the
// stream name, Amazon Resource Name (ARN), creation time, enhanced metric
// configuration, and shard map. The shard map is an array of shard objects. For
// each shard object, there is the hash key and sequence number ranges that the
// shard spans, and the IDs of any earlier shards that played in a role in creating
// the shard. Every record ingested in the stream is identified by a sequence
// number, which is assigned when the record is put into the stream. You can limit
// the number of shards returned by each call. For more information, see
// Retrieving Shards from a Stream (https://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-retrieve-shards.html)
// in the Amazon Kinesis Data Streams Developer Guide. There are no guarantees
// about the chronological order shards returned. To process shards in
// chronological order, use the ID of the parent shard to track the lineage to the
// oldest shard. This operation has a limit of 10 transactions per second per
// account.
func (c *Client) DescribeStream(ctx context.Context, params *DescribeStreamInput, optFns ...func(*Options)) (*DescribeStreamOutput, error) {
	if params == nil {
		params = &DescribeStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStream", params, optFns, c.addOperationDescribeStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for DescribeStream .
type DescribeStreamInput struct {

	// The shard ID of the shard to start with. Specify this parameter to indicate
	// that you want to describe the stream starting with the shard whose ID
	// immediately follows ExclusiveStartShardId . If you don't specify this parameter,
	// the default behavior for DescribeStream is to describe the stream starting with
	// the first shard in the stream.
	ExclusiveStartShardId *string

	// The maximum number of shards to return in a single call. The default value is
	// 100. If you specify a value greater than 100, at most 100 results are returned.
	Limit *int32

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream to describe.
	StreamName *string

	noSmithyDocumentSerde
}

func (in *DescribeStreamInput) bindEndpointParams(p *EndpointParameters) {
	p.StreamARN = in.StreamARN
	p.OperationType = ptr.String("control")
}

// Represents the output for DescribeStream .
type DescribeStreamOutput struct {

	// The current status of the stream, the stream Amazon Resource Name (ARN), an
	// array of shard objects that comprise the stream, and whether there are more
	// shards available.
	//
	// This member is required.
	StreamDescription *types.StreamDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStream"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStream(options.Region), middleware.Before); err != nil {
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

// DescribeStreamAPIClient is a client that implements the DescribeStream
// operation.
type DescribeStreamAPIClient interface {
	DescribeStream(context.Context, *DescribeStreamInput, ...func(*Options)) (*DescribeStreamOutput, error)
}

var _ DescribeStreamAPIClient = (*Client)(nil)

// StreamExistsWaiterOptions are waiter options for StreamExistsWaiter
type StreamExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// StreamExistsWaiter will use default minimum delay of 10 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, StreamExistsWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeStreamInput, *DescribeStreamOutput, error) (bool, error)
}

// StreamExistsWaiter defines the waiters for StreamExists
type StreamExistsWaiter struct {
	client DescribeStreamAPIClient

	options StreamExistsWaiterOptions
}

// NewStreamExistsWaiter constructs a StreamExistsWaiter.
func NewStreamExistsWaiter(client DescribeStreamAPIClient, optFns ...func(*StreamExistsWaiterOptions)) *StreamExistsWaiter {
	options := StreamExistsWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = streamExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &StreamExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for StreamExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *StreamExistsWaiter) Wait(ctx context.Context, params *DescribeStreamInput, maxWaitDur time.Duration, optFns ...func(*StreamExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for StreamExists waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *StreamExistsWaiter) WaitForOutput(ctx context.Context, params *DescribeStreamInput, maxWaitDur time.Duration, optFns ...func(*StreamExistsWaiterOptions)) (*DescribeStreamOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeStream(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for StreamExists waiter")
}

func streamExistsStateRetryable(ctx context.Context, input *DescribeStreamInput, output *DescribeStreamOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("StreamDescription.StreamStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ACTIVE"
		value, ok := pathValue.(types.StreamStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StreamStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	return true, nil
}

// StreamNotExistsWaiterOptions are waiter options for StreamNotExistsWaiter
type StreamNotExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// StreamNotExistsWaiter will use default minimum delay of 10 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, StreamNotExistsWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeStreamInput, *DescribeStreamOutput, error) (bool, error)
}

// StreamNotExistsWaiter defines the waiters for StreamNotExists
type StreamNotExistsWaiter struct {
	client DescribeStreamAPIClient

	options StreamNotExistsWaiterOptions
}

// NewStreamNotExistsWaiter constructs a StreamNotExistsWaiter.
func NewStreamNotExistsWaiter(client DescribeStreamAPIClient, optFns ...func(*StreamNotExistsWaiterOptions)) *StreamNotExistsWaiter {
	options := StreamNotExistsWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = streamNotExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &StreamNotExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for StreamNotExists waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *StreamNotExistsWaiter) Wait(ctx context.Context, params *DescribeStreamInput, maxWaitDur time.Duration, optFns ...func(*StreamNotExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for StreamNotExists waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *StreamNotExistsWaiter) WaitForOutput(ctx context.Context, params *DescribeStreamInput, maxWaitDur time.Duration, optFns ...func(*StreamNotExistsWaiterOptions)) (*DescribeStreamOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeStream(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for StreamNotExists waiter")
}

func streamNotExistsStateRetryable(ctx context.Context, input *DescribeStreamInput, output *DescribeStreamOutput, err error) (bool, error) {

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStream",
	}
}
