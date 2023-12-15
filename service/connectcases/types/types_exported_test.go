// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/connectcases/types"
)

func ExampleCaseFilter_outputUsage() {
	var union types.CaseFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CaseFilterMemberAndAll:
		_ = v.Value // Value is []types.CaseFilter

	case *types.CaseFilterMemberField:
		_ = v.Value // Value is types.FieldFilter

	case *types.CaseFilterMemberNot:
		_ = v.Value // Value is types.CaseFilter

	case *types.CaseFilterMemberOrAll:
		_ = v.Value // Value is []types.CaseFilter

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []types.CaseFilter
var _ types.CaseFilter
var _ types.FieldFilter

func ExampleFieldFilter_outputUsage() {
	var union types.FieldFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FieldFilterMemberContains:
		_ = v.Value // Value is types.FieldValue

	case *types.FieldFilterMemberEqualTo:
		_ = v.Value // Value is types.FieldValue

	case *types.FieldFilterMemberGreaterThan:
		_ = v.Value // Value is types.FieldValue

	case *types.FieldFilterMemberGreaterThanOrEqualTo:
		_ = v.Value // Value is types.FieldValue

	case *types.FieldFilterMemberLessThan:
		_ = v.Value // Value is types.FieldValue

	case *types.FieldFilterMemberLessThanOrEqualTo:
		_ = v.Value // Value is types.FieldValue

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FieldValue

func ExampleFieldValueUnion_outputUsage() {
	var union types.FieldValueUnion
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FieldValueUnionMemberBooleanValue:
		_ = v.Value // Value is bool

	case *types.FieldValueUnionMemberDoubleValue:
		_ = v.Value // Value is float64

	case *types.FieldValueUnionMemberEmptyValue:
		_ = v.Value // Value is types.EmptyFieldValue

	case *types.FieldValueUnionMemberStringValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EmptyFieldValue
var _ *string
var _ *bool
var _ *float64

func ExampleLayoutContent_outputUsage() {
	var union types.LayoutContent
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.LayoutContentMemberBasic:
		_ = v.Value // Value is types.BasicLayout

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.BasicLayout

func ExampleRelatedItemContent_outputUsage() {
	var union types.RelatedItemContent
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RelatedItemContentMemberComment:
		_ = v.Value // Value is types.CommentContent

	case *types.RelatedItemContentMemberContact:
		_ = v.Value // Value is types.ContactContent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ContactContent
var _ *types.CommentContent

func ExampleRelatedItemInputContent_outputUsage() {
	var union types.RelatedItemInputContent
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RelatedItemInputContentMemberComment:
		_ = v.Value // Value is types.CommentContent

	case *types.RelatedItemInputContentMemberContact:
		_ = v.Value // Value is types.Contact

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.CommentContent
var _ *types.Contact

func ExampleRelatedItemTypeFilter_outputUsage() {
	var union types.RelatedItemTypeFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.RelatedItemTypeFilterMemberComment:
		_ = v.Value // Value is types.CommentFilter

	case *types.RelatedItemTypeFilterMemberContact:
		_ = v.Value // Value is types.ContactFilter

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ContactFilter
var _ *types.CommentFilter

func ExampleSection_outputUsage() {
	var union types.Section
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SectionMemberFieldGroup:
		_ = v.Value // Value is types.FieldGroup

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FieldGroup

func ExampleUserUnion_outputUsage() {
	var union types.UserUnion
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.UserUnionMemberUserArn:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
