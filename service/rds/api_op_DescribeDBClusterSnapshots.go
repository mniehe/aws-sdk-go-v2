// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"strconv"
	"time"
)

// Returns information about DB cluster snapshots. This API action supports
// pagination. For more information on Amazon Aurora DB clusters, see What is
// Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
// see Multi-AZ DB cluster deployments (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
// in the Amazon RDS User Guide.
func (c *Client) DescribeDBClusterSnapshots(ctx context.Context, params *DescribeDBClusterSnapshotsInput, optFns ...func(*Options)) (*DescribeDBClusterSnapshotsOutput, error) {
	if params == nil {
		params = &DescribeDBClusterSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterSnapshots", params, optFns, c.addOperationDescribeDBClusterSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBClusterSnapshotsInput struct {

	// The ID of the DB cluster to retrieve the list of DB cluster snapshots for. This
	// parameter can't be used in conjunction with the DBClusterSnapshotIdentifier
	// parameter. This parameter isn't case-sensitive. Constraints:
	//   - If supplied, must match the identifier of an existing DBCluster.
	DBClusterIdentifier *string

	// A specific DB cluster snapshot identifier to describe. This parameter can't be
	// used in conjunction with the DBClusterIdentifier parameter. This value is
	// stored as a lowercase string. Constraints:
	//   - If supplied, must match the identifier of an existing DBClusterSnapshot.
	//   - If this identifier is for an automated snapshot, the SnapshotType parameter
	//   must also be specified.
	DBClusterSnapshotIdentifier *string

	// A specific DB cluster resource ID to describe.
	DbClusterResourceId *string

	// A filter that specifies one or more DB cluster snapshots to describe. Supported
	// filters:
	//   - db-cluster-id - Accepts DB cluster identifiers and DB cluster Amazon
	//   Resource Names (ARNs).
	//   - db-cluster-snapshot-id - Accepts DB cluster snapshot identifiers.
	//   - snapshot-type - Accepts types of DB cluster snapshots.
	//   - engine - Accepts names of database engines.
	Filters []types.Filter

	// Specifies whether to include manual DB cluster snapshots that are public and
	// can be copied or restored by any Amazon Web Services account. By default, the
	// public snapshots are not included. You can share a manual DB cluster snapshot as
	// public by using the ModifyDBClusterSnapshotAttribute API action.
	IncludePublic *bool

	// Specifies whether to include shared manual DB cluster snapshots from other
	// Amazon Web Services accounts that this Amazon Web Services account has been
	// given permission to copy or restore. By default, these snapshots are not
	// included. You can give an Amazon Web Services account permission to restore a
	// manual DB cluster snapshot from another Amazon Web Services account by the
	// ModifyDBClusterSnapshotAttribute API action.
	IncludeShared *bool

	// An optional pagination token provided by a previous DescribeDBClusterSnapshots
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The type of DB cluster snapshots to be returned. You can specify one of the
	// following values:
	//   - automated - Return all DB cluster snapshots that have been automatically
	//   taken by Amazon RDS for my Amazon Web Services account.
	//   - manual - Return all DB cluster snapshots that have been taken by my Amazon
	//   Web Services account.
	//   - shared - Return all manual DB cluster snapshots that have been shared to my
	//   Amazon Web Services account.
	//   - public - Return all DB cluster snapshots that have been marked as public.
	// If you don't specify a SnapshotType value, then both automated and manual DB
	// cluster snapshots are returned. You can include shared DB cluster snapshots with
	// these results by enabling the IncludeShared parameter. You can include public
	// DB cluster snapshots with these results by enabling the IncludePublic
	// parameter. The IncludeShared and IncludePublic parameters don't apply for
	// SnapshotType values of manual or automated . The IncludePublic parameter
	// doesn't apply when SnapshotType is set to shared . The IncludeShared parameter
	// doesn't apply when SnapshotType is set to public .
	SnapshotType *string

	noSmithyDocumentSerde
}

// Provides a list of DB cluster snapshots for the user as the result of a call to
// the DescribeDBClusterSnapshots action.
type DescribeDBClusterSnapshotsOutput struct {

	// Provides a list of DB cluster snapshots for the user.
	DBClusterSnapshots []types.DBClusterSnapshot

	// An optional pagination token provided by a previous DescribeDBClusterSnapshots
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBClusterSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDBClusterSnapshots"); err != nil {
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
	if err = addOpDescribeDBClusterSnapshotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterSnapshots(options.Region), middleware.Before); err != nil {
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

// DescribeDBClusterSnapshotsAPIClient is a client that implements the
// DescribeDBClusterSnapshots operation.
type DescribeDBClusterSnapshotsAPIClient interface {
	DescribeDBClusterSnapshots(context.Context, *DescribeDBClusterSnapshotsInput, ...func(*Options)) (*DescribeDBClusterSnapshotsOutput, error)
}

var _ DescribeDBClusterSnapshotsAPIClient = (*Client)(nil)

// DescribeDBClusterSnapshotsPaginatorOptions is the paginator options for
// DescribeDBClusterSnapshots
type DescribeDBClusterSnapshotsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBClusterSnapshotsPaginator is a paginator for
// DescribeDBClusterSnapshots
type DescribeDBClusterSnapshotsPaginator struct {
	options   DescribeDBClusterSnapshotsPaginatorOptions
	client    DescribeDBClusterSnapshotsAPIClient
	params    *DescribeDBClusterSnapshotsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBClusterSnapshotsPaginator returns a new
// DescribeDBClusterSnapshotsPaginator
func NewDescribeDBClusterSnapshotsPaginator(client DescribeDBClusterSnapshotsAPIClient, params *DescribeDBClusterSnapshotsInput, optFns ...func(*DescribeDBClusterSnapshotsPaginatorOptions)) *DescribeDBClusterSnapshotsPaginator {
	if params == nil {
		params = &DescribeDBClusterSnapshotsInput{}
	}

	options := DescribeDBClusterSnapshotsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBClusterSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBClusterSnapshotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBClusterSnapshots page.
func (p *DescribeDBClusterSnapshotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBClusterSnapshotsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDBClusterSnapshots(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DBClusterSnapshotAvailableWaiterOptions are waiter options for
// DBClusterSnapshotAvailableWaiter
type DBClusterSnapshotAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// DBClusterSnapshotAvailableWaiter will use default minimum delay of 30 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, DBClusterSnapshotAvailableWaiter will use default max delay of 120
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
	Retryable func(context.Context, *DescribeDBClusterSnapshotsInput, *DescribeDBClusterSnapshotsOutput, error) (bool, error)
}

// DBClusterSnapshotAvailableWaiter defines the waiters for
// DBClusterSnapshotAvailable
type DBClusterSnapshotAvailableWaiter struct {
	client DescribeDBClusterSnapshotsAPIClient

	options DBClusterSnapshotAvailableWaiterOptions
}

// NewDBClusterSnapshotAvailableWaiter constructs a
// DBClusterSnapshotAvailableWaiter.
func NewDBClusterSnapshotAvailableWaiter(client DescribeDBClusterSnapshotsAPIClient, optFns ...func(*DBClusterSnapshotAvailableWaiterOptions)) *DBClusterSnapshotAvailableWaiter {
	options := DBClusterSnapshotAvailableWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = dBClusterSnapshotAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &DBClusterSnapshotAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for DBClusterSnapshotAvailable waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *DBClusterSnapshotAvailableWaiter) Wait(ctx context.Context, params *DescribeDBClusterSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*DBClusterSnapshotAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for DBClusterSnapshotAvailable waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *DBClusterSnapshotAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeDBClusterSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*DBClusterSnapshotAvailableWaiterOptions)) (*DescribeDBClusterSnapshotsOutput, error) {
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

		out, err := w.client.DescribeDBClusterSnapshots(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for DBClusterSnapshotAvailable waiter")
}

func dBClusterSnapshotAvailableStateRetryable(ctx context.Context, input *DescribeDBClusterSnapshotsInput, output *DescribeDBClusterSnapshotsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("DBClusterSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusterSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusterSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleting"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusterSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "failed"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusterSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-restore"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusterSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-parameters"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

// DBClusterSnapshotDeletedWaiterOptions are waiter options for
// DBClusterSnapshotDeletedWaiter
type DBClusterSnapshotDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// DBClusterSnapshotDeletedWaiter will use default minimum delay of 30 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, DBClusterSnapshotDeletedWaiter will use default max delay of 120
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
	Retryable func(context.Context, *DescribeDBClusterSnapshotsInput, *DescribeDBClusterSnapshotsOutput, error) (bool, error)
}

// DBClusterSnapshotDeletedWaiter defines the waiters for DBClusterSnapshotDeleted
type DBClusterSnapshotDeletedWaiter struct {
	client DescribeDBClusterSnapshotsAPIClient

	options DBClusterSnapshotDeletedWaiterOptions
}

// NewDBClusterSnapshotDeletedWaiter constructs a DBClusterSnapshotDeletedWaiter.
func NewDBClusterSnapshotDeletedWaiter(client DescribeDBClusterSnapshotsAPIClient, optFns ...func(*DBClusterSnapshotDeletedWaiterOptions)) *DBClusterSnapshotDeletedWaiter {
	options := DBClusterSnapshotDeletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = dBClusterSnapshotDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &DBClusterSnapshotDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for DBClusterSnapshotDeleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *DBClusterSnapshotDeletedWaiter) Wait(ctx context.Context, params *DescribeDBClusterSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*DBClusterSnapshotDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for DBClusterSnapshotDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *DBClusterSnapshotDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeDBClusterSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*DBClusterSnapshotDeletedWaiterOptions)) (*DescribeDBClusterSnapshotsOutput, error) {
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

		out, err := w.client.DescribeDBClusterSnapshots(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for DBClusterSnapshotDeleted waiter")
}

func dBClusterSnapshotDeletedStateRetryable(ctx context.Context, input *DescribeDBClusterSnapshotsInput, output *DescribeDBClusterSnapshotsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("length(DBClusterSnapshots) == `0`", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, nil
		}
	}

	if err != nil {
		var errorType *types.DBClusterSnapshotNotFoundFault
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusterSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "creating"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusterSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "modifying"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusterSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "rebooting"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusterSnapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "resetting-master-credentials"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeDBClusterSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDBClusterSnapshots",
	}
}
