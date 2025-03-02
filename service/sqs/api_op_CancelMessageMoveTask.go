// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels a specified message movement task. A message movement can only be
// cancelled when the current status is RUNNING. Cancelling a message movement task
// does not revert the messages that have already been moved. It can only stop the
// messages that have not been moved yet.
//   - This action is currently limited to supporting message redrive from
//     dead-letter queues (DLQs) (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html)
//     only. In this context, the source queue is the dead-letter queue (DLQ), while
//     the destination queue can be the original source queue (from which the messages
//     were driven to the dead-letter-queue), or a custom destination queue.
//   - Currently, only standard queues are supported.
//   - Only one active message movement task is supported per queue at any given
//     time.
func (c *Client) CancelMessageMoveTask(ctx context.Context, params *CancelMessageMoveTaskInput, optFns ...func(*Options)) (*CancelMessageMoveTaskOutput, error) {
	if params == nil {
		params = &CancelMessageMoveTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelMessageMoveTask", params, optFns, c.addOperationCancelMessageMoveTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelMessageMoveTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelMessageMoveTaskInput struct {

	// An identifier associated with a message movement task.
	//
	// This member is required.
	TaskHandle *string

	noSmithyDocumentSerde
}

type CancelMessageMoveTaskOutput struct {

	// The approximate number of messages already moved to the destination queue.
	ApproximateNumberOfMessagesMoved int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelMessageMoveTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCancelMessageMoveTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCancelMessageMoveTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CancelMessageMoveTask"); err != nil {
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
	if err = addOpCancelMessageMoveTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelMessageMoveTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelMessageMoveTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CancelMessageMoveTask",
	}
}
