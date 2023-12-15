// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptography

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/paymentcryptography/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateAlias struct {
}

func (*validateOpCreateAlias) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAlias) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAliasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAliasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateKey struct {
}

func (*validateOpCreateKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAlias struct {
}

func (*validateOpDeleteAlias) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAlias) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAliasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAliasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteKey struct {
}

func (*validateOpDeleteKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExportKey struct {
}

func (*validateOpExportKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExportKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExportKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExportKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAlias struct {
}

func (*validateOpGetAlias) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAlias) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAliasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAliasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetKey struct {
}

func (*validateOpGetKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetParametersForExport struct {
}

func (*validateOpGetParametersForExport) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetParametersForExport) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetParametersForExportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetParametersForExportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetParametersForImport struct {
}

func (*validateOpGetParametersForImport) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetParametersForImport) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetParametersForImportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetParametersForImportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPublicKeyCertificate struct {
}

func (*validateOpGetPublicKeyCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPublicKeyCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPublicKeyCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPublicKeyCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpImportKey struct {
}

func (*validateOpImportKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpImportKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ImportKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpImportKeyInput(input); err != nil {
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

type validateOpRestoreKey struct {
}

func (*validateOpRestoreKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRestoreKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RestoreKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRestoreKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartKeyUsage struct {
}

func (*validateOpStartKeyUsage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartKeyUsage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartKeyUsageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartKeyUsageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopKeyUsage struct {
}

func (*validateOpStopKeyUsage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopKeyUsage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopKeyUsageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopKeyUsageInput(input); err != nil {
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

type validateOpUpdateAlias struct {
}

func (*validateOpUpdateAlias) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateAlias) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateAliasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateAliasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateAliasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAlias{}, middleware.After)
}

func addOpCreateKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateKey{}, middleware.After)
}

func addOpDeleteAliasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAlias{}, middleware.After)
}

func addOpDeleteKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteKey{}, middleware.After)
}

func addOpExportKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExportKey{}, middleware.After)
}

func addOpGetAliasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAlias{}, middleware.After)
}

func addOpGetKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetKey{}, middleware.After)
}

func addOpGetParametersForExportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetParametersForExport{}, middleware.After)
}

func addOpGetParametersForImportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetParametersForImport{}, middleware.After)
}

func addOpGetPublicKeyCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPublicKeyCertificate{}, middleware.After)
}

func addOpImportKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpImportKey{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpRestoreKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRestoreKey{}, middleware.After)
}

func addOpStartKeyUsageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartKeyUsage{}, middleware.After)
}

func addOpStopKeyUsageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopKeyUsage{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateAliasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateAlias{}, middleware.After)
}

