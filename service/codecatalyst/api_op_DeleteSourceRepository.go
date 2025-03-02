// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a source repository in Amazon CodeCatalyst. You cannot use this API to
// delete a linked repository. It can only be used to delete a Amazon CodeCatalyst
// source repository.
func (c *Client) DeleteSourceRepository(ctx context.Context, params *DeleteSourceRepositoryInput, optFns ...func(*Options)) (*DeleteSourceRepositoryOutput, error) {
	if params == nil {
		params = &DeleteSourceRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSourceRepository", params, optFns, c.addOperationDeleteSourceRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSourceRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSourceRepositoryInput struct {

	// The name of the source repository.
	//
	// This member is required.
	Name *string

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	noSmithyDocumentSerde
}

type DeleteSourceRepositoryOutput struct {

	// The name of the repository.
	//
	// This member is required.
	Name *string

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSourceRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSourceRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSourceRepository{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteSourceRepository"); err != nil {
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
	if err = addOpDeleteSourceRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSourceRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSourceRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteSourceRepository",
	}
}
