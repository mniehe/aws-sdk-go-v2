// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/migrationhuborchestrator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a step in the migration workflow.
func (c *Client) CreateWorkflowStep(ctx context.Context, params *CreateWorkflowStepInput, optFns ...func(*Options)) (*CreateWorkflowStepOutput, error) {
	if params == nil {
		params = &CreateWorkflowStepInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkflowStep", params, optFns, c.addOperationCreateWorkflowStepMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkflowStepOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkflowStepInput struct {

	// The name of the step.
	//
	// This member is required.
	Name *string

	// The action type of the step. You must run and update the status of a manual
	// step for the workflow to continue after the completion of the step.
	//
	// This member is required.
	StepActionType types.StepActionType

	// The ID of the step group.
	//
	// This member is required.
	StepGroupId *string

	// The ID of the migration workflow.
	//
	// This member is required.
	WorkflowId *string

	// The description of the step.
	Description *string

	// The next step.
	Next []string

	// The key value pairs added for the expected output.
	Outputs []types.WorkflowStepOutput

	// The previous step.
	Previous []string

	// The servers on which a step will be run.
	StepTarget []string

	// The custom script to run tests on source or target environments.
	WorkflowStepAutomationConfiguration *types.WorkflowStepAutomationConfiguration

	noSmithyDocumentSerde
}

type CreateWorkflowStepOutput struct {

	// The ID of the step.
	Id *string

	// The name of the step.
	Name *string

	// The ID of the step group.
	StepGroupId *string

	// The ID of the migration workflow.
	WorkflowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkflowStepMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorkflowStep{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorkflowStep{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWorkflowStep"); err != nil {
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
	if err = addOpCreateWorkflowStepValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkflowStep(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkflowStep(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWorkflowStep",
	}
}
