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

// StoriesSendReactionRequest represents TL type `stories.sendReaction#7fd736b2`.
// React to a story.
//
// See https://core.telegram.org/method/stories.sendReaction for reference.
type StoriesSendReactionRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to add this reaction to the recent reactions list »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/reactions#recent-reactions
	AddToRecent bool
	// The peer that sent the story
	Peer InputPeerClass
	// ID of the story to react to
	StoryID int
	// Reaction
	Reaction ReactionClass
}

// StoriesSendReactionRequestTypeID is TL type id of StoriesSendReactionRequest.
const StoriesSendReactionRequestTypeID = 0x7fd736b2

// Ensuring interfaces in compile-time for StoriesSendReactionRequest.
var (
	_ bin.Encoder     = &StoriesSendReactionRequest{}
	_ bin.Decoder     = &StoriesSendReactionRequest{}
	_ bin.BareEncoder = &StoriesSendReactionRequest{}
	_ bin.BareDecoder = &StoriesSendReactionRequest{}
)

func (s *StoriesSendReactionRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.AddToRecent == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.StoryID == 0) {
		return false
	}
	if !(s.Reaction == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoriesSendReactionRequest) String() string {
	if s == nil {
		return "StoriesSendReactionRequest(nil)"
	}
	type Alias StoriesSendReactionRequest
	return fmt.Sprintf("StoriesSendReactionRequest%+v", Alias(*s))
}

// FillFrom fills StoriesSendReactionRequest from given interface.
func (s *StoriesSendReactionRequest) FillFrom(from interface {
	GetAddToRecent() (value bool)
	GetPeer() (value InputPeerClass)
	GetStoryID() (value int)
	GetReaction() (value ReactionClass)
}) {
	s.AddToRecent = from.GetAddToRecent()
	s.Peer = from.GetPeer()
	s.StoryID = from.GetStoryID()
	s.Reaction = from.GetReaction()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesSendReactionRequest) TypeID() uint32 {
	return StoriesSendReactionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesSendReactionRequest) TypeName() string {
	return "stories.sendReaction"
}

// TypeInfo returns info about TL type.
func (s *StoriesSendReactionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.sendReaction",
		ID:   StoriesSendReactionRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AddToRecent",
			SchemaName: "add_to_recent",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
		{
			Name:       "Reaction",
			SchemaName: "reaction",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoriesSendReactionRequest) SetFlags() {
	if !(s.AddToRecent == false) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *StoriesSendReactionRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stories.sendReaction#7fd736b2 as nil")
	}
	b.PutID(StoriesSendReactionRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoriesSendReactionRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stories.sendReaction#7fd736b2 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.sendReaction#7fd736b2: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode stories.sendReaction#7fd736b2: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.sendReaction#7fd736b2: field peer: %w", err)
	}
	b.PutInt(s.StoryID)
	if s.Reaction == nil {
		return fmt.Errorf("unable to encode stories.sendReaction#7fd736b2: field reaction is nil")
	}
	if err := s.Reaction.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.sendReaction#7fd736b2: field reaction: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoriesSendReactionRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stories.sendReaction#7fd736b2 to nil")
	}
	if err := b.ConsumeID(StoriesSendReactionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.sendReaction#7fd736b2: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoriesSendReactionRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stories.sendReaction#7fd736b2 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.sendReaction#7fd736b2: field flags: %w", err)
		}
	}
	s.AddToRecent = s.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.sendReaction#7fd736b2: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.sendReaction#7fd736b2: field story_id: %w", err)
		}
		s.StoryID = value
	}
	{
		value, err := DecodeReaction(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.sendReaction#7fd736b2: field reaction: %w", err)
		}
		s.Reaction = value
	}
	return nil
}

// SetAddToRecent sets value of AddToRecent conditional field.
func (s *StoriesSendReactionRequest) SetAddToRecent(value bool) {
	if value {
		s.Flags.Set(0)
		s.AddToRecent = true
	} else {
		s.Flags.Unset(0)
		s.AddToRecent = false
	}
}

// GetAddToRecent returns value of AddToRecent conditional field.
func (s *StoriesSendReactionRequest) GetAddToRecent() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (s *StoriesSendReactionRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetStoryID returns value of StoryID field.
func (s *StoriesSendReactionRequest) GetStoryID() (value int) {
	if s == nil {
		return
	}
	return s.StoryID
}

// GetReaction returns value of Reaction field.
func (s *StoriesSendReactionRequest) GetReaction() (value ReactionClass) {
	if s == nil {
		return
	}
	return s.Reaction
}

// StoriesSendReaction invokes method stories.sendReaction#7fd736b2 returning error if any.
// React to a story.
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 REACTION_INVALID: The specified reaction is invalid.
//	400 STORY_ID_EMPTY: You specified no story IDs.
//
// See https://core.telegram.org/method/stories.sendReaction for reference.
func (c *Client) StoriesSendReaction(ctx context.Context, request *StoriesSendReactionRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
