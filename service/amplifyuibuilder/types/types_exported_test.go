// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/amplifyuibuilder/types"
)

func ExampleApiConfiguration_outputUsage() {
	var union types.ApiConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ApiConfigurationMemberDataStoreConfig:
		_ = v.Value // Value is types.DataStoreRenderConfig

	case *types.ApiConfigurationMemberGraphQLConfig:
		_ = v.Value // Value is types.GraphQLRenderConfig

	case *types.ApiConfigurationMemberNoApiConfig:
		_ = v.Value // Value is types.NoApiRenderConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.DataStoreRenderConfig
var _ *types.NoApiRenderConfig
var _ *types.GraphQLRenderConfig

func ExampleCodegenJobRenderConfig_outputUsage() {
	var union types.CodegenJobRenderConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CodegenJobRenderConfigMemberReact:
		_ = v.Value // Value is types.ReactStartCodegenJobData

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ReactStartCodegenJobData

func ExampleFieldPosition_outputUsage() {
	var union types.FieldPosition
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FieldPositionMemberBelow:
		_ = v.Value // Value is string

	case *types.FieldPositionMemberFixed:
		_ = v.Value // Value is types.FixedPosition

	case *types.FieldPositionMemberRightOf:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ types.FixedPosition

func ExampleFormStyleConfig_outputUsage() {
	var union types.FormStyleConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FormStyleConfigMemberTokenReference:
		_ = v.Value // Value is string

	case *types.FormStyleConfigMemberValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
