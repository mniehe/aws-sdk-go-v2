// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"bytes"
	"context"
	"github.com/mniehe/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/mniehe/aws-sdk-go-v2/internal/protocoltest"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	smithyprivateprotocol "github.com/aws/smithy-go/private/protocol"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_StreamingTraitsRequireLength_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *StreamingTraitsRequireLengthInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		Host          *url.URL
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes a blob in the HTTP payload with a required length
		"RestJsonStreamingTraitsRequireLengthWithBlob": {
			Params: &StreamingTraitsRequireLengthInput{
				Foo:  ptr.String("Foo"),
				Blob: smithyio.ReadSeekNopCloser{ReadSeeker: bytes.NewReader([]byte("blobby blob blob"))},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/StreamingTraitsRequireLength",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/octet-stream"},
				"X-Foo":        []string{"Foo"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/octet-stream",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderBytes(actual, []byte(`blobby blob blob`))
			},
		},
		// Serializes an empty blob in the HTTP payload
		"RestJsonStreamingTraitsRequireLengthWithNoBlobBody": {
			Params: &StreamingTraitsRequireLengthInput{
				Foo: ptr.String("Foo"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/StreamingTraitsRequireLength",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-Foo": []string{"Foo"},
			},
			BodyMediaType: "application/octet-stream",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actualReq := &http.Request{}
			serverURL := "http://localhost:8888/"
			if c.Host != nil {
				u, err := url.Parse(serverURL)
				if err != nil {
					t.Fatalf("expect no error, got %v", err)
				}
				u.Path = c.Host.Path
				u.RawPath = c.Host.RawPath
				u.RawQuery = c.Host.RawQuery
				serverURL = u.String()
			}
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               protocoltesthttp.NewClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.StreamingTraitsRequireLength(context.Background(), c.Params, func(options *Options) {
				options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
					return smithyprivateprotocol.AddCaptureRequestMiddleware(stack, actualReq)
				})
			})
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}
