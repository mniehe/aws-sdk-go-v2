// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

func (c *Client) OperationWithRequiredMembers(ctx context.Context, params *OperationWithRequiredMembersInput, optFns ...func(*Options)) (*OperationWithRequiredMembersOutput, error) {
	if params == nil {
		params = &OperationWithRequiredMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "OperationWithRequiredMembers", params, optFns, c.addOperationOperationWithRequiredMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*OperationWithRequiredMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type OperationWithRequiredMembersInput struct {
	noSmithyDocumentSerde
}

type OperationWithRequiredMembersOutput struct {

	// This member is required.
	RequiredBlob []byte

	// This member is required.
	RequiredBoolean *bool

	// This member is required.
	RequiredByte *int8

	// This member is required.
	RequiredDouble *float64

	// This member is required.
	RequiredFloat *float32

	// This member is required.
	RequiredInteger *int32

	// This member is required.
	RequiredList []string

	// This member is required.
	RequiredLong *int64

	// This member is required.
	RequiredMap map[string]string

	// This member is required.
	RequiredShort *int16

	// This member is required.
	RequiredString *string

	// This member is required.
	RequiredTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationOperationWithRequiredMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpOperationWithRequiredMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpOperationWithRequiredMembers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "OperationWithRequiredMembers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opOperationWithRequiredMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opOperationWithRequiredMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "OperationWithRequiredMembers",
	}
}
