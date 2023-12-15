// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the specified WorkSpace bundle. For more information about creating
// WorkSpace bundles, see Create a Custom WorkSpaces Image and Bundle (https://docs.aws.amazon.com/workspaces/latest/adminguide/create-custom-bundle.html)
// .
func (c *Client) CreateWorkspaceBundle(ctx context.Context, params *CreateWorkspaceBundleInput, optFns ...func(*Options)) (*CreateWorkspaceBundleOutput, error) {
	if params == nil {
		params = &CreateWorkspaceBundleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkspaceBundle", params, optFns, c.addOperationCreateWorkspaceBundleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkspaceBundleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkspaceBundleInput struct {

	// The description of the bundle.
	//
	// This member is required.
	BundleDescription *string

	// The name of the bundle.
	//
	// This member is required.
	BundleName *string

	// Describes the compute type of the bundle.
	//
	// This member is required.
	ComputeType *types.ComputeType

	// The identifier of the image that is used to create the bundle.
	//
	// This member is required.
	ImageId *string

	// Describes the user volume for a WorkSpace bundle.
	//
	// This member is required.
	UserStorage *types.UserStorage

	// Describes the root volume for a WorkSpace bundle.
	RootStorage *types.RootStorage

	// The tags associated with the bundle. To add tags at the same time when you're
	// creating the bundle, you must create an IAM policy that grants your IAM user
	// permissions to use workspaces:CreateTags .
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateWorkspaceBundleOutput struct {

	// Describes a WorkSpace bundle.
	WorkspaceBundle *types.WorkspaceBundle

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkspaceBundleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWorkspaceBundle{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWorkspaceBundle{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWorkspaceBundle"); err != nil {
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
	if err = addOpCreateWorkspaceBundleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkspaceBundle(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkspaceBundle(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWorkspaceBundle",
	}
}
