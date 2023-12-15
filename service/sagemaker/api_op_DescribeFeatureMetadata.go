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
	"time"
)

// Shows the metadata for a feature within a feature group.
func (c *Client) DescribeFeatureMetadata(ctx context.Context, params *DescribeFeatureMetadataInput, optFns ...func(*Options)) (*DescribeFeatureMetadataOutput, error) {
	if params == nil {
		params = &DescribeFeatureMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFeatureMetadata", params, optFns, c.addOperationDescribeFeatureMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFeatureMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFeatureMetadataInput struct {

	// The name or Amazon Resource Name (ARN) of the feature group containing the
	// feature.
	//
	// This member is required.
	FeatureGroupName *string

	// The name of the feature.
	//
	// This member is required.
	FeatureName *string

	noSmithyDocumentSerde
}

type DescribeFeatureMetadataOutput struct {

	// A timestamp indicating when the feature was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Number (ARN) of the feature group that contains the feature.
	//
	// This member is required.
	FeatureGroupArn *string

	// The name of the feature group that you've specified.
	//
	// This member is required.
	FeatureGroupName *string

	// The name of the feature that you've specified.
	//
	// This member is required.
	FeatureName *string

	// The data type of the feature.
	//
	// This member is required.
	FeatureType types.FeatureType

	// A timestamp indicating when the metadata for the feature group was modified.
	// For example, if you add a parameter describing the feature, the timestamp
	// changes to reflect the last time you
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The description you added to describe the feature.
	Description *string

	// The key-value pairs that you added to describe the feature.
	Parameters []types.FeatureParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFeatureMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFeatureMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFeatureMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFeatureMetadata"); err != nil {
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
	if err = addOpDescribeFeatureMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFeatureMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFeatureMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFeatureMetadata",
	}
}
