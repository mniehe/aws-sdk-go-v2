// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the runtime information that was logged for a specific runtime instance of
// the workflow.
func (c *Client) GetWorkflowExecution(ctx context.Context, params *GetWorkflowExecutionInput, optFns ...func(*Options)) (*GetWorkflowExecutionOutput, error) {
	if params == nil {
		params = &GetWorkflowExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflowExecution", params, optFns, c.addOperationGetWorkflowExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowExecutionInput struct {

	// Use the unique identifier for a runtime instance of the workflow to get runtime
	// details.
	//
	// This member is required.
	WorkflowExecutionId *string

	noSmithyDocumentSerde
}

type GetWorkflowExecutionOutput struct {

	// The timestamp when the specified runtime instance of the workflow finished.
	EndTime *string

	// The Amazon Resource Name (ARN) of the image resource build version that the
	// specified runtime instance of the workflow created.
	ImageBuildVersionArn *string

	// The output message from the specified runtime instance of the workflow, if
	// applicable.
	Message *string

	// Test workflows are defined within named runtime groups. The parallel group is a
	// named group that contains one or more test workflows.
	ParallelGroup *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// The timestamp when the specified runtime instance of the workflow started.
	StartTime *string

	// The current runtime status for the specified runtime instance of the workflow.
	Status types.WorkflowExecutionStatus

	// The total number of steps in the specified runtime instance of the workflow
	// that ran. This number should equal the sum of the step counts for steps that
	// succeeded, were skipped, and failed.
	TotalStepCount int32

	// A runtime count for the number of steps that failed in the specified runtime
	// instance of the workflow.
	TotalStepsFailed int32

	// A runtime count for the number of steps that were skipped in the specified
	// runtime instance of the workflow.
	TotalStepsSkipped int32

	// A runtime count for the number of steps that ran successfully in the specified
	// runtime instance of the workflow.
	TotalStepsSucceeded int32

	// The type of workflow that Image Builder ran for the specified runtime instance
	// of the workflow.
	Type types.WorkflowType

	// The Amazon Resource Name (ARN) of the build version for the Image Builder
	// workflow resource that defines the specified runtime instance of the workflow.
	WorkflowBuildVersionArn *string

	// The unique identifier that Image Builder assigned to keep track of runtime
	// details when it ran the workflow.
	WorkflowExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkflowExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWorkflowExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWorkflowExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetWorkflowExecution"); err != nil {
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
	if err = addOpGetWorkflowExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWorkflowExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetWorkflowExecution",
	}
}
