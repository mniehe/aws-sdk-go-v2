// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfrontkeyvaluestore

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/cloudfrontkeyvaluestore/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpDeleteKey struct {
}

func (*awsRestjson1_serializeOpDeleteKey) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteKeyInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/key-value-stores/{KvsARN}/keys/{Key}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDeleteKeyInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeleteKeyInput(v *DeleteKeyInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.IfMatch != nil && len(*v.IfMatch) > 0 {
		locationName := "If-Match"
		encoder.SetHeader(locationName).String(*v.IfMatch)
	}

	if v.Key == nil || len(*v.Key) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Key must not be empty")}
	}
	if v.Key != nil {
		if err := encoder.SetURI("Key").String(*v.Key); err != nil {
			return err
		}
	}

	if v.KvsARN == nil || len(*v.KvsARN) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member KvsARN must not be empty")}
	}
	if v.KvsARN != nil {
		if err := encoder.SetURI("KvsARN").String(*v.KvsARN); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDescribeKeyValueStore struct {
}

func (*awsRestjson1_serializeOpDescribeKeyValueStore) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeKeyValueStore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeKeyValueStoreInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/key-value-stores/{KvsARN}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDescribeKeyValueStoreInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDescribeKeyValueStoreInput(v *DescribeKeyValueStoreInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.KvsARN == nil || len(*v.KvsARN) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member KvsARN must not be empty")}
	}
	if v.KvsARN != nil {
		if err := encoder.SetURI("KvsARN").String(*v.KvsARN); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetKey struct {
}

func (*awsRestjson1_serializeOpGetKey) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetKeyInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/key-value-stores/{KvsARN}/keys/{Key}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetKeyInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetKeyInput(v *GetKeyInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Key == nil || len(*v.Key) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Key must not be empty")}
	}
	if v.Key != nil {
		if err := encoder.SetURI("Key").String(*v.Key); err != nil {
			return err
		}
	}

	if v.KvsARN == nil || len(*v.KvsARN) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member KvsARN must not be empty")}
	}
	if v.KvsARN != nil {
		if err := encoder.SetURI("KvsARN").String(*v.KvsARN); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListKeys struct {
}

func (*awsRestjson1_serializeOpListKeys) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListKeys) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListKeysInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/key-value-stores/{KvsARN}/keys")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListKeysInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListKeysInput(v *ListKeysInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.KvsARN == nil || len(*v.KvsARN) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member KvsARN must not be empty")}
	}
	if v.KvsARN != nil {
		if err := encoder.SetURI("KvsARN").String(*v.KvsARN); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		encoder.SetQuery("MaxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("NextToken").String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpPutKey struct {
}

func (*awsRestjson1_serializeOpPutKey) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutKey) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutKeyInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/key-value-stores/{KvsARN}/keys/{Key}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PUT"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsPutKeyInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentPutKeyInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsPutKeyInput(v *PutKeyInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.IfMatch != nil && len(*v.IfMatch) > 0 {
		locationName := "If-Match"
		encoder.SetHeader(locationName).String(*v.IfMatch)
	}

	if v.Key == nil || len(*v.Key) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Key must not be empty")}
	}
	if v.Key != nil {
		if err := encoder.SetURI("Key").String(*v.Key); err != nil {
			return err
		}
	}

	if v.KvsARN == nil || len(*v.KvsARN) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member KvsARN must not be empty")}
	}
	if v.KvsARN != nil {
		if err := encoder.SetURI("KvsARN").String(*v.KvsARN); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentPutKeyInput(v *PutKeyInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

type awsRestjson1_serializeOpUpdateKeys struct {
}

func (*awsRestjson1_serializeOpUpdateKeys) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateKeys) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateKeysInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/key-value-stores/{KvsARN}/keys")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUpdateKeysInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdateKeysInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsUpdateKeysInput(v *UpdateKeysInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.IfMatch != nil && len(*v.IfMatch) > 0 {
		locationName := "If-Match"
		encoder.SetHeader(locationName).String(*v.IfMatch)
	}

	if v.KvsARN == nil || len(*v.KvsARN) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member KvsARN must not be empty")}
	}
	if v.KvsARN != nil {
		if err := encoder.SetURI("KvsARN").String(*v.KvsARN); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdateKeysInput(v *UpdateKeysInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Deletes != nil {
		ok := object.Key("Deletes")
		if err := awsRestjson1_serializeDocumentDeleteKeyRequestsList(v.Deletes, ok); err != nil {
			return err
		}
	}

	if v.Puts != nil {
		ok := object.Key("Puts")
		if err := awsRestjson1_serializeDocumentPutKeyRequestsList(v.Puts, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentDeleteKeyRequestListItem(v *types.DeleteKeyRequestListItem, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	return nil
}

func awsRestjson1_serializeDocumentDeleteKeyRequestsList(v []types.DeleteKeyRequestListItem, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentDeleteKeyRequestListItem(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentPutKeyRequestListItem(v *types.PutKeyRequestListItem, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Key != nil {
		ok := object.Key("Key")
		ok.String(*v.Key)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsRestjson1_serializeDocumentPutKeyRequestsList(v []types.PutKeyRequestListItem, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentPutKeyRequestListItem(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}
