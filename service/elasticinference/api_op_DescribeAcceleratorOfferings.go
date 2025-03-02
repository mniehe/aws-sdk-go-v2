// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticinference

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/elasticinference/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the locations in which a given accelerator type or set of types is
// present in a given region. February 15, 2023: Starting April 15, 2023, AWS will
// not onboard new customers to Amazon Elastic Inference (EI), and will help
// current customers migrate their workloads to options that offer better price and
// performance. After April 15, 2023, new customers will not be able to launch
// instances with Amazon EI accelerators in Amazon SageMaker, Amazon ECS, or Amazon
// EC2. However, customers who have used Amazon EI at least once during the past
// 30-day period are considered current customers and will be able to continue
// using the service.
func (c *Client) DescribeAcceleratorOfferings(ctx context.Context, params *DescribeAcceleratorOfferingsInput, optFns ...func(*Options)) (*DescribeAcceleratorOfferingsOutput, error) {
	if params == nil {
		params = &DescribeAcceleratorOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAcceleratorOfferings", params, optFns, c.addOperationDescribeAcceleratorOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAcceleratorOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAcceleratorOfferingsInput struct {

	// The location type that you want to describe accelerator type offerings for. It
	// can assume the following values: region: will return the accelerator type
	// offering at the regional level. availability-zone: will return the accelerator
	// type offering at the availability zone level. availability-zone-id: will return
	// the accelerator type offering at the availability zone level returning the
	// availability zone id.
	//
	// This member is required.
	LocationType types.LocationType

	// The list of accelerator types to describe.
	AcceleratorTypes []string

	noSmithyDocumentSerde
}

type DescribeAcceleratorOfferingsOutput struct {

	// The list of accelerator type offerings for a specific location.
	AcceleratorTypeOfferings []types.AcceleratorTypeOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAcceleratorOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAcceleratorOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAcceleratorOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAcceleratorOfferings"); err != nil {
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
	if err = addOpDescribeAcceleratorOfferingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAcceleratorOfferings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAcceleratorOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAcceleratorOfferings",
	}
}
