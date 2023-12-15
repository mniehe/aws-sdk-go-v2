// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Begins setup of time-based one-time password (TOTP) multi-factor authentication
// (MFA) for a user, with a unique private key that Amazon Cognito generates and
// returns in the API response. You can authorize an AssociateSoftwareToken
// request with either the user's access token, or a session string from a
// challenge response that you received from Amazon Cognito. Amazon Cognito
// disassociates an existing software token when you verify the new token in a
// VerifySoftwareToken (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerifySoftwareToken.html)
// API request. If you don't verify the software token and your user pool doesn't
// require MFA, the user can then authenticate with user name and password
// credentials alone. If your user pool requires TOTP MFA, Amazon Cognito generates
// an MFA_SETUP or SOFTWARE_TOKEN_SETUP challenge each time your user signs.
// Complete setup with AssociateSoftwareToken and VerifySoftwareToken . After you
// set up software token MFA for your user, Amazon Cognito generates a
// SOFTWARE_TOKEN_MFA challenge when they authenticate. Respond to this challenge
// with your user's TOTP. Amazon Cognito doesn't evaluate Identity and Access
// Management (IAM) policies in requests for this API operation. For this
// operation, you can't use IAM credentials to authorize requests, and you can't
// grant IAM permissions in policies. For more information about authorization
// models in Amazon Cognito, see Using the Amazon Cognito native and OIDC APIs (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
// .
func (c *Client) AssociateSoftwareToken(ctx context.Context, params *AssociateSoftwareTokenInput, optFns ...func(*Options)) (*AssociateSoftwareTokenOutput, error) {
	if params == nil {
		params = &AssociateSoftwareTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateSoftwareToken", params, optFns, c.addOperationAssociateSoftwareTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateSoftwareTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateSoftwareTokenInput struct {

	// A valid access token that Amazon Cognito issued to the user whose software
	// token you want to generate.
	AccessToken *string

	// The session that should be passed both ways in challenge-response calls to the
	// service. This allows authentication of the user as part of the MFA setup
	// process.
	Session *string

	noSmithyDocumentSerde
}

type AssociateSoftwareTokenOutput struct {

	// A unique generated shared secret code that is used in the TOTP algorithm to
	// generate a one-time code.
	SecretCode *string

	// The session that should be passed both ways in challenge-response calls to the
	// service. This allows authentication of the user as part of the MFA setup
	// process.
	Session *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateSoftwareTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateSoftwareToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateSoftwareToken{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateSoftwareToken"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateSoftwareToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateSoftwareToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateSoftwareToken",
	}
}
