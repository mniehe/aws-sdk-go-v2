// Code generated by smithy-go-codegen DO NOT EDIT.

package amp

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/amp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a rule group namespace.
func (c *Client) CreateRuleGroupsNamespace(ctx context.Context, params *CreateRuleGroupsNamespaceInput, optFns ...func(*Options)) (*CreateRuleGroupsNamespaceOutput, error) {
	if params == nil {
		params = &CreateRuleGroupsNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRuleGroupsNamespace", params, optFns, c.addOperationCreateRuleGroupsNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRuleGroupsNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateRuleGroupsNamespace operation.
type CreateRuleGroupsNamespaceInput struct {

	// The namespace data that define the rule groups.
	//
	// This member is required.
	Data []byte

	// The rule groups namespace name.
	//
	// This member is required.
	Name *string

	// The ID of the workspace in which to create the rule group namespace.
	//
	// This member is required.
	WorkspaceId *string

	// Optional, unique, case-sensitive, user-provided identifier to ensure the
	// idempotency of the request.
	ClientToken *string

	// Optional, user-provided tags for this rule groups namespace.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents the output of a CreateRuleGroupsNamespace operation.
type CreateRuleGroupsNamespaceOutput struct {

	// The Amazon Resource Name (ARN) of this rule groups namespace.
	//
	// This member is required.
	Arn *string

	// The rule groups namespace name.
	//
	// This member is required.
	Name *string

	// The status of rule groups namespace.
	//
	// This member is required.
	Status *types.RuleGroupsNamespaceStatus

	// The tags of this rule groups namespace.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRuleGroupsNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRuleGroupsNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRuleGroupsNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRuleGroupsNamespace"); err != nil {
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
	if err = addIdempotencyToken_opCreateRuleGroupsNamespaceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateRuleGroupsNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRuleGroupsNamespace(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateRuleGroupsNamespace struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateRuleGroupsNamespace) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateRuleGroupsNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateRuleGroupsNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateRuleGroupsNamespaceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateRuleGroupsNamespaceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateRuleGroupsNamespace{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateRuleGroupsNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRuleGroupsNamespace",
	}
}
