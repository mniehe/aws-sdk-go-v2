// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the Secure Sockets Layer (SSL) certificate for the custom domain for
// your user pool. You can use this operation to provide the Amazon Resource Name
// (ARN) of a new certificate to Amazon Cognito. You can't use it to change the
// domain for a user pool. A custom domain is used to host the Amazon Cognito
// hosted UI, which provides sign-up and sign-in pages for your application. When
// you set up a custom domain, you provide a certificate that you manage with
// Certificate Manager (ACM). When necessary, you can use this operation to change
// the certificate that you applied to your custom domain. Usually, this is
// unnecessary following routine certificate renewal with ACM. When you renew your
// existing certificate in ACM, the ARN for your certificate remains the same, and
// your custom domain uses the new certificate automatically. However, if you
// replace your existing certificate with a new one, ACM gives the new certificate
// a new ARN. To apply the new certificate to your custom domain, you must provide
// this ARN to Amazon Cognito. When you add your new certificate in ACM, you must
// choose US East (N. Virginia) as the Amazon Web Services Region. After you submit
// your request, Amazon Cognito requires up to 1 hour to distribute your new
// certificate to your custom domain. For more information about adding a custom
// domain to your user pool, see Using Your Own Domain for the Hosted UI (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html)
// . Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) UpdateUserPoolDomain(ctx context.Context, params *UpdateUserPoolDomainInput, optFns ...func(*Options)) (*UpdateUserPoolDomainOutput, error) {
	if params == nil {
		params = &UpdateUserPoolDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUserPoolDomain", params, optFns, c.addOperationUpdateUserPoolDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserPoolDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The UpdateUserPoolDomain request input.
type UpdateUserPoolDomainInput struct {

	// The configuration for a custom domain that hosts the sign-up and sign-in pages
	// for your application. Use this object to specify an SSL certificate that is
	// managed by ACM.
	//
	// This member is required.
	CustomDomainConfig *types.CustomDomainConfigType

	// The domain name for the custom domain that hosts the sign-up and sign-in pages
	// for your application. One example might be auth.example.com . This string can
	// include only lowercase letters, numbers, and hyphens. Don't use a hyphen for the
	// first or last character. Use periods to separate subdomain names.
	//
	// This member is required.
	Domain *string

	// The ID of the user pool that is associated with the custom domain whose
	// certificate you're updating.
	//
	// This member is required.
	UserPoolId *string

	noSmithyDocumentSerde
}

// The UpdateUserPoolDomain response output.
type UpdateUserPoolDomainOutput struct {

	// The Amazon CloudFront endpoint that Amazon Cognito set up when you added the
	// custom domain to your user pool.
	CloudFrontDomain *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUserPoolDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUserPoolDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUserPoolDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateUserPoolDomain"); err != nil {
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
	if err = addOpUpdateUserPoolDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserPoolDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUserPoolDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateUserPoolDomain",
	}
}
