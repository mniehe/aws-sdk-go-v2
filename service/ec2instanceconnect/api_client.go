// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2instanceconnect

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/logging"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"net/http"
	"time"
)

const ServiceID = "EC2 Instance Connect"
const ServiceAPIVersion = "2018-04-02"

// Client provides the API client to make operations call for AWS EC2 Instance
// Connect.
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveDefaultLogger(&options)

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// Configures the events that will be sent to the configured logger.
	ClientLogMode aws.ClientLogMode

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions EndpointResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// The logger writer interface to write logging messages to.
	Logger logging.Logger

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}
func (c *Client) invokeOperation(ctx context.Context, opID string, params interface{}, optFns []func(*Options), stackFns ...func(*middleware.Stack, Options) error) (result interface{}, metadata middleware.Metadata, err error) {
	stack := middleware.NewStack(opID, smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}

	for _, fn := range stackFns {
		if err := fn(stack, options); err != nil {
			return nil, metadata, err
		}
	}

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, metadata, err
		}
	}

	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err = handler.Handle(ctx, params)
	if err != nil {
		err = &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: opID,
			Err:           err,
		}
	}
	return result, metadata, err
}

func resolveDefaultLogger(o *Options) {
	if o.Logger != nil {
		return
	}
	o.Logger = logging.Nop{}
}

func addSetLoggerMiddleware(stack *middleware.Stack, o Options) error {
	return middleware.AddSetLoggerMiddleware(stack, o.Logger)
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:        cfg.Region,
		Retryer:       cfg.Retryer,
		HTTPClient:    cfg.HTTPClient,
		Credentials:   cfg.Credentials,
		APIOptions:    cfg.APIOptions,
		Logger:        cfg.Logger,
		ClientLogMode: cfg.ClientLogMode,
	}
	resolveAWSEndpointResolver(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = awshttp.NewBuildableClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil {
		return
	}
	o.EndpointResolver = WithEndpointResolver(cfg.EndpointResolver, NewDefaultEndpointResolver())
}

func addClientUserAgent(stack *middleware.Stack) error {
	return awsmiddleware.AddUserAgentKey("ec2instanceconnect")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) error {
	return stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner(
		func(s *v4.Signer) {
			s.Logger = o.Logger
			s.LogSigning = o.ClientLogMode.IsSigning()
		},
	)
}

func addRetryMiddlewares(stack *middleware.Stack, o Options) error {
	mo := retry.AddRetryMiddlewaresOptions{
		Retryer:          o.Retryer,
		LogRetryAttempts: o.ClientLogMode.IsRetries(),
	}
	return retry.AddRetryMiddlewares(stack, mo)
}

func addRequestIDRetrieverMiddleware(stack *middleware.Stack) error {
	return awsmiddleware.AddRequestIDRetrieverMiddleware(stack)
}

func addResponseErrorMiddleware(stack *middleware.Stack) error {
	return awshttp.AddResponseErrorMiddleware(stack)
}

func addRequestResponseLogging(stack *middleware.Stack, o Options) error {
	return stack.Deserialize.Add(&smithyhttp.RequestResponseLogger{
		LogRequest:          o.ClientLogMode.IsRequest(),
		LogRequestWithBody:  o.ClientLogMode.IsRequestWithBody(),
		LogResponse:         o.ClientLogMode.IsResponse(),
		LogResponseWithBody: o.ClientLogMode.IsResponseWithBody(),
	}, middleware.After)
}
