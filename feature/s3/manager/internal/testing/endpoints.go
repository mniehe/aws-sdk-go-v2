package testing

import (
	"github.com/mniehe/aws-sdk-go-v2/aws"
	"github.com/mniehe/aws-sdk-go-v2/service/s3"
)

// EndpointResolverFunc is a mock s3 endpoint resolver that wraps the given function
type EndpointResolverFunc func(region string, options s3.EndpointResolverOptions) (aws.Endpoint, error)

// ResolveEndpoint returns the results from the wrapped function.
func (m EndpointResolverFunc) ResolveEndpoint(region string, options s3.EndpointResolverOptions) (aws.Endpoint, error) {
	return m(region, options)
}
