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

// StoriesGetAllReadPeerStoriesRequest represents TL type `stories.getAllReadPeerStories#9b5ae7f9`.
// Obtain the latest read story ID for all peers when first logging in, returned as a
// list of updateReadStories¹ updates, see here »² for more info.
//
// Links:
//  1. https://core.telegram.org/constructor/updateReadStories
//  2. https://core.telegram.org/api/stories#watching-stories
//
// See https://core.telegram.org/method/stories.getAllReadPeerStories for reference.
type StoriesGetAllReadPeerStoriesRequest struct {
}

// StoriesGetAllReadPeerStoriesRequestTypeID is TL type id of StoriesGetAllReadPeerStoriesRequest.
const StoriesGetAllReadPeerStoriesRequestTypeID = 0x9b5ae7f9

// Ensuring interfaces in compile-time for StoriesGetAllReadPeerStoriesRequest.
var (
	_ bin.Encoder     = &StoriesGetAllReadPeerStoriesRequest{}
	_ bin.Decoder     = &StoriesGetAllReadPeerStoriesRequest{}
	_ bin.BareEncoder = &StoriesGetAllReadPeerStoriesRequest{}
	_ bin.BareDecoder = &StoriesGetAllReadPeerStoriesRequest{}
)

func (g *StoriesGetAllReadPeerStoriesRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *StoriesGetAllReadPeerStoriesRequest) String() string {
	if g == nil {
		return "StoriesGetAllReadPeerStoriesRequest(nil)"
	}
	type Alias StoriesGetAllReadPeerStoriesRequest
	return fmt.Sprintf("StoriesGetAllReadPeerStoriesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesGetAllReadPeerStoriesRequest) TypeID() uint32 {
	return StoriesGetAllReadPeerStoriesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesGetAllReadPeerStoriesRequest) TypeName() string {
	return "stories.getAllReadPeerStories"
}

// TypeInfo returns info about TL type.
func (g *StoriesGetAllReadPeerStoriesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.getAllReadPeerStories",
		ID:   StoriesGetAllReadPeerStoriesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *StoriesGetAllReadPeerStoriesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getAllReadPeerStories#9b5ae7f9 as nil")
	}
	b.PutID(StoriesGetAllReadPeerStoriesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StoriesGetAllReadPeerStoriesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getAllReadPeerStories#9b5ae7f9 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *StoriesGetAllReadPeerStoriesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getAllReadPeerStories#9b5ae7f9 to nil")
	}
	if err := b.ConsumeID(StoriesGetAllReadPeerStoriesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.getAllReadPeerStories#9b5ae7f9: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StoriesGetAllReadPeerStoriesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getAllReadPeerStories#9b5ae7f9 to nil")
	}
	return nil
}

// StoriesGetAllReadPeerStories invokes method stories.getAllReadPeerStories#9b5ae7f9 returning error if any.
// Obtain the latest read story ID for all peers when first logging in, returned as a
// list of updateReadStories¹ updates, see here »² for more info.
//
// Links:
//  1. https://core.telegram.org/constructor/updateReadStories
//  2. https://core.telegram.org/api/stories#watching-stories
//
// See https://core.telegram.org/method/stories.getAllReadPeerStories for reference.
func (c *Client) StoriesGetAllReadPeerStories(ctx context.Context) (UpdatesClass, error) {
	var result UpdatesBox

	request := &StoriesGetAllReadPeerStoriesRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}