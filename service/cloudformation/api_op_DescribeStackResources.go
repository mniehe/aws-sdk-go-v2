// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns Amazon Web Services resource descriptions for running and deleted
// stacks. If StackName is specified, all the associated resources that are part
// of the stack are returned. If PhysicalResourceId is specified, the associated
// resources of the stack that the resource belongs to are returned. Only the first
// 100 resources will be returned. If your stack has more resources than this, you
// should use ListStackResources instead. For deleted stacks,
// DescribeStackResources returns resource information for up to 90 days after the
// stack has been deleted. You must specify either StackName or PhysicalResourceId
// , but not both. In addition, you can specify LogicalResourceId to filter the
// returned result. For more information about resources, the LogicalResourceId
// and PhysicalResourceId , go to the CloudFormation User Guide (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/)
// . A ValidationError is returned if you specify both StackName and
// PhysicalResourceId in the same request.
func (c *Client) DescribeStackResources(ctx context.Context, params *DescribeStackResourcesInput, optFns ...func(*Options)) (*DescribeStackResourcesOutput, error) {
	if params == nil {
		params = &DescribeStackResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStackResources", params, optFns, c.addOperationDescribeStackResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStackResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for DescribeStackResources action.
type DescribeStackResourcesInput struct {

	// The logical name of the resource as specified in the template. Default: There
	// is no default value.
	LogicalResourceId *string

	// The name or unique identifier that corresponds to a physical instance ID of a
	// resource supported by CloudFormation. For example, for an Amazon Elastic Compute
	// Cloud (EC2) instance, PhysicalResourceId corresponds to the InstanceId . You can
	// pass the EC2 InstanceId to DescribeStackResources to find which stack the
	// instance belongs to and what other resources are part of the stack. Required:
	// Conditional. If you don't specify PhysicalResourceId , you must specify
	// StackName . Default: There is no default value.
	PhysicalResourceId *string

	// The name or the unique stack ID that is associated with the stack, which aren't
	// always interchangeable:
	//   - Running stacks: You can specify either the stack's name or its unique stack
	//   ID.
	//   - Deleted stacks: You must specify the unique stack ID.
	// Default: There is no default value. Required: Conditional. If you don't specify
	// StackName , you must specify PhysicalResourceId .
	StackName *string

	noSmithyDocumentSerde
}

// The output for a DescribeStackResources action.
type DescribeStackResourcesOutput struct {

	// A list of StackResource structures.
	StackResources []types.StackResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStackResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStackResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStackResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStackResources"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStackResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStackResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStackResources",
	}
}
