// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the SchemaMapping of a given name.
func (c *Client) GetSchemaMapping(ctx context.Context, params *GetSchemaMappingInput, optFns ...func(*Options)) (*GetSchemaMappingOutput, error) {
	if params == nil {
		params = &GetSchemaMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSchemaMapping", params, optFns, c.addOperationGetSchemaMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSchemaMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSchemaMappingInput struct {

	// The name of the schema to be retrieved.
	//
	// This member is required.
	SchemaName *string

	noSmithyDocumentSerde
}

type GetSchemaMappingOutput struct {

	// The timestamp of when the SchemaMapping was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// Specifies whether the schema mapping has been applied to a workflow.
	//
	// This member is required.
	HasWorkflows *bool

	// A list of MappedInputFields . Each MappedInputField corresponds to a column the
	// source data table, and contains column name plus additional information Venice
	// uses for matching.
	//
	// This member is required.
	MappedInputFields []types.SchemaInputAttribute

	// The ARN (Amazon Resource Name) that Entity Resolution generated for the
	// SchemaMapping.
	//
	// This member is required.
	SchemaArn *string

	// The name of the schema.
	//
	// This member is required.
	SchemaName *string

	// The timestamp of when the SchemaMapping was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// A description of the schema.
	Description *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSchemaMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSchemaMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSchemaMapping{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSchemaMapping"); err != nil {
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
	if err = addOpGetSchemaMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSchemaMapping(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSchemaMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSchemaMapping",
	}
}
