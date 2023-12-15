// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns detailed information about an extension that has been registered. If
// you specify a VersionId , DescribeType returns information about that specific
// extension version. Otherwise, it returns information about the default extension
// version.
func (c *Client) DescribeType(ctx context.Context, params *DescribeTypeInput, optFns ...func(*Options)) (*DescribeTypeOutput, error) {
	if params == nil {
		params = &DescribeTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeType", params, optFns, c.addOperationDescribeTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTypeInput struct {

	// The Amazon Resource Name (ARN) of the extension. Conditional: You must specify
	// either TypeName and Type , or Arn .
	Arn *string

	// The version number of a public third-party extension.
	PublicVersionNumber *string

	// The publisher ID of the extension publisher. Extensions provided by Amazon Web
	// Services are not assigned a publisher ID.
	PublisherId *string

	// The kind of extension. Conditional: You must specify either TypeName and Type ,
	// or Arn .
	Type types.RegistryType

	// The name of the extension. Conditional: You must specify either TypeName and
	// Type , or Arn .
	TypeName *string

	// The ID of a specific version of the extension. The version ID is the value at
	// the end of the Amazon Resource Name (ARN) assigned to the extension version when
	// it is registered. If you specify a VersionId , DescribeType returns information
	// about that specific extension version. Otherwise, it returns information about
	// the default extension version.
	VersionId *string

	noSmithyDocumentSerde
}

type DescribeTypeOutput struct {

	// The Amazon Resource Name (ARN) of the extension.
	Arn *string

	// Whether CloudFormation automatically updates the extension in this account and
	// Region when a new minor version is published by the extension publisher. Major
	// versions released by the publisher must be manually updated. For more
	// information, see Activating public extensions for use in your account (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public.html#registry-public-enable)
	// in the CloudFormation User Guide.
	AutoUpdate *bool

	// A JSON string that represent the current configuration data for the extension
	// in this account and Region. To set the configuration data for an extension, use
	// SetTypeConfiguration (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html)
	// . For more information, see Configuring extensions at the account level (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-register.html#registry-set-configuration)
	// in the CloudFormation User Guide.
	ConfigurationSchema *string

	// The ID of the default version of the extension. The default version is used
	// when the extension version isn't specified. This applies only to private
	// extensions you have registered in your account. For public extensions, both
	// those provided by Amazon Web Services and published by third parties,
	// CloudFormation returns null . For more information, see RegisterType (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterType.html)
	// . To set the default version of an extension, use SetTypeDefaultVersion .
	DefaultVersionId *string

	// The deprecation status of the extension version. Valid values include:
	//   - LIVE : The extension is activated or registered and can be used in
	//   CloudFormation operations, dependent on its provisioning behavior and visibility
	//   scope.
	//   - DEPRECATED : The extension has been deactivated or deregistered and can no
	//   longer be used in CloudFormation operations.
	// For public third-party extensions, CloudFormation returns null .
	DeprecatedStatus types.DeprecatedStatus

	// The description of the extension.
	Description *string

	// The URL of a page providing detailed documentation for this extension.
	DocumentationUrl *string

	// The Amazon Resource Name (ARN) of the IAM execution role used to register the
	// extension. This applies only to private extensions you have registered in your
	// account. For more information, see RegisterType (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterType.html)
	// . If the registered extension calls any Amazon Web Services APIs, you must
	// create an IAM execution role (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)
	// that includes the necessary permissions to call those Amazon Web Services APIs,
	// and provision that execution role in your account. CloudFormation then assumes
	// that execution role to provide your extension with the appropriate credentials.
	ExecutionRoleArn *string

	// Whether the extension is activated in the account and Region. This only applies
	// to public third-party extensions. For all other extensions, CloudFormation
	// returns null .
	IsActivated *bool

	// Whether the specified extension version is set as the default version. This
	// applies only to private extensions you have registered in your account, and
	// extensions published by Amazon Web Services. For public third-party extensions,
	// whether they are activated in your account, CloudFormation returns null .
	IsDefaultVersion *bool

	// When the specified extension version was registered. This applies only to:
	//   - Private extensions you have registered in your account. For more
	//   information, see RegisterType (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterType.html)
	//   .
	//   - Public extensions you have activated in your account with auto-update
	//   specified. For more information, see ActivateType (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ActivateType.html)
	//   .
	LastUpdated *time.Time

	// The latest version of a public extension that is available for use. This only
	// applies if you specify a public extension, and you don't specify a version. For
	// all other requests, CloudFormation returns null .
	LatestPublicVersion *string

	// Contains logging configuration information for private extensions. This applies
	// only to private extensions you have registered in your account. For public
	// extensions, both those provided by Amazon Web Services and published by third
	// parties, CloudFormation returns null . For more information, see RegisterType (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterType.html)
	// .
	LoggingConfig *types.LoggingConfig

	// For public extensions that have been activated for this account and Region, the
	// Amazon Resource Name (ARN) of the public extension.
	OriginalTypeArn *string

	// For public extensions that have been activated for this account and Region, the
	// type name of the public extension. If you specified a TypeNameAlias when
	// enabling the extension in this account and Region, CloudFormation treats that
	// alias as the extension's type name within the account and Region, not the type
	// name of the public extension. For more information, see Specifying aliases to
	// refer to extensions (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public.html#registry-public-enable-alias)
	// in the CloudFormation User Guide.
	OriginalTypeName *string

	// For resource type extensions, the provisioning behavior of the resource type.
	// CloudFormation determines the provisioning type during registration, based on
	// the types of handlers in the schema handler package submitted. Valid values
	// include:
	//   - FULLY_MUTABLE : The resource type includes an update handler to process
	//   updates to the type during stack update operations.
	//   - IMMUTABLE : The resource type doesn't include an update handler, so the type
	//   can't be updated and must instead be replaced during stack update operations.
	//   - NON_PROVISIONABLE : The resource type doesn't include all the following
	//   handlers, and therefore can't actually be provisioned.
	//   - create
	//   - read
	//   - delete
	ProvisioningType types.ProvisioningType

	// The version number of a public third-party extension. This applies only if you
	// specify a public extension you have activated in your account, or specify a
	// public extension without specifying a version. For all other extensions,
	// CloudFormation returns null .
	PublicVersionNumber *string

	// The publisher ID of the extension publisher. This applies only to public
	// third-party extensions. For private registered extensions, and extensions
	// provided by Amazon Web Services, CloudFormation returns null .
	PublisherId *string

	// For extensions that are modules, the public third-party extensions that must be
	// activated in your account in order for the module itself to be activated.
	RequiredActivatedTypes []types.RequiredActivatedType

	// The schema that defines the extension. For more information about extension
	// schemas, see Resource Provider Schema (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html)
	// in the CloudFormation CLI User Guide.
	Schema *string

	// The URL of the source code for the extension.
	SourceUrl *string

	// When the specified private extension version was registered or activated in
	// your account.
	TimeCreated *time.Time

	// The kind of extension.
	Type types.RegistryType

	// The name of the extension. If the extension is a public third-party type you
	// have activated with a type name alias, CloudFormation returns the type name
	// alias. For more information, see ActivateType (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ActivateType.html)
	// .
	TypeName *string

	// The contract test status of the registered extension version. To return the
	// extension test status of a specific extension version, you must specify
	// VersionId . This applies only to registered private extension versions.
	// CloudFormation doesn't return this information for public extensions, whether
	// they are activated in your account.
	//   - PASSED : The extension has passed all its contract tests. An extension must
	//   have a test status of PASSED before it can be published. For more information,
	//   see Publishing extensions to make them available for public use (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-publish.html)
	//   in the CloudFormation Command Line Interface User Guide.
	//   - FAILED : The extension has failed one or more contract tests.
	//   - IN_PROGRESS : Contract tests are currently being performed on the extension.
	//   - NOT_TESTED : Contract tests haven't been performed on the extension.
	TypeTestsStatus types.TypeTestsStatus

	// The description of the test status. To return the extension test status of a
	// specific extension version, you must specify VersionId . This applies only to
	// registered private extension versions. CloudFormation doesn't return this
	// information for public extensions, whether they are activated in your account.
	TypeTestsStatusDescription *string

	// The scope at which the extension is visible and usable in CloudFormation
	// operations. Valid values include:
	//   - PRIVATE : The extension is only visible and usable within the account in
	//   which it is registered. CloudFormation marks any extensions you register as
	//   PRIVATE .
	//   - PUBLIC : The extension is publicly visible and usable within any Amazon Web
	//   Services account.
	Visibility types.Visibility

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeType"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeType",
	}
}
