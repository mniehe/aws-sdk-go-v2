// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
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

func TestClient_XmlNamespaces_awsEc2queryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *XmlNamespacesOutput
	}{
		// Serializes XML namespaces
		"Ec2XmlNamespaces": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml;charset=UTF-8"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlNamespacesResponse xmlns="http://foo.com" xmlns="https://example.com/">
			    <nested>
			        <foo xmlns:baz="http://baz.com">Foo</foo>
			        <values xmlns="http://qux.com">
			            <member xmlns="http://bux.com">Bar</member>
			            <member xmlns="http://bux.com">Baz</member>
			        </values>
			    </nested>
			    <RequestId>requestid</RequestId>
			</XmlNamespacesResponse>
			`),
			ExpectResult: &XmlNamespacesOutput{
				Nested: &types.XmlNamespaceNested{
					Foo: ptr.String("Foo"),
					Values: []*string{
						ptr.String("Bar"),
						ptr.String("Baz"),
					},
				},
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
			var params XmlNamespacesInput
			result, err := client.XmlNamespaces(context.Background(), &params)
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
