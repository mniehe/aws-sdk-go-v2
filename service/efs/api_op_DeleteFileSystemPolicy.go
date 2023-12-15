// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the FileSystemPolicy for the specified file system. The default
// FileSystemPolicy goes into effect once the existing policy is deleted. For more
// information about the default file system policy, see Using Resource-based
// Policies with EFS (https://docs.aws.amazon.com/efs/latest/ug/res-based-policies-efs.html)
// . This operation requires permissions for the
// elasticfilesystem:DeleteFileSystemPolicy action.
func (c *Client) DeleteFileSystemPolicy(ctx context.Context, params *DeleteFileSystemPolicyInput, optFns ...func(*Options)) (*DeleteFileSystemPolicyOutput, error) {
	if params == nil {
		params = &DeleteFileSystemPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFileSystemPolicy", params, optFns, c.addOperationDeleteFileSystemPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFileSystemPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFileSystemPolicyInput struct {

	// Specifies the EFS file system for which to delete the FileSystemPolicy .
	//
	// This member is required.
	FileSystemId *string

	noSmithyDocumentSerde
}

type DeleteFileSystemPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFileSystemPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteFileSystemPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteFileSystemPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteFileSystemPolicy"); err != nil {
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
	if err = addOpDeleteFileSystemPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFileSystemPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFileSystemPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteFileSystemPolicy",
	}
}
