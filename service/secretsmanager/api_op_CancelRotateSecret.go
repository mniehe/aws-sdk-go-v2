// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Turns off automatic rotation, and if a rotation is currently in progress,
// cancels the rotation. If you cancel a rotation in progress, it can leave the
// VersionStage labels in an unexpected state. You might need to remove the staging
// label AWSPENDING from the partially created version. You also need to determine
// whether to roll back to the previous version of the secret by moving the staging
// label AWSCURRENT to the version that has AWSPENDING . To determine which version
// has a specific staging label, call ListSecretVersionIds . Then use
// UpdateSecretVersionStage to change staging labels. For more information, see
// How rotation works (https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_how.html)
// . To turn on automatic rotation again, call RotateSecret . Secrets Manager
// generates a CloudTrail log entry when you call this action. Do not include
// sensitive information in request parameters because it might be logged. For more
// information, see Logging Secrets Manager events with CloudTrail (https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html)
// . Required permissions: secretsmanager:CancelRotateSecret . For more
// information, see IAM policy actions for Secrets Manager (https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions)
// and Authentication and access control in Secrets Manager (https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html)
// .
func (c *Client) CancelRotateSecret(ctx context.Context, params *CancelRotateSecretInput, optFns ...func(*Options)) (*CancelRotateSecretOutput, error) {
	if params == nil {
		params = &CancelRotateSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelRotateSecret", params, optFns, c.addOperationCancelRotateSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelRotateSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelRotateSecretInput struct {

	// The ARN or name of the secret. For an ARN, we recommend that you specify a
	// complete ARN rather than a partial ARN. See Finding a secret from a partial ARN (https://docs.aws.amazon.com/secretsmanager/latest/userguide/troubleshoot.html#ARN_secretnamehyphen)
	// .
	//
	// This member is required.
	SecretId *string

	noSmithyDocumentSerde
}

type CancelRotateSecretOutput struct {

	// The ARN of the secret.
	ARN *string

	// The name of the secret.
	Name *string

	// The unique identifier of the version of the secret created during the rotation.
	// This version might not be complete, and should be evaluated for possible
	// deletion. We recommend that you remove the VersionStage value AWSPENDING from
	// this version so that Secrets Manager can delete it. Failing to clean up a
	// cancelled rotation can block you from starting future rotations.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelRotateSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelRotateSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelRotateSecret{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CancelRotateSecret"); err != nil {
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
	if err = addOpCancelRotateSecretValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelRotateSecret(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelRotateSecret(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CancelRotateSecret",
	}
}
