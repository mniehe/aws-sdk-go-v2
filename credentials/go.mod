module github.com/mniehe/aws-sdk-go-v2/credentials

go 1.19

require (
	github.com/mniehe/aws-sdk-go-v2 v1.24.0
	github.com/mniehe/aws-sdk-go-v2/feature/ec2/imds v1.14.10
	github.com/mniehe/aws-sdk-go-v2/service/sso v1.18.5
	github.com/mniehe/aws-sdk-go-v2/service/ssooidc v1.21.5
	github.com/mniehe/aws-sdk-go-v2/service/sts v1.26.5
	github.com/aws/smithy-go v1.19.0
	github.com/google/go-cmp v0.5.8
)

require (
	github.com/mniehe/aws-sdk-go-v2/internal/configsources v1.2.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2 v2.5.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/accept-encoding v1.10.4 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/presigned-url v1.10.9 // indirect
)

replace github.com/mniehe/aws-sdk-go-v2 => ../

replace github.com/mniehe/aws-sdk-go-v2/feature/ec2/imds => ../feature/ec2/imds/

replace github.com/mniehe/aws-sdk-go-v2/internal/configsources => ../internal/configsources/

replace github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2 => ../internal/endpoints/v2/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/accept-encoding => ../service/internal/accept-encoding/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/presigned-url => ../service/internal/presigned-url/

replace github.com/mniehe/aws-sdk-go-v2/service/sso => ../service/sso/

replace github.com/mniehe/aws-sdk-go-v2/service/ssooidc => ../service/ssooidc/

replace github.com/mniehe/aws-sdk-go-v2/service/sts => ../service/sts/
