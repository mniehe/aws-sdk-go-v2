// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/sqs/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAddPermission struct {
}

func (*validateOpAddPermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddPermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddPermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddPermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCancelMessageMoveTask struct {
}

func (*validateOpCancelMessageMoveTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelMessageMoveTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelMessageMoveTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelMessageMoveTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpChangeMessageVisibilityBatch struct {
}

func (*validateOpChangeMessageVisibilityBatch) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpChangeMessageVisibilityBatch) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ChangeMessageVisibilityBatchInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpChangeMessageVisibilityBatchInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpChangeMessageVisibility struct {
}

func (*validateOpChangeMessageVisibility) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpChangeMessageVisibility) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ChangeMessageVisibilityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpChangeMessageVisibilityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateQueue struct {
}

func (*validateOpCreateQueue) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateQueue) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateQueueInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateQueueInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMessageBatch struct {
}

func (*validateOpDeleteMessageBatch) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMessageBatch) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMessageBatchInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMessageBatchInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMessage struct {
}

func (*validateOpDeleteMessage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMessage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMessageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMessageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteQueue struct {
}

func (*validateOpDeleteQueue) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteQueue) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteQueueInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteQueueInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetQueueAttributes struct {
}

func (*validateOpGetQueueAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetQueueAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetQueueAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetQueueAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetQueueUrl struct {
}

func (*validateOpGetQueueUrl) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetQueueUrl) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetQueueUrlInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetQueueUrlInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListDeadLetterSourceQueues struct {
}

func (*validateOpListDeadLetterSourceQueues) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListDeadLetterSourceQueues) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListDeadLetterSourceQueuesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListDeadLetterSourceQueuesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListMessageMoveTasks struct {
}

func (*validateOpListMessageMoveTasks) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListMessageMoveTasks) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListMessageMoveTasksInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListMessageMoveTasksInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListQueueTags struct {
}

func (*validateOpListQueueTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListQueueTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListQueueTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListQueueTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPurgeQueue struct {
}

func (*validateOpPurgeQueue) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPurgeQueue) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PurgeQueueInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPurgeQueueInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpReceiveMessage struct {
}

func (*validateOpReceiveMessage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpReceiveMessage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ReceiveMessageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpReceiveMessageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemovePermission struct {
}

func (*validateOpRemovePermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemovePermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemovePermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemovePermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendMessageBatch struct {
}

func (*validateOpSendMessageBatch) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendMessageBatch) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendMessageBatchInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendMessageBatchInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendMessage struct {
}

func (*validateOpSendMessage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendMessage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendMessageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendMessageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetQueueAttributes struct {
}

func (*validateOpSetQueueAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetQueueAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetQueueAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetQueueAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartMessageMoveTask struct {
}

func (*validateOpStartMessageMoveTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartMessageMoveTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartMessageMoveTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartMessageMoveTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagQueue struct {
}

func (*validateOpTagQueue) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagQueue) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagQueueInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagQueueInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagQueue struct {
}

func (*validateOpUntagQueue) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagQueue) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagQueueInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagQueueInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAddPermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddPermission{}, middleware.After)
}

func addOpCancelMessageMoveTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelMessageMoveTask{}, middleware.After)
}

func addOpChangeMessageVisibilityBatchValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpChangeMessageVisibilityBatch{}, middleware.After)
}

func addOpChangeMessageVisibilityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpChangeMessageVisibility{}, middleware.After)
}

func addOpCreateQueueValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateQueue{}, middleware.After)
}

func addOpDeleteMessageBatchValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMessageBatch{}, middleware.After)
}

func addOpDeleteMessageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMessage{}, middleware.After)
}

func addOpDeleteQueueValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteQueue{}, middleware.After)
}

func addOpGetQueueAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetQueueAttributes{}, middleware.After)
}

func addOpGetQueueUrlValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetQueueUrl{}, middleware.After)
}

func addOpListDeadLetterSourceQueuesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListDeadLetterSourceQueues{}, middleware.After)
}

func addOpListMessageMoveTasksValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListMessageMoveTasks{}, middleware.After)
}

func addOpListQueueTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListQueueTags{}, middleware.After)
}

func addOpPurgeQueueValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPurgeQueue{}, middleware.After)
}

func addOpReceiveMessageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpReceiveMessage{}, middleware.After)
}

func addOpRemovePermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemovePermission{}, middleware.After)
}

func addOpSendMessageBatchValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSendMessageBatch{}, middleware.After)
}

func addOpSendMessageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSendMessage{}, middleware.After)
}

func addOpSetQueueAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetQueueAttributes{}, middleware.After)
}

func addOpStartMessageMoveTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartMessageMoveTask{}, middleware.After)
}

func addOpTagQueueValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagQueue{}, middleware.After)
}

func addOpUntagQueueValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagQueue{}, middleware.After)
}

