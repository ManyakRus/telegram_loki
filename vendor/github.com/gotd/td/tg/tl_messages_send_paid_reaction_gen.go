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

// MessagesSendPaidReactionRequest represents TL type `messages.sendPaidReaction#9dd6a67b`.
// Sends one or more paid Telegram Star reactions »¹, transferring Telegram Stars »²
// to a channel's balance.
//
// Links:
//  1. https://core.telegram.org/api/reactions#paid-reactions
//  2. https://core.telegram.org/api/stars
//
// See https://core.telegram.org/method/messages.sendPaidReaction for reference.
type MessagesSendPaidReactionRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// The channel
	Peer InputPeerClass
	// The message to react to
	MsgID int
	// The number of stars¹ to send (each will increment the reaction counter by one).
	//
	// Links:
	//  1) https://core.telegram.org/api/stars
	Count int
	// Unique client message ID required to prevent message resending
	RandomID int64
	// Each post with star reactions has a leaderboard with the top senders, but users can
	// opt out of appearing there if they prefer more privacy.  If the user explicitly chose
	// to make their paid reaction(s) private, pass boolTrue¹ to messages.sendPaidReaction²
	// private.  If the user explicitly chose to make their paid reaction(s) private, pass
	// boolFalse³ to messages.sendPaidReaction⁴.private.  If the user did not make any
	// explicit choice about the privacy of their paid reaction(s) (i.e. when reacting by
	// clicking on an existing star reaction on a message), do not populate the messages
	// sendPaidReaction⁵.private flag.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/boolTrue
	//  2) https://core.telegram.org/method/messages.sendPaidReaction
	//  3) https://core.telegram.org/constructor/boolTrue
	//  4) https://core.telegram.org/method/messages.sendPaidReaction
	//  5) https://core.telegram.org/method/messages.sendPaidReaction
	//
	// Use SetPrivate and GetPrivate helpers.
	Private bool
}

// MessagesSendPaidReactionRequestTypeID is TL type id of MessagesSendPaidReactionRequest.
const MessagesSendPaidReactionRequestTypeID = 0x9dd6a67b

// Ensuring interfaces in compile-time for MessagesSendPaidReactionRequest.
var (
	_ bin.Encoder     = &MessagesSendPaidReactionRequest{}
	_ bin.Decoder     = &MessagesSendPaidReactionRequest{}
	_ bin.BareEncoder = &MessagesSendPaidReactionRequest{}
	_ bin.BareDecoder = &MessagesSendPaidReactionRequest{}
)

func (s *MessagesSendPaidReactionRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.MsgID == 0) {
		return false
	}
	if !(s.Count == 0) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.Private == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendPaidReactionRequest) String() string {
	if s == nil {
		return "MessagesSendPaidReactionRequest(nil)"
	}
	type Alias MessagesSendPaidReactionRequest
	return fmt.Sprintf("MessagesSendPaidReactionRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendPaidReactionRequest from given interface.
func (s *MessagesSendPaidReactionRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetCount() (value int)
	GetRandomID() (value int64)
	GetPrivate() (value bool, ok bool)
}) {
	s.Peer = from.GetPeer()
	s.MsgID = from.GetMsgID()
	s.Count = from.GetCount()
	s.RandomID = from.GetRandomID()
	if val, ok := from.GetPrivate(); ok {
		s.Private = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendPaidReactionRequest) TypeID() uint32 {
	return MessagesSendPaidReactionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendPaidReactionRequest) TypeName() string {
	return "messages.sendPaidReaction"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendPaidReactionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendPaidReaction",
		ID:   MessagesSendPaidReactionRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "Private",
			SchemaName: "private",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSendPaidReactionRequest) SetFlags() {
	if !(s.Private == false) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSendPaidReactionRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendPaidReaction#9dd6a67b as nil")
	}
	b.PutID(MessagesSendPaidReactionRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendPaidReactionRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendPaidReaction#9dd6a67b as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendPaidReaction#9dd6a67b: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendPaidReaction#9dd6a67b: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendPaidReaction#9dd6a67b: field peer: %w", err)
	}
	b.PutInt(s.MsgID)
	b.PutInt(s.Count)
	b.PutLong(s.RandomID)
	if s.Flags.Has(0) {
		b.PutBool(s.Private)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendPaidReactionRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendPaidReaction#9dd6a67b to nil")
	}
	if err := b.ConsumeID(MessagesSendPaidReactionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendPaidReaction#9dd6a67b: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendPaidReactionRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendPaidReaction#9dd6a67b to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendPaidReaction#9dd6a67b: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendPaidReaction#9dd6a67b: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendPaidReaction#9dd6a67b: field msg_id: %w", err)
		}
		s.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendPaidReaction#9dd6a67b: field count: %w", err)
		}
		s.Count = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendPaidReaction#9dd6a67b: field random_id: %w", err)
		}
		s.RandomID = value
	}
	if s.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendPaidReaction#9dd6a67b: field private: %w", err)
		}
		s.Private = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSendPaidReactionRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetMsgID returns value of MsgID field.
func (s *MessagesSendPaidReactionRequest) GetMsgID() (value int) {
	if s == nil {
		return
	}
	return s.MsgID
}

// GetCount returns value of Count field.
func (s *MessagesSendPaidReactionRequest) GetCount() (value int) {
	if s == nil {
		return
	}
	return s.Count
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendPaidReactionRequest) GetRandomID() (value int64) {
	if s == nil {
		return
	}
	return s.RandomID
}

// SetPrivate sets value of Private conditional field.
func (s *MessagesSendPaidReactionRequest) SetPrivate(value bool) {
	s.Flags.Set(0)
	s.Private = value
}

// GetPrivate returns value of Private conditional field and
// boolean which is true if field was set.
func (s *MessagesSendPaidReactionRequest) GetPrivate() (value bool, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Private, true
}

// MessagesSendPaidReaction invokes method messages.sendPaidReaction#9dd6a67b returning error if any.
// Sends one or more paid Telegram Star reactions »¹, transferring Telegram Stars »²
// to a channel's balance.
//
// Links:
//  1. https://core.telegram.org/api/reactions#paid-reactions
//  2. https://core.telegram.org/api/stars
//
// See https://core.telegram.org/method/messages.sendPaidReaction for reference.
// Can be used by bots.
func (c *Client) MessagesSendPaidReaction(ctx context.Context, request *MessagesSendPaidReactionRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
