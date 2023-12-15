// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a partner integration from a cluster. Data can still flow to the
// cluster until the integration is deleted at the partner's website.
func (c *Client) DeletePartner(ctx context.Context, params *DeletePartnerInput, optFns ...func(*Options)) (*DeletePartnerOutput, error) {
	if params == nil {
		params = &DeletePartnerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePartner", params, optFns, c.addOperationDeletePartnerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePartnerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePartnerInput struct {

	// The Amazon Web Services account ID that owns the cluster.
	//
	// This member is required.
	AccountId *string

	// The cluster identifier of the cluster that receives data from the partner.
	//
	// This member is required.
	ClusterIdentifier *string

	// The name of the database that receives data from the partner.
	//
	// This member is required.
	DatabaseName *string

	// The name of the partner that is authorized to send data.
	//
	// This member is required.
	PartnerName *string

	noSmithyDocumentSerde
}

type DeletePartnerOutput struct {

	// The name of the database that receives data from the partner.
	DatabaseName *string

	// The name of the partner that is authorized to send data.
	PartnerName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePartnerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeletePartner{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeletePartner{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeletePartner"); err != nil {
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
	if err = addOpDeletePartnerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePartner(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePartner(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeletePartner",
	}
}
