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

// Retrieves a specified import task.
func (c *Client) GetImportTask(ctx context.Context, params *GetImportTaskInput, optFns ...func(*Options)) (*GetImportTaskOutput, error) {
	if params == nil {
		params = &GetImportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetImportTask", params, optFns, c.addOperationGetImportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetImportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetImportTaskInput struct {

	// The unique identifier of the import task.
	//
	// This member is required.
	TaskIdentifier *string

	noSmithyDocumentSerde
}

func (in *GetImportTaskInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type GetImportTaskOutput struct {

	// The ARN of the IAM role that will allow access to the data that is to be
	// imported.
	//
	// This member is required.
	RoleArn *string

	// A URL identifying to the location of the data to be imported. This can be an
	// Amazon S3 path, or can point to a Neptune database endpoint or snapshot
	//
	// This member is required.
	Source *string

	// The status of the import task:
	//   - INITIALIZING – The necessary resources needed to create the graph are being
	//   prepared.
	//   - ANALYZING_DATA – The data is being analyzed to determine the optimal
	//   infrastructure configuration for the new graph.
	//   - RE_PROVISIONING – The data did not fit into the provisioned graph, so it is
	//   being re-provisioned with more capacity.
	//   - IMPORTING – The data is being loaded.
	//   - ERROR_ENCOUNTERED – An error has been encountered while trying to create
	//   the graph and import the data.
	//   - ERROR_ENCOUNTERED_ROLLING_BACK – Because of the error that was encountered,
	//   the graph is being rolled back and all its resources released.
	//   - SUCCEEDED – Graph creation and data loading succeeded.
	//   - FAILED – Graph creation or data loading failed. When the status is FAILED ,
	//   you can use get-graphs to get more information about the state of the graph.
	//   - CANCELLING – Because you cancelled the import task, cancellation is in
	//   progress.
	//   - CANCELLED – You have successfully cancelled the import task.
	//
	// This member is required.
	Status types.ImportTaskStatus

	// The unique identifier of the import task.
	//
	// This member is required.
	TaskId *string

	// The number of the current attempt to execute the import task.
	AttemptNumber *int32

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

	// Contains details about the specified import task.
	ImportTaskDetails *types.ImportTaskDetails

	// The reason that the import task has this status value.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetImportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetImportTask"); err != nil {
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
	if err = addOpGetImportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetImportTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetImportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetImportTask",
	}
}