func validateExportAttributes(v *types.ExportAttributes) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportAttributes"}
	if v.ExportDukptInitialKey != nil {
		if err := validateExportDukptInitialKey(v.ExportDukptInitialKey); err != nil {
			invalidParams.AddNested("ExportDukptInitialKey", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExportDukptInitialKey(v *types.ExportDukptInitialKey) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportDukptInitialKey"}
	if v.KeySerialNumber == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeySerialNumber"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExportKeyMaterial(v types.ExportKeyMaterial) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportKeyMaterial"}
	switch uv := v.(type) {
	case *types.ExportKeyMaterialMemberTr31KeyBlock:
		if err := validateExportTr31KeyBlock(&uv.Value); err != nil {
			invalidParams.AddNested("[Tr31KeyBlock]", err.(smithy.InvalidParamsError))
		}

	case *types.ExportKeyMaterialMemberTr34KeyBlock:
		if err := validateExportTr34KeyBlock(&uv.Value); err != nil {
			invalidParams.AddNested("[Tr34KeyBlock]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExportTr31KeyBlock(v *types.ExportTr31KeyBlock) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportTr31KeyBlock"}
	if v.WrappingKeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WrappingKeyIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExportTr34KeyBlock(v *types.ExportTr34KeyBlock) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportTr34KeyBlock"}
	if v.CertificateAuthorityPublicKeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityPublicKeyIdentifier"))
	}
	if v.WrappingKeyCertificate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WrappingKeyCertificate"))
	}
	if v.ExportToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExportToken"))
	}
	if len(v.KeyBlockFormat) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("KeyBlockFormat"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateImportKeyMaterial(v types.ImportKeyMaterial) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportKeyMaterial"}
	switch uv := v.(type) {
	case *types.ImportKeyMaterialMemberRootCertificatePublicKey:
		if err := validateRootCertificatePublicKey(&uv.Value); err != nil {
			invalidParams.AddNested("[RootCertificatePublicKey]", err.(smithy.InvalidParamsError))
		}

	case *types.ImportKeyMaterialMemberTr31KeyBlock:
		if err := validateImportTr31KeyBlock(&uv.Value); err != nil {
			invalidParams.AddNested("[Tr31KeyBlock]", err.(smithy.InvalidParamsError))
		}

	case *types.ImportKeyMaterialMemberTr34KeyBlock:
		if err := validateImportTr34KeyBlock(&uv.Value); err != nil {
			invalidParams.AddNested("[Tr34KeyBlock]", err.(smithy.InvalidParamsError))
		}

	case *types.ImportKeyMaterialMemberTrustedCertificatePublicKey:
		if err := validateTrustedCertificatePublicKey(&uv.Value); err != nil {
			invalidParams.AddNested("[TrustedCertificatePublicKey]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateImportTr31KeyBlock(v *types.ImportTr31KeyBlock) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportTr31KeyBlock"}
	if v.WrappingKeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WrappingKeyIdentifier"))
	}
	if v.WrappedKeyBlock == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WrappedKeyBlock"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateImportTr34KeyBlock(v *types.ImportTr34KeyBlock) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportTr34KeyBlock"}
	if v.CertificateAuthorityPublicKeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityPublicKeyIdentifier"))
	}
	if v.SigningKeyCertificate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SigningKeyCertificate"))
	}
	if v.ImportToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ImportToken"))
	}
	if v.WrappedKeyBlock == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WrappedKeyBlock"))
	}
	if len(v.KeyBlockFormat) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("KeyBlockFormat"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateKeyAttributes(v *types.KeyAttributes) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "KeyAttributes"}
	if len(v.KeyUsage) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("KeyUsage"))
	}
	if len(v.KeyClass) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("KeyClass"))
	}
	if len(v.KeyAlgorithm) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("KeyAlgorithm"))
	}
	if v.KeyModesOfUse == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyModesOfUse"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRootCertificatePublicKey(v *types.RootCertificatePublicKey) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RootCertificatePublicKey"}
	if v.KeyAttributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyAttributes"))
	} else if v.KeyAttributes != nil {
		if err := validateKeyAttributes(v.KeyAttributes); err != nil {
			invalidParams.AddNested("KeyAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if v.PublicKeyCertificate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PublicKeyCertificate"))
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

func validateTrustedCertificatePublicKey(v *types.TrustedCertificatePublicKey) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TrustedCertificatePublicKey"}
	if v.KeyAttributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyAttributes"))
	} else if v.KeyAttributes != nil {
		if err := validateKeyAttributes(v.KeyAttributes); err != nil {
			invalidParams.AddNested("KeyAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if v.PublicKeyCertificate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PublicKeyCertificate"))
	}
	if v.CertificateAuthorityPublicKeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityPublicKeyIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAliasInput(v *CreateAliasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAliasInput"}
	if v.AliasName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AliasName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateKeyInput(v *CreateKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateKeyInput"}
	if v.KeyAttributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyAttributes"))
	} else if v.KeyAttributes != nil {
		if err := validateKeyAttributes(v.KeyAttributes); err != nil {
			invalidParams.AddNested("KeyAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if v.Exportable == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Exportable"))
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

func validateOpDeleteAliasInput(v *DeleteAliasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAliasInput"}
	if v.AliasName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AliasName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteKeyInput(v *DeleteKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteKeyInput"}
	if v.KeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpExportKeyInput(v *ExportKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportKeyInput"}
	if v.KeyMaterial == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyMaterial"))
	} else if v.KeyMaterial != nil {
		if err := validateExportKeyMaterial(v.KeyMaterial); err != nil {
			invalidParams.AddNested("KeyMaterial", err.(smithy.InvalidParamsError))
		}
	}
	if v.ExportKeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExportKeyIdentifier"))
	}
	if v.ExportAttributes != nil {
		if err := validateExportAttributes(v.ExportAttributes); err != nil {
			invalidParams.AddNested("ExportAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAliasInput(v *GetAliasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAliasInput"}
	if v.AliasName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AliasName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetKeyInput(v *GetKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetKeyInput"}
	if v.KeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetParametersForExportInput(v *GetParametersForExportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetParametersForExportInput"}
	if len(v.KeyMaterialType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("KeyMaterialType"))
	}
	if len(v.SigningKeyAlgorithm) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("SigningKeyAlgorithm"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetParametersForImportInput(v *GetParametersForImportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetParametersForImportInput"}
	if len(v.KeyMaterialType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("KeyMaterialType"))
	}
	if len(v.WrappingKeyAlgorithm) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("WrappingKeyAlgorithm"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPublicKeyCertificateInput(v *GetPublicKeyCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPublicKeyCertificateInput"}
	if v.KeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpImportKeyInput(v *ImportKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportKeyInput"}
	if v.KeyMaterial == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyMaterial"))
	} else if v.KeyMaterial != nil {
		if err := validateImportKeyMaterial(v.KeyMaterial); err != nil {
			invalidParams.AddNested("KeyMaterial", err.(smithy.InvalidParamsError))
		}
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

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRestoreKeyInput(v *RestoreKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RestoreKeyInput"}
	if v.KeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartKeyUsageInput(v *StartKeyUsageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartKeyUsageInput"}
	if v.KeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopKeyUsageInput(v *StopKeyUsageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopKeyUsageInput"}
	if v.KeyIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyIdentifier"))
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

func validateOpUpdateAliasInput(v *UpdateAliasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateAliasInput"}
	if v.AliasName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AliasName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
