// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a voice tone analysis task. For more information about voice tone
// analysis, see Using Amazon Chime SDK voice analytics
// (https://docs.aws.amazon.com/chime-sdk/latest/dg/pstn-voice-analytics.html) in
// the Amazon Chime SDK Developer Guide. Before starting any voice tone analysis
// tasks, you must provide all notices and obtain all consents from the speaker as
// required under applicable privacy and biometrics laws, and as required under the
// AWS service terms (https://aws.amazon.com/service-terms/) for the Amazon Chime
// SDK.
func (c *Client) StartVoiceToneAnalysisTask(ctx context.Context, params *StartVoiceToneAnalysisTaskInput, optFns ...func(*Options)) (*StartVoiceToneAnalysisTaskOutput, error) {
	if params == nil {
		params = &StartVoiceToneAnalysisTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartVoiceToneAnalysisTask", params, optFns, c.addOperationStartVoiceToneAnalysisTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartVoiceToneAnalysisTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartVoiceToneAnalysisTaskInput struct {

	// The language code.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The transaction ID.
	//
	// This member is required.
	TransactionId *string

	// The Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// The unique identifier for the client request. Use a different token for
	// different voice tone analysis tasks.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type StartVoiceToneAnalysisTaskOutput struct {

	// The details of the voice tone analysis task.
	VoiceToneAnalysisTask *types.VoiceToneAnalysisTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartVoiceToneAnalysisTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartVoiceToneAnalysisTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartVoiceToneAnalysisTask{}, middleware.After)
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
	if err = addOpStartVoiceToneAnalysisTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartVoiceToneAnalysisTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartVoiceToneAnalysisTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "StartVoiceToneAnalysisTask",
	}
}
