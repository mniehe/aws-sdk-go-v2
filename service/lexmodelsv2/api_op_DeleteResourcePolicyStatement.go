// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a policy statement from a resource policy. If you delete the last
// statement from a policy, the policy is deleted. If you specify a statement ID
// that doesn't exist in the policy, or if the bot or bot alias doesn't have a
// policy attached, Amazon Lex returns an exception.
func (c *Client) DeleteResourcePolicyStatement(ctx context.Context, params *DeleteResourcePolicyStatementInput, optFns ...func(*Options)) (*DeleteResourcePolicyStatementOutput, error) {
	if params == nil {
		params = &DeleteResourcePolicyStatementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteResourcePolicyStatement", params, optFns, c.addOperationDeleteResourcePolicyStatementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteResourcePolicyStatementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteResourcePolicyStatementInput struct {

	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy
	// is attached to.
	//
	// This member is required.
	ResourceArn *string

	// The name of the statement (SID) to delete from the policy.
	//
	// This member is required.
	StatementId *string

	// The identifier of the revision of the policy to delete the statement from. If
	// this revision ID doesn't match the current revision ID, Amazon Lex throws an
	// exception. If you don't specify a revision, Amazon Lex removes the current
	// contents of the statement.
	ExpectedRevisionId *string

	noSmithyDocumentSerde
}

type DeleteResourcePolicyStatementOutput struct {

	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy
	// statement was removed from.
	ResourceArn *string

	// The current revision of the resource policy. Use the revision ID to make sure
	// that you are updating the most current version of a resource policy when you add
	// a policy statement to a resource, delete a resource, or update a resource.
	RevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteResourcePolicyStatementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteResourcePolicyStatement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteResourcePolicyStatement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteResourcePolicyStatement"); err != nil {
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
	if err = addOpDeleteResourcePolicyStatementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteResourcePolicyStatement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteResourcePolicyStatement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteResourcePolicyStatement",
	}
}
