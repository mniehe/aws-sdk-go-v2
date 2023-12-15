// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a specified run of a workflow.
func (c *Client) GetWorkflowRun(ctx context.Context, params *GetWorkflowRunInput, optFns ...func(*Options)) (*GetWorkflowRunOutput, error) {
	if params == nil {
		params = &GetWorkflowRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflowRun", params, optFns, c.addOperationGetWorkflowRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowRunInput struct {

	// The ID of the workflow run. To retrieve a list of workflow run IDs, use
	// ListWorkflowRuns .
	//
	// This member is required.
	Id *string

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	noSmithyDocumentSerde
}

type GetWorkflowRunOutput struct {

	// The ID of the workflow run.
	//
	// This member is required.
	Id *string

	// The date and time the workflow run status was last updated, in coordinated
	// universal time (UTC) timestamp format as specified in RFC 3339 (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// The date and time the workflow run began, in coordinated universal time (UTC)
	// timestamp format as specified in RFC 3339 (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
	//
	// This member is required.
	StartTime *time.Time

	// The status of the workflow run.
	//
	// This member is required.
	Status types.WorkflowRunStatus

	// The ID of the workflow.
	//
	// This member is required.
	WorkflowId *string

	// The date and time the workflow run ended, in coordinated universal time (UTC)
	// timestamp format as specified in RFC 3339 (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
	// .
	EndTime *time.Time

	// Information about the reasons for the status of the workflow run.
	StatusReasons []types.WorkflowRunStatusReason

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkflowRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWorkflowRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWorkflowRun{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetWorkflowRun"); err != nil {
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
	if err = addOpGetWorkflowRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWorkflowRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetWorkflowRun",
	}
}
