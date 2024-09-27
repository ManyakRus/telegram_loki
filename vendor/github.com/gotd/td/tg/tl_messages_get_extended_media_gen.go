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

// MessagesGetExtendedMediaRequest represents TL type `messages.getExtendedMedia#84f80814`.
// Fetch updated information about paid media, see here »¹ for the full flow.
// This method will return an array of updateMessageExtendedMedia¹ updates, only for
// messages containing already bought paid media.
// No information will be returned for messages containing not yet bought paid media.
//
// Links:
//  1. https://core.telegram.org/api/paid-media
//  2. https://core.telegram.org/constructor/updateMessageExtendedMedia
//
// See https://core.telegram.org/method/messages.getExtendedMedia for reference.
type MessagesGetExtendedMediaRequest struct {
	// Peer with visible paid media messages.
	Peer InputPeerClass
	// IDs of currently visible messages containing paid media.
	ID []int
}

// MessagesGetExtendedMediaRequestTypeID is TL type id of MessagesGetExtendedMediaRequest.
const MessagesGetExtendedMediaRequestTypeID = 0x84f80814

// Ensuring interfaces in compile-time for MessagesGetExtendedMediaRequest.
var (
	_ bin.Encoder     = &MessagesGetExtendedMediaRequest{}
	_ bin.Decoder     = &MessagesGetExtendedMediaRequest{}
	_ bin.BareEncoder = &MessagesGetExtendedMediaRequest{}
	_ bin.BareDecoder = &MessagesGetExtendedMediaRequest{}
)

func (g *MessagesGetExtendedMediaRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetExtendedMediaRequest) String() string {
	if g == nil {
		return "MessagesGetExtendedMediaRequest(nil)"
	}
	type Alias MessagesGetExtendedMediaRequest
	return fmt.Sprintf("MessagesGetExtendedMediaRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetExtendedMediaRequest from given interface.
func (g *MessagesGetExtendedMediaRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value []int)
}) {
	g.Peer = from.GetPeer()
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetExtendedMediaRequest) TypeID() uint32 {
	return MessagesGetExtendedMediaRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetExtendedMediaRequest) TypeName() string {
	return "messages.getExtendedMedia"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetExtendedMediaRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getExtendedMedia",
		ID:   MessagesGetExtendedMediaRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetExtendedMediaRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getExtendedMedia#84f80814 as nil")
	}
	b.PutID(MessagesGetExtendedMediaRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetExtendedMediaRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getExtendedMedia#84f80814 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getExtendedMedia#84f80814: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getExtendedMedia#84f80814: field peer: %w", err)
	}
	b.PutVectorHeader(len(g.ID))
	for _, v := range g.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetExtendedMediaRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getExtendedMedia#84f80814 to nil")
	}
	if err := b.ConsumeID(MessagesGetExtendedMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getExtendedMedia#84f80814: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetExtendedMediaRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getExtendedMedia#84f80814 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getExtendedMedia#84f80814: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getExtendedMedia#84f80814: field id: %w", err)
		}

		if headerLen > 0 {
			g.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getExtendedMedia#84f80814: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetExtendedMediaRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetID returns value of ID field.
func (g *MessagesGetExtendedMediaRequest) GetID() (value []int) {
	if g == nil {
		return
	}
	return g.ID
}

// MessagesGetExtendedMedia invokes method messages.getExtendedMedia#84f80814 returning error if any.
// Fetch updated information about paid media, see here »¹ for the full flow.
// This method will return an array of updateMessageExtendedMedia¹ updates, only for
// messages containing already bought paid media.
// No information will be returned for messages containing not yet bought paid media.
//
// Links:
//  1. https://core.telegram.org/api/paid-media
//  2. https://core.telegram.org/constructor/updateMessageExtendedMedia
//
// See https://core.telegram.org/method/messages.getExtendedMedia for reference.
func (c *Client) MessagesGetExtendedMedia(ctx context.Context, request *MessagesGetExtendedMediaRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
