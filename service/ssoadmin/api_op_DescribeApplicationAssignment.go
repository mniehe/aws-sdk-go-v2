// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a direct assignment of a user or group to an application. If the user
// doesn’t have a direct assignment to the application, the user may still have
// access to the application through a group. Therefore, don’t use this API to test
// access to an application for a user. Instead use
// ListApplicationAssignmentsForPrincipal .
func (c *Client) DescribeApplicationAssignment(ctx context.Context, params *DescribeApplicationAssignmentInput, optFns ...func(*Options)) (*DescribeApplicationAssignmentOutput, error) {
	if params == nil {
		params = &DescribeApplicationAssignmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApplicationAssignment", params, optFns, c.addOperationDescribeApplicationAssignmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeApplicationAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeApplicationAssignmentInput struct {

	// Specifies the ARN of the application. For more information about ARNs, see
	// Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces in the
	// Amazon Web Services General Reference.
	//
	// This member is required.
	ApplicationArn *string

	// An identifier for an object in IAM Identity Center, such as a user or group.
	// PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). For
	// more information about PrincipalIds in IAM Identity Center, see the IAM
	// Identity Center Identity Store API Reference .
	//
	// This member is required.
	PrincipalId *string

	// The entity type for which the assignment will be created.
	//
	// This member is required.
	PrincipalType types.PrincipalType

	noSmithyDocumentSerde
}

type DescribeApplicationAssignmentOutput struct {

	// Specifies the ARN of the application. For more information about ARNs, see
	// Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces in the
	// Amazon Web Services General Reference.
	ApplicationArn *string

	// An identifier for an object in IAM Identity Center, such as a user or group.
	// PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). For
	// more information about PrincipalIds in IAM Identity Center, see the IAM
	// Identity Center Identity Store API Reference .
	PrincipalId *string

	// The entity type for which the assignment will be created.
	PrincipalType types.PrincipalType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeApplicationAssignmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeApplicationAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeApplicationAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeApplicationAssignment"); err != nil {
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
	if err = addOpDescribeApplicationAssignmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplicationAssignment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeApplicationAssignment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeApplicationAssignment",
	}
}
