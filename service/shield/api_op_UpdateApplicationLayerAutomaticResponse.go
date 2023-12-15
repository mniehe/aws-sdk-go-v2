// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/shield/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Shield Advanced automatic application layer DDoS mitigation
// configuration for the specified resource.
func (c *Client) UpdateApplicationLayerAutomaticResponse(ctx context.Context, params *UpdateApplicationLayerAutomaticResponseInput, optFns ...func(*Options)) (*UpdateApplicationLayerAutomaticResponseOutput, error) {
	if params == nil {
		params = &UpdateApplicationLayerAutomaticResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApplicationLayerAutomaticResponse", params, optFns, c.addOperationUpdateApplicationLayerAutomaticResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApplicationLayerAutomaticResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApplicationLayerAutomaticResponseInput struct {

	// Specifies the action setting that Shield Advanced should use in the WAF rules
	// that it creates on behalf of the protected resource in response to DDoS attacks.
	// You specify this as part of the configuration for the automatic application
	// layer DDoS mitigation feature, when you enable or update automatic mitigation.
	// Shield Advanced creates the WAF rules in a Shield Advanced-managed rule group,
	// inside the web ACL that you have associated with the resource.
	//
	// This member is required.
	Action *types.ResponseAction

	// The ARN (Amazon Resource Name) of the resource.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type UpdateApplicationLayerAutomaticResponseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateApplicationLayerAutomaticResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateApplicationLayerAutomaticResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateApplicationLayerAutomaticResponse{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateApplicationLayerAutomaticResponse"); err != nil {
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
	if err = addOpUpdateApplicationLayerAutomaticResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApplicationLayerAutomaticResponse(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateApplicationLayerAutomaticResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateApplicationLayerAutomaticResponse",
	}
}