func validateChangeMessageVisibilityBatchRequestEntry(v *types.ChangeMessageVisibilityBatchRequestEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ChangeMessageVisibilityBatchRequestEntry"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.ReceiptHandle == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReceiptHandle"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateChangeMessageVisibilityBatchRequestEntryList(v []types.ChangeMessageVisibilityBatchRequestEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ChangeMessageVisibilityBatchRequestEntryList"}
	for i := range v {
		if err := validateChangeMessageVisibilityBatchRequestEntry(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDeleteMessageBatchRequestEntry(v *types.DeleteMessageBatchRequestEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMessageBatchRequestEntry"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.ReceiptHandle == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReceiptHandle"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDeleteMessageBatchRequestEntryList(v []types.DeleteMessageBatchRequestEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMessageBatchRequestEntryList"}
	for i := range v {
		if err := validateDeleteMessageBatchRequestEntry(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMessageAttributeValue(v *types.MessageAttributeValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MessageAttributeValue"}
	if v.DataType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMessageBodyAttributeMap(v map[string]types.MessageAttributeValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MessageBodyAttributeMap"}
	for key := range v {
		value := v[key]
		if err := validateMessageAttributeValue(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMessageBodySystemAttributeMap(v map[string]types.MessageSystemAttributeValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MessageBodySystemAttributeMap"}
	for key := range v {
		value := v[key]
		if err := validateMessageSystemAttributeValue(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMessageSystemAttributeValue(v *types.MessageSystemAttributeValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MessageSystemAttributeValue"}
	if v.DataType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSendMessageBatchRequestEntry(v *types.SendMessageBatchRequestEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendMessageBatchRequestEntry"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.MessageBody == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MessageBody"))
	}
	if v.MessageAttributes != nil {
		if err := validateMessageBodyAttributeMap(v.MessageAttributes); err != nil {
			invalidParams.AddNested("MessageAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if v.MessageSystemAttributes != nil {
		if err := validateMessageBodySystemAttributeMap(v.MessageSystemAttributes); err != nil {
			invalidParams.AddNested("MessageSystemAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSendMessageBatchRequestEntryList(v []types.SendMessageBatchRequestEntry) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendMessageBatchRequestEntryList"}
	for i := range v {
		if err := validateSendMessageBatchRequestEntry(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddPermissionInput(v *AddPermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddPermissionInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if v.Label == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Label"))
	}
	if v.AWSAccountIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AWSAccountIds"))
	}
	if v.Actions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Actions"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelMessageMoveTaskInput(v *CancelMessageMoveTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelMessageMoveTaskInput"}
	if v.TaskHandle == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskHandle"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpChangeMessageVisibilityBatchInput(v *ChangeMessageVisibilityBatchInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ChangeMessageVisibilityBatchInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if v.Entries == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Entries"))
	} else if v.Entries != nil {
		if err := validateChangeMessageVisibilityBatchRequestEntryList(v.Entries); err != nil {
			invalidParams.AddNested("Entries", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpChangeMessageVisibilityInput(v *ChangeMessageVisibilityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ChangeMessageVisibilityInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if v.ReceiptHandle == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReceiptHandle"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateQueueInput(v *CreateQueueInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateQueueInput"}
	if v.QueueName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMessageBatchInput(v *DeleteMessageBatchInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMessageBatchInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if v.Entries == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Entries"))
	} else if v.Entries != nil {
		if err := validateDeleteMessageBatchRequestEntryList(v.Entries); err != nil {
			invalidParams.AddNested("Entries", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMessageInput(v *DeleteMessageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMessageInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if v.ReceiptHandle == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReceiptHandle"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteQueueInput(v *DeleteQueueInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteQueueInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetQueueAttributesInput(v *GetQueueAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetQueueAttributesInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetQueueUrlInput(v *GetQueueUrlInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetQueueUrlInput"}
	if v.QueueName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListDeadLetterSourceQueuesInput(v *ListDeadLetterSourceQueuesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListDeadLetterSourceQueuesInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListMessageMoveTasksInput(v *ListMessageMoveTasksInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListMessageMoveTasksInput"}
	if v.SourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListQueueTagsInput(v *ListQueueTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListQueueTagsInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPurgeQueueInput(v *PurgeQueueInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PurgeQueueInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpReceiveMessageInput(v *ReceiveMessageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ReceiveMessageInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemovePermissionInput(v *RemovePermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemovePermissionInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if v.Label == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Label"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendMessageBatchInput(v *SendMessageBatchInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendMessageBatchInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if v.Entries == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Entries"))
	} else if v.Entries != nil {
		if err := validateSendMessageBatchRequestEntryList(v.Entries); err != nil {
			invalidParams.AddNested("Entries", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendMessageInput(v *SendMessageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendMessageInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if v.MessageBody == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MessageBody"))
	}
	if v.MessageAttributes != nil {
		if err := validateMessageBodyAttributeMap(v.MessageAttributes); err != nil {
			invalidParams.AddNested("MessageAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if v.MessageSystemAttributes != nil {
		if err := validateMessageBodySystemAttributeMap(v.MessageSystemAttributes); err != nil {
			invalidParams.AddNested("MessageSystemAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetQueueAttributesInput(v *SetQueueAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetQueueAttributesInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if v.Attributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Attributes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartMessageMoveTaskInput(v *StartMessageMoveTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartMessageMoveTaskInput"}
	if v.SourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagQueueInput(v *TagQueueInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagQueueInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagQueueInput(v *UntagQueueInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagQueueInput"}
	if v.QueueUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QueueUrl"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
