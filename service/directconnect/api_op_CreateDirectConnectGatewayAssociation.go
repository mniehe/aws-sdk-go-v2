// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an association between a Direct Connect gateway and a virtual private
// gateway. The virtual private gateway must be attached to a VPC and must not be
// associated with another Direct Connect gateway.
func (c *Client) CreateDirectConnectGatewayAssociation(ctx context.Context, params *CreateDirectConnectGatewayAssociationInput, optFns ...func(*Options)) (*CreateDirectConnectGatewayAssociationOutput, error) {
	if params == nil {
		params = &CreateDirectConnectGatewayAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDirectConnectGatewayAssociation", params, optFns, c.addOperationCreateDirectConnectGatewayAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDirectConnectGatewayAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDirectConnectGatewayAssociationInput struct {

	// The ID of the Direct Connect gateway.
	//
	// This member is required.
	DirectConnectGatewayId *string

	// The Amazon VPC prefixes to advertise to the Direct Connect gateway This
	// parameter is required when you create an association to a transit gateway. For
	// information about how to set the prefixes, see Allowed Prefixes (https://docs.aws.amazon.com/directconnect/latest/UserGuide/multi-account-associate-vgw.html#allowed-prefixes)
	// in the Direct Connect User Guide.
	AddAllowedPrefixesToDirectConnectGateway []types.RouteFilterPrefix

	// The ID of the virtual private gateway or transit gateway.
	GatewayId *string

	// The ID of the virtual private gateway.
	VirtualGatewayId *string

	noSmithyDocumentSerde
}

type CreateDirectConnectGatewayAssociationOutput struct {

	// The association to be created.
	DirectConnectGatewayAssociation *types.DirectConnectGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDirectConnectGatewayAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDirectConnectGatewayAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDirectConnectGatewayAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDirectConnectGatewayAssociation"); err != nil {
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
	if err = addOpCreateDirectConnectGatewayAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDirectConnectGatewayAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDirectConnectGatewayAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDirectConnectGatewayAssociation",
	}
}
