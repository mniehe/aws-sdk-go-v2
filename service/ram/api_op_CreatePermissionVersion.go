// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new version of the specified customer managed permission. The new
// version is automatically set as the default version of the customer managed
// permission. New resource shares automatically use the default permission.
// Existing resource shares continue to use their original permission versions, but
// you can use ReplacePermissionAssociations to update them. If the specified
// customer managed permission already has the maximum of 5 versions, then you must
// delete one of the existing versions before you can create a new one.
func (c *Client) CreatePermissionVersion(ctx context.Context, params *CreatePermissionVersionInput, optFns ...func(*Options)) (*CreatePermissionVersionOutput, error) {
	if params == nil {
		params = &CreatePermissionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePermissionVersion", params, optFns, c.addOperationCreatePermissionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePermissionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePermissionVersionInput struct {

	// Specifies the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the customer managed permission you're creating a new version for.
	//
	// This member is required.
	PermissionArn *string

	// A string in JSON format string that contains the following elements of a
	// resource-based policy:
	//   - Effect: must be set to ALLOW .
	//   - Action: specifies the actions that are allowed by this customer managed
	//   permission. The list must contain only actions that are supported by the
	//   specified resource type. For a list of all actions supported by each resource
	//   type, see Actions, resources, and condition keys for Amazon Web Services
	//   services (https://docs.aws.amazon.com/service-authorization/latest/reference/reference_policies_actions-resources-contextkeys.html)
	//   in the Identity and Access Management User Guide.
	//   - Condition: (optional) specifies conditional parameters that must evaluate
	//   to true when a user attempts an action for that action to be allowed. For more
	//   information about the Condition element, see IAM policies: Condition element (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html)
	//   in the Identity and Access Management User Guide.
	// This template can't include either the Resource or Principal elements. Those
	// are both filled in by RAM when it instantiates the resource-based policy on each
	// resource shared using this managed permission. The Resource comes from the ARN
	// of the specific resource that you are sharing. The Principal comes from the
	// list of identities added to the resource share.
	//
	// This member is required.
	PolicyTemplate *string

	// Specifies a unique, case-sensitive identifier that you provide to ensure the
	// idempotency of the request. This lets you safely retry the request without
	// accidentally performing the same operation a second time. Passing the same value
	// to a later call to an operation requires that you also pass the same value for
	// all other parameters. We recommend that you use a UUID type of value. (https://wikipedia.org/wiki/Universally_unique_identifier)
	// . If you don't provide this value, then Amazon Web Services generates a random
	// one for you. If you retry the operation with the same ClientToken , but with
	// different parameters, the retry fails with an IdempotentParameterMismatch error.
	ClientToken *string

	noSmithyDocumentSerde
}

type CreatePermissionVersionOutput struct {

	// The idempotency identifier associated with this request. If you want to repeat
	// the same operation in an idempotent manner then you must include this value in
	// the clientToken request parameter of that later call. All other parameters must
	// also have the same values that you used in the first call.
	ClientToken *string

	// Information about a RAM managed permission.
	Permission *types.ResourceSharePermissionDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePermissionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePermissionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePermissionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePermissionVersion"); err != nil {
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
	if err = addOpCreatePermissionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePermissionVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePermissionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePermissionVersion",
	}
}
