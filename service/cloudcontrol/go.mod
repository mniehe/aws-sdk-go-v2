module github.com/mniehe/aws-sdk-go-v2/service/cloudcontrol

go 1.19

require (
	github.com/mniehe/aws-sdk-go-v2 v1.24.0
	github.com/mniehe/aws-sdk-go-v2/internal/configsources v1.2.9
	github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2 v2.5.9
	github.com/aws/smithy-go v1.19.0
	github.com/google/go-cmp v0.5.8
	github.com/jmespath/go-jmespath v0.4.0
)

replace github.com/mniehe/aws-sdk-go-v2 => ../../

replace github.com/mniehe/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
