module github.com/mniehe/aws-sdk-go-v2/service/internal/eventstreamtesting

go 1.19

require (
	github.com/mniehe/aws-sdk-go-v2 v1.24.0
	github.com/mniehe/aws-sdk-go-v2/aws/protocol/eventstream v1.5.4
	github.com/mniehe/aws-sdk-go-v2/credentials v1.16.12
	golang.org/x/net v0.19.0
)

require (
	github.com/aws/smithy-go v1.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/mniehe/aws-sdk-go-v2 => ../../../

replace github.com/mniehe/aws-sdk-go-v2/aws/protocol/eventstream => ../../../aws/protocol/eventstream/

replace github.com/mniehe/aws-sdk-go-v2/credentials => ../../../credentials/

replace github.com/mniehe/aws-sdk-go-v2/feature/ec2/imds => ../../../feature/ec2/imds/

replace github.com/mniehe/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/presigned-url => ../../../service/internal/presigned-url/

replace github.com/mniehe/aws-sdk-go-v2/service/sso => ../../../service/sso/

replace github.com/mniehe/aws-sdk-go-v2/service/ssooidc => ../../../service/ssooidc/

replace github.com/mniehe/aws-sdk-go-v2/service/sts => ../../../service/sts/
