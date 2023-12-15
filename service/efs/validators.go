// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/efs/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateAccessPoint struct {
}

func (*validateOpCreateAccessPoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAccessPoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAccessPointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateFileSystem struct {
}

func (*validateOpCreateFileSystem) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateFileSystemInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateMountTarget struct {
}

func (*validateOpCreateMountTarget) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMountTarget) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMountTargetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMountTargetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateReplicationConfiguration struct {
}

func (*validateOpCreateReplicationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateReplicationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateReplicationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateReplicationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateTags struct {
}

func (*validateOpCreateTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAccessPoint struct {
}

func (*validateOpDeleteAccessPoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAccessPoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAccessPointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteFileSystem struct {
}

func (*validateOpDeleteFileSystem) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteFileSystemInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteFileSystemPolicy struct {
}

func (*validateOpDeleteFileSystemPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteFileSystemPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteFileSystemPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteFileSystemPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMountTarget struct {
}

func (*validateOpDeleteMountTarget) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMountTarget) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMountTargetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMountTargetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteReplicationConfiguration struct {
}

func (*validateOpDeleteReplicationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteReplicationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteReplicationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteReplicationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTags struct {
}

func (*validateOpDeleteTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeBackupPolicy struct {
}

func (*validateOpDescribeBackupPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeBackupPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeBackupPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeBackupPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeFileSystemPolicy struct {
}

func (*validateOpDescribeFileSystemPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeFileSystemPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeFileSystemPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeFileSystemPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeLifecycleConfiguration struct {
}

func (*validateOpDescribeLifecycleConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeLifecycleConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeLifecycleConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeLifecycleConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeMountTargetSecurityGroups struct {
}

func (*validateOpDescribeMountTargetSecurityGroups) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeMountTargetSecurityGroups) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeMountTargetSecurityGroupsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeMountTargetSecurityGroupsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeTags struct {
}

func (*validateOpDescribeTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpModifyMountTargetSecurityGroups struct {
}

func (*validateOpModifyMountTargetSecurityGroups) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpModifyMountTargetSecurityGroups) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ModifyMountTargetSecurityGroupsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpModifyMountTargetSecurityGroupsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutAccountPreferences struct {
}

func (*validateOpPutAccountPreferences) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutAccountPreferences) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutAccountPreferencesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutAccountPreferencesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutBackupPolicy struct {
}

func (*validateOpPutBackupPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutBackupPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutBackupPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutBackupPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutFileSystemPolicy struct {
}

func (*validateOpPutFileSystemPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutFileSystemPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutFileSystemPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutFileSystemPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutLifecycleConfiguration struct {
}

func (*validateOpPutLifecycleConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutLifecycleConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutLifecycleConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutLifecycleConfigurationInput(input); err != nil {
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

type validateOpUpdateFileSystem struct {
}

func (*validateOpUpdateFileSystem) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateFileSystemInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateFileSystemProtection struct {
}

func (*validateOpUpdateFileSystemProtection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateFileSystemProtection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateFileSystemProtectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateFileSystemProtectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateAccessPointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAccessPoint{}, middleware.After)
}

func addOpCreateFileSystemValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateFileSystem{}, middleware.After)
}

func addOpCreateMountTargetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMountTarget{}, middleware.After)
}

func addOpCreateReplicationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateReplicationConfiguration{}, middleware.After)
}

func addOpCreateTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateTags{}, middleware.After)
}

func addOpDeleteAccessPointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAccessPoint{}, middleware.After)
}

func addOpDeleteFileSystemValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteFileSystem{}, middleware.After)
}

func addOpDeleteFileSystemPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteFileSystemPolicy{}, middleware.After)
}

func addOpDeleteMountTargetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMountTarget{}, middleware.After)
}

func addOpDeleteReplicationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteReplicationConfiguration{}, middleware.After)
}

func addOpDeleteTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTags{}, middleware.After)
}

func addOpDescribeBackupPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeBackupPolicy{}, middleware.After)
}

func addOpDescribeFileSystemPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeFileSystemPolicy{}, middleware.After)
}

func addOpDescribeLifecycleConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeLifecycleConfiguration{}, middleware.After)
}

func addOpDescribeMountTargetSecurityGroupsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeMountTargetSecurityGroups{}, middleware.After)
}

func addOpDescribeTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeTags{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpModifyMountTargetSecurityGroupsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpModifyMountTargetSecurityGroups{}, middleware.After)
}

func addOpPutAccountPreferencesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutAccountPreferences{}, middleware.After)
}

func addOpPutBackupPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutBackupPolicy{}, middleware.After)
}

func addOpPutFileSystemPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutFileSystemPolicy{}, middleware.After)
}

func addOpPutLifecycleConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutLifecycleConfiguration{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateFileSystemValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateFileSystem{}, middleware.After)
}

func addOpUpdateFileSystemProtectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateFileSystemProtection{}, middleware.After)
}

func validateBackupPolicy(v *types.BackupPolicy) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BackupPolicy"}
	if len(v.Status) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Status"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreationInfo(v *types.CreationInfo) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreationInfo"}
	if v.OwnerUid == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OwnerUid"))
	}
	if v.OwnerGid == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OwnerGid"))
	}
	if v.Permissions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Permissions"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePosixUser(v *types.PosixUser) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PosixUser"}
	if v.Uid == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Uid"))
	}
	if v.Gid == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Gid"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRootDirectory(v *types.RootDirectory) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RootDirectory"}
	if v.CreationInfo != nil {
		if err := validateCreationInfo(v.CreationInfo); err != nil {
			invalidParams.AddNested("CreationInfo", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTags(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tags"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAccessPointInput(v *CreateAccessPointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAccessPointInput"}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if v.PosixUser != nil {
		if err := validatePosixUser(v.PosixUser); err != nil {
			invalidParams.AddNested("PosixUser", err.(smithy.InvalidParamsError))
		}
	}
	if v.RootDirectory != nil {
		if err := validateRootDirectory(v.RootDirectory); err != nil {
			invalidParams.AddNested("RootDirectory", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateFileSystemInput(v *CreateFileSystemInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateFileSystemInput"}
	if v.CreationToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CreationToken"))
	}
	if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateMountTargetInput(v *CreateMountTargetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMountTargetInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if v.SubnetId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateReplicationConfigurationInput(v *CreateReplicationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateReplicationConfigurationInput"}
	if v.SourceFileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceFileSystemId"))
	}
	if v.Destinations == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Destinations"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateTagsInput(v *CreateTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateTagsInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAccessPointInput(v *DeleteAccessPointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAccessPointInput"}
	if v.AccessPointId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessPointId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteFileSystemInput(v *DeleteFileSystemInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteFileSystemInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteFileSystemPolicyInput(v *DeleteFileSystemPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteFileSystemPolicyInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMountTargetInput(v *DeleteMountTargetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMountTargetInput"}
	if v.MountTargetId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MountTargetId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteReplicationConfigurationInput(v *DeleteReplicationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteReplicationConfigurationInput"}
	if v.SourceFileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceFileSystemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTagsInput(v *DeleteTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTagsInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
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

func validateOpDescribeBackupPolicyInput(v *DescribeBackupPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeBackupPolicyInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeFileSystemPolicyInput(v *DescribeFileSystemPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeFileSystemPolicyInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeLifecycleConfigurationInput(v *DescribeLifecycleConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeLifecycleConfigurationInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeMountTargetSecurityGroupsInput(v *DescribeMountTargetSecurityGroupsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeMountTargetSecurityGroupsInput"}
	if v.MountTargetId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MountTargetId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeTagsInput(v *DescribeTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeTagsInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpModifyMountTargetSecurityGroupsInput(v *ModifyMountTargetSecurityGroupsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ModifyMountTargetSecurityGroupsInput"}
	if v.MountTargetId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MountTargetId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutAccountPreferencesInput(v *PutAccountPreferencesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutAccountPreferencesInput"}
	if len(v.ResourceIdType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceIdType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutBackupPolicyInput(v *PutBackupPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutBackupPolicyInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if v.BackupPolicy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackupPolicy"))
	} else if v.BackupPolicy != nil {
		if err := validateBackupPolicy(v.BackupPolicy); err != nil {
			invalidParams.AddNested("BackupPolicy", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutFileSystemPolicyInput(v *PutFileSystemPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutFileSystemPolicyInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if v.Policy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Policy"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutLifecycleConfigurationInput(v *PutLifecycleConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutLifecycleConfigurationInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if v.LifecyclePolicies == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LifecyclePolicies"))
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
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
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
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
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

func validateOpUpdateFileSystemInput(v *UpdateFileSystemInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateFileSystemInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateFileSystemProtectionInput(v *UpdateFileSystemProtectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateFileSystemProtectionInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
