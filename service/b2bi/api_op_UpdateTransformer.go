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

// Updates the specified parameters for a transformer. A transformer describes how
// to process the incoming EDI documents and extract the necessary information to
// the output file.
func (c *Client) UpdateTransformer(ctx context.Context, params *UpdateTransformerInput, optFns ...func(*Options)) (*UpdateTransformerOutput, error) {
	if params == nil {
		params = &UpdateTransformerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTransformer", params, optFns, c.addOperationUpdateTransformerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTransformerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTransformerInput struct {

	// Specifies the system-assigned unique identifier for the transformer.
	//
	// This member is required.
	TransformerId *string

	// Specifies the details for the EDI standard that is being used for the
	// transformer. Currently, only X12 is supported. X12 is a set of standards and
	// corresponding messages that define specific business documents.
	EdiType types.EdiType

	// Specifies that the currently supported file formats for EDI transformations are
	// JSON and XML .
	FileFormat types.FileFormat

	// Specifies the name of the mapping template for the transformer. This template
	// is used to convert the input document into the correct set of objects.
	MappingTemplate *string

	// Specify a new name for the transformer, if you want to update it.
	Name *string

	// Specifies a sample EDI document that is used by a transformer as a guide for
	// processing the EDI data.
	SampleDocument *string

	// Specifies the transformer's status. You can update the state of the
	// transformer, from active to inactive , or inactive to active .
	Status types.TransformerStatus

	noSmithyDocumentSerde
}

type UpdateTransformerOutput struct {

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

	// Returns a timestamp for last time the transformer was modified.
	//
	// This member is required.
	ModifiedAt *time.Time

	// Returns the name of the transformer.
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

	// Returns a sample EDI document that is used by a transformer as a guide for
	// processing the EDI data.
	SampleDocument *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTransformerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateTransformer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateTransformer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTransformer"); err != nil {
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
	if err = addOpUpdateTransformerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTransformer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTransformer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTransformer",
	}
}
