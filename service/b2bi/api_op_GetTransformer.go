// Code generated by smithy-go-codegen DO NOT EDIT.

package b2bi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/b2bi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the details for the transformer specified by the transformer ID. A
// transformer describes how to process the incoming EDI documents and extract the
// necessary information to the output file.
func (c *Client) GetTransformer(ctx context.Context, params *GetTransformerInput, optFns ...func(*Options)) (*GetTransformerOutput, error) {
	if params == nil {
		params = &GetTransformerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransformer", params, optFns, c.addOperationGetTransformerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransformerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransformerInput struct {

	// Specifies the system-assigned unique identifier for the transformer.
	//
	// This member is required.
	TransformerId *string

	noSmithyDocumentSerde
}

type GetTransformerOutput struct {

	// Returns a timestamp for creation date and time of the transformer.
	//
	// This member is required.
	CreatedAt *time.Time

	// Returns the details for the EDI standard that is being used for the
	// transformer. Currently, only X12 is supported. X12 is a set of standards and
	// corresponding messages that define specific business documents.
	//
	// This member is required.
	EdiType types.EdiType

	// Returns that the currently supported file formats for EDI transformations are
	// JSON and XML .
	//
	// This member is required.
	FileFormat types.FileFormat

	// Returns the name of the mapping template for the transformer. This template is
	// used to convert the input document into the correct set of objects.
	//
	// This member is required.
	MappingTemplate *string

	// Returns the name of the transformer, used to identify it.
	//
	// This member is required.
	Name *string

	// Returns the state of the newly created transformer. The transformer can be
	// either active or inactive . For the transformer to be used in a capability, its
	// status must active .
	//
	// This member is required.
	Status types.TransformerStatus

	// Returns an Amazon Resource Name (ARN) for a specific Amazon Web Services
	// resource, such as a capability, partnership, profile, or transformer.
	//
	// This member is required.
	TransformerArn *string

	// Returns the system-assigned unique identifier for the transformer.
	//
	// This member is required.
	TransformerId *string

	// Returns a timestamp for last time the transformer was modified.
	ModifiedAt *time.Time

	// Returns a sample EDI document that is used by a transformer as a guide for
	// processing the EDI data.
	SampleDocument *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTransformerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetTransformer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetTransformer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTransformer"); err != nil {
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
	if err = addOpGetTransformerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransformer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTransformer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTransformer",
	}
}
