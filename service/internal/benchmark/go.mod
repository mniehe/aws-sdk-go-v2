module github.com/mniehe/aws-sdk-go-v2/service/internal/benchmark

go 1.19

require (
	github.com/aws/aws-sdk-go v1.44.28
	github.com/mniehe/aws-sdk-go-v2 v1.24.0
	github.com/mniehe/aws-sdk-go-v2/service/dynamodb v1.26.6
	github.com/mniehe/aws-sdk-go-v2/service/lexruntimeservice v1.18.5
	github.com/mniehe/aws-sdk-go-v2/service/s3 v1.47.5
	github.com/mniehe/aws-sdk-go-v2/service/schemas v1.21.5
	github.com/aws/smithy-go v1.19.0
)

require (
	github.com/mniehe/aws-sdk-go-v2/aws/protocol/eventstream v1.5.4 // indirect
	github.com/mniehe/aws-sdk-go-v2/internal/configsources v1.2.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2 v2.5.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/internal/v4a v1.2.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/accept-encoding v1.10.4 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/checksum v1.2.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.10 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/presigned-url v1.10.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/s3shared v1.16.9 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace github.com/mniehe/aws-sdk-go-v2 => ../../../

replace github.com/mniehe/aws-sdk-go-v2/aws/protocol/eventstream => ../../../aws/protocol/eventstream/

replace github.com/mniehe/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace github.com/mniehe/aws-sdk-go-v2/internal/v4a => ../../../internal/v4a/

replace github.com/mniehe/aws-sdk-go-v2/service/dynamodb => ../../../service/dynamodb/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/checksum => ../../../service/internal/checksum/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../../service/internal/endpoint-discovery/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/presigned-url => ../../../service/internal/presigned-url/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/s3shared => ../../../service/internal/s3shared/

replace github.com/mniehe/aws-sdk-go-v2/service/lexruntimeservice => ../../../service/lexruntimeservice/

replace github.com/mniehe/aws-sdk-go-v2/service/s3 => ../../../service/s3/

replace github.com/mniehe/aws-sdk-go-v2/service/schemas => ../../../service/schemas/
