// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The details of an AppInstance, an instance of an Amazon Chime SDK messaging
// application.
type AppInstance struct {

	// The ARN of the messaging instance.
	AppInstanceArn *string

	// The time at which an AppInstance was created. In epoch milliseconds.
	CreatedTimestamp *time.Time

	// The time an AppInstance was last updated. In epoch milliseconds.
	LastUpdatedTimestamp *time.Time

	// The metadata of an AppInstance.
	Metadata *string

	// The name of an AppInstance.
	Name *string

	noSmithyDocumentSerde
}

// The name and ARN of the admin for the AppInstance.
type AppInstanceAdmin struct {

	// The AppInstanceAdmin data.
	Admin *Identity

	// The ARN of the AppInstance for which the user is an administrator.
	AppInstanceArn *string

	// The time at which an administrator was created.
	CreatedTimestamp *time.Time

	noSmithyDocumentSerde
}

// Summary of the details of an AppInstanceAdmin.
type AppInstanceAdminSummary struct {

	// The details of the AppInstanceAdmin.
	Admin *Identity

	noSmithyDocumentSerde
}

// An Amazon Lex V2 chat bot created under an AppInstance.
type AppInstanceBot struct {

	// The ARN of the AppInstanceBot.
	AppInstanceBotArn *string

	// The data processing instructions for an AppInstanceBot.
	Configuration *Configuration

	// The time at which the AppInstanceBot was created.
	CreatedTimestamp *time.Time

	// The time at which the AppInstanceBot was last updated.
	LastUpdatedTimestamp *time.Time

	// The metadata for an AppInstanceBot.
	Metadata *string

	// The name of the AppInstanceBot.
	Name *string

	noSmithyDocumentSerde
}

// High-level information about an AppInstanceBot.
type AppInstanceBotSummary struct {

	// The ARN of the AppInstanceBot.
	AppInstanceBotArn *string

	// The metadata of the AppInstanceBot.
	Metadata *string

	// The name of the AppInstanceBox.
	Name *string

	noSmithyDocumentSerde
}

// The details of the data-retention settings for an AppInstance.
type AppInstanceRetentionSettings struct {

	// The length of time in days to retain the messages in a channel.
	ChannelRetentionSettings *ChannelRetentionSettings

	noSmithyDocumentSerde
}

// Summary of the data for an AppInstance.
type AppInstanceSummary struct {

	// The AppInstance ARN.
	AppInstanceArn *string

	// The metadata of the AppInstance.
	Metadata *string

	// The name of the AppInstance.
	Name *string

	noSmithyDocumentSerde
}

// The details of an AppInstanceUser.
type AppInstanceUser struct {

	// The ARN of the AppInstanceUser.
	AppInstanceUserArn *string

	// The time at which the AppInstanceUser was created.
	CreatedTimestamp *time.Time

	// The interval after which an AppInstanceUser is automatically deleted.
	ExpirationSettings *ExpirationSettings

	// The time at which the AppInstanceUser was last updated.
	LastUpdatedTimestamp *time.Time

	// The metadata of the AppInstanceUser.
	Metadata *string

	// The name of the AppInstanceUser.
	Name *string

	noSmithyDocumentSerde
}

// An endpoint under an Amazon Chime AppInstanceUser that receives messages for a
// user. For push notifications, the endpoint is a mobile device used to receive
// mobile push notifications for a user.
type AppInstanceUserEndpoint struct {

	// Boolean that controls whether the AppInstanceUserEndpoint is opted in to receive
	// messages. ALL indicates the endpoint will receive all messages. NONE indicates
	// the endpoint will receive no messages.
	AllowMessages AllowMessages

	// The ARN of the AppInstanceUser.
	AppInstanceUserArn *string

	// The time at which an AppInstanceUserEndpoint was created.
	CreatedTimestamp *time.Time

	// The attributes of an Endpoint.
	EndpointAttributes *EndpointAttributes

	// The unique identifier of the AppInstanceUserEndpoint.
	EndpointId *string

	// A read-only field that represents the state of an AppInstanceUserEndpoint.
	// Supported values:
	//
	// * ACTIVE: The AppInstanceUserEndpoint is active and able to
	// receive messages. When ACTIVE, the EndpointStatusReason remains empty.
	//
	// *
	// INACTIVE: The AppInstanceUserEndpoint is inactive and can't receive message.
	// When INACTIVE, the corresponding reason will be conveyed through
	// EndpointStatusReason.
	//
	// * INVALID_DEVICE_TOKEN indicates that an
	// AppInstanceUserEndpoint is INACTIVE due to invalid device token
	//
	// *
	// INVALID_PINPOINT_ARN indicates that an AppInstanceUserEndpoint is INACTIVE due
	// to an invalid pinpoint ARN that was input through the ResourceArn field.
	EndpointState *EndpointState

	// The time at which an AppInstanceUserEndpoint was last updated.
	LastUpdatedTimestamp *time.Time

	// The name of the AppInstanceUserEndpoint.
	Name *string

	// The ARN of the resource to which the endpoint belongs.
	ResourceArn *string

	// The type of the AppInstanceUserEndpoint.
	Type AppInstanceUserEndpointType

	noSmithyDocumentSerde
}

