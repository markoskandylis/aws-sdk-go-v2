// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamwrite

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateBatchLoadTask struct {
}

func (*validateOpCreateBatchLoadTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateBatchLoadTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateBatchLoadTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateBatchLoadTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDatabase struct {
}

func (*validateOpCreateDatabase) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDatabase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDatabaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDatabaseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateTable struct {
}

func (*validateOpCreateTable) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateTable) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateTableInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateTableInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDatabase struct {
}

func (*validateOpDeleteDatabase) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDatabase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDatabaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDatabaseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTable struct {
}

func (*validateOpDeleteTable) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTable) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTableInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTableInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeBatchLoadTask struct {
}

func (*validateOpDescribeBatchLoadTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeBatchLoadTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeBatchLoadTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeBatchLoadTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeDatabase struct {
}

func (*validateOpDescribeDatabase) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeDatabase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeDatabaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeDatabaseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeTable struct {
}

func (*validateOpDescribeTable) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeTable) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeTableInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeTableInput(input); err != nil {
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

type validateOpResumeBatchLoadTask struct {
}

func (*validateOpResumeBatchLoadTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpResumeBatchLoadTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ResumeBatchLoadTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpResumeBatchLoadTaskInput(input); err != nil {
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

type validateOpUpdateDatabase struct {
}

func (*validateOpUpdateDatabase) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDatabase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDatabaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDatabaseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateTable struct {
}

func (*validateOpUpdateTable) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateTable) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateTableInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateTableInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpWriteRecords struct {
}

func (*validateOpWriteRecords) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpWriteRecords) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*WriteRecordsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpWriteRecordsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateBatchLoadTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateBatchLoadTask{}, middleware.After)
}

func addOpCreateDatabaseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDatabase{}, middleware.After)
}

func addOpCreateTableValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateTable{}, middleware.After)
}

func addOpDeleteDatabaseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDatabase{}, middleware.After)
}

func addOpDeleteTableValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTable{}, middleware.After)
}

func addOpDescribeBatchLoadTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeBatchLoadTask{}, middleware.After)
}

func addOpDescribeDatabaseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeDatabase{}, middleware.After)
}

func addOpDescribeTableValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeTable{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpResumeBatchLoadTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpResumeBatchLoadTask{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateDatabaseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDatabase{}, middleware.After)
}

func addOpUpdateTableValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateTable{}, middleware.After)
}

func addOpWriteRecordsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpWriteRecords{}, middleware.After)
}

