// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a VPC connection.
func (c *Client) DeleteVPCConnection(ctx context.Context, params *DeleteVPCConnectionInput, optFns ...func(*Options)) (*DeleteVPCConnectionOutput, error) {
	if params == nil {
		params = &DeleteVPCConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVPCConnection", params, optFns, c.addOperationDeleteVPCConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVPCConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVPCConnectionInput struct {

	// The Amazon Web Services account ID of the account where you want to delete a
	// VPC connection.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the VPC connection that you're creating. This ID is a unique
	// identifier for each Amazon Web Services Region in an Amazon Web Services
	// account.
	//
	// This member is required.
	VPCConnectionId *string

	noSmithyDocumentSerde
}

type DeleteVPCConnectionOutput struct {

	// The Amazon Resource Name (ARN) of the deleted VPC connection.
	Arn *string

	// The availability status of the VPC connection.
	AvailabilityStatus types.VPCConnectionAvailabilityStatus

	// The deletion status of the VPC connection.
	DeletionStatus types.VPCConnectionResourceStatus

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The ID of the VPC connection that you're creating. This ID is a unique
	// identifier for each Amazon Web Services Region in an Amazon Web Services
	// account.
	VPCConnectionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteVPCConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteVPCConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteVPCConnection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteVPCConnection"); err != nil {
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
	if err = addOpDeleteVPCConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVPCConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVPCConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteVPCConnection",
	}
}
