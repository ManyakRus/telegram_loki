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

// PrepaidGiveaway represents TL type `prepaidGiveaway#b2539d54`.
// Contains info about a prepaid giveaway »¹.
//
// Links:
//  1. https://core.telegram.org/api/giveaways
//
// See https://core.telegram.org/constructor/prepaidGiveaway for reference.
type PrepaidGiveaway struct {
	// Prepaid giveaway ID.
	ID int64
	// Duration in months of each gifted Telegram Premium¹ subscription.
	//
	// Links:
	//  1) https://core.telegram.org/api/premium
	Months int
	// Number of given away Telegram Premium¹ subscriptions.
	//
	// Links:
	//  1) https://core.telegram.org/api/premium
	Quantity int
	// Payment date.
	Date int
}

// PrepaidGiveawayTypeID is TL type id of PrepaidGiveaway.
const PrepaidGiveawayTypeID = 0xb2539d54

// Ensuring interfaces in compile-time for PrepaidGiveaway.
var (
	_ bin.Encoder     = &PrepaidGiveaway{}
	_ bin.Decoder     = &PrepaidGiveaway{}
	_ bin.BareEncoder = &PrepaidGiveaway{}
	_ bin.BareDecoder = &PrepaidGiveaway{}
)

func (p *PrepaidGiveaway) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == 0) {
		return false
	}
	if !(p.Months == 0) {
		return false
	}
	if !(p.Quantity == 0) {
		return false
	}
	if !(p.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrepaidGiveaway) String() string {
	if p == nil {
		return "PrepaidGiveaway(nil)"
	}
	type Alias PrepaidGiveaway
	return fmt.Sprintf("PrepaidGiveaway%+v", Alias(*p))
}

// FillFrom fills PrepaidGiveaway from given interface.
func (p *PrepaidGiveaway) FillFrom(from interface {
	GetID() (value int64)
	GetMonths() (value int)
	GetQuantity() (value int)
	GetDate() (value int)
}) {
	p.ID = from.GetID()
	p.Months = from.GetMonths()
	p.Quantity = from.GetQuantity()
	p.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PrepaidGiveaway) TypeID() uint32 {
	return PrepaidGiveawayTypeID
}

// TypeName returns name of type in TL schema.
func (*PrepaidGiveaway) TypeName() string {
	return "prepaidGiveaway"
}

// TypeInfo returns info about TL type.
func (p *PrepaidGiveaway) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "prepaidGiveaway",
		ID:   PrepaidGiveawayTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Months",
			SchemaName: "months",
		},
		{
			Name:       "Quantity",
			SchemaName: "quantity",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PrepaidGiveaway) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode prepaidGiveaway#b2539d54 as nil")
	}
	b.PutID(PrepaidGiveawayTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PrepaidGiveaway) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode prepaidGiveaway#b2539d54 as nil")
	}
	b.PutLong(p.ID)
	b.PutInt(p.Months)
	b.PutInt(p.Quantity)
	b.PutInt(p.Date)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrepaidGiveaway) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode prepaidGiveaway#b2539d54 to nil")
	}
	if err := b.ConsumeID(PrepaidGiveawayTypeID); err != nil {
		return fmt.Errorf("unable to decode prepaidGiveaway#b2539d54: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PrepaidGiveaway) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode prepaidGiveaway#b2539d54 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode prepaidGiveaway#b2539d54: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode prepaidGiveaway#b2539d54: field months: %w", err)
		}
		p.Months = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode prepaidGiveaway#b2539d54: field quantity: %w", err)
		}
		p.Quantity = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode prepaidGiveaway#b2539d54: field date: %w", err)
		}
		p.Date = value
	}
	return nil
}

// GetID returns value of ID field.
func (p *PrepaidGiveaway) GetID() (value int64) {
	if p == nil {
		return
	}
	return p.ID
}

// GetMonths returns value of Months field.
func (p *PrepaidGiveaway) GetMonths() (value int) {
	if p == nil {
		return
	}
	return p.Months
}

// GetQuantity returns value of Quantity field.
func (p *PrepaidGiveaway) GetQuantity() (value int) {
	if p == nil {
		return
	}
	return p.Quantity
}

// GetDate returns value of Date field.
func (p *PrepaidGiveaway) GetDate() (value int) {
	if p == nil {
		return
	}
	return p.Date
}