func validateDataModel(v *types.DataModel) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DataModel"}
	if v.DimensionMappings == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DimensionMappings"))
	}
	if v.MultiMeasureMappings != nil {
		if err := validateMultiMeasureMappings(v.MultiMeasureMappings); err != nil {
			invalidParams.AddNested("MultiMeasureMappings", err.(smithy.InvalidParamsError))
		}
	}
	if v.MixedMeasureMappings != nil {
		if err := validateMixedMeasureMappingList(v.MixedMeasureMappings); err != nil {
			invalidParams.AddNested("MixedMeasureMappings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDataModelConfiguration(v *types.DataModelConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DataModelConfiguration"}
	if v.DataModel != nil {
		if err := validateDataModel(v.DataModel); err != nil {
			invalidParams.AddNested("DataModel", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDataSourceConfiguration(v *types.DataSourceConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DataSourceConfiguration"}
	if v.DataSourceS3Configuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSourceS3Configuration"))
	} else if v.DataSourceS3Configuration != nil {
		if err := validateDataSourceS3Configuration(v.DataSourceS3Configuration); err != nil {
			invalidParams.AddNested("DataSourceS3Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.DataFormat) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("DataFormat"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDataSourceS3Configuration(v *types.DataSourceS3Configuration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DataSourceS3Configuration"}
	if v.BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDimension(v *types.Dimension) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Dimension"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
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

func validateDimensions(v []types.Dimension) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Dimensions"}
	for i := range v {
		if err := validateDimension(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMagneticStoreWriteProperties(v *types.MagneticStoreWriteProperties) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MagneticStoreWriteProperties"}
	if v.EnableMagneticStoreWrites == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnableMagneticStoreWrites"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMeasureValue(v *types.MeasureValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MeasureValue"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
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

func validateMeasureValues(v []types.MeasureValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MeasureValues"}
	for i := range v {
		if err := validateMeasureValue(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMixedMeasureMapping(v *types.MixedMeasureMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MixedMeasureMapping"}
	if len(v.MeasureValueType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("MeasureValueType"))
	}
	if v.MultiMeasureAttributeMappings != nil {
		if err := validateMultiMeasureAttributeMappingList(v.MultiMeasureAttributeMappings); err != nil {
			invalidParams.AddNested("MultiMeasureAttributeMappings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMixedMeasureMappingList(v []types.MixedMeasureMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MixedMeasureMappingList"}
	for i := range v {
		if err := validateMixedMeasureMapping(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMultiMeasureAttributeMapping(v *types.MultiMeasureAttributeMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MultiMeasureAttributeMapping"}
	if v.SourceColumn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceColumn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMultiMeasureAttributeMappingList(v []types.MultiMeasureAttributeMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MultiMeasureAttributeMappingList"}
	for i := range v {
		if err := validateMultiMeasureAttributeMapping(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMultiMeasureMappings(v *types.MultiMeasureMappings) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MultiMeasureMappings"}
	if v.MultiMeasureAttributeMappings == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MultiMeasureAttributeMappings"))
	} else if v.MultiMeasureAttributeMappings != nil {
		if err := validateMultiMeasureAttributeMappingList(v.MultiMeasureAttributeMappings); err != nil {
			invalidParams.AddNested("MultiMeasureAttributeMappings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRecord(v *types.Record) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Record"}
	if v.Dimensions != nil {
		if err := validateDimensions(v.Dimensions); err != nil {
			invalidParams.AddNested("Dimensions", err.(smithy.InvalidParamsError))
		}
	}
	if v.MeasureValues != nil {
		if err := validateMeasureValues(v.MeasureValues); err != nil {
			invalidParams.AddNested("MeasureValues", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRecords(v []types.Record) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Records"}
	for i := range v {
		if err := validateRecord(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateReportConfiguration(v *types.ReportConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ReportConfiguration"}
	if v.ReportS3Configuration != nil {
		if err := validateReportS3Configuration(v.ReportS3Configuration); err != nil {
			invalidParams.AddNested("ReportS3Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateReportS3Configuration(v *types.ReportS3Configuration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ReportS3Configuration"}
	if v.BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRetentionProperties(v *types.RetentionProperties) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetentionProperties"}
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

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
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

func validateOpCreateBatchLoadTaskInput(v *CreateBatchLoadTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBatchLoadTaskInput"}
	if v.DataModelConfiguration != nil {
		if err := validateDataModelConfiguration(v.DataModelConfiguration); err != nil {
			invalidParams.AddNested("DataModelConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.DataSourceConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSourceConfiguration"))
	} else if v.DataSourceConfiguration != nil {
		if err := validateDataSourceConfiguration(v.DataSourceConfiguration); err != nil {
			invalidParams.AddNested("DataSourceConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.ReportConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportConfiguration"))
	} else if v.ReportConfiguration != nil {
		if err := validateReportConfiguration(v.ReportConfiguration); err != nil {
			invalidParams.AddNested("ReportConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.TargetDatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetDatabaseName"))
	}
	if v.TargetTableName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetTableName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDatabaseInput(v *CreateDatabaseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDatabaseInput"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateTableInput(v *CreateTableInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateTableInput"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if v.TableName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TableName"))
	}
	if v.RetentionProperties != nil {
		if err := validateRetentionProperties(v.RetentionProperties); err != nil {
			invalidParams.AddNested("RetentionProperties", err.(smithy.InvalidParamsError))
		}
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.MagneticStoreWriteProperties != nil {
		if err := validateMagneticStoreWriteProperties(v.MagneticStoreWriteProperties); err != nil {
			invalidParams.AddNested("MagneticStoreWriteProperties", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDatabaseInput(v *DeleteDatabaseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDatabaseInput"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTableInput(v *DeleteTableInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTableInput"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if v.TableName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TableName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeBatchLoadTaskInput(v *DescribeBatchLoadTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeBatchLoadTaskInput"}
	if v.TaskId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeDatabaseInput(v *DescribeDatabaseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeDatabaseInput"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeTableInput(v *DescribeTableInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeTableInput"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if v.TableName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TableName"))
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
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpResumeBatchLoadTaskInput(v *ResumeBatchLoadTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResumeBatchLoadTaskInput"}
	if v.TaskId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskId"))
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
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
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
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
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

func validateOpUpdateDatabaseInput(v *UpdateDatabaseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDatabaseInput"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if v.KmsKeyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KmsKeyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateTableInput(v *UpdateTableInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateTableInput"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if v.TableName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TableName"))
	}
	if v.RetentionProperties != nil {
		if err := validateRetentionProperties(v.RetentionProperties); err != nil {
			invalidParams.AddNested("RetentionProperties", err.(smithy.InvalidParamsError))
		}
	}
	if v.MagneticStoreWriteProperties != nil {
		if err := validateMagneticStoreWriteProperties(v.MagneticStoreWriteProperties); err != nil {
			invalidParams.AddNested("MagneticStoreWriteProperties", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpWriteRecordsInput(v *WriteRecordsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "WriteRecordsInput"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if v.TableName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TableName"))
	}
	if v.CommonAttributes != nil {
		if err := validateRecord(v.CommonAttributes); err != nil {
			invalidParams.AddNested("CommonAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if v.Records == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Records"))
	} else if v.Records != nil {
		if err := validateRecords(v.Records); err != nil {
			invalidParams.AddNested("Records", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
