module github.com/mniehe/aws-sdk-go-v2/feature/dynamodb/expression

go 1.19

require (
	github.com/mniehe/aws-sdk-go-v2 v1.24.0
	github.com/mniehe/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.12.12
	github.com/mniehe/aws-sdk-go-v2/service/dynamodb v1.26.6
	github.com/google/go-cmp v0.5.8
)

require (
	github.com/mniehe/aws-sdk-go-v2/internal/configsources v1.2.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2 v2.5.9 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/dynamodbstreams v1.18.5 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/accept-encoding v1.10.4 // indirect
	github.com/mniehe/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.10 // indirect
	github.com/aws/smithy-go v1.19.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace github.com/mniehe/aws-sdk-go-v2 => ../../../

replace github.com/mniehe/aws-sdk-go-v2/feature/dynamodb/attributevalue => ../../../feature/dynamodb/attributevalue/

replace github.com/mniehe/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/mniehe/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace github.com/mniehe/aws-sdk-go-v2/service/dynamodb => ../../../service/dynamodb/

replace github.com/mniehe/aws-sdk-go-v2/service/dynamodbstreams => ../../../service/dynamodbstreams/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace github.com/mniehe/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../../service/internal/endpoint-discovery/
