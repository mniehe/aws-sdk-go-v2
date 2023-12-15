// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/redshift/types"
)

func ExampleLakeFormationScopeUnion_outputUsage() {
	var union types.LakeFormationScopeUnion
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.LakeFormationScopeUnionMemberLakeFormationQuery:
		_ = v.Value // Value is types.LakeFormationQuery

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.LakeFormationQuery

func ExampleServiceIntegrationsUnion_outputUsage() {
	var union types.ServiceIntegrationsUnion
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ServiceIntegrationsUnionMemberLakeFormation:
		_ = v.Value // Value is []types.LakeFormationScopeUnion

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []types.LakeFormationScopeUnion
