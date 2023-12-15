// Code generated by smithy-go-codegen DO NOT EDIT.

package amp

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/amp/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describe an existing scraper.
func (c *Client) DescribeScraper(ctx context.Context, params *DescribeScraperInput, optFns ...func(*Options)) (*DescribeScraperOutput, error) {
	if params == nil {
		params = &DescribeScraperInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScraper", params, optFns, c.addOperationDescribeScraperMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScraperOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeScraper operation.
type DescribeScraperInput struct {

	// The IDs of the scraper to describe.
	//
	// This member is required.
	ScraperId *string

	noSmithyDocumentSerde
}

// Represents the output of a DescribeScraper operation.
type DescribeScraperOutput struct {

	// The properties of the selected scrapers.
	//
	// This member is required.
	Scraper *types.ScraperDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScraperMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeScraper{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeScraper{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeScraper"); err != nil {
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
	if err = addOpDescribeScraperValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScraper(options.Region), middleware.Before); err != nil {
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

// DescribeScraperAPIClient is a client that implements the DescribeScraper
// operation.
type DescribeScraperAPIClient interface {
	DescribeScraper(context.Context, *DescribeScraperInput, ...func(*Options)) (*DescribeScraperOutput, error)
}

var _ DescribeScraperAPIClient = (*Client)(nil)

// ScraperActiveWaiterOptions are waiter options for ScraperActiveWaiter
type ScraperActiveWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ScraperActiveWaiter will use default minimum delay of 2 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ScraperActiveWaiter will use default max delay of 120 seconds. Note
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
	Retryable func(context.Context, *DescribeScraperInput, *DescribeScraperOutput, error) (bool, error)
}

// ScraperActiveWaiter defines the waiters for ScraperActive
type ScraperActiveWaiter struct {
	client DescribeScraperAPIClient

	options ScraperActiveWaiterOptions
}

// NewScraperActiveWaiter constructs a ScraperActiveWaiter.
func NewScraperActiveWaiter(client DescribeScraperAPIClient, optFns ...func(*ScraperActiveWaiterOptions)) *ScraperActiveWaiter {
	options := ScraperActiveWaiterOptions{}
	options.MinDelay = 2 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = scraperActiveStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ScraperActiveWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ScraperActive waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ScraperActiveWaiter) Wait(ctx context.Context, params *DescribeScraperInput, maxWaitDur time.Duration, optFns ...func(*ScraperActiveWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ScraperActive waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *ScraperActiveWaiter) WaitForOutput(ctx context.Context, params *DescribeScraperInput, maxWaitDur time.Duration, optFns ...func(*ScraperActiveWaiterOptions)) (*DescribeScraperOutput, error) {
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

		out, err := w.client.DescribeScraper(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ScraperActive waiter")
}

func scraperActiveStateRetryable(ctx context.Context, input *DescribeScraperInput, output *DescribeScraperOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("scraper.status.statusCode", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ACTIVE"
		value, ok := pathValue.(types.ScraperStatusCode)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ScraperStatusCode value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("scraper.status.statusCode", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATION_FAILED"
		value, ok := pathValue.(types.ScraperStatusCode)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ScraperStatusCode value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// ScraperDeletedWaiterOptions are waiter options for ScraperDeletedWaiter
type ScraperDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ScraperDeletedWaiter will use default minimum delay of 2 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ScraperDeletedWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeScraperInput, *DescribeScraperOutput, error) (bool, error)
}

// ScraperDeletedWaiter defines the waiters for ScraperDeleted
type ScraperDeletedWaiter struct {
	client DescribeScraperAPIClient

	options ScraperDeletedWaiterOptions
}

// NewScraperDeletedWaiter constructs a ScraperDeletedWaiter.
func NewScraperDeletedWaiter(client DescribeScraperAPIClient, optFns ...func(*ScraperDeletedWaiterOptions)) *ScraperDeletedWaiter {
	options := ScraperDeletedWaiterOptions{}
	options.MinDelay = 2 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = scraperDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ScraperDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ScraperDeleted waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ScraperDeletedWaiter) Wait(ctx context.Context, params *DescribeScraperInput, maxWaitDur time.Duration, optFns ...func(*ScraperDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ScraperDeleted waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *ScraperDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeScraperInput, maxWaitDur time.Duration, optFns ...func(*ScraperDeletedWaiterOptions)) (*DescribeScraperOutput, error) {
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

		out, err := w.client.DescribeScraper(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ScraperDeleted waiter")
}

func scraperDeletedStateRetryable(ctx context.Context, input *DescribeScraperInput, output *DescribeScraperOutput, err error) (bool, error) {

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("scraper.status.statusCode", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETION_FAILED"
		value, ok := pathValue.(types.ScraperStatusCode)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ScraperStatusCode value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeScraper(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeScraper",
	}
}
