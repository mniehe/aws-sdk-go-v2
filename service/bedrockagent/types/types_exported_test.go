// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/bedrockagent/types"
)

func ExampleActionGroupExecutor_outputUsage() {
	var union types.ActionGroupExecutor
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ActionGroupExecutorMemberLambda:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleAPISchema_outputUsage() {
	var union types.APISchema
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.APISchemaMemberPayload:
		_ = v.Value // Value is string

	case *types.APISchemaMemberS3:
		_ = v.Value // Value is types.S3Identifier

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *types.S3Identifier
