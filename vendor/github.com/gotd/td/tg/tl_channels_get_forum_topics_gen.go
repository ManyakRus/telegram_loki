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

// ChannelsGetForumTopicsRequest represents TL type `channels.getForumTopics#de560d1`.
// Get topics of a forum¹
//
// Links:
//  1. https://core.telegram.org/api/forum
//
// See https://core.telegram.org/method/channels.getForumTopics for reference.
type ChannelsGetForumTopicsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Supergroup
	Channel InputChannelClass
	// Search query
	//
	// Use SetQ and GetQ helpers.
	Q string
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetDate int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetTopic int
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// ChannelsGetForumTopicsRequestTypeID is TL type id of ChannelsGetForumTopicsRequest.
const ChannelsGetForumTopicsRequestTypeID = 0xde560d1

// Ensuring interfaces in compile-time for ChannelsGetForumTopicsRequest.
var (
	_ bin.Encoder     = &ChannelsGetForumTopicsRequest{}
	_ bin.Decoder     = &ChannelsGetForumTopicsRequest{}
	_ bin.BareEncoder = &ChannelsGetForumTopicsRequest{}
	_ bin.BareDecoder = &ChannelsGetForumTopicsRequest{}
)

func (g *ChannelsGetForumTopicsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Channel == nil) {
		return false
	}
	if !(g.Q == "") {
		return false
	}
	if !(g.OffsetDate == 0) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.OffsetTopic == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetForumTopicsRequest) String() string {
	if g == nil {
		return "ChannelsGetForumTopicsRequest(nil)"
	}
	type Alias ChannelsGetForumTopicsRequest
	return fmt.Sprintf("ChannelsGetForumTopicsRequest%+v", Alias(*g))
}

// FillFrom fills ChannelsGetForumTopicsRequest from given interface.
func (g *ChannelsGetForumTopicsRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetQ() (value string, ok bool)
	GetOffsetDate() (value int)
	GetOffsetID() (value int)
	GetOffsetTopic() (value int)
	GetLimit() (value int)
}) {
	g.Channel = from.GetChannel()
	if val, ok := from.GetQ(); ok {
		g.Q = val
	}

	g.OffsetDate = from.GetOffsetDate()
	g.OffsetID = from.GetOffsetID()
	g.OffsetTopic = from.GetOffsetTopic()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetForumTopicsRequest) TypeID() uint32 {
	return ChannelsGetForumTopicsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetForumTopicsRequest) TypeName() string {
	return "channels.getForumTopics"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetForumTopicsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getForumTopics",
		ID:   ChannelsGetForumTopicsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Q",
			SchemaName: "q",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "OffsetDate",
			SchemaName: "offset_date",
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "OffsetTopic",
			SchemaName: "offset_topic",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *ChannelsGetForumTopicsRequest) SetFlags() {
	if !(g.Q == "") {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *ChannelsGetForumTopicsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getForumTopics#de560d1 as nil")
	}
	b.PutID(ChannelsGetForumTopicsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetForumTopicsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getForumTopics#de560d1 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getForumTopics#de560d1: field flags: %w", err)
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode channels.getForumTopics#de560d1: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getForumTopics#de560d1: field channel: %w", err)
	}
	if g.Flags.Has(0) {
		b.PutString(g.Q)
	}
	b.PutInt(g.OffsetDate)
	b.PutInt(g.OffsetID)
	b.PutInt(g.OffsetTopic)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetForumTopicsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getForumTopics#de560d1 to nil")
	}
	if err := b.ConsumeID(ChannelsGetForumTopicsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getForumTopics#de560d1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetForumTopicsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getForumTopics#de560d1 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.getForumTopics#de560d1: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getForumTopics#de560d1: field channel: %w", err)
		}
		g.Channel = value
	}
	if g.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getForumTopics#de560d1: field q: %w", err)
		}
		g.Q = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getForumTopics#de560d1: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getForumTopics#de560d1: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getForumTopics#de560d1: field offset_topic: %w", err)
		}
		g.OffsetTopic = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getForumTopics#de560d1: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (g *ChannelsGetForumTopicsRequest) GetChannel() (value InputChannelClass) {
	if g == nil {
		return
	}
	return g.Channel
}

// SetQ sets value of Q conditional field.
func (g *ChannelsGetForumTopicsRequest) SetQ(value string) {
	g.Flags.Set(0)
	g.Q = value
}

// GetQ returns value of Q conditional field and
// boolean which is true if field was set.
func (g *ChannelsGetForumTopicsRequest) GetQ() (value string, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Q, true
}

// GetOffsetDate returns value of OffsetDate field.
func (g *ChannelsGetForumTopicsRequest) GetOffsetDate() (value int) {
	if g == nil {
		return
	}
	return g.OffsetDate
}

// GetOffsetID returns value of OffsetID field.
func (g *ChannelsGetForumTopicsRequest) GetOffsetID() (value int) {
	if g == nil {
		return
	}
	return g.OffsetID
}

// GetOffsetTopic returns value of OffsetTopic field.
func (g *ChannelsGetForumTopicsRequest) GetOffsetTopic() (value int) {
	if g == nil {
		return
	}
	return g.OffsetTopic
}

// GetLimit returns value of Limit field.
func (g *ChannelsGetForumTopicsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (g *ChannelsGetForumTopicsRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return g.Channel.AsNotEmpty()
}

// ChannelsGetForumTopics invokes method channels.getForumTopics#de560d1 returning error if any.
// Get topics of a forum¹
//
// Links:
//  1. https://core.telegram.org/api/forum
//
// Possible errors:
//
//	400 CHANNEL_FORUM_MISSING: This supergroup is not a forum.
//	400 CHANNEL_INVALID: The provided channel is invalid.
//
// See https://core.telegram.org/method/channels.getForumTopics for reference.
// Can be used by bots.
func (c *Client) ChannelsGetForumTopics(ctx context.Context, request *ChannelsGetForumTopicsRequest) (*MessagesForumTopics, error) {
	var result MessagesForumTopics

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
