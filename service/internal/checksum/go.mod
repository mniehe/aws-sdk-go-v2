module github.com/mniehe/aws-sdk-go-v2/service/internal/checksum

go 1.19

require (
	github.com/mniehe/aws-sdk-go-v2 v1.24.0
	github.com/mniehe/aws-sdk-go-v2/service/internal/presigned-url v1.10.9
	github.com/aws/smithy-go v1.19.0
	github.com/google/go-cmp v0.5.8
)

replace github.com/mniehe/aws-sdk-go-v2 => ../../../

replace github.com/mniehe/aws-sdk-go-v2/service/internal/presigned-url => ../../../service/internal/presigned-url/
