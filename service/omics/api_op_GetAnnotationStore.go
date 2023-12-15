// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Gets information about an annotation store.
func (c *Client) GetAnnotationStore(ctx context.Context, params *GetAnnotationStoreInput, optFns ...func(*Options)) (*GetAnnotationStoreOutput, error) {
	if params == nil {
		params = &GetAnnotationStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAnnotationStore", params, optFns, c.addOperationGetAnnotationStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAnnotationStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAnnotationStoreInput struct {

	// The store's name.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetAnnotationStoreOutput struct {

	// When the store was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The store's description.
	//
	// This member is required.
	Description *string

	// The store's ID.
	//
	// This member is required.
	Id *string

	// The store's name.
	//
	// This member is required.
	Name *string

	// An integer indicating how many versions of an annotation store exist.
	//
	// This member is required.
	NumVersions *int32

	// The store's genome reference.
	//
	// This member is required.
	Reference types.ReferenceItem

	// The store's server-side encryption (SSE) settings.
	//
	// This member is required.
	SseConfig *types.SseConfig

	// The store's status.
	//
	// This member is required.
	Status types.StoreStatus

	// A status message.
	//
	// This member is required.
	StatusMessage *string

	// The store's ARN.
	//
	// This member is required.
	StoreArn *string

	// The store's size in bytes.
	//
	// This member is required.
	StoreSizeBytes *int64

	// The store's tags.
	//
	// This member is required.
	Tags map[string]string

	// When the store was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The store's annotation file format.
	StoreFormat types.StoreFormat

	// The store's parsing options.
	StoreOptions types.StoreOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAnnotationStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAnnotationStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAnnotationStore{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAnnotationStore"); err != nil {
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
	if err = addEndpointPrefix_opGetAnnotationStoreMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAnnotationStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAnnotationStore(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetAnnotationStoreMiddleware struct {
}

func (*endpointPrefix_opGetAnnotationStoreMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAnnotationStoreMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "analytics-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetAnnotationStoreMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetAnnotationStoreMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// GetAnnotationStoreAPIClient is a client that implements the GetAnnotationStore
// operation.
type GetAnnotationStoreAPIClient interface {
	GetAnnotationStore(context.Context, *GetAnnotationStoreInput, ...func(*Options)) (*GetAnnotationStoreOutput, error)
}

var _ GetAnnotationStoreAPIClient = (*Client)(nil)

// AnnotationStoreCreatedWaiterOptions are waiter options for
// AnnotationStoreCreatedWaiter
type AnnotationStoreCreatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// AnnotationStoreCreatedWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, AnnotationStoreCreatedWaiter will use default max delay of 600
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *GetAnnotationStoreInput, *GetAnnotationStoreOutput, error) (bool, error)
}

// AnnotationStoreCreatedWaiter defines the waiters for AnnotationStoreCreated
type AnnotationStoreCreatedWaiter struct {
	client GetAnnotationStoreAPIClient

	options AnnotationStoreCreatedWaiterOptions
}

// NewAnnotationStoreCreatedWaiter constructs a AnnotationStoreCreatedWaiter.
func NewAnnotationStoreCreatedWaiter(client GetAnnotationStoreAPIClient, optFns ...func(*AnnotationStoreCreatedWaiterOptions)) *AnnotationStoreCreatedWaiter {
	options := AnnotationStoreCreatedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 600 * time.Second
	options.Retryable = annotationStoreCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AnnotationStoreCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AnnotationStoreCreated waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *AnnotationStoreCreatedWaiter) Wait(ctx context.Context, params *GetAnnotationStoreInput, maxWaitDur time.Duration, optFns ...func(*AnnotationStoreCreatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for AnnotationStoreCreated waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *AnnotationStoreCreatedWaiter) WaitForOutput(ctx context.Context, params *GetAnnotationStoreInput, maxWaitDur time.Duration, optFns ...func(*AnnotationStoreCreatedWaiterOptions)) (*GetAnnotationStoreOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 600 * time.Second
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

		out, err := w.client.GetAnnotationStore(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for AnnotationStoreCreated waiter")
}

func annotationStoreCreatedStateRetryable(ctx context.Context, input *GetAnnotationStoreInput, output *GetAnnotationStoreOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ACTIVE"
		value, ok := pathValue.(types.StoreStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StoreStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATING"
		value, ok := pathValue.(types.StoreStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StoreStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "UPDATING"
		value, ok := pathValue.(types.StoreStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StoreStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.StoreStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StoreStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// AnnotationStoreDeletedWaiterOptions are waiter options for
// AnnotationStoreDeletedWaiter
type AnnotationStoreDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// AnnotationStoreDeletedWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, AnnotationStoreDeletedWaiter will use default max delay of 600
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *GetAnnotationStoreInput, *GetAnnotationStoreOutput, error) (bool, error)
}

// AnnotationStoreDeletedWaiter defines the waiters for AnnotationStoreDeleted
type AnnotationStoreDeletedWaiter struct {
	client GetAnnotationStoreAPIClient

	options AnnotationStoreDeletedWaiterOptions
}

// NewAnnotationStoreDeletedWaiter constructs a AnnotationStoreDeletedWaiter.
func NewAnnotationStoreDeletedWaiter(client GetAnnotationStoreAPIClient, optFns ...func(*AnnotationStoreDeletedWaiterOptions)) *AnnotationStoreDeletedWaiter {
	options := AnnotationStoreDeletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 600 * time.Second
	options.Retryable = annotationStoreDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AnnotationStoreDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AnnotationStoreDeleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *AnnotationStoreDeletedWaiter) Wait(ctx context.Context, params *GetAnnotationStoreInput, maxWaitDur time.Duration, optFns ...func(*AnnotationStoreDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for AnnotationStoreDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *AnnotationStoreDeletedWaiter) WaitForOutput(ctx context.Context, params *GetAnnotationStoreInput, maxWaitDur time.Duration, optFns ...func(*AnnotationStoreDeletedWaiterOptions)) (*GetAnnotationStoreOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 600 * time.Second
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

		out, err := w.client.GetAnnotationStore(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for AnnotationStoreDeleted waiter")
}

func annotationStoreDeletedStateRetryable(ctx context.Context, input *GetAnnotationStoreInput, output *GetAnnotationStoreOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETED"
		value, ok := pathValue.(types.StoreStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StoreStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETING"
		value, ok := pathValue.(types.StoreStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StoreStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetAnnotationStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAnnotationStore",
	}
}
