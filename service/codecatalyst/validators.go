// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/codecatalyst/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateAccessToken struct {
}

func (*validateOpCreateAccessToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAccessToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAccessTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAccessTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDevEnvironment struct {
}

func (*validateOpCreateDevEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDevEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDevEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDevEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateProject struct {
}

func (*validateOpCreateProject) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateProject) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateProjectInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateProjectInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSourceRepositoryBranch struct {
}

func (*validateOpCreateSourceRepositoryBranch) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSourceRepositoryBranch) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSourceRepositoryBranchInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSourceRepositoryBranchInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAccessToken struct {
}

func (*validateOpDeleteAccessToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAccessToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAccessTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAccessTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDevEnvironment struct {
}

func (*validateOpDeleteDevEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDevEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDevEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDevEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDevEnvironment struct {
}

func (*validateOpGetDevEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDevEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDevEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDevEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetProject struct {
}

func (*validateOpGetProject) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetProject) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetProjectInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetProjectInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSourceRepositoryCloneUrls struct {
}

func (*validateOpGetSourceRepositoryCloneUrls) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSourceRepositoryCloneUrls) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSourceRepositoryCloneUrlsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSourceRepositoryCloneUrlsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSpace struct {
}

func (*validateOpGetSpace) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSpace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSpaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSpaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSubscription struct {
}

func (*validateOpGetSubscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSubscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSubscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSubscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListDevEnvironments struct {
}

func (*validateOpListDevEnvironments) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListDevEnvironments) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListDevEnvironmentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListDevEnvironmentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListEventLogs struct {
}

func (*validateOpListEventLogs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListEventLogs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListEventLogsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListEventLogsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListProjects struct {
}

func (*validateOpListProjects) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListProjects) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListProjectsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListProjectsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSourceRepositories struct {
}

func (*validateOpListSourceRepositories) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSourceRepositories) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSourceRepositoriesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSourceRepositoriesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSourceRepositoryBranches struct {
}

func (*validateOpListSourceRepositoryBranches) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSourceRepositoryBranches) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSourceRepositoryBranchesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSourceRepositoryBranchesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartDevEnvironment struct {
}

func (*validateOpStartDevEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartDevEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartDevEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartDevEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartDevEnvironmentSession struct {
}

func (*validateOpStartDevEnvironmentSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartDevEnvironmentSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartDevEnvironmentSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartDevEnvironmentSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopDevEnvironment struct {
}

func (*validateOpStopDevEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopDevEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopDevEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopDevEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopDevEnvironmentSession struct {
}

func (*validateOpStopDevEnvironmentSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopDevEnvironmentSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopDevEnvironmentSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopDevEnvironmentSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateDevEnvironment struct {
}

func (*validateOpUpdateDevEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDevEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDevEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDevEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateAccessTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAccessToken{}, middleware.After)
}

func addOpCreateDevEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDevEnvironment{}, middleware.After)
}

func addOpCreateProjectValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateProject{}, middleware.After)
}

func addOpCreateSourceRepositoryBranchValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSourceRepositoryBranch{}, middleware.After)
}

func addOpDeleteAccessTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAccessToken{}, middleware.After)
}

func addOpDeleteDevEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDevEnvironment{}, middleware.After)
}

func addOpGetDevEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDevEnvironment{}, middleware.After)
}

func addOpGetProjectValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetProject{}, middleware.After)
}

func addOpGetSourceRepositoryCloneUrlsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSourceRepositoryCloneUrls{}, middleware.After)
}

func addOpGetSpaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSpace{}, middleware.After)
}

func addOpGetSubscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSubscription{}, middleware.After)
}

func addOpListDevEnvironmentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListDevEnvironments{}, middleware.After)
}

func addOpListEventLogsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListEventLogs{}, middleware.After)
}

func addOpListProjectsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListProjects{}, middleware.After)
}

func addOpListSourceRepositoriesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSourceRepositories{}, middleware.After)
}

func addOpListSourceRepositoryBranchesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSourceRepositoryBranches{}, middleware.After)
}

func addOpStartDevEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartDevEnvironment{}, middleware.After)
}

func addOpStartDevEnvironmentSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartDevEnvironmentSession{}, middleware.After)
}

func addOpStopDevEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopDevEnvironment{}, middleware.After)
}

func addOpStopDevEnvironmentSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopDevEnvironmentSession{}, middleware.After)
}

func addOpUpdateDevEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDevEnvironment{}, middleware.After)
}

func validateDevEnvironmentSessionConfiguration(v *types.DevEnvironmentSessionConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DevEnvironmentSessionConfiguration"}
	if len(v.SessionType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("SessionType"))
	}
	if v.ExecuteCommandSessionConfiguration != nil {
		if err := validateExecuteCommandSessionConfiguration(v.ExecuteCommandSessionConfiguration); err != nil {
			invalidParams.AddNested("ExecuteCommandSessionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExecuteCommandSessionConfiguration(v *types.ExecuteCommandSessionConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExecuteCommandSessionConfiguration"}
	if v.Command == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Command"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFilter(v *types.Filter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Filter"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
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

func validateFilters(v []types.Filter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Filters"}
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

func validatePersistentStorageConfiguration(v *types.PersistentStorageConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PersistentStorageConfiguration"}
	if v.SizeInGiB == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SizeInGiB"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProjectListFilter(v *types.ProjectListFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProjectListFilter"}
	if len(v.Key) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
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

func validateProjectListFilters(v []types.ProjectListFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProjectListFilters"}
	for i := range v {
		if err := validateProjectListFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRepositoriesInput(v []types.RepositoryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RepositoriesInput"}
	for i := range v {
		if err := validateRepositoryInput(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRepositoryInput(v *types.RepositoryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RepositoryInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAccessTokenInput(v *CreateAccessTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAccessTokenInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDevEnvironmentInput(v *CreateDevEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDevEnvironmentInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.Repositories != nil {
		if err := validateRepositoriesInput(v.Repositories); err != nil {
			invalidParams.AddNested("Repositories", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.InstanceType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceType"))
	}
	if v.PersistentStorage == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PersistentStorage"))
	} else if v.PersistentStorage != nil {
		if err := validatePersistentStorageConfiguration(v.PersistentStorage); err != nil {
			invalidParams.AddNested("PersistentStorage", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateProjectInput(v *CreateProjectInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateProjectInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.DisplayName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DisplayName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSourceRepositoryBranchInput(v *CreateSourceRepositoryBranchInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSourceRepositoryBranchInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.SourceRepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceRepositoryName"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAccessTokenInput(v *DeleteAccessTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAccessTokenInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDevEnvironmentInput(v *DeleteDevEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDevEnvironmentInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDevEnvironmentInput(v *GetDevEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDevEnvironmentInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetProjectInput(v *GetProjectInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetProjectInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSourceRepositoryCloneUrlsInput(v *GetSourceRepositoryCloneUrlsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSourceRepositoryCloneUrlsInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.SourceRepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceRepositoryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSpaceInput(v *GetSpaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSpaceInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSubscriptionInput(v *GetSubscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSubscriptionInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListDevEnvironmentsInput(v *ListDevEnvironmentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListDevEnvironmentsInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.Filters != nil {
		if err := validateFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListEventLogsInput(v *ListEventLogsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListEventLogsInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListProjectsInput(v *ListProjectsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListProjectsInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.Filters != nil {
		if err := validateProjectListFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSourceRepositoriesInput(v *ListSourceRepositoriesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSourceRepositoriesInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSourceRepositoryBranchesInput(v *ListSourceRepositoryBranchesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSourceRepositoryBranchesInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.SourceRepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceRepositoryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartDevEnvironmentInput(v *StartDevEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartDevEnvironmentInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartDevEnvironmentSessionInput(v *StartDevEnvironmentSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartDevEnvironmentSessionInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.SessionConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionConfiguration"))
	} else if v.SessionConfiguration != nil {
		if err := validateDevEnvironmentSessionConfiguration(v.SessionConfiguration); err != nil {
			invalidParams.AddNested("SessionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopDevEnvironmentInput(v *StopDevEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopDevEnvironmentInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopDevEnvironmentSessionInput(v *StopDevEnvironmentSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopDevEnvironmentSessionInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateDevEnvironmentInput(v *UpdateDevEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDevEnvironmentInput"}
	if v.SpaceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpaceName"))
	}
	if v.ProjectName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
