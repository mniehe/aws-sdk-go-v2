// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a model compilation job. After the model has been compiled, Amazon
// SageMaker saves the resulting model artifacts to an Amazon Simple Storage
// Service (Amazon S3) bucket that you specify. If you choose to host your model
// using Amazon SageMaker hosting services, you can use the resulting model
// artifacts as part of the model. You can also use the artifacts with Amazon Web
// Services IoT Greengrass. In that case, deploy them as an ML resource. In the
// request body, you provide the following:
//   - A name for the compilation job
//   - Information about the input model artifacts
//   - The output location for the compiled model and the device (target) that the
//     model runs on
//   - The Amazon Resource Name (ARN) of the IAM role that Amazon SageMaker
//     assumes to perform the model compilation job.
//
// You can also provide a Tag to track the model compilation job's resource use
// and costs. The response body contains the CompilationJobArn for the compiled
// job. To stop a model compilation job, use StopCompilationJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StopCompilationJob.html)
// . To get information about a particular model compilation job, use
// DescribeCompilationJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeCompilationJob.html)
// . To get information about multiple model compilation jobs, use
// ListCompilationJobs (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListCompilationJobs.html)
// .
func (c *Client) CreateCompilationJob(ctx context.Context, params *CreateCompilationJobInput, optFns ...func(*Options)) (*CreateCompilationJobOutput, error) {
	if params == nil {
		params = &CreateCompilationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCompilationJob", params, optFns, c.addOperationCreateCompilationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCompilationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCompilationJobInput struct {

	// A name for the model compilation job. The name must be unique within the Amazon
	// Web Services Region and within your Amazon Web Services account.
	//
	// This member is required.
	CompilationJobName *string

	// Provides information about the output location for the compiled model and the
	// target device the model runs on.
	//
	// This member is required.
	OutputConfig *types.OutputConfig

	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to
	// perform tasks on your behalf. During model compilation, Amazon SageMaker needs
	// your permission to:
	//   - Read input data from an S3 bucket
	//   - Write model artifacts to an S3 bucket
	//   - Write logs to Amazon CloudWatch Logs
	//   - Publish metrics to Amazon CloudWatch
	// You grant permissions for all of these tasks to an IAM role. To pass this role
	// to Amazon SageMaker, the caller of this API must have the iam:PassRole
	// permission. For more information, see Amazon SageMaker Roles. (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html)
	//
	// This member is required.
	RoleArn *string

	// Specifies a limit to how long a model compilation job can run. When the job
	// reaches the time limit, Amazon SageMaker ends the compilation job. Use this API
	// to cap model training costs.
	//
	// This member is required.
	StoppingCondition *types.StoppingCondition

	// Provides information about the location of input model artifacts, the name and
	// shape of the expected data inputs, and the framework in which the model was
	// trained.
	InputConfig *types.InputConfig

	// The Amazon Resource Name (ARN) of a versioned model package. Provide either a
	// ModelPackageVersionArn or an InputConfig object in the request syntax. The
	// presence of both objects in the CreateCompilationJob request will return an
	// exception.
	ModelPackageVersionArn *string

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// .
	Tags []types.Tag

	// A VpcConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VpcConfig.html)
	// object that specifies the VPC that you want your compilation job to connect to.
	// Control access to your models by configuring the VPC. For more information, see
	// Protect Compilation Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/neo-vpc.html)
	// .
	VpcConfig *types.NeoVpcConfig

	noSmithyDocumentSerde
}

type CreateCompilationJobOutput struct {

	// If the action is successful, the service sends back an HTTP 200 response.
	// Amazon SageMaker returns the following data in JSON format:
	//   - CompilationJobArn : The Amazon Resource Name (ARN) of the compiled job.
	//
	// This member is required.
	CompilationJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCompilationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCompilationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCompilationJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCompilationJob"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateCompilationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCompilationJob(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateCompilationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCompilationJob",
	}
}
