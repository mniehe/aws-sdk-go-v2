// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"github.com/mniehe/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	internalauthsmithy "github.com/mniehe/aws-sdk-go-v2/internal/auth/smithy"
	smithyauth "github.com/aws/smithy-go/auth"
	"github.com/aws/smithy-go/logging"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net/http"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// The optional application specific identifier appended to the User-Agent header.
	AppID string

	// This endpoint will be given as input to an EndpointResolverV2. It is used for
	// providing a custom base endpoint that is subject to modifications by the
	// processing EndpointResolverV2.
	BaseEndpoint *string

	// Configures the events that will be sent to the configured logger.
	ClientLogMode aws.ClientLogMode

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The configuration DefaultsMode that the SDK should use when constructing the
	// clients initial default settings.
	DefaultsMode aws.DefaultsMode

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions EndpointResolverOptions

	// The service endpoint resolver.
	//
	// Deprecated: Deprecated: EndpointResolver and WithEndpointResolver. Providing a
	// value for this field will likely prevent you from using any endpoint-related
	// service features released after the introduction of EndpointResolverV2 and
	// BaseEndpoint. To migrate an EndpointResolver implementation that uses a custom
	// endpoint, set the client option BaseEndpoint instead.
	EndpointResolver EndpointResolver

	// Resolves the endpoint used for a particular service operation. This should be
	// used over the deprecated EndpointResolver.
	EndpointResolverV2 EndpointResolverV2

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// The logger writer interface to write logging messages to.
	Logger logging.Logger

	// The region to send requests to. (Required)
	Region string

	// RetryMaxAttempts specifies the maximum number attempts an API client will call
	// an operation that fails with a retryable error. A value of 0 is ignored, and
	// will not be used to configure the API client created default retryer, or modify
	// per operation call's retry max attempts. If specified in an operation call's
	// functional options with a value that is different than the constructed client's
	// Options, the Client's Retryer will be wrapped to use the operation's specific
	// RetryMaxAttempts value.
	RetryMaxAttempts int

	// RetryMode specifies the retry mode the API client will be created with, if
	// Retryer option is not also specified. When creating a new API Clients this
	// member will only be used if the Retryer Options member is nil. This value will
	// be ignored if Retryer is not nil. Currently does not support per operation call
	// overrides, may in the future.
	RetryMode aws.RetryMode

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer. The kind of
	// default retry created by the API client can be changed with the RetryMode
	// option.
	Retryer aws.Retryer

	// The RuntimeEnvironment configuration, only populated if the DefaultsMode is set
	// to DefaultsModeAuto and is initialized using config.LoadDefaultConfig . You
	// should not populate this structure programmatically, or rely on the values here
	// within your applications.
	RuntimeEnvironment aws.RuntimeEnvironment

	// The initial DefaultsMode used when the client options were constructed. If the
	// DefaultsMode was set to aws.DefaultsModeAuto this will store what the resolved
	// value was at that point in time. Currently does not support per operation call
	// overrides, may in the future.
	resolvedDefaultsMode aws.DefaultsMode

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient

	// The auth scheme resolver which determines how to authenticate for each
	// operation.
	AuthSchemeResolver AuthSchemeResolver

	// The list of auth schemes supported by the client.
	AuthSchemes []smithyhttp.AuthScheme
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)

	return to
}

func (o Options) GetIdentityResolver(schemeID string) smithyauth.IdentityResolver {
	if schemeID == "aws.auth#sigv4" {
		return getSigV4IdentityResolver(o)
	}
	if schemeID == "smithy.api#noAuth" {
		return &smithyauth.AnonymousIdentityResolver{}
	}
	return nil
}

// WithAPIOptions returns a functional option for setting the Client's APIOptions
// option.
func WithAPIOptions(optFns ...func(*middleware.Stack) error) func(*Options) {
	return func(o *Options) {
		o.APIOptions = append(o.APIOptions, optFns...)
	}
}

// Deprecated: EndpointResolver and WithEndpointResolver. Providing a value for
// this field will likely prevent you from using any endpoint-related service
// features released after the introduction of EndpointResolverV2 and BaseEndpoint.
// To migrate an EndpointResolver implementation that uses a custom endpoint, set
// the client option BaseEndpoint instead.
func WithEndpointResolver(v EndpointResolver) func(*Options) {
	return func(o *Options) {
		o.EndpointResolver = v
	}
}

// WithEndpointResolverV2 returns a functional option for setting the Client's
// EndpointResolverV2 option.
func WithEndpointResolverV2(v EndpointResolverV2) func(*Options) {
	return func(o *Options) {
		o.EndpointResolverV2 = v
	}
}

func getSigV4IdentityResolver(o Options) smithyauth.IdentityResolver {
	if o.Credentials != nil {
		return &internalauthsmithy.CredentialsProviderAdapter{Provider: o.Credentials}
	}
	return nil
}

// WithSigV4SigningName applies an override to the authentication workflow to
// use the given signing name for SigV4-authenticated operations.
//
// This is an advanced setting. The value here is FINAL, taking precedence over
// the resolved signing name from both auth scheme resolution and endpoint
// resolution.
func WithSigV4SigningName(name string) func(*Options) {
	fn := func(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
		out middleware.InitializeOutput, metadata middleware.Metadata, err error,
	) {
		return next.HandleInitialize(awsmiddleware.SetSigningName(ctx, name), in)
	}
	return func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(s *middleware.Stack) error {
			return s.Initialize.Add(
				middleware.InitializeMiddlewareFunc("withSigV4SigningName", fn),
				middleware.Before,
			)
		})
	}
}

// WithSigV4SigningRegion applies an override to the authentication workflow to
// use the given signing region for SigV4-authenticated operations.
//
// This is an advanced setting. The value here is FINAL, taking precedence over
// the resolved signing region from both auth scheme resolution and endpoint
// resolution.
func WithSigV4SigningRegion(region string) func(*Options) {
	fn := func(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
		out middleware.InitializeOutput, metadata middleware.Metadata, err error,
	) {
		return next.HandleInitialize(awsmiddleware.SetSigningRegion(ctx, region), in)
	}
	return func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(s *middleware.Stack) error {
			return s.Initialize.Add(
				middleware.InitializeMiddlewareFunc("withSigV4SigningRegion", fn),
				middleware.Before,
			)
		})
	}
}

func ignoreAnonymousAuth(options *Options) {
	if aws.IsCredentialsProvider(options.Credentials, (*aws.AnonymousCredentials)(nil)) {
		options.Credentials = nil
	}
}
