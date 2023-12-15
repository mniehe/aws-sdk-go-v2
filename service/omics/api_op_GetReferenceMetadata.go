// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a genome reference's metadata.
func (c *Client) GetReferenceMetadata(ctx context.Context, params *GetReferenceMetadataInput, optFns ...func(*Options)) (*GetReferenceMetadataOutput, error) {
	if params == nil {
		params = &GetReferenceMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReferenceMetadata", params, optFns, c.addOperationGetReferenceMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReferenceMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReferenceMetadataInput struct {

	// The reference's ID.
	//
	// This member is required.
	Id *string

	// The reference's reference store ID.
	//
	// This member is required.
	ReferenceStoreId *string

	noSmithyDocumentSerde
}

type GetReferenceMetadataOutput struct {

	// The reference's ARN.
	//
	// This member is required.
	Arn *string

	// When the reference was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The reference's ID.
	//
	// This member is required.
	Id *string

	// The reference's MD5 checksum.
	//
	// This member is required.
	Md5 *string

	// The reference's reference store ID.
	//
	// This member is required.
	ReferenceStoreId *string

	// When the reference was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The reference's description.
	Description *string

	// The reference's files.
	Files *types.ReferenceFiles

	// The reference's name.
	Name *string

	// The reference's status.
	Status types.ReferenceStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReferenceMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetReferenceMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetReferenceMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetReferenceMetadata"); err != nil {
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
	if err = addEndpointPrefix_opGetReferenceMetadataMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetReferenceMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReferenceMetadata(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetReferenceMetadataMiddleware struct {
}

func (*endpointPrefix_opGetReferenceMetadataMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetReferenceMetadataMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetReferenceMetadataMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetReferenceMetadataMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetReferenceMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetReferenceMetadata",
	}
}
