// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"github.com/mniehe/aws-sdk-go-v2/aws"
	endpoints "github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2"
	"github.com/aws/smithy-go/logging"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	// Logger is a logging implementation that log events should be sent to.
	Logger logging.Logger

	// LogDeprecated indicates that deprecated endpoints should be logged to the
	// provided logger.
	LogDeprecated bool

	// ResolvedRegion is used to override the region to be resolved, rather then the
	// using the value passed to the ResolveEndpoint method. This value is used by the
	// SDK to translate regions like fips-us-east-1 or us-east-1-fips to an alternative
	// name. You must not set this value directly in your application.
	ResolvedRegion string

	// DisableHTTPS informs the resolver to return an endpoint that does not use the
	// HTTPS scheme.
	DisableHTTPS bool

	// UseDualStackEndpoint specifies the resolver must resolve a dual-stack endpoint.
	UseDualStackEndpoint aws.DualStackEndpointState

	// UseFIPSEndpoint specifies the resolver must resolve a FIPS endpoint.
	UseFIPSEndpoint aws.FIPSEndpointState
}

func (o Options) GetResolvedRegion() string {
	return o.ResolvedRegion
}

func (o Options) GetDisableHTTPS() bool {
	return o.DisableHTTPS
}

func (o Options) GetUseDualStackEndpoint() aws.DualStackEndpointState {
	return o.UseDualStackEndpoint
}

func (o Options) GetUseFIPSEndpoint() aws.FIPSEndpointState {
	return o.UseFIPSEndpoint
}

func transformToSharedOptions(options Options) endpoints.Options {
	return endpoints.Options{
		Logger:               options.Logger,
		LogDeprecated:        options.LogDeprecated,
		ResolvedRegion:       options.ResolvedRegion,
		DisableHTTPS:         options.DisableHTTPS,
		UseDualStackEndpoint: options.UseDualStackEndpoint,
		UseFIPSEndpoint:      options.UseFIPSEndpoint,
	}
}

// Resolver neptunedata endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint aws.Endpoint, err error) {
	if len(region) == 0 {
		return endpoint, &aws.MissingRegionError{}
	}

	opt := transformToSharedOptions(options)
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var partitionRegexp = struct {
	Aws      *regexp.Regexp
	AwsCn    *regexp.Regexp
	AwsIso   *regexp.Regexp
	AwsIsoB  *regexp.Regexp
	AwsIsoE  *regexp.Regexp
	AwsIsoF  *regexp.Regexp
	AwsUsGov *regexp.Regexp
}{

	Aws:      regexp.MustCompile("^(us|eu|ap|sa|ca|me|af|il)\\-\\w+\\-\\d+$"),
	AwsCn:    regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
	AwsIso:   regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
	AwsIsoB:  regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
	AwsIsoE:  regexp.MustCompile("^eu\\-isoe\\-\\w+\\-\\d+$"),
	AwsIsoF:  regexp.MustCompile("^us\\-isof\\-\\w+\\-\\d+$"),
	AwsUsGov: regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "neptune-db.{region}.api.aws",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "neptune-db-fips.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "neptune-db-fips.{region}.api.aws",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "neptune-db.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.Aws,
		IsRegionalized: true,
	},
	{
		ID: "aws-cn",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "neptune-db.{region}.api.amazonwebservices.com.cn",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "neptune-db-fips.{region}.amazonaws.com.cn",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "neptune-db-fips.{region}.api.amazonwebservices.com.cn",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "neptune-db.{region}.amazonaws.com.cn",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsCn,
		IsRegionalized: true,
	},
	{
		ID: "aws-iso",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "neptune-db-fips.{region}.c2s.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "neptune-db.{region}.c2s.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIso,
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-b",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "neptune-db-fips.{region}.sc2s.sgov.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "neptune-db.{region}.sc2s.sgov.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIsoB,
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-e",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "neptune-db-fips.{region}.cloud.adc-e.uk",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "neptune-db.{region}.cloud.adc-e.uk",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIsoE,
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-f",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "neptune-db-fips.{region}.csp.hci.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "neptune-db.{region}.csp.hci.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIsoF,
		IsRegionalized: true,
	},
	{
		ID: "aws-us-gov",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "neptune-db.{region}.api.aws",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "neptune-db-fips.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "neptune-db-fips.{region}.api.aws",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "neptune-db.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsUsGov,
		IsRegionalized: true,
	},
}
