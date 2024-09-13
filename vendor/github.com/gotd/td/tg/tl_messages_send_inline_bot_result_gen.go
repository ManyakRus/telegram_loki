// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// MessagesSendInlineBotResultRequest represents TL type `messages.sendInlineBotResult#3ebee86a`.
// Send a result obtained using messages.getInlineBotResults¹.
//
// Links:
//  1. https://core.telegram.org/method/messages.getInlineBotResults
//
// See https://core.telegram.org/method/messages.sendInlineBotResult for reference.
type MessagesSendInlineBotResultRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to send the message silently (no notification will be triggered on the other
	// client)
	Silent bool
	// Whether to send the message in background
	Background bool
	// Whether to clear the draft¹
	//
	// Links:
	//  1) https://core.telegram.org/api/drafts
	ClearDraft bool
	// Whether to hide the via @botname in the resulting message (only for bot usernames
	// encountered in the config¹)
	//
	// Links:
	//  1) https://core.telegram.org/constructor/config
	HideVia bool
	// Destination
	Peer InputPeerClass
	// If set, indicates that the message should be sent in reply to the specified message or
	// story.
	//
	// Use SetReplyTo and GetReplyTo helpers.
	ReplyTo InputReplyToClass
	// Random ID to avoid resending the same query
	RandomID int64
	// Query ID from messages.getInlineBotResults¹
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.getInlineBotResults
	QueryID int64
	// Result ID from messages.getInlineBotResults¹
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.getInlineBotResults
	ID string
	// Scheduled message date for scheduled messages
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
	// Send this message as the specified peer
	//
	// Use SetSendAs and GetSendAs helpers.
	SendAs InputPeerClass
	// QuickReplyShortcut field of MessagesSendInlineBotResultRequest.
	//
	// Use SetQuickReplyShortcut and GetQuickReplyShortcut helpers.
	QuickReplyShortcut InputQuickReplyShortcutClass
}

// MessagesSendInlineBotResultRequestTypeID is TL type id of MessagesSendInlineBotResultRequest.
const MessagesSendInlineBotResultRequestTypeID = 0x3ebee86a

// Ensuring interfaces in compile-time for MessagesSendInlineBotResultRequest.
var (
	_ bin.Encoder     = &MessagesSendInlineBotResultRequest{}
	_ bin.Decoder     = &MessagesSendInlineBotResultRequest{}
	_ bin.BareEncoder = &MessagesSendInlineBotResultRequest{}
	_ bin.BareDecoder = &MessagesSendInlineBotResultRequest{}
)

func (s *MessagesSendInlineBotResultRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Silent == false) {
		return false
	}
	if !(s.Background == false) {
		return false
	}
	if !(s.ClearDraft == false) {
		return false
	}
	if !(s.HideVia == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ReplyTo == nil) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.QueryID == 0) {
		return false
	}
	if !(s.ID == "") {
		return false
	}
	if !(s.ScheduleDate == 0) {
		return false
	}
	if !(s.SendAs == nil) {
		return false
	}
	if !(s.QuickReplyShortcut == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendInlineBotResultRequest) String() string {
	if s == nil {
		return "MessagesSendInlineBotResultRequest(nil)"
	}
	type Alias MessagesSendInlineBotResultRequest
	return fmt.Sprintf("MessagesSendInlineBotResultRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendInlineBotResultRequest from given interface.
func (s *MessagesSendInlineBotResultRequest) FillFrom(from interface {
	GetSilent() (value bool)
	GetBackground() (value bool)
	GetClearDraft() (value bool)
	GetHideVia() (value bool)
	GetPeer() (value InputPeerClass)
	GetReplyTo() (value InputReplyToClass, ok bool)
	GetRandomID() (value int64)
	GetQueryID() (value int64)
	GetID() (value string)
	GetScheduleDate() (value int, ok bool)
	GetSendAs() (value InputPeerClass, ok bool)
	GetQuickReplyShortcut() (value InputQuickReplyShortcutClass, ok bool)
}) {
	s.Silent = from.GetSilent()
	s.Background = from.GetBackground()
	s.ClearDraft = from.GetClearDraft()
	s.HideVia = from.GetHideVia()
	s.Peer = from.GetPeer()
	if val, ok := from.GetReplyTo(); ok {
		s.ReplyTo = val
	}

	s.RandomID = from.GetRandomID()
	s.QueryID = from.GetQueryID()
	s.ID = from.GetID()
	if val, ok := from.GetScheduleDate(); ok {
		s.ScheduleDate = val
	}

	if val, ok := from.GetSendAs(); ok {
		s.SendAs = val
	}

	if val, ok := from.GetQuickReplyShortcut(); ok {
		s.QuickReplyShortcut = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendInlineBotResultRequest) TypeID() uint32 {
	return MessagesSendInlineBotResultRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendInlineBotResultRequest) TypeName() string {
	return "messages.sendInlineBotResult"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendInlineBotResultRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendInlineBotResult",
		ID:   MessagesSendInlineBotResultRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "Background",
			SchemaName: "background",
			Null:       !s.Flags.Has(6),
		},
		{
			Name:       "ClearDraft",
			SchemaName: "clear_draft",
			Null:       !s.Flags.Has(7),
		},
		{
			Name:       "HideVia",
			SchemaName: "hide_via",
			Null:       !s.Flags.Has(11),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ReplyTo",
			SchemaName: "reply_to",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "ScheduleDate",
			SchemaName: "schedule_date",
			Null:       !s.Flags.Has(10),
		},
		{
			Name:       "SendAs",
			SchemaName: "send_as",
			Null:       !s.Flags.Has(13),
		},
		{
			Name:       "QuickReplyShortcut",
			SchemaName: "quick_reply_shortcut",
			Null:       !s.Flags.Has(17),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSendInlineBotResultRequest) SetFlags() {
	if !(s.Silent == false) {
		s.Flags.Set(5)
	}
	if !(s.Background == false) {
		s.Flags.Set(6)
	}
	if !(s.ClearDraft == false) {
		s.Flags.Set(7)
	}
	if !(s.HideVia == false) {
		s.Flags.Set(11)
	}
	if !(s.ReplyTo == nil) {
		s.Flags.Set(0)
	}
	if !(s.ScheduleDate == 0) {
		s.Flags.Set(10)
	}
	if !(s.SendAs == nil) {
		s.Flags.Set(13)
	}
	if !(s.QuickReplyShortcut == nil) {
		s.Flags.Set(17)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSendInlineBotResultRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendInlineBotResult#3ebee86a as nil")
	}
	b.PutID(MessagesSendInlineBotResultRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendInlineBotResultRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendInlineBotResult#3ebee86a as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#3ebee86a: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#3ebee86a: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#3ebee86a: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		if s.ReplyTo == nil {
			return fmt.Errorf("unable to encode messages.sendInlineBotResult#3ebee86a: field reply_to is nil")
		}
		if err := s.ReplyTo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendInlineBotResult#3ebee86a: field reply_to: %w", err)
		}
	}
	b.PutLong(s.RandomID)
	b.PutLong(s.QueryID)
	b.PutString(s.ID)
	if s.Flags.Has(10) {
		b.PutInt(s.ScheduleDate)
	}
	if s.Flags.Has(13) {
		if s.SendAs == nil {
			return fmt.Errorf("unable to encode messages.sendInlineBotResult#3ebee86a: field send_as is nil")
		}
		if err := s.SendAs.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendInlineBotResult#3ebee86a: field send_as: %w", err)
		}
	}
	if s.Flags.Has(17) {
		if s.QuickReplyShortcut == nil {
			return fmt.Errorf("unable to encode messages.sendInlineBotResult#3ebee86a: field quick_reply_shortcut is nil")
		}
		if err := s.QuickReplyShortcut.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendInlineBotResult#3ebee86a: field quick_reply_shortcut: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendInlineBotResultRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendInlineBotResult#3ebee86a to nil")
	}
	if err := b.ConsumeID(MessagesSendInlineBotResultRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendInlineBotResult#3ebee86a: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendInlineBotResultRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendInlineBotResult#3ebee86a to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#3ebee86a: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(5)
	s.Background = s.Flags.Has(6)
	s.ClearDraft = s.Flags.Has(7)
	s.HideVia = s.Flags.Has(11)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#3ebee86a: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := DecodeInputReplyTo(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#3ebee86a: field reply_to: %w", err)
		}
		s.ReplyTo = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#3ebee86a: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#3ebee86a: field query_id: %w", err)
		}
		s.QueryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#3ebee86a: field id: %w", err)
		}
		s.ID = value
	}
	if s.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#3ebee86a: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	if s.Flags.Has(13) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#3ebee86a: field send_as: %w", err)
		}
		s.SendAs = value
	}
	if s.Flags.Has(17) {
		value, err := DecodeInputQuickReplyShortcut(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#3ebee86a: field quick_reply_shortcut: %w", err)
		}
		s.QuickReplyShortcut = value
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendInlineBotResultRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(5)
		s.Silent = true
	} else {
		s.Flags.Unset(5)
		s.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (s *MessagesSendInlineBotResultRequest) GetSilent() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(5)
}

