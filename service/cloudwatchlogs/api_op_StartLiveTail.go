// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"sync"
)

// Starts a Live Tail streaming session for one or more log groups. A Live Tail
// session returns a stream of log events that have been recently ingested in the
// log groups. For more information, see Use Live Tail to view logs in near real
// time (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs_LiveTail.html)
// . The response to this operation is a response stream, over which the server
// sends live log events and the client receives them. The following objects are
// sent over the stream:
//   - A single LiveTailSessionStart (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LiveTailSessionStart.html)
//     object is sent at the start of the session.
//   - Every second, a LiveTailSessionUpdate (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LiveTailSessionUpdate.html)
//     object is sent. Each of these objects contains an array of the actual log
//     events. If no new log events were ingested in the past second, the
//     LiveTailSessionUpdate object will contain an empty array. The array of log
//     events contained in a LiveTailSessionUpdate can include as many as 500 log
//     events. If the number of log events matching the request exceeds 500 per second,
//     the log events are sampled down to 500 log events to be included in each
//     LiveTailSessionUpdate object. If your client consumes the log events slower
//     than the server produces them, CloudWatch Logs buffers up to 10
//     LiveTailSessionUpdate events or 5000 log events, after which it starts
//     dropping the oldest events.
//   - A SessionStreamingException (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_SessionStreamingException.html)
//     object is returned if an unknown error occurs on the server side.
//   - A SessionTimeoutException (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_SessionTimeoutException.html)
//     object is returned when the session times out, after it has been kept open for
//     three hours.
//
// You can end a session before it times out by closing the session stream or by
// closing the client that is receiving the stream. The session also ends if the
// established connection between the client and the server breaks.
func (c *Client) StartLiveTail(ctx context.Context, params *StartLiveTailInput, optFns ...func(*Options)) (*StartLiveTailOutput, error) {
	if params == nil {
		params = &StartLiveTailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartLiveTail", params, optFns, c.addOperationStartLiveTailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartLiveTailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartLiveTailInput struct {

	// An array where each item in the array is a log group to include in the Live
	// Tail session. Specify each log group by its ARN. If you specify an ARN, the ARN
	// can't end with an asterisk (*). You can include up to 10 log groups.
	//
	// This member is required.
	LogGroupIdentifiers []string

	// An optional pattern to use to filter the results to include only log events
	// that match the pattern. For example, a filter pattern of error 404 causes only
	// log events that include both error and 404 to be included in the Live Tail
	// stream. Regular expression filter patterns are supported. For more information
	// about filter pattern syntax, see Filter and Pattern Syntax (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html)
	// .
	LogEventFilterPattern *string

	// If you specify this parameter, then only log events in the log streams that
	// have names that start with the prefixes that you specify here are included in
	// the Live Tail session. You can specify this parameter only if you specify only
	// one log group in logGroupIdentifiers .
	LogStreamNamePrefixes []string

	// If you specify this parameter, then only log events in the log streams that you
	// specify here are included in the Live Tail session. You can specify this
	// parameter only if you specify only one log group in logGroupIdentifiers .
	LogStreamNames []string

	noSmithyDocumentSerde
}

type StartLiveTailOutput struct {
	eventStream *StartLiveTailEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *StartLiveTailOutput) GetStream() *StartLiveTailEventStream {
	return o.eventStream
}

func (c *Client) addOperationStartLiveTailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartLiveTail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartLiveTail{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartLiveTail"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addEventStreamStartLiveTailMiddleware(stack, options); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opStartLiveTailMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartLiveTailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartLiveTail(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opStartLiveTailMiddleware struct {
}

func (*endpointPrefix_opStartLiveTailMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opStartLiveTailMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "streaming-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opStartLiveTailMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opStartLiveTailMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opStartLiveTail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartLiveTail",
	}
}

// StartLiveTailEventStream provides the event stream handling for the StartLiveTail operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewStartLiveTailEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type StartLiveTailEventStream struct {
	// StartLiveTailResponseStreamReader is the EventStream reader for the
	// StartLiveTailResponseStream events. This value is automatically set by the SDK
	// when the API call is made Use this member when unit testing your code with the
	// SDK to mock out the EventStream Reader.
	//
	// Must not be nil.
	Reader StartLiveTailResponseStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewStartLiveTailEventStream initializes an StartLiveTailEventStream.
// This function should only be used for testing and mocking the StartLiveTailEventStream
// stream within your application.
//
// The Reader member must be set before reading events from the stream.
func NewStartLiveTailEventStream(optFns ...func(*StartLiveTailEventStream)) *StartLiveTailEventStream {
	es := &StartLiveTailEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Events returns a channel to read events from.
func (es *StartLiveTailEventStream) Events() <-chan types.StartLiveTailResponseStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *StartLiveTailEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *StartLiveTailEventStream) safeClose() {
	close(es.done)

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *StartLiveTailEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *StartLiveTailEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}