// Summary of the details of an AppInstanceUserEndpoint.
type AppInstanceUserEndpointSummary struct {

	// BBoolean that controls whether the AppInstanceUserEndpoint is opted in to
	// receive messages. ALL indicates the endpoint will receive all messages. NONE
	// indicates the endpoint will receive no messages.
	AllowMessages AllowMessages

	// The ARN of the AppInstanceUser.
	AppInstanceUserArn *string

	// The unique identifier of the AppInstanceUserEndpoint.
	EndpointId *string

	// A read-only field that represent the state of an AppInstanceUserEndpoint.
	EndpointState *EndpointState

	// The name of the AppInstanceUserEndpoint.
	Name *string

	// The type of the AppInstanceUserEndpoint.
	Type AppInstanceUserEndpointType

	noSmithyDocumentSerde
}

// Summary of the details of an AppInstanceUser.
type AppInstanceUserSummary struct {

	// The ARN of the AppInstanceUser.
	AppInstanceUserArn *string

	// The metadata of the AppInstanceUser.
	Metadata *string

	// The name of an AppInstanceUser.
	Name *string

	noSmithyDocumentSerde
}

// The details of the retention settings for a channel.
type ChannelRetentionSettings struct {

	// The time in days to retain the messages in a channel.
	RetentionDays *int32

	noSmithyDocumentSerde
}

// A structure that contains configuration data.
type Configuration struct {

	// The configuration for an Amazon Lex V2 bot.
	//
	// This member is required.
	Lex *LexConfiguration

	noSmithyDocumentSerde
}

// The attributes of an Endpoint.
type EndpointAttributes struct {

	// The device token for the GCM, APNS, and APNS_SANDBOX endpoint types.
	//
	// This member is required.
	DeviceToken *string

	// The VOIP device token for the APNS and APNS_SANDBOX endpoint types.
	VoipDeviceToken *string

	noSmithyDocumentSerde
}

// A read-only field that represents the state of an AppInstanceUserEndpoint.
// Supported values:
//
// * ACTIVE: The AppInstanceUserEndpoint is active and able to
// receive messages. When ACTIVE, the EndpointStatusReason remains empty.
//
// *
// INACTIVE: The AppInstanceUserEndpoint is inactive and can't receive message.
// When INACTIVE, the corresponding reason will be conveyed through
// EndpointStatusReason.
//
// * INVALID_DEVICE_TOKEN indicates that an
// AppInstanceUserEndpoint is INACTIVE due to invalid device token
//
// *
// INVALID_PINPOINT_ARN indicates that an AppInstanceUserEndpoint is INACTIVE due
// to an invalid pinpoint ARN that was input through the ResourceArn field.
type EndpointState struct {

	// Enum that indicates the Status of an AppInstanceUserEndpoint.
	//
	// This member is required.
	Status EndpointStatus

	// The reason for the EndpointStatus.
	StatusReason EndpointStatusReason

	noSmithyDocumentSerde
}

// Determines the interval after which an AppInstanceUser is automatically deleted.
type ExpirationSettings struct {

	// Specifies the conditions under which an AppInstanceUser will expire.
	//
	// This member is required.
	ExpirationCriterion ExpirationCriterion

	// The period in days after which an AppInstanceUser will be automatically deleted.
	//
	// This member is required.
	ExpirationDays *int32

	noSmithyDocumentSerde
}

// The details of a user or bot.
type Identity struct {

	// The ARN in an Identity.
	Arn *string

	// The name in an Identity.
	Name *string

	noSmithyDocumentSerde
}

// The configuration for an Amazon Lex V2 bot.
type LexConfiguration struct {

	// The ARN of the Amazon Lex V2 bot's alias. The ARN uses this format:
	// arn:aws:lex:REGION:ACCOUNT:bot-alias/MYBOTID/MYBOTALIAS
	//
	// This member is required.
	LexBotAliasArn *string

	// Identifies the Amazon Lex V2 bot's language and locale. The string must match
	// one of the supported locales in Amazon Lex V2. All of the intents, slot types,
	// and slots used in the bot must have the same locale. For more information, see
	// Supported languages
	// (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html) in the Amazon
	// Lex V2 Developer Guide.
	//
	// This member is required.
	LocaleId *string

	// Determines whether the Amazon Lex V2 bot responds to all standard messages.
	// Control messages are not supported.
	//
	// This member is required.
	RespondsTo RespondsTo

	// The name of the welcome intent configured in the Amazon Lex V2 bot.
	WelcomeIntent *string

	noSmithyDocumentSerde
}

// A tag object containing a key-value pair.
type Tag struct {

	// The key in a tag.
	//
	// This member is required.
	Key *string

	// The value in a tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
