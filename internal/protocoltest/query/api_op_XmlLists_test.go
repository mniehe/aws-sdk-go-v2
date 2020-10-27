// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithyrand "github.com/awslabs/smithy-go/rand"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestClient_XmlLists_awsAwsqueryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *XmlListsOutput
	}{
		// Tests for XML list serialization
		"QueryXmlLists": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlListsResponse xmlns="https://example.com/">
			    <XmlListsResult>
			        <stringList>
			            <member>foo</member>
			            <member>bar</member>
			        </stringList>
			        <stringSet>
			            <member>foo</member>
			            <member>bar</member>
			        </stringSet>
			        <integerList>
			            <member>1</member>
			            <member>2</member>
			        </integerList>
			        <booleanList>
			            <member>true</member>
			            <member>false</member>
			        </booleanList>
			        <timestampList>
			            <member>2014-04-29T18:30:38Z</member>
			            <member>2014-04-29T18:30:38Z</member>
			        </timestampList>
			        <enumList>
			            <member>Foo</member>
			            <member>0</member>
			        </enumList>
			        <nestedStringList>
			            <member>
			                <member>foo</member>
			                <member>bar</member>
			            </member>
			            <member>
			                <member>baz</member>
			                <member>qux</member>
			            </member>
			        </nestedStringList>
			        <renamed>
			            <item>foo</item>
			            <item>bar</item>
			        </renamed>
			        <flattenedList>hi</flattenedList>
			        <flattenedList>bye</flattenedList>
			        <customName>yep</customName>
			        <customName>nope</customName>
			        <myStructureList>
			            <item>
			                <value>1</value>
			                <other>2</other>
			            </item>
			            <item>
			                <value>3</value>
			                <other>4</other>
			            </item>
			        </myStructureList>
			    </XmlListsResult>
			</XmlListsResponse>
			`),
			ExpectResult: &XmlListsOutput{
				StringList: []*string{
					ptr.String("foo"),
					ptr.String("bar"),
				},
				StringSet: []*string{
					ptr.String("foo"),
					ptr.String("bar"),
				},
				IntegerList: []*int32{
					ptr.Int32(1),
					ptr.Int32(2),
				},
				BooleanList: []*bool{
					ptr.Bool(true),
					ptr.Bool(false),
				},
				TimestampList: []*time.Time{
					ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
					ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
				},
				EnumList: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("0"),
				},
				NestedStringList: [][]*string{
					{
						ptr.String("foo"),
						ptr.String("bar"),
					},
					{
						ptr.String("baz"),
						ptr.String("qux"),
					},
				},
				RenamedListMembers: []*string{
					ptr.String("foo"),
					ptr.String("bar"),
				},
				FlattenedList: []*string{
					ptr.String("hi"),
					ptr.String("bye"),
				},
				FlattenedList2: []*string{
					ptr.String("yep"),
					ptr.String("nope"),
				},
				StructureList: []*types.StructureListMember{
					{
						A: ptr.String("1"),
						B: ptr.String("2"),
					},
					{
						A: ptr.String("3"),
						B: ptr.String("4"),
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
			var params XmlListsInput
			result, err := client.XmlLists(context.Background(), &params)
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