// SetBackground sets value of Background conditional field.
func (s *MessagesSendInlineBotResultRequest) SetBackground(value bool) {
	if value {
		s.Flags.Set(6)
		s.Background = true
	} else {
		s.Flags.Unset(6)
		s.Background = false
	}
}

// GetBackground returns value of Background conditional field.
func (s *MessagesSendInlineBotResultRequest) GetBackground() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(6)
}

// SetClearDraft sets value of ClearDraft conditional field.
func (s *MessagesSendInlineBotResultRequest) SetClearDraft(value bool) {
	if value {
		s.Flags.Set(7)
		s.ClearDraft = true
	} else {
		s.Flags.Unset(7)
		s.ClearDraft = false
	}
}

// GetClearDraft returns value of ClearDraft conditional field.
func (s *MessagesSendInlineBotResultRequest) GetClearDraft() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(7)
}

// SetHideVia sets value of HideVia conditional field.
func (s *MessagesSendInlineBotResultRequest) SetHideVia(value bool) {
	if value {
		s.Flags.Set(11)
		s.HideVia = true
	} else {
		s.Flags.Unset(11)
		s.HideVia = false
	}
}

// GetHideVia returns value of HideVia conditional field.
func (s *MessagesSendInlineBotResultRequest) GetHideVia() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(11)
}

// GetPeer returns value of Peer field.
func (s *MessagesSendInlineBotResultRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// SetReplyTo sets value of ReplyTo conditional field.
func (s *MessagesSendInlineBotResultRequest) SetReplyTo(value InputReplyToClass) {
	s.Flags.Set(0)
	s.ReplyTo = value
}

// GetReplyTo returns value of ReplyTo conditional field and
// boolean which is true if field was set.
func (s *MessagesSendInlineBotResultRequest) GetReplyTo() (value InputReplyToClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyTo, true
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendInlineBotResultRequest) GetRandomID() (value int64) {
	if s == nil {
		return
	}
	return s.RandomID
}

// GetQueryID returns value of QueryID field.
func (s *MessagesSendInlineBotResultRequest) GetQueryID() (value int64) {
	if s == nil {
		return
	}
	return s.QueryID
}

// GetID returns value of ID field.
func (s *MessagesSendInlineBotResultRequest) GetID() (value string) {
	if s == nil {
		return
	}
	return s.ID
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *MessagesSendInlineBotResultRequest) SetScheduleDate(value int) {
	s.Flags.Set(10)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *MessagesSendInlineBotResultRequest) GetScheduleDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(10) {
		return value, false
	}
	return s.ScheduleDate, true
}

