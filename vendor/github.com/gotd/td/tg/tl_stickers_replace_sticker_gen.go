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

// StickersReplaceStickerRequest represents TL type `stickers.replaceSticker#4696459a`.
// Replace a sticker in a stickerset »¹.
//
// Links:
//  1. https://core.telegram.org/api/stickers
//
// See https://core.telegram.org/method/stickers.replaceSticker for reference.
type StickersReplaceStickerRequest struct {
	// Old sticker document.
	Sticker InputDocumentClass
	// New sticker.
	NewSticker InputStickerSetItem
}

// StickersReplaceStickerRequestTypeID is TL type id of StickersReplaceStickerRequest.
const StickersReplaceStickerRequestTypeID = 0x4696459a

// Ensuring interfaces in compile-time for StickersReplaceStickerRequest.
var (
	_ bin.Encoder     = &StickersReplaceStickerRequest{}
	_ bin.Decoder     = &StickersReplaceStickerRequest{}
	_ bin.BareEncoder = &StickersReplaceStickerRequest{}
	_ bin.BareDecoder = &StickersReplaceStickerRequest{}
)

func (r *StickersReplaceStickerRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Sticker == nil) {
		return false
	}
	if !(r.NewSticker.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *StickersReplaceStickerRequest) String() string {
	if r == nil {
		return "StickersReplaceStickerRequest(nil)"
	}
	type Alias StickersReplaceStickerRequest
	return fmt.Sprintf("StickersReplaceStickerRequest%+v", Alias(*r))
}

// FillFrom fills StickersReplaceStickerRequest from given interface.
func (r *StickersReplaceStickerRequest) FillFrom(from interface {
	GetSticker() (value InputDocumentClass)
	GetNewSticker() (value InputStickerSetItem)
}) {
	r.Sticker = from.GetSticker()
	r.NewSticker = from.GetNewSticker()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickersReplaceStickerRequest) TypeID() uint32 {
	return StickersReplaceStickerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StickersReplaceStickerRequest) TypeName() string {
	return "stickers.replaceSticker"
}

// TypeInfo returns info about TL type.
func (r *StickersReplaceStickerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickers.replaceSticker",
		ID:   StickersReplaceStickerRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
		{
			Name:       "NewSticker",
			SchemaName: "new_sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *StickersReplaceStickerRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode stickers.replaceSticker#4696459a as nil")
	}
	b.PutID(StickersReplaceStickerRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *StickersReplaceStickerRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode stickers.replaceSticker#4696459a as nil")
	}
	if r.Sticker == nil {
		return fmt.Errorf("unable to encode stickers.replaceSticker#4696459a: field sticker is nil")
	}
	if err := r.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.replaceSticker#4696459a: field sticker: %w", err)
	}
	if err := r.NewSticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.replaceSticker#4696459a: field new_sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *StickersReplaceStickerRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode stickers.replaceSticker#4696459a to nil")
	}
	if err := b.ConsumeID(StickersReplaceStickerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.replaceSticker#4696459a: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *StickersReplaceStickerRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode stickers.replaceSticker#4696459a to nil")
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.replaceSticker#4696459a: field sticker: %w", err)
		}
		r.Sticker = value
	}
	{
		if err := r.NewSticker.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickers.replaceSticker#4696459a: field new_sticker: %w", err)
		}
	}
	return nil
}

// GetSticker returns value of Sticker field.
func (r *StickersReplaceStickerRequest) GetSticker() (value InputDocumentClass) {
	if r == nil {
		return
	}
	return r.Sticker
}

// GetNewSticker returns value of NewSticker field.
func (r *StickersReplaceStickerRequest) GetNewSticker() (value InputStickerSetItem) {
	if r == nil {
		return
	}
	return r.NewSticker
}

// GetStickerAsNotEmpty returns mapped value of Sticker field.
func (r *StickersReplaceStickerRequest) GetStickerAsNotEmpty() (*InputDocument, bool) {
	return r.Sticker.AsNotEmpty()
}

// StickersReplaceSticker invokes method stickers.replaceSticker#4696459a returning error if any.
// Replace a sticker in a stickerset »¹.
//
// Links:
//  1. https://core.telegram.org/api/stickers
//
// Possible errors:
//
//	400 STICKER_INVALID: The provided sticker is invalid.
//
// See https://core.telegram.org/method/stickers.replaceSticker for reference.
// Can be used by bots.
func (c *Client) StickersReplaceSticker(ctx context.Context, request *StickersReplaceStickerRequest) (MessagesStickerSetClass, error) {
	var result MessagesStickerSetBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.StickerSet, nil
}