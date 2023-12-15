// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/ssoadmin/types"
)

func ExampleAuthenticationMethod_outputUsage() {
	var union types.AuthenticationMethod
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AuthenticationMethodMemberIam:
		_ = v.Value // Value is types.IamAuthenticationMethod

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.IamAuthenticationMethod

func ExampleGrant_outputUsage() {
	var union types.Grant
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.GrantMemberAuthorizationCode:
		_ = v.Value // Value is types.AuthorizationCodeGrant

	case *types.GrantMemberJwtBearer:
		_ = v.Value // Value is types.JwtBearerGrant

	case *types.GrantMemberRefreshToken:
		_ = v.Value // Value is types.RefreshTokenGrant

	case *types.GrantMemberTokenExchange:
		_ = v.Value // Value is types.TokenExchangeGrant

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.JwtBearerGrant
var _ *types.AuthorizationCodeGrant
var _ *types.RefreshTokenGrant
var _ *types.TokenExchangeGrant

func ExampleTrustedTokenIssuerConfiguration_outputUsage() {
	var union types.TrustedTokenIssuerConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TrustedTokenIssuerConfigurationMemberOidcJwtConfiguration:
		_ = v.Value // Value is types.OidcJwtConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.OidcJwtConfiguration

func ExampleTrustedTokenIssuerUpdateConfiguration_outputUsage() {
	var union types.TrustedTokenIssuerUpdateConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TrustedTokenIssuerUpdateConfigurationMemberOidcJwtConfiguration:
		_ = v.Value // Value is types.OidcJwtUpdateConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.OidcJwtUpdateConfiguration