// SetSendAs sets value of SendAs conditional field.
func (s *MessagesSendInlineBotResultRequest) SetSendAs(value InputPeerClass) {
	s.Flags.Set(13)
	s.SendAs = value
}

// GetSendAs returns value of SendAs conditional field and
// boolean which is true if field was set.
func (s *MessagesSendInlineBotResultRequest) GetSendAs() (value InputPeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(13) {
		return value, false
	}
	return s.SendAs, true
}

// SetQuickReplyShortcut sets value of QuickReplyShortcut conditional field.
func (s *MessagesSendInlineBotResultRequest) SetQuickReplyShortcut(value InputQuickReplyShortcutClass) {
	s.Flags.Set(17)
	s.QuickReplyShortcut = value
}

// GetQuickReplyShortcut returns value of QuickReplyShortcut conditional field and
// boolean which is true if field was set.
func (s *MessagesSendInlineBotResultRequest) GetQuickReplyShortcut() (value InputQuickReplyShortcutClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(17) {
		return value, false
	}
	return s.QuickReplyShortcut, true
}

// MessagesSendInlineBotResult invokes method messages.sendInlineBotResult#3ebee86a returning error if any.
// Send a result obtained using messages.getInlineBotResults¹.
//
// Links:
//  1. https://core.telegram.org/method/messages.getInlineBotResults
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	403 CHAT_GUEST_SEND_FORBIDDEN: You join the discussion group before commenting, see here » for more info.
//	400 CHAT_RESTRICTED: You can't send messages in this chat, you were restricted.
//	403 CHAT_SEND_AUDIOS_FORBIDDEN: You can't send audio messages in this chat.
//	403 CHAT_SEND_GAME_FORBIDDEN: You can't send a game to this chat.
//	403 CHAT_SEND_GIFS_FORBIDDEN: You can't send gifs in this chat.
//	403 CHAT_SEND_INLINE_FORBIDDEN: You can't send inline messages in this group.
//	403 CHAT_SEND_MEDIA_FORBIDDEN: You can't send media in this chat.
//	403 CHAT_SEND_PHOTOS_FORBIDDEN: You can't send photos in this chat.
//	403 CHAT_SEND_PLAIN_FORBIDDEN: You can't send non-media (text) messages in this chat.
//	403 CHAT_SEND_STICKERS_FORBIDDEN: You can't send stickers in this chat.
//	403 CHAT_SEND_VOICES_FORBIDDEN: You can't send voice recordings in this chat.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 ENTITY_BOUNDS_INVALID: A specified entity offset or length is invalid, see here » for info on how to properly compute the entity offset/length.
//	400 INLINE_RESULT_EXPIRED: The inline query expired.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 MEDIA_EMPTY: The provided media object is invalid.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 QUERY_ID_EMPTY: The query ID is empty.
//	500 RANDOM_ID_DUPLICATE: You provided a random ID that was already used.
//	400 RESULT_ID_EMPTY: Result ID empty.
//	400 RESULT_ID_INVALID: One of the specified result IDs is invalid.
//	400 SCHEDULE_DATE_TOO_LATE: You can't schedule a message this far in the future.
//	400 SCHEDULE_TOO_MUCH: There are too many scheduled messages.
//	500 SEND_MEDIA_INVALID: The specified media is invalid.
//	420 SLOWMODE_WAIT_%d: Slowmode is enabled in this chat: wait %d seconds before sending another message to this chat.
//	400 TOPIC_DELETED: The specified topic was deleted.
//	400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels.
//	400 VOICE_MESSAGES_FORBIDDEN: This user's privacy settings forbid you from sending voice messages.
//	400 WEBPAGE_CURL_FAILED: Failure while fetching the webpage with cURL.
//	400 WEBPAGE_MEDIA_EMPTY: Webpage media empty.
//	400 YOU_BLOCKED_USER: You blocked this user.
//
// See https://core.telegram.org/method/messages.sendInlineBotResult for reference.
func (c *Client) MessagesSendInlineBotResult(ctx context.Context, request *MessagesSendInlineBotResultRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
