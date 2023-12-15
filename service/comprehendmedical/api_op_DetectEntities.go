// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The DetectEntities operation is deprecated. You should use the DetectEntitiesV2
// operation instead. Inspects the clinical text for a variety of medical entities
// and returns specific information about them such as entity category, location,
// and confidence score on that information.
//
// Deprecated: This operation is deprecated, use DetectEntitiesV2 instead.
func (c *Client) DetectEntities(ctx context.Context, params *DetectEntitiesInput, optFns ...func(*Options)) (*DetectEntitiesOutput, error) {
	if params == nil {
		params = &DetectEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectEntities", params, optFns, c.addOperationDetectEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectEntitiesInput struct {

	// A UTF-8 text string containing the clinical content being examined for entities.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

type DetectEntitiesOutput struct {

	// The collection of medical entities extracted from the input text and their
	// associated information. For each entity, the response provides the entity text,
	// the entity category, where the entity text begins and ends, and the level of
	// confidence that Amazon Comprehend Medical has in the detection and analysis.
	// Attributes and traits of the entity are also returned.
	//
	// This member is required.
	Entities []types.Entity

	// The version of the model used to analyze the documents. The version number
	// looks like X.X.X. You can use this information to track the model used for a
	// particular batch of documents.
	//
	// This member is required.
	ModelVersion *string

	// If the result of the previous request to DetectEntities was truncated, include
	// the PaginationToken to fetch the next page of entities.
	PaginationToken *string

	// Attributes extracted from the input text that we were unable to relate to an
	// entity.
	UnmappedAttributes []types.UnmappedAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectEntities{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DetectEntities"); err != nil {
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
	if err = addOpDetectEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectEntities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DetectEntities",
	}
}
