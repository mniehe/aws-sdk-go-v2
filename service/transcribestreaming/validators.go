// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribestreaming

import (
	"context"
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/transcribestreaming/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpStartCallAnalyticsStreamTranscription struct {
}

func (*validateOpStartCallAnalyticsStreamTranscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartCallAnalyticsStreamTranscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartCallAnalyticsStreamTranscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartCallAnalyticsStreamTranscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartMedicalStreamTranscription struct {
}

func (*validateOpStartMedicalStreamTranscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartMedicalStreamTranscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartMedicalStreamTranscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartMedicalStreamTranscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartStreamTranscription struct {
}

func (*validateOpStartStreamTranscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartStreamTranscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartStreamTranscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartStreamTranscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpStartCallAnalyticsStreamTranscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartCallAnalyticsStreamTranscription{}, middleware.After)
}

func addOpStartMedicalStreamTranscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartMedicalStreamTranscription{}, middleware.After)
}

func addOpStartStreamTranscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartStreamTranscription{}, middleware.After)
}

func validateAudioStream(v types.AudioStream) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AudioStream"}
	switch uv := v.(type) {
	case *types.AudioStreamMemberConfigurationEvent:
		if err := validateConfigurationEvent(&uv.Value); err != nil {
			invalidParams.AddNested("[ConfigurationEvent]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateChannelDefinition(v *types.ChannelDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ChannelDefinition"}
	if len(v.ParticipantRole) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ParticipantRole"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateChannelDefinitions(v []types.ChannelDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ChannelDefinitions"}
	for i := range v {
		if err := validateChannelDefinition(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateConfigurationEvent(v *types.ConfigurationEvent) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConfigurationEvent"}
	if v.ChannelDefinitions != nil {
		if err := validateChannelDefinitions(v.ChannelDefinitions); err != nil {
			invalidParams.AddNested("ChannelDefinitions", err.(smithy.InvalidParamsError))
		}
	}
	if v.PostCallAnalyticsSettings != nil {
		if err := validatePostCallAnalyticsSettings(v.PostCallAnalyticsSettings); err != nil {
			invalidParams.AddNested("PostCallAnalyticsSettings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePostCallAnalyticsSettings(v *types.PostCallAnalyticsSettings) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PostCallAnalyticsSettings"}
	if v.OutputLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputLocation"))
	}
	if v.DataAccessRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataAccessRoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartCallAnalyticsStreamTranscriptionInput(v *StartCallAnalyticsStreamTranscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartCallAnalyticsStreamTranscriptionInput"}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if v.MediaSampleRateHertz == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MediaSampleRateHertz"))
	}
	if len(v.MediaEncoding) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("MediaEncoding"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartMedicalStreamTranscriptionInput(v *StartMedicalStreamTranscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartMedicalStreamTranscriptionInput"}
	if len(v.LanguageCode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LanguageCode"))
	}
	if v.MediaSampleRateHertz == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MediaSampleRateHertz"))
	}
	if len(v.MediaEncoding) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("MediaEncoding"))
	}
	if len(v.Specialty) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Specialty"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartStreamTranscriptionInput(v *StartStreamTranscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartStreamTranscriptionInput"}
	if v.MediaSampleRateHertz == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MediaSampleRateHertz"))
	}
	if len(v.MediaEncoding) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("MediaEncoding"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
