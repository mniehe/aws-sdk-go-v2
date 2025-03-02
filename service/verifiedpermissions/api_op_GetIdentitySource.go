// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the details about the specified identity source.
func (c *Client) GetIdentitySource(ctx context.Context, params *GetIdentitySourceInput, optFns ...func(*Options)) (*GetIdentitySourceOutput, error) {
	if params == nil {
		params = &GetIdentitySourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdentitySource", params, optFns, c.addOperationGetIdentitySourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdentitySourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIdentitySourceInput struct {

	// Specifies the ID of the identity source you want information about.
	//
	// This member is required.
	IdentitySourceId *string

	// Specifies the ID of the policy store that contains the identity source you want
	// information about.
	//
	// This member is required.
	PolicyStoreId *string

	noSmithyDocumentSerde
}

type GetIdentitySourceOutput struct {

	// The date and time that the identity source was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// A structure that describes the configuration of the identity source.
	//
	// This member is required.
	Details *types.IdentitySourceDetails

	// The ID of the identity source.
	//
	// This member is required.
	IdentitySourceId *string

	// The date and time that the identity source was most recently updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The ID of the policy store that contains the identity source.
	//
	// This member is required.
	PolicyStoreId *string

	// The data type of principals generated for identities authenticated by this
	// identity source.
	//
	// This member is required.
	PrincipalEntityType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIdentitySourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetIdentitySource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetIdentitySource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetIdentitySource"); err != nil {
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
	if err = addOpGetIdentitySourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentitySource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIdentitySource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetIdentitySource",
	}
}
