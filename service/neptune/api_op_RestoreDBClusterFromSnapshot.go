// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new DB cluster from a DB snapshot or DB cluster snapshot. If a DB
// snapshot is specified, the target DB cluster is created from the source DB
// snapshot with a default configuration and default security group. If a DB
// cluster snapshot is specified, the target DB cluster is created from the source
// DB cluster restore point with the same configuration as the original source DB
// cluster, except that the new DB cluster is created with the default security
// group.
func (c *Client) RestoreDBClusterFromSnapshot(ctx context.Context, params *RestoreDBClusterFromSnapshotInput, optFns ...func(*Options)) (*RestoreDBClusterFromSnapshotOutput, error) {
	if params == nil {
		params = &RestoreDBClusterFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBClusterFromSnapshot", params, optFns, c.addOperationRestoreDBClusterFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBClusterFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreDBClusterFromSnapshotInput struct {

	// The name of the DB cluster to create from the DB snapshot or DB cluster
	// snapshot. This parameter isn't case-sensitive. Constraints:
	//   - Must contain from 1 to 63 letters, numbers, or hyphens
	//   - First character must be a letter
	//   - Cannot end with a hyphen or contain two consecutive hyphens
	// Example: my-snapshot-id
	//
	// This member is required.
	DBClusterIdentifier *string

	// The database engine to use for the new DB cluster. Default: The same as source
	// Constraint: Must be compatible with the engine of the source
	//
	// This member is required.
	Engine *string

	// The identifier for the DB snapshot or DB cluster snapshot to restore from. You
	// can use either the name or the Amazon Resource Name (ARN) to specify a DB
	// cluster snapshot. However, you can use only the ARN to specify a DB snapshot.
	// Constraints:
	//   - Must match the identifier of an existing Snapshot.
	//
	// This member is required.
	SnapshotIdentifier *string

	// Provides the list of EC2 Availability Zones that instances in the restored DB
	// cluster can be created in.
	AvailabilityZones []string

	// If set to true , tags are copied to any snapshot of the restored DB cluster that
	// is created.
	CopyTagsToSnapshot *bool

	// The name of the DB cluster parameter group to associate with the new DB
	// cluster. Constraints:
	//   - If supplied, must match the name of an existing DBClusterParameterGroup.
	DBClusterParameterGroupName *string

	// The name of the DB subnet group to use for the new DB cluster. Constraints: If
	// supplied, must match the name of an existing DBSubnetGroup. Example:
	// mySubnetgroup
	DBSubnetGroupName *string

	// Not supported.
	DatabaseName *string

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled.
	DeletionProtection *bool

	// The list of logs that the restored DB cluster is to export to Amazon CloudWatch
	// Logs.
	EnableCloudwatchLogsExports []string

	// True to enable mapping of Amazon Identity and Access Management (IAM) accounts
	// to database accounts, and otherwise false. Default: false
	EnableIAMDatabaseAuthentication *bool

	// The version of the database engine to use for the new DB cluster.
	EngineVersion *string

	// The Amazon KMS key identifier to use when restoring an encrypted DB cluster
	// from a DB snapshot or DB cluster snapshot. The KMS key identifier is the Amazon
	// Resource Name (ARN) for the KMS encryption key. If you are restoring a DB
	// cluster with the same Amazon account that owns the KMS encryption key used to
	// encrypt the new DB cluster, then you can use the KMS key alias instead of the
	// ARN for the KMS encryption key. If you do not specify a value for the KmsKeyId
	// parameter, then the following will occur:
	//   - If the DB snapshot or DB cluster snapshot in SnapshotIdentifier is
	//   encrypted, then the restored DB cluster is encrypted using the KMS key that was
	//   used to encrypt the DB snapshot or DB cluster snapshot.
	//   - If the DB snapshot or DB cluster snapshot in SnapshotIdentifier is not
	//   encrypted, then the restored DB cluster is not encrypted.
	KmsKeyId *string

	// (Not supported by Neptune)
	OptionGroupName *string

	// The port number on which the new DB cluster accepts connections. Constraints:
	// Value must be 1150-65535 Default: The same port as the original DB cluster.
	Port *int32

	// Contains the scaling configuration of a Neptune Serverless DB cluster. For more
	// information, see Using Amazon Neptune Serverless (https://docs.aws.amazon.com/neptune/latest/userguide/neptune-serverless-using.html)
	// in the Amazon Neptune User Guide.
	ServerlessV2ScalingConfiguration *types.ServerlessV2ScalingConfiguration

	// Specifies the storage type to be associated with the DB cluster. Valid values:
	// standard , iopt1 Default: standard
	StorageType *string

	// The tags to be assigned to the restored DB cluster.
	Tags []types.Tag

	// A list of VPC security groups that the new DB cluster will belong to.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreDBClusterFromSnapshotOutput struct {

	// Contains the details of an Amazon Neptune DB cluster. This data type is used as
	// a response element in the DescribeDBClusters .
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreDBClusterFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBClusterFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBClusterFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreDBClusterFromSnapshot"); err != nil {
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
	if err = addOpRestoreDBClusterFromSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBClusterFromSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreDBClusterFromSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreDBClusterFromSnapshot",
	}
}
