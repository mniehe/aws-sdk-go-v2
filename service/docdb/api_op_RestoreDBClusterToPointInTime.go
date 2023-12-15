// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Restores a cluster to an arbitrary point in time. Users can restore to any
// point in time before LatestRestorableTime for up to BackupRetentionPeriod days.
// The target cluster is created from the source cluster with the same
// configuration as the original cluster, except that the new cluster is created
// with the default security group.
func (c *Client) RestoreDBClusterToPointInTime(ctx context.Context, params *RestoreDBClusterToPointInTimeInput, optFns ...func(*Options)) (*RestoreDBClusterToPointInTimeOutput, error) {
	if params == nil {
		params = &RestoreDBClusterToPointInTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBClusterToPointInTime", params, optFns, c.addOperationRestoreDBClusterToPointInTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBClusterToPointInTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to RestoreDBClusterToPointInTime .
type RestoreDBClusterToPointInTimeInput struct {

	// The name of the new cluster to be created. Constraints:
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//   - The first character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The identifier of the source cluster from which to restore. Constraints:
	//   - Must match the identifier of an existing DBCluster .
	//
	// This member is required.
	SourceDBClusterIdentifier *string

	// The subnet group name to use for the new cluster. Constraints: If provided,
	// must match the name of an existing DBSubnetGroup . Example: mySubnetgroup
	DBSubnetGroupName *string

	// Specifies whether this cluster can be deleted. If DeletionProtection is
	// enabled, the cluster cannot be deleted unless it is modified and
	// DeletionProtection is disabled. DeletionProtection protects clusters from being
	// accidentally deleted.
	DeletionProtection *bool

	// A list of log types that must be enabled for exporting to Amazon CloudWatch
	// Logs.
	EnableCloudwatchLogsExports []string

	// The KMS key identifier to use when restoring an encrypted cluster from an
	// encrypted cluster. The KMS key identifier is the Amazon Resource Name (ARN) for
	// the KMS encryption key. If you are restoring a cluster with the same Amazon Web
	// Services account that owns the KMS encryption key used to encrypt the new
	// cluster, then you can use the KMS key alias instead of the ARN for the KMS
	// encryption key. You can restore to a new cluster and encrypt the new cluster
	// with an KMS key that is different from the KMS key used to encrypt the source
	// cluster. The new DB cluster is encrypted with the KMS key identified by the
	// KmsKeyId parameter. If you do not specify a value for the KmsKeyId parameter,
	// then the following occurs:
	//   - If the cluster is encrypted, then the restored cluster is encrypted using
	//   the KMS key that was used to encrypt the source cluster.
	//   - If the cluster is not encrypted, then the restored cluster is not
	//   encrypted.
	// If DBClusterIdentifier refers to a cluster that is not encrypted, then the
	// restore request is rejected.
	KmsKeyId *string

	// The port number on which the new cluster accepts connections. Constraints: Must
	// be a value from 1150 to 65535 . Default: The default port for the engine.
	Port *int32

	// The date and time to restore the cluster to. Valid values: A time in Universal
	// Coordinated Time (UTC) format. Constraints:
	//   - Must be before the latest restorable time for the instance.
	//   - Must be specified if the UseLatestRestorableTime parameter is not provided.
	//   - Cannot be specified if the UseLatestRestorableTime parameter is true .
	//   - Cannot be specified if the RestoreType parameter is copy-on-write .
	// Example: 2015-03-07T23:45:00Z
	RestoreToTime *time.Time

	// The type of restore to be performed. You can specify one of the following
	// values:
	//   - full-copy - The new DB cluster is restored as a full copy of the source DB
	//   cluster.
	//   - copy-on-write - The new DB cluster is restored as a clone of the source DB
	//   cluster.
	// Constraints: You can't specify copy-on-write if the engine version of the
	// source DB cluster is earlier than 1.11. If you don't specify a RestoreType
	// value, then the new DB cluster is restored as a full copy of the source DB
	// cluster.
	RestoreType *string

	// The storage type to associate with the DB cluster. For information on storage
	// types for Amazon DocumentDB clusters, see Cluster storage configurations in the
	// Amazon DocumentDB Developer Guide. Valid values for storage type - standard |
	// iopt1 Default value is standard
	StorageType *string

	// The tags to be assigned to the restored cluster.
	Tags []types.Tag

	// A value that is set to true to restore the cluster to the latest restorable
	// backup time, and false otherwise. Default: false Constraints: Cannot be
	// specified if the RestoreToTime parameter is provided.
	UseLatestRestorableTime *bool

	// A list of VPC security groups that the new cluster belongs to.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreDBClusterToPointInTimeOutput struct {

	// Detailed information about a cluster.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreDBClusterToPointInTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBClusterToPointInTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBClusterToPointInTime{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreDBClusterToPointInTime"); err != nil {
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
	if err = addOpRestoreDBClusterToPointInTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBClusterToPointInTime(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreDBClusterToPointInTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreDBClusterToPointInTime",
	}
}
