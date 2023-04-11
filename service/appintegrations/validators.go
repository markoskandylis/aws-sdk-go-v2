// Code generated by smithy-go-codegen DO NOT EDIT.

package appintegrations

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/appintegrations/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateDataIntegration struct {
}

func (*validateOpCreateDataIntegration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDataIntegration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDataIntegrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDataIntegrationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateEventIntegration struct {
}

func (*validateOpCreateEventIntegration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateEventIntegration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateEventIntegrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateEventIntegrationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDataIntegration struct {
}

func (*validateOpDeleteDataIntegration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDataIntegration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDataIntegrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDataIntegrationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEventIntegration struct {
}

func (*validateOpDeleteEventIntegration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEventIntegration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEventIntegrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEventIntegrationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDataIntegration struct {
}

func (*validateOpGetDataIntegration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDataIntegration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDataIntegrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDataIntegrationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetEventIntegration struct {
}

func (*validateOpGetEventIntegration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetEventIntegration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetEventIntegrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetEventIntegrationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListDataIntegrationAssociations struct {
}

func (*validateOpListDataIntegrationAssociations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListDataIntegrationAssociations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListDataIntegrationAssociationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListDataIntegrationAssociationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListEventIntegrationAssociations struct {
}

func (*validateOpListEventIntegrationAssociations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListEventIntegrationAssociations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListEventIntegrationAssociationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListEventIntegrationAssociationsInput(input); err != nil {
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

type validateOpUpdateDataIntegration struct {
}

func (*validateOpUpdateDataIntegration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDataIntegration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDataIntegrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDataIntegrationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateEventIntegration struct {
}

func (*validateOpUpdateEventIntegration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateEventIntegration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateEventIntegrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateEventIntegrationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateDataIntegrationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDataIntegration{}, middleware.After)
}

func addOpCreateEventIntegrationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateEventIntegration{}, middleware.After)
}

func addOpDeleteDataIntegrationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDataIntegration{}, middleware.After)
}

func addOpDeleteEventIntegrationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteEventIntegration{}, middleware.After)
}

func addOpGetDataIntegrationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDataIntegration{}, middleware.After)
}

func addOpGetEventIntegrationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetEventIntegration{}, middleware.After)
}

func addOpListDataIntegrationAssociationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListDataIntegrationAssociations{}, middleware.After)
}

func addOpListEventIntegrationAssociationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListEventIntegrationAssociations{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateDataIntegrationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDataIntegration{}, middleware.After)
}

func addOpUpdateEventIntegrationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateEventIntegration{}, middleware.After)
}

func validateEventFilter(v *types.EventFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EventFilter"}
	if v.Source == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Source"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFileConfiguration(v *types.FileConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FileConfiguration"}
	if v.Folders == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Folders"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateScheduleConfiguration(v *types.ScheduleConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ScheduleConfiguration"}
	if v.ScheduleExpression == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduleExpression"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDataIntegrationInput(v *CreateDataIntegrationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDataIntegrationInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.KmsKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KmsKey"))
	}
	if v.SourceURI == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceURI"))
	}
	if v.ScheduleConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduleConfig"))
	} else if v.ScheduleConfig != nil {
		if err := validateScheduleConfiguration(v.ScheduleConfig); err != nil {
			invalidParams.AddNested("ScheduleConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.FileConfiguration != nil {
		if err := validateFileConfiguration(v.FileConfiguration); err != nil {
			invalidParams.AddNested("FileConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateEventIntegrationInput(v *CreateEventIntegrationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateEventIntegrationInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.EventFilter == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventFilter"))
	} else if v.EventFilter != nil {
		if err := validateEventFilter(v.EventFilter); err != nil {
			invalidParams.AddNested("EventFilter", err.(smithy.InvalidParamsError))
		}
	}
	if v.EventBridgeBus == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventBridgeBus"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDataIntegrationInput(v *DeleteDataIntegrationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDataIntegrationInput"}
	if v.DataIntegrationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataIntegrationIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEventIntegrationInput(v *DeleteEventIntegrationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEventIntegrationInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDataIntegrationInput(v *GetDataIntegrationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDataIntegrationInput"}
	if v.Identifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Identifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetEventIntegrationInput(v *GetEventIntegrationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetEventIntegrationInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListDataIntegrationAssociationsInput(v *ListDataIntegrationAssociationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListDataIntegrationAssociationsInput"}
	if v.DataIntegrationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataIntegrationIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListEventIntegrationAssociationsInput(v *ListEventIntegrationAssociationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListEventIntegrationAssociationsInput"}
	if v.EventIntegrationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventIntegrationName"))
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

func validateOpUpdateDataIntegrationInput(v *UpdateDataIntegrationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDataIntegrationInput"}
	if v.Identifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Identifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateEventIntegrationInput(v *UpdateEventIntegrationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateEventIntegrationInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
