// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a message. ConnectionToken is used for invoking this API instead of
// ParticipantToken. The Amazon Connect Participant Service APIs do not use
// Signature Version 4 authentication
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
func (c *Client) SendMessage(ctx context.Context, params *SendMessageInput, optFns ...func(*Options)) (*SendMessageOutput, error) {
	if params == nil {
		params = &SendMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendMessage", params, optFns, c.addOperationSendMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendMessageInput struct {

	// The authentication token associated with the connection.
	//
	// This member is required.
	ConnectionToken *string

	// The content of the message.
	//
	// * For text/plain and text/markdown, the Length
	// Constraints are Minimum of 1, Maximum of 1024.
	//
	// * For application/json, the
	// Length Constraints are Minimum of 1, Maximum of 12000.
	//
	// * For
	// application/vnd.amazonaws.connect.message.interactive.response, the Length
	// Constraints are Minimum of 1, Maximum of 12288.
	//
	// This member is required.
	Content *string

	// The type of the content. Supported types are text/plain, text/markdown,
	// application/json, and
	// application/vnd.amazonaws.connect.message.interactive.response.
	//
	// This member is required.
	ContentType *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see Making retries safe with
	// idempotent APIs
	// (https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/).
	ClientToken *string

	noSmithyDocumentSerde
}

type SendMessageOutput struct {

	// The time when the message was sent. It's specified in ISO 8601 format:
	// yyyy-MM-ddThh:mm:ss.SSSZ. For example, 2019-11-08T02:41:28.172Z.
	AbsoluteTime *string

	// The ID of the message.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendMessage{}, middleware.After)
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
	if err = addIdempotencyToken_opSendMessageMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSendMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendMessage(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpSendMessage struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendMessage) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendMessage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendMessageInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendMessageInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opSendMessageMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpSendMessage{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "SendMessage",
	}
}
