// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates the VPC from the service network. You can't disassociate the VPC
// if there is a create or update association in progress.
func (c *Client) DeleteServiceNetworkVpcAssociation(ctx context.Context, params *DeleteServiceNetworkVpcAssociationInput, optFns ...func(*Options)) (*DeleteServiceNetworkVpcAssociationOutput, error) {
	if params == nil {
		params = &DeleteServiceNetworkVpcAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteServiceNetworkVpcAssociation", params, optFns, c.addOperationDeleteServiceNetworkVpcAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteServiceNetworkVpcAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteServiceNetworkVpcAssociationInput struct {

	// The ID or Amazon Resource Name (ARN) of the association.
	//
	// This member is required.
	ServiceNetworkVpcAssociationIdentifier *string

	noSmithyDocumentSerde
}

type DeleteServiceNetworkVpcAssociationOutput struct {

	// The Amazon Resource Name (ARN) of the association.
	Arn *string

	// The ID of the association.
	Id *string

	// The status. You can retry the operation if the status is DELETE_FAILED .
	// However, if you retry it when the status is DELETE_IN_PROGRESS , there is no
	// change in the status.
	Status types.ServiceNetworkVpcAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteServiceNetworkVpcAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteServiceNetworkVpcAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteServiceNetworkVpcAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteServiceNetworkVpcAssociation"); err != nil {
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
	if err = addOpDeleteServiceNetworkVpcAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteServiceNetworkVpcAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteServiceNetworkVpcAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteServiceNetworkVpcAssociation",
	}
}
