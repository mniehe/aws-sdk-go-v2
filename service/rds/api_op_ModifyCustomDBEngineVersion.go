// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Modifies the status of a custom engine version (CEV). You can find CEVs to
// modify by calling DescribeDBEngineVersions . The MediaImport service that
// imports files from Amazon S3 to create CEVs isn't integrated with Amazon Web
// Services CloudTrail. If you turn on data logging for Amazon RDS in CloudTrail,
// calls to the ModifyCustomDbEngineVersion event aren't logged. However, you
// might see calls from the API gateway that accesses your Amazon S3 bucket. These
// calls originate from the MediaImport service for the ModifyCustomDbEngineVersion
// event. For more information, see Modifying CEV status (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.modify)
// in the Amazon RDS User Guide.
func (c *Client) ModifyCustomDBEngineVersion(ctx context.Context, params *ModifyCustomDBEngineVersionInput, optFns ...func(*Options)) (*ModifyCustomDBEngineVersionOutput, error) {
	if params == nil {
		params = &ModifyCustomDBEngineVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCustomDBEngineVersion", params, optFns, c.addOperationModifyCustomDBEngineVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyCustomDBEngineVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyCustomDBEngineVersionInput struct {

	// The DB engine. The only supported values are custom-oracle-ee and
	// custom-oracle-ee-cdb .
	//
	// This member is required.
	Engine *string

	// The custom engine version (CEV) that you want to modify. This option is
	// required for RDS Custom for Oracle, but optional for Amazon RDS. The combination
	// of Engine and EngineVersion is unique per customer per Amazon Web Services
	// Region.
	//
	// This member is required.
	EngineVersion *string

	// An optional description of your CEV.
	Description *string

	// The availability status to be assigned to the CEV. Valid values are as follows:
	// available You can use this CEV to create a new RDS Custom DB instance. inactive
	// You can create a new RDS Custom instance by restoring a DB snapshot with this
	// CEV. You can't patch or create new instances with this CEV. You can change any
	// status to any status. A typical reason to change status is to prevent the
	// accidental use of a CEV, or to make a deprecated CEV eligible for use again. For
	// example, you might change the status of your CEV from available to inactive ,
	// and from inactive back to available . To change the availability status of the
	// CEV, it must not currently be in use by an RDS Custom instance, snapshot, or
	// automated backup.
	Status types.CustomEngineVersionStatus

	noSmithyDocumentSerde
}

// This data type is used as a response element in the action
// DescribeDBEngineVersions .
type ModifyCustomDBEngineVersionOutput struct {

	// The creation time of the DB engine version.
	CreateTime *time.Time

	// JSON string that lists the installation files and parameters that RDS Custom
	// uses to create a custom engine version (CEV). RDS Custom applies the patches in
	// the order in which they're listed in the manifest. You can set the Oracle home,
	// Oracle base, and UNIX/Linux user and group using the installation parameters.
	// For more information, see JSON fields in the CEV manifest (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.preparing.html#custom-cev.preparing.manifest.fields)
	// in the Amazon RDS User Guide.
	CustomDBEngineVersionManifest *string

	// The description of the database engine.
	DBEngineDescription *string

	// A value that indicates the source media provider of the AMI based on the usage
	// operation. Applicable for RDS Custom for SQL Server.
	DBEngineMediaType *string

	// The ARN of the custom engine version.
	DBEngineVersionArn *string

	// The description of the database engine version.
	DBEngineVersionDescription *string

	// The name of the DB parameter group family for the database engine.
	DBParameterGroupFamily *string

	// The name of the Amazon S3 bucket that contains your database installation files.
	DatabaseInstallationFilesS3BucketName *string

	// The Amazon S3 directory that contains the database installation files. If not
	// specified, then no prefix is assumed.
	DatabaseInstallationFilesS3Prefix *string

	// The default character set for new instances of this engine version, if the
	// CharacterSetName parameter of the CreateDBInstance API isn't specified.
	DefaultCharacterSet *types.CharacterSet

	// The name of the database engine.
	Engine *string

	// The version number of the database engine.
	EngineVersion *string

	// The types of logs that the database engine has available for export to
	// CloudWatch Logs.
	ExportableLogTypes []string

	// The EC2 image
	Image *types.CustomDBEngineVersionAMI

	// The Amazon Web Services KMS key identifier for an encrypted CEV. This parameter
	// is required for RDS Custom, but optional for Amazon RDS.
	KMSKeyId *string

	// The major engine version of the CEV.
	MajorEngineVersion *string

	// The status of the DB engine version, either available or deprecated .
	Status *string

	// A list of the supported CA certificate identifiers. For more information, see
	// Using SSL/TLS to encrypt a connection to a DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html)
	// in the Amazon RDS User Guide and Using SSL/TLS to encrypt a connection to a DB
	// cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html)
	// in the Amazon Aurora User Guide.
	SupportedCACertificateIdentifiers []string

	// A list of the character sets supported by this engine for the CharacterSetName
	// parameter of the CreateDBInstance operation.
	SupportedCharacterSets []types.CharacterSet

	// A list of the supported DB engine modes.
	SupportedEngineModes []string

	// A list of features supported by the DB engine. The supported features vary by
	// DB engine and DB engine version. To determine the supported features for a
	// specific DB engine and DB engine version using the CLI, use the following
	// command: aws rds describe-db-engine-versions --engine --engine-version  For
	// example, to determine the supported features for RDS for PostgreSQL version 13.3
	// using the CLI, use the following command: aws rds describe-db-engine-versions
	// --engine postgres --engine-version 13.3 The supported features are listed under
	// SupportedFeatureNames in the output.
	SupportedFeatureNames []string

	// A list of the character sets supported by the Oracle DB engine for the
	// NcharCharacterSetName parameter of the CreateDBInstance operation.
	SupportedNcharCharacterSets []types.CharacterSet

	// A list of the time zones supported by this engine for the Timezone parameter of
	// the CreateDBInstance action.
	SupportedTimezones []types.Timezone

	// Indicates whether the engine version supports Babelfish for Aurora PostgreSQL.
	SupportsBabelfish *bool

	// Indicates whether the engine version supports rotating the server certificate
	// without rebooting the DB instance.
	SupportsCertificateRotationWithoutRestart *bool

	// Indicates whether you can use Aurora global databases with a specific DB engine
	// version.
	SupportsGlobalDatabases *bool

	// Indicates whether the DB engine version supports zero-ETL integrations with
	// Amazon Redshift.
	SupportsIntegrations *bool

	// Indicates whether the DB engine version supports forwarding write operations
	// from reader DB instances to the writer DB instance in the DB cluster. By
	// default, write operations aren't allowed on reader DB instances. Valid for:
	// Aurora DB clusters only
	SupportsLocalWriteForwarding *bool

	// Indicates whether the engine version supports exporting the log types specified
	// by ExportableLogTypes to CloudWatch Logs.
	SupportsLogExportsToCloudwatchLogs *bool

	// Indicates whether you can use Aurora parallel query with a specific DB engine
	// version.
	SupportsParallelQuery *bool

	// Indicates whether the database engine version supports read replicas.
	SupportsReadReplica *bool

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	TagList []types.Tag

	// A list of engine versions that this database engine version can be upgraded to.
	ValidUpgradeTarget []types.UpgradeTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyCustomDBEngineVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyCustomDBEngineVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyCustomDBEngineVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyCustomDBEngineVersion"); err != nil {
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
	if err = addOpModifyCustomDBEngineVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCustomDBEngineVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyCustomDBEngineVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyCustomDBEngineVersion",
	}
}
