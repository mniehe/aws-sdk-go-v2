module github.com/mniehe/aws-sdk-go-v2/example/service/dynamodb/createTable

go 1.19

require (
	github.com/mniehe/aws-sdk-go-v2 v1.24.0
	github.com/mniehe/aws-sdk-go-v2/config v1.26.1
	github.com/mniehe/aws-sdk-go-v2/service/dynamodb v1.26.6
)

require (
	github.com/mniehe/aws-sdk-go-v2/credentials v1.16.12 // indirect
	github.com/mniehe/aws-sdk-go-v2/feature/ec2/imds v1.14.10 // indirect
	github.com/mniehe/aws-sdk-go-v2/internal/configsources v1.2.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2 v2.5.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/internal/ini v1.7.2 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/accept-encoding v1.10.4 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.10 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/presigned-url v1.10.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/sso v1.18.5 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/ssooidc v1.21.5 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/sts v1.26.5 // indirect
	github.com/aws/smithy-go v1.19.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace github.com/mniehe/aws-sdk-go-v2 => ../../../../

replace github.com/mniehe/aws-sdk-go-v2/config => ../../../../config/

replace github.com/mniehe/aws-sdk-go-v2/credentials => ../../../../credentials/

replace github.com/mniehe/aws-sdk-go-v2/feature/ec2/imds => ../../../../feature/ec2/imds/

replace github.com/mniehe/aws-sdk-go-v2/internal/configsources => ../../../../internal/configsources/

replace github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2 => ../../../../internal/endpoints/v2/

replace github.com/mniehe/aws-sdk-go-v2/internal/ini => ../../../../internal/ini/

replace github.com/mniehe/aws-sdk-go-v2/service/dynamodb => ../../../../service/dynamodb/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/accept-encoding => ../../../../service/internal/accept-encoding/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../../../service/internal/endpoint-discovery/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/presigned-url => ../../../../service/internal/presigned-url/

replace github.com/mniehe/aws-sdk-go-v2/service/sso => ../../../../service/sso/

replace github.com/mniehe/aws-sdk-go-v2/service/ssooidc => ../../../../service/ssooidc/

replace github.com/mniehe/aws-sdk-go-v2/service/sts => ../../../../service/sts/
