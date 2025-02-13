// THIS FILE IS AUTOGENERATED. DO NOT EDIT.
// Regen by running 'go generate' in the repo root.

package gotgbot

// The consts listed below represent all the update types that can be requested from telegram.
const (
	UpdateTypeMessage              = "message"
	UpdateTypeEditedMessage        = "edited_message"
	UpdateTypeChannelPost          = "channel_post"
	UpdateTypeEditedChannelPost    = "edited_channel_post"
	UpdateTypeMessageReaction      = "message_reaction"
	UpdateTypeMessageReactionCount = "message_reaction_count"
	UpdateTypeInlineQuery          = "inline_query"
	UpdateTypeChosenInlineResult   = "chosen_inline_result"
	UpdateTypeCallbackQuery        = "callback_query"
	UpdateTypeShippingQuery        = "shipping_query"
	UpdateTypePreCheckoutQuery     = "pre_checkout_query"
	UpdateTypePoll                 = "poll"
	UpdateTypePollAnswer           = "poll_answer"
	UpdateTypeMyChatMember         = "my_chat_member"
	UpdateTypeChatMember           = "chat_member"
	UpdateTypeChatJoinRequest      = "chat_join_request"
	UpdateTypeChatBoost            = "chat_boost"
	UpdateTypeRemovedChatBoost     = "removed_chat_boost"
)

// The consts listed below represent all the parse_mode options that can be sent to telegram.
const (
	ParseModeHTML       = "HTML"
	ParseModeMarkdownV2 = "MarkdownV2"
	ParseModeMarkdown   = "Markdown"
	ParseModeNone       = ""
)

// The consts listed below represent all the sticker types that can be obtained from telegram.
const (
	StickerTypeRegular     = "regular"
	StickerTypeMask        = "mask"
	StickerTypeCustomEmoji = "custom_emoji"
)

// The consts listed below represent all the chat types that can be obtained from telegram.
const (
	ChatTypePrivate    = "private"
	ChatTypeGroup      = "group"
	ChatTypeSupergroup = "supergroup"
	ChatTypeChannel    = "channel"
)
