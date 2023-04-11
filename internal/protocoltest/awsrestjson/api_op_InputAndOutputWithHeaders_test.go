// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestClient_InputAndOutputWithHeaders_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *InputAndOutputWithHeadersInput
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
		// Tests requests with string header bindings
		"RestJsonInputAndOutputWithStringHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderString: ptr.String("Hello"),
				HeaderStringList: []string{
					"a",
					"b",
					"c",
				},
				HeaderStringSet: []string{
					"a",
					"b",
					"c",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-String":     []string{"Hello"},
				"X-StringList": []string{"a, b, c"},
				"X-StringSet":  []string{"a, b, c"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Tests requests with string list header bindings that require quoting
		"RestJsonInputAndOutputWithQuotedStringHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderStringList: []string{
					"b,c",
					"\"def\"",
					"a",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-StringList": []string{"\"b,c\", \"\\\"def\\\"\", a"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Tests requests with numeric header bindings
		"RestJsonInputAndOutputWithNumericHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderByte:    ptr.Int8(1),
				HeaderShort:   ptr.Int16(123),
				HeaderInteger: ptr.Int32(123),
				HeaderLong:    ptr.Int64(123),
				HeaderFloat:   ptr.Float32(1.1),
				HeaderDouble:  ptr.Float64(1.1),
				HeaderIntegerList: []int32{
					1,
					2,
					3,
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-Byte":        []string{"1"},
				"X-Double":      []string{"1.1"},
				"X-Float":       []string{"1.1"},
				"X-Integer":     []string{"123"},
				"X-IntegerList": []string{"1, 2, 3"},
				"X-Long":        []string{"123"},
				"X-Short":       []string{"123"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Tests requests with boolean header bindings
		"RestJsonInputAndOutputWithBooleanHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderTrueBool:  ptr.Bool(true),
				HeaderFalseBool: ptr.Bool(false),
				HeaderBooleanList: []bool{
					true,
					false,
					true,
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-Boolean1":    []string{"true"},
				"X-Boolean2":    []string{"false"},
				"X-BooleanList": []string{"true, false, true"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Tests requests with timestamp header bindings
		"RestJsonInputAndOutputWithTimestampHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderTimestampList: []time.Time{
					smithytime.ParseEpochSeconds(1576540098),
					smithytime.ParseEpochSeconds(1576540098),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-TimestampList": []string{"Mon, 16 Dec 2019 23:48:18 GMT, Mon, 16 Dec 2019 23:48:18 GMT"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Tests requests with enum header bindings
		"RestJsonInputAndOutputWithEnumHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderEnum: types.FooEnum("Foo"),
				HeaderEnumList: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("Bar"),
					types.FooEnum("Baz"),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-Enum":     []string{"Foo"},
				"X-EnumList": []string{"Foo, Bar, Baz"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Tests requests with intEnum header bindings
		"RestJsonInputAndOutputWithIntEnumHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderIntegerEnum: 1,
				HeaderIntegerEnumList: []types.IntegerEnum{
					1,
					2,
					3,
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-IntegerEnum":     []string{"1"},
				"X-IntegerEnumList": []string{"1, 2, 3"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Supports handling NaN float header values.
		"RestJsonSupportsNaNFloatHeaderInputs": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderFloat:  ptr.Float32(float32(math.NaN())),
				HeaderDouble: ptr.Float64(math.NaN()),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-Double": []string{"NaN"},
				"X-Float":  []string{"NaN"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Supports handling Infinity float header values.
		"RestJsonSupportsInfinityFloatHeaderInputs": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderFloat:  ptr.Float32(float32(math.Inf(1))),
				HeaderDouble: ptr.Float64(math.Inf(1)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-Double": []string{"Infinity"},
				"X-Float":  []string{"Infinity"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Supports handling -Infinity float header values.
		"RestJsonSupportsNegativeInfinityFloatHeaderInputs": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderFloat:  ptr.Float32(float32(math.Inf(-1))),
				HeaderDouble: ptr.Float64(math.Inf(-1)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-Double": []string{"-Infinity"},
				"X-Float":  []string{"-Infinity"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			serverURL := server.URL
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
				HTTPClient:               awshttp.NewBuildableClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.InputAndOutputWithHeaders(context.Background(), c.Params)
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

func TestClient_InputAndOutputWithHeaders_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *InputAndOutputWithHeadersOutput
	}{
		// Tests responses with string header bindings
		"RestJsonInputAndOutputWithStringHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-String":     []string{"Hello"},
				"X-StringList": []string{"a, b, c"},
				"X-StringSet":  []string{"a, b, c"},
			},
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderString: ptr.String("Hello"),
				HeaderStringList: []string{
					"a",
					"b",
					"c",
				},
				HeaderStringSet: []string{
					"a",
					"b",
					"c",
				},
			},
		},
		// Tests responses with string list header bindings that require quoting
		"RestJsonInputAndOutputWithQuotedStringHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-StringList": []string{"\"b,c\", \"\\\"def\\\"\", a"},
			},
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderStringList: []string{
					"b,c",
					"\"def\"",
					"a",
				},
			},
		},
		// Tests responses with numeric header bindings
		"RestJsonInputAndOutputWithNumericHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-Byte":        []string{"1"},
				"X-Double":      []string{"1.1"},
				"X-Float":       []string{"1.1"},
				"X-Integer":     []string{"123"},
				"X-IntegerList": []string{"1, 2, 3"},
				"X-Long":        []string{"123"},
				"X-Short":       []string{"123"},
			},
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderByte:    ptr.Int8(1),
				HeaderShort:   ptr.Int16(123),
				HeaderInteger: ptr.Int32(123),
				HeaderLong:    ptr.Int64(123),
				HeaderFloat:   ptr.Float32(1.1),
				HeaderDouble:  ptr.Float64(1.1),
				HeaderIntegerList: []int32{
					1,
					2,
					3,
				},
			},
		},
		// Tests responses with boolean header bindings
		"RestJsonInputAndOutputWithBooleanHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-Boolean1":    []string{"true"},
				"X-Boolean2":    []string{"false"},
				"X-BooleanList": []string{"true, false, true"},
			},
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderTrueBool:  ptr.Bool(true),
				HeaderFalseBool: ptr.Bool(false),
				HeaderBooleanList: []bool{
					true,
					false,
					true,
				},
			},
		},
		// Tests responses with timestamp header bindings
		"RestJsonInputAndOutputWithTimestampHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-TimestampList": []string{"Mon, 16 Dec 2019 23:48:18 GMT, Mon, 16 Dec 2019 23:48:18 GMT"},
			},
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderTimestampList: []time.Time{
					smithytime.ParseEpochSeconds(1576540098),
					smithytime.ParseEpochSeconds(1576540098),
				},
			},
		},
		// Tests responses with enum header bindings
		"RestJsonInputAndOutputWithEnumHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-Enum":     []string{"Foo"},
				"X-EnumList": []string{"Foo, Bar, Baz"},
			},
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderEnum: types.FooEnum("Foo"),
				HeaderEnumList: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("Bar"),
					types.FooEnum("Baz"),
				},
			},
		},
		// Tests responses with intEnum header bindings
		"RestJsonInputAndOutputWithIntEnumHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-IntegerEnum":     []string{"1"},
				"X-IntegerEnumList": []string{"1, 2, 3"},
			},
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderIntegerEnum: 1,
				HeaderIntegerEnumList: []types.IntegerEnum{
					1,
					2,
					3,
				},
			},
		},
		// Supports handling NaN float header values.
		"RestJsonSupportsNaNFloatHeaderOutputs": {
			StatusCode: 200,
			Header: http.Header{
				"X-Double": []string{"NaN"},
				"X-Float":  []string{"NaN"},
			},
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderFloat:  ptr.Float32(float32(math.NaN())),
				HeaderDouble: ptr.Float64(math.NaN()),
			},
		},
		// Supports handling Infinity float header values.
		"RestJsonSupportsInfinityFloatHeaderOutputs": {
			StatusCode: 200,
			Header: http.Header{
				"X-Double": []string{"Infinity"},
				"X-Float":  []string{"Infinity"},
			},
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderFloat:  ptr.Float32(float32(math.Inf(1))),
				HeaderDouble: ptr.Float64(math.Inf(1)),
			},
		},
		// Supports handling -Infinity float header values.
		"RestJsonSupportsNegativeInfinityFloatHeaderOutputs": {
			StatusCode: 200,
			Header: http.Header{
				"X-Double": []string{"-Infinity"},
				"X-Float":  []string{"-Infinity"},
			},
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderFloat:  ptr.Float32(float32(math.Inf(-1))),
				HeaderDouble: ptr.Float64(math.Inf(-1)),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
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
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params InputAndOutputWithHeadersInput
			result, err := client.InputAndOutputWithHeaders(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(
					middleware.Metadata{},
				),
				cmp.FilterValues(func(x, y float64) bool {
					return math.IsNaN(x) && math.IsNaN(y)
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmp.FilterValues(func(x, y float32) bool {
					return math.IsNaN(float64(x)) && math.IsNaN(float64(y))
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmpopts.IgnoreTypes(smithydocument.NoSerde{}),
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, opts...); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
