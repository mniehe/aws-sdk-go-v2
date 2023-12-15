// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the specified user's password in a user pool as an administrator. Works on
// any user. The password can be temporary or permanent. If it is temporary, the
// user status enters the FORCE_CHANGE_PASSWORD state. When the user next tries to
// sign in, the InitiateAuth/AdminInitiateAuth response will contain the
// NEW_PASSWORD_REQUIRED challenge. If the user doesn't sign in before it expires,
// the user won't be able to sign in, and an administrator must reset their
// password. Once the user has set a new password, or the password is permanent,
// the user status is set to Confirmed . AdminSetUserPassword can set a password
// for the user profile that Amazon Cognito creates for third-party federated
// users. When you set a password, the federated user's status changes from
// EXTERNAL_PROVIDER to CONFIRMED . A user in this state can sign in as a federated
// user, and initiate authentication flows in the API like a linked native user.
// They can also modify their password and attributes in token-authenticated API
// requests like ChangePassword and UpdateUserAttributes . As a best security
// practice and to keep users in sync with your external IdP, don't set passwords
// on federated user profiles. To set up a federated user for native sign-in with a
// linked native user, refer to Linking federated users to an existing user profile (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation-consolidate-users.html)
// . Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) AdminSetUserPassword(ctx context.Context, params *AdminSetUserPasswordInput, optFns ...func(*Options)) (*AdminSetUserPasswordOutput, error) {
	if params == nil {
		params = &AdminSetUserPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminSetUserPassword", params, optFns, c.addOperationAdminSetUserPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminSetUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminSetUserPasswordInput struct {

	// The password for the user.
	//
	// This member is required.
	Password *string

	// The user pool ID for the user pool where you want to set the user's password.
	//
	// This member is required.
	UserPoolId *string

	// The user name of the user whose password you want to set.
	//
	// This member is required.
	Username *string

	// True if the password is permanent, False if it is temporary.
	Permanent bool

	noSmithyDocumentSerde
}

type AdminSetUserPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminSetUserPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminSetUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminSetUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AdminSetUserPassword"); err != nil {
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
	if err = addOpAdminSetUserPasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminSetUserPassword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminSetUserPassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AdminSetUserPassword",
	}
}
