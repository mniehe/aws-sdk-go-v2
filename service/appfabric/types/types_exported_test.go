// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/appfabric/types"
)

func ExampleCredential_outputUsage() {
	var union types.Credential
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CredentialMemberApiKeyCredential:
		_ = v.Value // Value is types.ApiKeyCredential

	case *types.CredentialMemberOauth2Credential:
		_ = v.Value // Value is types.Oauth2Credential

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ApiKeyCredential
var _ *types.Oauth2Credential

func ExampleDestination_outputUsage() {
	var union types.Destination
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DestinationMemberFirehoseStream:
		_ = v.Value // Value is types.FirehoseStream

	case *types.DestinationMemberS3Bucket:
		_ = v.Value // Value is types.S3Bucket

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FirehoseStream
var _ *types.S3Bucket

func ExampleDestinationConfiguration_outputUsage() {
	var union types.DestinationConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DestinationConfigurationMemberAuditLog:
		_ = v.Value // Value is types.AuditLogDestinationConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AuditLogDestinationConfiguration

func ExampleProcessingConfiguration_outputUsage() {
	var union types.ProcessingConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ProcessingConfigurationMemberAuditLog:
		_ = v.Value // Value is types.AuditLogProcessingConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AuditLogProcessingConfiguration
