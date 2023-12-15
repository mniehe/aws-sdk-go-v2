// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the VPCs that were created by other accounts and that can be
// associated with a specified hosted zone because you've submitted one or more
// CreateVPCAssociationAuthorization requests. The response includes a VPCs
// element with a VPC child element for each VPC that can be associated with the
// hosted zone.
func (c *Client) ListVPCAssociationAuthorizations(ctx context.Context, params *ListVPCAssociationAuthorizationsInput, optFns ...func(*Options)) (*ListVPCAssociationAuthorizationsOutput, error) {
	if params == nil {
		params = &ListVPCAssociationAuthorizationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVPCAssociationAuthorizations", params, optFns, c.addOperationListVPCAssociationAuthorizationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVPCAssociationAuthorizationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about that can be associated with your
// hosted zone.
type ListVPCAssociationAuthorizationsInput struct {

	// The ID of the hosted zone for which you want a list of VPCs that can be
	// associated with the hosted zone.
	//
	// This member is required.
	HostedZoneId *string

	// Optional: An integer that specifies the maximum number of VPCs that you want
	// Amazon Route 53 to return. If you don't specify a value for MaxResults , Route
	// 53 returns up to 50 VPCs per page.
	MaxResults *int32

	// Optional: If a response includes a NextToken element, there are more VPCs that
	// can be associated with the specified hosted zone. To get the next page of
	// results, submit another request, and include the value of NextToken from the
	// response in the nexttoken parameter in another ListVPCAssociationAuthorizations
	// request.
	NextToken *string

	noSmithyDocumentSerde
}

// A complex type that contains the response information for the request.
type ListVPCAssociationAuthorizationsOutput struct {

	// The ID of the hosted zone that you can associate the listed VPCs with.
	//
	// This member is required.
	HostedZoneId *string

	// The list of VPCs that are authorized to be associated with the specified hosted
	// zone.
	//
	// This member is required.
	VPCs []types.VPC

	// When the response includes a NextToken element, there are more VPCs that can be
	// associated with the specified hosted zone. To get the next page of VPCs, submit
	// another ListVPCAssociationAuthorizations request, and include the value of the
	// NextToken element from the response in the nexttoken request parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVPCAssociationAuthorizationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListVPCAssociationAuthorizations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListVPCAssociationAuthorizations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVPCAssociationAuthorizations"); err != nil {
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
	if err = addOpListVPCAssociationAuthorizationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVPCAssociationAuthorizations(options.Region), middleware.Before); err != nil {
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
	if err = addSanitizeURLMiddleware(stack); err != nil {
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

func newServiceMetadataMiddleware_opListVPCAssociationAuthorizations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVPCAssociationAuthorizations",
	}
}
