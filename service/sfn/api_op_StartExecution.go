// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts a state machine execution. A qualified state machine ARN can either
// refer to a Distributed Map state defined within a state machine, a version ARN,
// or an alias ARN. The following are some examples of qualified and unqualified
// state machine ARNs:
//   - The following qualified state machine ARN refers to a Distributed Map state
//     with a label mapStateLabel in a state machine named myStateMachine .
//     arn:partition:states:region:account-id:stateMachine:myStateMachine/mapStateLabel
//     If you provide a qualified state machine ARN that refers to a Distributed Map
//     state, the request fails with ValidationException .
//   - The following qualified state machine ARN refers to an alias named PROD .
//     arn::states:::stateMachine: If you provide a qualified state machine ARN that
//     refers to a version ARN or an alias ARN, the request starts execution for that
//     version or alias.
//   - The following unqualified state machine ARN refers to a state machine named
//     myStateMachine . arn::states:::stateMachine:
//
// If you start an execution with an unqualified state machine ARN, Step Functions
// uses the latest revision of the state machine for the execution. To start
// executions of a state machine version (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html)
// , call StartExecution and provide the version ARN or the ARN of an alias (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html)
// that points to the version. StartExecution is idempotent for STANDARD
// workflows. For a STANDARD workflow, if you call StartExecution with the same
// name and input as a running execution, the call succeeds and return the same
// response as the original request. If the execution is closed or if the input is
// different, it returns a 400 ExecutionAlreadyExists error. You can reuse names
// after 90 days. StartExecution isn't idempotent for EXPRESS workflows.
func (c *Client) StartExecution(ctx context.Context, params *StartExecutionInput, optFns ...func(*Options)) (*StartExecutionOutput, error) {
	if params == nil {
		params = &StartExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartExecution", params, optFns, c.addOperationStartExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartExecutionInput struct {

	// The Amazon Resource Name (ARN) of the state machine to execute. The
	// stateMachineArn parameter accepts one of the following inputs:
	//   - An unqualified state machine ARN – Refers to a state machine ARN that isn't
	//   qualified with a version or alias ARN. The following is an example of an
	//   unqualified state machine ARN. arn::states:::stateMachine: Step Functions
	//   doesn't associate state machine executions that you start with an unqualified
	//   ARN with a version. This is true even if that version uses the same revision
	//   that the execution used.
	//   - A state machine version ARN – Refers to a version ARN, which is a
	//   combination of state machine ARN and the version number separated by a colon
	//   (:). The following is an example of the ARN for version 10.
	//   arn::states:::stateMachine::10 Step Functions doesn't associate executions
	//   that you start with a version ARN with any aliases that point to that version.
	//   - A state machine alias ARN – Refers to an alias ARN, which is a combination
	//   of state machine ARN and the alias name separated by a colon (:). The following
	//   is an example of the ARN for an alias named PROD . arn::states:::stateMachine:
	//   Step Functions associates executions that you start with an alias ARN with that
	//   alias and the state machine version used for that execution.
	//
	// This member is required.
	StateMachineArn *string

	// The string that contains the JSON input data for the execution, for example:
	// "input": "{\"first_name\" : \"test\"}" If you don't include any JSON input data,
	// you still must include the two braces, for example: "input": "{}" Length
	// constraints apply to the payload size, and are expressed as bytes in UTF-8
	// encoding.
	Input *string

	// Optional name of the execution. This name must be unique for your Amazon Web
	// Services account, Region, and state machine for 90 days. For more information,
	// see Limits Related to State Machine Executions (https://docs.aws.amazon.com/step-functions/latest/dg/limits.html#service-limits-state-machine-executions)
	// in the Step Functions Developer Guide. If you don't provide a name for the
	// execution, Step Functions automatically generates a universally unique
	// identifier (UUID) as the execution name. A name must not contain:
	//   - white space
	//   - brackets < > { } [ ]
	//   - wildcard characters ? *
	//   - special characters " # % \ ^ | ~ ` $ & , ; : /
	//   - control characters ( U+0000-001F , U+007F-009F )
	// To enable logging with CloudWatch Logs, the name should only contain 0-9, A-Z,
	// a-z, - and _.
	Name *string

	// Passes the X-Ray trace header. The trace header can also be passed in the
	// request payload.
	TraceHeader *string

	noSmithyDocumentSerde
}

type StartExecutionOutput struct {

	// The Amazon Resource Name (ARN) that identifies the execution.
	//
	// This member is required.
	ExecutionArn *string

	// The date the execution is started.
	//
	// This member is required.
	StartDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartExecution"); err != nil {
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
	if err = addOpStartExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartExecution",
	}
}
