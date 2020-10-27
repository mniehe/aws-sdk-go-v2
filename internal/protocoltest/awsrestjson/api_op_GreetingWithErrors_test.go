// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"bytes"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithyrand "github.com/awslabs/smithy-go/rand"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GreetingWithErrors_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *GreetingWithErrorsOutput
	}{
		// Ensures that operations with errors successfully know how to deserialize the
		// successful response
		"RestJsonGreetingWithErrors": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
				"X-Greeting":   []string{"Hello"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "greeting": "Hello"
			}`),
			ExpectResult: &GreetingWithErrorsOutput{
				Greeting: ptr.String("Hello"),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}

func TestClient_GreetingWithErrors_FooError_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.FooError
	}{
		// Serializes the X-Amzn-ErrorType header. For an example service, see Amazon EKS.
		"RestJsonFooErrorUsingXAmznErrorType": {
			StatusCode: 500,
			Header: http.Header{
				"X-Amzn-Errortype": []string{"FooError"},
			},
			ExpectError: &types.FooError{},
		},
		// Some X-Amzn-Errortype headers contain URLs. Clients need to split the URL on ':'
		// and take only the first half of the string. For example,
		// 'ValidationException:http://internal.amazon.com/coral/com.amazon.coral.validate/'
		// is to be interpreted as 'ValidationException'. For an example service see Amazon
		// Polly.
		"RestJsonFooErrorUsingXAmznErrorTypeWithUri": {
			StatusCode: 500,
			Header: http.Header{
				"X-Amzn-Errortype": []string{"FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"},
			},
			ExpectError: &types.FooError{},
		},
		// X-Amzn-Errortype might contain a URL and a namespace. Client should extract only
		// the shape name. This is a pathalogical case that might not actually happen in
		// any deployed AWS service.
		"RestJsonFooErrorUsingXAmznErrorTypeWithUriAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"X-Amzn-Errortype": []string{"aws.protocoltests.restjson#FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"},
			},
			ExpectError: &types.FooError{},
		},
		// This example uses the 'code' property in the output rather than
		// X-Amzn-Errortype. Some services do this though it's preferable to send the
		// X-Amzn-Errortype. Client implementations must first check for the
		// X-Amzn-Errortype and then check for a top-level 'code' property. For example
		// service see Amazon S3 Glacier.
		"RestJsonFooErrorUsingCode": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "code": "FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using code, and it might contain a namespace.
		// Clients should just take the last part of the string after '#'.
		"RestJsonFooErrorUsingCodeAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "code": "aws.protocoltests.restjson#FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using code, and it might contain a namespace. It
		// also might contain a URI. Clients should just take the last part of the string
		// after '#' and before ":". This is a pathalogical case that might not occur in
		// any deployed AWS service.
		"RestJsonFooErrorUsingCodeUriAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "code": "aws.protocoltests.restjson#FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using __type.
		"RestJsonFooErrorWithDunderType": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using __type, and it might contain a namespace.
		// Clients should just take the last part of the string after '#'.
		"RestJsonFooErrorWithDunderTypeAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "aws.protocoltests.restjson#FooError"
			}`),
			ExpectError: &types.FooError{},
		},
		// Some services serialize errors using __type, and it might contain a namespace.
		// It also might contain a URI. Clients should just take the last part of the
		// string after '#' and before ":". This is a pathalogical case that might not
		// occur in any deployed AWS service.
		"RestJsonFooErrorWithDunderTypeUriAndNamespace": {
			StatusCode: 500,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "__type": "aws.protocoltests.restjson#FooError:http://internal.amazon.com/coral/com.amazon.coral.validate/"
			}`),
			ExpectError: &types.FooError{},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.FooError operation error, got %T", err)
			}
			if e, a := ServiceID, opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.FooError
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.FooError result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}

func TestClient_GreetingWithErrors_ComplexError_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.ComplexError
	}{
		// Serializes a complex error with no message member
		"RestJsonComplexErrorWithNoMessage": {
			StatusCode: 403,
			Header: http.Header{
				"Content-Type":     []string{"application/json"},
				"X-Amzn-Errortype": []string{"ComplexError"},
				"X-Header":         []string{"Header"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "TopLevel": "Top level",
			    "Nested": {
			        "Fooooo": "bar"
			    }
			}`),
			ExpectError: &types.ComplexError{
				Header:   ptr.String("Header"),
				TopLevel: ptr.String("Top level"),
				Nested: &types.ComplexNestedErrorData{
					Foo: ptr.String("bar"),
				},
			},
		},
		"RestJsonEmptyComplexErrorWithNoMessage": {
			StatusCode: 403,
			Header: http.Header{
				"Content-Type":     []string{"application/json"},
				"X-Amzn-Errortype": []string{"ComplexError"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{}`),
			ExpectError:   &types.ComplexError{},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.ComplexError operation error, got %T", err)
			}
			if e, a := ServiceID, opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.ComplexError
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.ComplexError result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}

func TestClient_GreetingWithErrors_InvalidGreeting_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.InvalidGreeting
	}{
		// Parses simple JSON errors
		"RestJsonInvalidGreetingError": {
			StatusCode: 400,
			Header: http.Header{
				"Content-Type":     []string{"application/json"},
				"X-Amzn-Errortype": []string{"InvalidGreeting"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "Message": "Hi"
			}`),
			ExpectError: &types.InvalidGreeting{
				Message: ptr.String("Hi"),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.InvalidGreeting operation error, got %T", err)
			}
			if e, a := ServiceID, opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.InvalidGreeting
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.InvalidGreeting result error, got %T", err)
			}
			if err := smithytesting.CompareValues(c.ExpectError, actualErr, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectError value match:\n%v", err)
			}
		})
	}
}
