// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a node.
func (c *Client) DescribeNode(ctx context.Context, params *DescribeNodeInput, optFns ...func(*Options)) (*DescribeNodeOutput, error) {
	if params == nil {
		params = &DescribeNodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNode", params, optFns, c.addOperationDescribeNodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNodeInput struct {

	// The node's ID.
	//
	// This member is required.
	NodeId *string

	// The account ID of the node's owner.
	OwnerAccount *string

	noSmithyDocumentSerde
}

type DescribeNodeOutput struct {

	// The node's category.
	//
	// This member is required.
	Category types.NodeCategory

	// When the node was created.
	//
	// This member is required.
	CreatedTime *time.Time

	// The node's description.
	//
	// This member is required.
	Description *string

	// When the node was updated.
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// The node's name.
	//
	// This member is required.
	Name *string

	// The node's ID.
	//
	// This member is required.
	NodeId *string

	// The node's interface.
	//
	// This member is required.
	NodeInterface *types.NodeInterface

	// The account ID of the node's owner.
	//
	// This member is required.
	OwnerAccount *string

	// The node's package ID.
	//
	// This member is required.
	PackageId *string

	// The node's package name.
	//
	// This member is required.
	PackageName *string

	// The node's package version.
	//
	// This member is required.
	PackageVersion *string

	// The node's patch version.
	//
	// This member is required.
	PatchVersion *string

	// The node's asset name.
	AssetName *string

	// The node's ARN.
	PackageArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeNode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeNode{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeNode"); err != nil {
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
	if err = addOpDescribeNodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeNode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeNode",
	}
}
