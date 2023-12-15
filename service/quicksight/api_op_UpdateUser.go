// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Amazon QuickSight user.
func (c *Client) UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) {
	if params == nil {
		params = &UpdateUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUser", params, optFns, c.addOperationUpdateUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserInput struct {

	// The ID for the Amazon Web Services account that the user is in. Currently, you
	// use the ID for the Amazon Web Services account that contains your Amazon
	// QuickSight account.
	//
	// This member is required.
	AwsAccountId *string

	// The email address of the user that you want to update.
	//
	// This member is required.
	Email *string

	// The namespace. Currently, you should set this to default .
	//
	// This member is required.
	Namespace *string

	// The Amazon QuickSight role of the user. The role can be one of the following
	// default security cohorts:
	//   - READER : A user who has read-only access to dashboards.
	//   - AUTHOR : A user who can create data sources, datasets, analyses, and
	//   dashboards.
	//   - ADMIN : A user who is an author, who can also manage Amazon QuickSight
	//   settings.
	// The name of the Amazon QuickSight role is invisible to the user except for the
	// console screens dealing with permissions.
	//
	// This member is required.
	Role types.UserRole

	// The Amazon QuickSight user name that you want to update.
	//
	// This member is required.
	UserName *string

	// The URL of the custom OpenID Connect (OIDC) provider that provides identity to
	// let a user federate into Amazon QuickSight with an associated Identity and
	// Access Management(IAM) role. This parameter should only be used when
	// ExternalLoginFederationProviderType parameter is set to CUSTOM_OIDC .
	CustomFederationProviderUrl *string

	// (Enterprise edition only) The name of the custom permissions profile that you
	// want to assign to this user. Customized permissions allows you to control a
	// user's access by restricting access the following operations:
	//   - Create and update data sources
	//   - Create and update datasets
	//   - Create and update email reports
	//   - Subscribe to email reports
	// A set of custom permissions includes any combination of these restrictions.
	// Currently, you need to create the profile names for custom permission sets by
	// using the Amazon QuickSight console. Then, you use the RegisterUser API
	// operation to assign the named set of permissions to a Amazon QuickSight user.
	// Amazon QuickSight custom permissions are applied through IAM policies.
	// Therefore, they override the permissions typically granted by assigning Amazon
	// QuickSight users to one of the default security cohorts in Amazon QuickSight
	// (admin, author, reader). This feature is available only to Amazon QuickSight
	// Enterprise edition subscriptions.
	CustomPermissionsName *string

	// The type of supported external login provider that provides identity to let a
	// user federate into Amazon QuickSight with an associated Identity and Access
	// Management(IAM) role. The type of supported external login provider can be one
	// of the following.
	//   - COGNITO : Amazon Cognito. The provider URL is
	//   cognito-identity.amazonaws.com. When choosing the COGNITO provider type, don’t
	//   use the "CustomFederationProviderUrl" parameter which is only needed when the
	//   external provider is custom.
	//   - CUSTOM_OIDC : Custom OpenID Connect (OIDC) provider. When choosing
	//   CUSTOM_OIDC type, use the CustomFederationProviderUrl parameter to provide the
	//   custom OIDC provider URL.
	//   - NONE : This clears all the previously saved external login information for a
	//   user. Use the DescribeUser (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeUser.html)
	//   API operation to check the external login information.
	ExternalLoginFederationProviderType *string

	// The identity ID for a user in the external login provider.
	ExternalLoginId *string

	// A flag that you use to indicate that you want to remove all custom permissions
	// from this user. Using this parameter resets the user to the state it was in
	// before a custom permissions profile was applied. This parameter defaults to NULL
	// and it doesn't accept any other value.
	UnapplyCustomPermissions bool

	noSmithyDocumentSerde
}

type UpdateUserOutput struct {

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The Amazon QuickSight user.
	User *types.User

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateUser"); err != nil {
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
	if err = addOpUpdateUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateUser",
	}
}
