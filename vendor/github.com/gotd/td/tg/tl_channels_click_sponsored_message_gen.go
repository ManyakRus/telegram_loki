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

// ChannelsClickSponsoredMessageRequest represents TL type `channels.clickSponsoredMessage#18afbc93`.
// Informs the server that the user has either:
//
// See https://core.telegram.org/method/channels.clickSponsoredMessage for reference.
type ChannelsClickSponsoredMessageRequest struct {
	// Channel where the sponsored message was posted
	Channel InputChannelClass
	// Message ID
	RandomID []byte
}

// ChannelsClickSponsoredMessageRequestTypeID is TL type id of ChannelsClickSponsoredMessageRequest.
const ChannelsClickSponsoredMessageRequestTypeID = 0x18afbc93

// Ensuring interfaces in compile-time for ChannelsClickSponsoredMessageRequest.
var (
	_ bin.Encoder     = &ChannelsClickSponsoredMessageRequest{}
	_ bin.Decoder     = &ChannelsClickSponsoredMessageRequest{}
	_ bin.BareEncoder = &ChannelsClickSponsoredMessageRequest{}
	_ bin.BareDecoder = &ChannelsClickSponsoredMessageRequest{}
)

func (c *ChannelsClickSponsoredMessageRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Channel == nil) {
		return false
	}
	if !(c.RandomID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelsClickSponsoredMessageRequest) String() string {
	if c == nil {
		return "ChannelsClickSponsoredMessageRequest(nil)"
	}
	type Alias ChannelsClickSponsoredMessageRequest
	return fmt.Sprintf("ChannelsClickSponsoredMessageRequest%+v", Alias(*c))
}

// FillFrom fills ChannelsClickSponsoredMessageRequest from given interface.
func (c *ChannelsClickSponsoredMessageRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetRandomID() (value []byte)
}) {
	c.Channel = from.GetChannel()
	c.RandomID = from.GetRandomID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsClickSponsoredMessageRequest) TypeID() uint32 {
	return ChannelsClickSponsoredMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsClickSponsoredMessageRequest) TypeName() string {
	return "channels.clickSponsoredMessage"
}

// TypeInfo returns info about TL type.
func (c *ChannelsClickSponsoredMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.clickSponsoredMessage",
		ID:   ChannelsClickSponsoredMessageRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelsClickSponsoredMessageRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.clickSponsoredMessage#18afbc93 as nil")
	}
	b.PutID(ChannelsClickSponsoredMessageRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelsClickSponsoredMessageRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.clickSponsoredMessage#18afbc93 as nil")
	}
	if c.Channel == nil {
		return fmt.Errorf("unable to encode channels.clickSponsoredMessage#18afbc93: field channel is nil")
	}
	if err := c.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.clickSponsoredMessage#18afbc93: field channel: %w", err)
	}
	b.PutBytes(c.RandomID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelsClickSponsoredMessageRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.clickSponsoredMessage#18afbc93 to nil")
	}
	if err := b.ConsumeID(ChannelsClickSponsoredMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.clickSponsoredMessage#18afbc93: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelsClickSponsoredMessageRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.clickSponsoredMessage#18afbc93 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.clickSponsoredMessage#18afbc93: field channel: %w", err)
		}
		c.Channel = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode channels.clickSponsoredMessage#18afbc93: field random_id: %w", err)
		}
		c.RandomID = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (c *ChannelsClickSponsoredMessageRequest) GetChannel() (value InputChannelClass) {
	if c == nil {
		return
	}
	return c.Channel
}

// GetRandomID returns value of RandomID field.
func (c *ChannelsClickSponsoredMessageRequest) GetRandomID() (value []byte) {
	if c == nil {
		return
	}
	return c.RandomID
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (c *ChannelsClickSponsoredMessageRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return c.Channel.AsNotEmpty()
}

// ChannelsClickSponsoredMessage invokes method channels.clickSponsoredMessage#18afbc93 returning error if any.
// Informs the server that the user has either:
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//
// See https://core.telegram.org/method/channels.clickSponsoredMessage for reference.
func (c *Client) ChannelsClickSponsoredMessage(ctx context.Context, request *ChannelsClickSponsoredMessageRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
