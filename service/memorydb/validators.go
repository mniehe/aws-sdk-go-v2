// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/memorydb/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchUpdateCluster struct {
}

func (*validateOpBatchUpdateCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchUpdateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchUpdateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchUpdateClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCopySnapshot struct {
}

func (*validateOpCopySnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCopySnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CopySnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCopySnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateACL struct {
}

func (*validateOpCreateACL) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateACL) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateACLInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateACLInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCluster struct {
}

func (*validateOpCreateCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateParameterGroup struct {
}

func (*validateOpCreateParameterGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateParameterGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateParameterGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateParameterGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSnapshot struct {
}

func (*validateOpCreateSnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSubnetGroup struct {
}

func (*validateOpCreateSubnetGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSubnetGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSubnetGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSubnetGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateUser struct {
}

func (*validateOpCreateUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteACL struct {
}

func (*validateOpDeleteACL) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteACL) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteACLInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteACLInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCluster struct {
}

func (*validateOpDeleteCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteParameterGroup struct {
}

func (*validateOpDeleteParameterGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteParameterGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteParameterGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteParameterGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSnapshot struct {
}

func (*validateOpDeleteSnapshot) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSnapshotInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSubnetGroup struct {
}

func (*validateOpDeleteSubnetGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSubnetGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSubnetGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSubnetGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteUser struct {
}

func (*validateOpDeleteUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeParameters struct {
}

func (*validateOpDescribeParameters) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeParameters) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeParametersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeParametersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeUsers struct {
}

func (*validateOpDescribeUsers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeUsers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeUsersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeUsersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpFailoverShard struct {
}

func (*validateOpFailoverShard) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpFailoverShard) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*FailoverShardInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpFailoverShardInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAllowedNodeTypeUpdates struct {
}

func (*validateOpListAllowedNodeTypeUpdates) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAllowedNodeTypeUpdates) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAllowedNodeTypeUpdatesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAllowedNodeTypeUpdatesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTags struct {
}

func (*validateOpListTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPurchaseReservedNodesOffering struct {
}

func (*validateOpPurchaseReservedNodesOffering) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPurchaseReservedNodesOffering) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PurchaseReservedNodesOfferingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPurchaseReservedNodesOfferingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpResetParameterGroup struct {
}

func (*validateOpResetParameterGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpResetParameterGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ResetParameterGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpResetParameterGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateACL struct {
}

func (*validateOpUpdateACL) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateACL) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateACLInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateACLInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCluster struct {
}

func (*validateOpUpdateCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateParameterGroup struct {
}

func (*validateOpUpdateParameterGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateParameterGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateParameterGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateParameterGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSubnetGroup struct {
}

func (*validateOpUpdateSubnetGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSubnetGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSubnetGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSubnetGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateUser struct {
}

func (*validateOpUpdateUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchUpdateClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchUpdateCluster{}, middleware.After)
}

func addOpCopySnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCopySnapshot{}, middleware.After)
}

func addOpCreateACLValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateACL{}, middleware.After)
}

func addOpCreateClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCluster{}, middleware.After)
}

func addOpCreateParameterGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateParameterGroup{}, middleware.After)
}

func addOpCreateSnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSnapshot{}, middleware.After)
}

func addOpCreateSubnetGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSubnetGroup{}, middleware.After)
}

func addOpCreateUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateUser{}, middleware.After)
}

func addOpDeleteACLValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteACL{}, middleware.After)
}

func addOpDeleteClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCluster{}, middleware.After)
}

func addOpDeleteParameterGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteParameterGroup{}, middleware.After)
}

func addOpDeleteSnapshotValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSnapshot{}, middleware.After)
}

func addOpDeleteSubnetGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSubnetGroup{}, middleware.After)
}

func addOpDeleteUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteUser{}, middleware.After)
}

func addOpDescribeParametersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeParameters{}, middleware.After)
}

func addOpDescribeUsersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeUsers{}, middleware.After)
}

func addOpFailoverShardValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpFailoverShard{}, middleware.After)
}

func addOpListAllowedNodeTypeUpdatesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAllowedNodeTypeUpdates{}, middleware.After)
}

func addOpListTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTags{}, middleware.After)
}

func addOpPurchaseReservedNodesOfferingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPurchaseReservedNodesOffering{}, middleware.After)
}

func addOpResetParameterGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpResetParameterGroup{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateACLValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateACL{}, middleware.After)
}

func addOpUpdateClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCluster{}, middleware.After)
}

func addOpUpdateParameterGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateParameterGroup{}, middleware.After)
}

func addOpUpdateSubnetGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSubnetGroup{}, middleware.After)
}

func addOpUpdateUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateUser{}, middleware.After)
}

func validateFilter(v *types.Filter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Filter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFilterList(v []types.Filter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FilterList"}
	for i := range v {
		if err := validateFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchUpdateClusterInput(v *BatchUpdateClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchUpdateClusterInput"}
	if v.ClusterNames == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterNames"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCopySnapshotInput(v *CopySnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CopySnapshotInput"}
	if v.SourceSnapshotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceSnapshotName"))
	}
	if v.TargetSnapshotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetSnapshotName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateACLInput(v *CreateACLInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateACLInput"}
	if v.ACLName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ACLName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateClusterInput(v *CreateClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateClusterInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.NodeType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodeType"))
	}
	if v.ACLName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ACLName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateParameterGroupInput(v *CreateParameterGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateParameterGroupInput"}
	if v.ParameterGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ParameterGroupName"))
	}
	if v.Family == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Family"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSnapshotInput(v *CreateSnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSnapshotInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.SnapshotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSubnetGroupInput(v *CreateSubnetGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSubnetGroupInput"}
	if v.SubnetGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetGroupName"))
	}
	if v.SubnetIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateUserInput(v *CreateUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateUserInput"}
	if v.UserName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserName"))
	}
	if v.AuthenticationMode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AuthenticationMode"))
	}
	if v.AccessString == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessString"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteACLInput(v *DeleteACLInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteACLInput"}
	if v.ACLName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ACLName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteClusterInput(v *DeleteClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteClusterInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteParameterGroupInput(v *DeleteParameterGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteParameterGroupInput"}
	if v.ParameterGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ParameterGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSnapshotInput(v *DeleteSnapshotInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSnapshotInput"}
	if v.SnapshotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSubnetGroupInput(v *DeleteSubnetGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSubnetGroupInput"}
	if v.SubnetGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteUserInput(v *DeleteUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteUserInput"}
	if v.UserName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeParametersInput(v *DescribeParametersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeParametersInput"}
	if v.ParameterGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ParameterGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeUsersInput(v *DescribeUsersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeUsersInput"}
	if v.Filters != nil {
		if err := validateFilterList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpFailoverShardInput(v *FailoverShardInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FailoverShardInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.ShardName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShardName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAllowedNodeTypeUpdatesInput(v *ListAllowedNodeTypeUpdatesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAllowedNodeTypeUpdatesInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsInput(v *ListTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPurchaseReservedNodesOfferingInput(v *PurchaseReservedNodesOfferingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PurchaseReservedNodesOfferingInput"}
	if v.ReservedNodesOfferingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReservedNodesOfferingId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpResetParameterGroupInput(v *ResetParameterGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResetParameterGroupInput"}
	if v.ParameterGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ParameterGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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

func validateOpUpdateACLInput(v *UpdateACLInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateACLInput"}
	if v.ACLName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ACLName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateClusterInput(v *UpdateClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateClusterInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateParameterGroupInput(v *UpdateParameterGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateParameterGroupInput"}
	if v.ParameterGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ParameterGroupName"))
	}
	if v.ParameterNameValues == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ParameterNameValues"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSubnetGroupInput(v *UpdateSubnetGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSubnetGroupInput"}
	if v.SubnetGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateUserInput(v *UpdateUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateUserInput"}
	if v.UserName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
