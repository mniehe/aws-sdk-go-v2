// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified sending authorization policy for the given identity (an
// email address or a domain). This operation returns successfully even if a policy
// with the specified name does not exist. This operation is for the identity owner
// only. If you have not verified the identity, it returns an error. Sending
// authorization is a feature that enables an identity owner to authorize other
// senders to use its identities. For information about using sending
// authorization, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html)
// . You can execute this operation no more than once per second.
func (c *Client) DeleteIdentityPolicy(ctx context.Context, params *DeleteIdentityPolicyInput, optFns ...func(*Options)) (*DeleteIdentityPolicyOutput, error) {
	if params == nil {
		params = &DeleteIdentityPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIdentityPolicy", params, optFns, c.addOperationDeleteIdentityPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIdentityPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to delete a sending authorization policy for an identity.
// Sending authorization is an Amazon SES feature that enables you to authorize
// other senders to use your identities. For information, see the Amazon SES
// Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html)
// .
type DeleteIdentityPolicyInput struct {

	// The identity that is associated with the policy to delete. You can specify the
	// identity by using its name or by using its Amazon Resource Name (ARN). Examples:
	// user@example.com , example.com ,
	// arn:aws:ses:us-east-1:123456789012:identity/example.com . To successfully call
	// this operation, you must own the identity.
	//
	// This member is required.
	Identity *string

	// The name of the policy to be deleted.
	//
	// This member is required.
	PolicyName *string

	noSmithyDocumentSerde
}

// An empty element returned on a successful request.
type DeleteIdentityPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteIdentityPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteIdentityPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteIdentityPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteIdentityPolicy"); err != nil {
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
	if err = addOpDeleteIdentityPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIdentityPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteIdentityPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteIdentityPolicy",
	}
}
