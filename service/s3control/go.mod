module github.com/aws/aws-sdk-go-v2/service/s3control

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v0.29.1-0.20201112231636-9ae467d8157d
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v0.3.1-0.20201112231636-9ae467d8157d
	github.com/awslabs/smithy-go v0.3.1-0.20201108010311-62c2a93810b4
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/service/internal/s3shared => ../../service/internal/s3shared/
