// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/neptunegraph/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Neptune Analytics graph and imports data into it, either from
// Amazon Simple Storage Service (S3) or from a Neptune database or a Neptune
// database snapshot. The data can be loaded from files in S3 that in either the
// Gremlin CSV format (https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-gremlin.html)
// or the openCypher load format (https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-opencypher.html)
// .
func (c *Client) CreateGraphUsingImportTask(ctx context.Context, params *CreateGraphUsingImportTaskInput, optFns ...func(*Options)) (*CreateGraphUsingImportTaskOutput, error) {
	if params == nil {
		params = &CreateGraphUsingImportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGraphUsingImportTask", params, optFns, c.addOperationCreateGraphUsingImportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGraphUsingImportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGraphUsingImportTaskInput struct {

	// A name for the new Neptune Analytics graph to be created. The name must contain
	// from 1 to 63 letters, numbers, or hyphens, and its first character must be a
	// letter. It cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	GraphName *string

	// The ARN of the IAM role that will allow access to the data that is to be
	// imported.
	//
	// This member is required.
	RoleArn *string

	// A URL identifying to the location of the data to be imported. This can be an
	// Amazon S3 path, or can point to a Neptune database endpoint or snapshot.
	//
	// This member is required.
	Source *string

	// Indicates whether or not to enable deletion protection on the graph. The graph
	// can’t be deleted when deletion protection is enabled. ( true or false ).
	DeletionProtection *bool

	// If set to true , the task halts when an import error is encountered. If set to
	// false , the task skips the data that caused the error and continues if possible.
	FailOnError *bool

	// Specifies the format of S3 data to be imported. Valid values are CSV , which
	// identifies the Gremlin CSV format (https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-gremlin.html)
	// or OPENCYPHER , which identies the openCypher load format (https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-opencypher.html)
	// .
	Format types.Format

	// Contains options for controlling the import process. For example, if the
	// failOnError key is set to false , the import skips problem data and attempts to
	// continue (whereas if set to true , the default, or if omitted, the import
	// operation halts immediately when an error is encountered.
	ImportOptions types.ImportOptions

	// Specifies a KMS key to use to encrypt data imported into the new graph.
	KmsKeyIdentifier *string

	// The maximum provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use
	// for the graph. Default: 1024, or the approved upper limit for your account. If
	// both the minimum and maximum values are specified, the max of the
	// min-provisioned-memory and max-provisioned memory is used to create the graph.
	// If neither value is specified 128 m-NCUs are used.
	MaxProvisionedMemory *int32

	// The minimum provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use
	// for the graph. Default: 128
	MinProvisionedMemory *int32

	// Specifies whether or not the graph can be reachable over the internet. All
	// access to graphs IAM authenticated. ( true to enable, or false to disable.
	PublicConnectivity *bool

	// The number of replicas in other AZs to provision on the new graph after import.
	// Default = 0, Min = 0, Max = 2.
	ReplicaCount *int32

	// Adds metadata tags to the new graph. These tags can also be used with cost
	// allocation reporting, or used in a Condition statement in an IAM policy.
	Tags map[string]string

	// Specifies the number of dimensions for vector embeddings that will be loaded
	// into the graph. The value is specified as dimension= value. Max = 65,535
	VectorSearchConfiguration *types.VectorSearchConfiguration

	noSmithyDocumentSerde
}

func (in *CreateGraphUsingImportTaskInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type CreateGraphUsingImportTaskOutput struct {

	// The ARN of the IAM role that will allow access to the data that is to be
	// imported.
	//
	// This member is required.
	RoleArn *string

	// A URL identifying to the location of the data to be imported. This can be an
	// Amazon S3 path, or can point to a Neptune database endpoint or snapshot.
	//
	// This member is required.
	Source *string

	// The status of the import task.
	//
	// This member is required.
	Status types.ImportTaskStatus

	// The unique identifier of the import task.
	//
	// This member is required.
	TaskId *string

	// Specifies the format of S3 data to be imported. Valid values are CSV , which
	// identifies the Gremlin CSV format (https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-gremlin.html)
	// or OPENCYPHER , which identies the openCypher load format (https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-opencypher.html)
	// .
	Format types.Format

	// The unique identifier of the Neptune Analytics graph.
	GraphId *string

	// Contains options for controlling the import process. For example, if the
	// failOnError key is set to false , the import skips problem data and attempts to
	// continue (whereas if set to true , the default, or if omitted, the import
	// operation halts immediately when an error is encountered.
	ImportOptions types.ImportOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGraphUsingImportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGraphUsingImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGraphUsingImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateGraphUsingImportTask"); err != nil {
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
	if err = addOpCreateGraphUsingImportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGraphUsingImportTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGraphUsingImportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateGraphUsingImportTask",
	}
}
