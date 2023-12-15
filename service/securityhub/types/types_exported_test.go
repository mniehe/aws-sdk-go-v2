// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/securityhub/types"
)

func ExampleConfigurationOptions_outputUsage() {
	var union types.ConfigurationOptions
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConfigurationOptionsMemberBoolean:
		_ = v.Value // Value is types.BooleanConfigurationOptions

	case *types.ConfigurationOptionsMemberDouble:
		_ = v.Value // Value is types.DoubleConfigurationOptions

	case *types.ConfigurationOptionsMemberEnum:
		_ = v.Value // Value is types.EnumConfigurationOptions

	case *types.ConfigurationOptionsMemberEnumList:
		_ = v.Value // Value is types.EnumListConfigurationOptions

	case *types.ConfigurationOptionsMemberInteger:
		_ = v.Value // Value is types.IntegerConfigurationOptions

	case *types.ConfigurationOptionsMemberIntegerList:
		_ = v.Value // Value is types.IntegerListConfigurationOptions

	case *types.ConfigurationOptionsMemberString:
		_ = v.Value // Value is types.StringConfigurationOptions

	case *types.ConfigurationOptionsMemberStringList:
		_ = v.Value // Value is types.StringListConfigurationOptions

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.IntegerConfigurationOptions
var _ *types.StringConfigurationOptions
var _ *types.StringListConfigurationOptions
var _ *types.DoubleConfigurationOptions
var _ *types.IntegerListConfigurationOptions
var _ *types.BooleanConfigurationOptions
var _ *types.EnumConfigurationOptions
var _ *types.EnumListConfigurationOptions

func ExampleParameterValue_outputUsage() {
	var union types.ParameterValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ParameterValueMemberBoolean:
		_ = v.Value // Value is bool

	case *types.ParameterValueMemberDouble:
		_ = v.Value // Value is float64

	case *types.ParameterValueMemberEnum:
		_ = v.Value // Value is string

	case *types.ParameterValueMemberEnumList:
		_ = v.Value // Value is []string

	case *types.ParameterValueMemberInteger:
		_ = v.Value // Value is int32

	case *types.ParameterValueMemberIntegerList:
		_ = v.Value // Value is []int32

	case *types.ParameterValueMemberString:
		_ = v.Value // Value is string

	case *types.ParameterValueMemberStringList:
		_ = v.Value // Value is []string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ []string
var _ *int32
var _ *bool
var _ *float64
var _ []int32

func ExamplePolicy_outputUsage() {
	var union types.Policy
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PolicyMemberSecurityHub:
		_ = v.Value // Value is types.SecurityHubPolicy

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SecurityHubPolicy

func ExampleTarget_outputUsage() {
	var union types.Target
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TargetMemberAccountId:
		_ = v.Value // Value is string

	case *types.TargetMemberOrganizationalUnitId:
		_ = v.Value // Value is string

	case *types.TargetMemberRootId:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
