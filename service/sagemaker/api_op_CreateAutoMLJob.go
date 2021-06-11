// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Autopilot job. Find the best performing model after you run an
// Autopilot job by calling . For information about how to use Autopilot, see
// Automate Model Development with Amazon SageMaker Autopilot
// (https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html).
func (c *Client) CreateAutoMLJob(ctx context.Context, params *CreateAutoMLJobInput, optFns ...func(*Options)) (*CreateAutoMLJobOutput, error) {
	if params == nil {
		params = &CreateAutoMLJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAutoMLJob", params, optFns, addOperationCreateAutoMLJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAutoMLJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAutoMLJobInput struct {

	// Identifies an Autopilot job. The name must be unique to your account and is
	// case-insensitive.
	//
	// This member is required.
	AutoMLJobName *string

	// An array of channel objects that describes the input data and its location. Each
	// channel is a named input source. Similar to InputDataConfig supported by .
	// Format(s) supported: CSV. Minimum of 500 rows.
	//
	// This member is required.
	InputDataConfig []types.AutoMLChannel

	// Provides information about encryption and the Amazon S3 output path needed to
	// store artifacts from an AutoML job. Format(s) supported: CSV.
	//
	// This member is required.
	OutputDataConfig *types.AutoMLOutputDataConfig

	// The ARN of the role that is used to access the data.
	//
	// This member is required.
	RoleArn *string

	// Contains CompletionCriteria and SecurityConfig settings for the AutoML job.
	AutoMLJobConfig *types.AutoMLJobConfig

	// Defines the objective metric used to measure the predictive quality of an AutoML
	// job. You provide an AutoMLJobObjective$MetricName and Autopilot infers whether
	// to minimize or maximize it.
	AutoMLJobObjective *types.AutoMLJobObjective

	// Generates possible candidates without training the models. A candidate is a
	// combination of data preprocessors, algorithms, and algorithm parameter settings.
	GenerateCandidateDefinitionsOnly bool

	// Specifies how to generate the endpoint name for an automatic one-click Autopilot
	// model deployment.
	ModelDeployConfig *types.ModelDeployConfig

	// Defines the type of supervised learning available for the candidates. Options
	// include: BinaryClassification, MulticlassClassification, and Regression. For
	// more information, see  Amazon SageMaker Autopilot problem types and algorithm
	// support
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development-problem-types.html).
	ProblemType types.ProblemType

	// Each tag consists of a key and an optional value. Tag keys must be unique per
	// resource.
	Tags []types.Tag
}

type CreateAutoMLJobOutput struct {

	// The unique ARN that is assigned to the AutoML job when it is created.
	//
	// This member is required.
	AutoMLJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateAutoMLJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAutoMLJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAutoMLJob{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAutoMLJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAutoMLJob(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opCreateAutoMLJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateAutoMLJob",
	}
}
