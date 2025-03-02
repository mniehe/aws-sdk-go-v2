// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration of a DataSync transfer task.
func (c *Client) UpdateTask(ctx context.Context, params *UpdateTaskInput, optFns ...func(*Options)) (*UpdateTaskOutput, error) {
	if params == nil {
		params = &UpdateTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTask", params, optFns, c.addOperationUpdateTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// UpdateTaskResponse
type UpdateTaskInput struct {

	// The Amazon Resource Name (ARN) of the resource name of the task to update.
	//
	// This member is required.
	TaskArn *string

	// The Amazon Resource Name (ARN) of the resource name of the Amazon CloudWatch
	// log group.
	CloudWatchLogGroupArn *string

	// Specifies a list of filter rules that exclude specific data during your
	// transfer. For more information and examples, see Filtering data transferred by
	// DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
	Excludes []types.FilterRule

	// Specifies a list of filter rules that include specific data during your
	// transfer. For more information and examples, see Filtering data transferred by
	// DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
	Includes []types.FilterRule

	// The name of the task to update.
	Name *string

	// Indicates how your transfer task is configured. These options include how
	// DataSync handles files, objects, and their associated metadata during your
	// transfer. You also can specify how to verify data integrity, set bandwidth
	// limits for your task, among other options. Each option has a default value.
	// Unless you need to, you don't have to configure any of these options before
	// starting your task.
	Options *types.Options

	// Specifies a schedule used to periodically transfer files from a source to a
	// destination location. You can configure your task to execute hourly, daily,
	// weekly or on specific days of the week. You control when in the day or hour you
	// want the task to execute. The time you specify is UTC time. For more
	// information, see Scheduling your task (https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html)
	// .
	Schedule *types.TaskSchedule

	// Specifies how you want to configure a task report, which provides detailed
	// information about for your DataSync transfer.
	TaskReportConfig *types.TaskReportConfig

	noSmithyDocumentSerde
}

type UpdateTaskOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTask"); err != nil {
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
	if err = addOpUpdateTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTask",
	}
}
