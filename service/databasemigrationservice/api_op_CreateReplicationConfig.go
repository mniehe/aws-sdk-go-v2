// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a configuration that you can later provide to configure and start an
// DMS Serverless replication. You can also provide options to validate the
// configuration inputs before you start the replication.
func (c *Client) CreateReplicationConfig(ctx context.Context, params *CreateReplicationConfigInput, optFns ...func(*Options)) (*CreateReplicationConfigOutput, error) {
	if params == nil {
		params = &CreateReplicationConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReplicationConfig", params, optFns, c.addOperationCreateReplicationConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReplicationConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReplicationConfigInput struct {

	// Configuration parameters for provisioning an DMS Serverless replication.
	//
	// This member is required.
	ComputeConfig *types.ComputeConfig

	// A unique identifier that you want to use to create a ReplicationConfigArn that
	// is returned as part of the output from this action. You can then pass this
	// output ReplicationConfigArn as the value of the ReplicationConfigArn option for
	// other actions to identify both DMS Serverless replications and replication
	// configurations that you want those actions to operate on. For some actions, you
	// can also use either this unique identifier or a corresponding ARN in action
	// filters to identify the specific replication and replication configuration to
	// operate on.
	//
	// This member is required.
	ReplicationConfigIdentifier *string

	// The type of DMS Serverless replication to provision using this replication
	// configuration. Possible values:
	//   - "full-load"
	//   - "cdc"
	//   - "full-load-and-cdc"
	//
	// This member is required.
	ReplicationType types.MigrationTypeValue

	// The Amazon Resource Name (ARN) of the source endpoint for this DMS Serverless
	// replication configuration.
	//
	// This member is required.
	SourceEndpointArn *string

	// JSON table mappings for DMS Serverless replications that are provisioned using
	// this replication configuration. For more information, see Specifying table
	// selection and transformations rules using JSON (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.SelectionTransformation.html)
	// .
	//
	// This member is required.
	TableMappings *string

	// The Amazon Resource Name (ARN) of the target endpoint for this DMS serverless
	// replication configuration.
	//
	// This member is required.
	TargetEndpointArn *string

	// Optional JSON settings for DMS Serverless replications that are provisioned
	// using this replication configuration. For example, see Change processing tuning
	// settings (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.ChangeProcessingTuning.html)
	// .
	ReplicationSettings *string

	// Optional unique value or name that you set for a given resource that can be
	// used to construct an Amazon Resource Name (ARN) for that resource. For more
	// information, see Fine-grained access control using resource names and tags (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#CHAP_Security.FineGrainedAccess)
	// .
	ResourceIdentifier *string

	// Optional JSON settings for specifying supplemental data. For more information,
	// see Specifying supplemental data for task settings (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.TaskData.html)
	// .
	SupplementalSettings *string

	// One or more optional tags associated with resources used by the DMS Serverless
	// replication. For more information, see Tagging resources in Database Migration
	// Service (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tagging.html) .
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateReplicationConfigOutput struct {

	// Configuration parameters returned from the DMS Serverless replication after it
	// is created.
	ReplicationConfig *types.ReplicationConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReplicationConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateReplicationConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateReplicationConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateReplicationConfig"); err != nil {
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
	if err = addOpCreateReplicationConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicationConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReplicationConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateReplicationConfig",
	}
}
