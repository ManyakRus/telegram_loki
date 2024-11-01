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

// QuickReply represents TL type `quickReply#697102b`.
// A quick reply shortcut¹.
//
// Links:
//  1. https://core.telegram.org/api/business#quick-reply-shortcuts
//
// See https://core.telegram.org/constructor/quickReply for reference.
type QuickReply struct {
	// Unique shortcut ID.
	ShortcutID int
	// Shortcut name.
	Shortcut string
	// ID of the last message in the shortcut.
	TopMessage int
	// Total number of messages in the shortcut.
	Count int
}

// QuickReplyTypeID is TL type id of QuickReply.
const QuickReplyTypeID = 0x697102b

// Ensuring interfaces in compile-time for QuickReply.
var (
	_ bin.Encoder     = &QuickReply{}
	_ bin.Decoder     = &QuickReply{}
	_ bin.BareEncoder = &QuickReply{}
	_ bin.BareDecoder = &QuickReply{}
)

func (q *QuickReply) Zero() bool {
	if q == nil {
		return true
	}
	if !(q.ShortcutID == 0) {
		return false
	}
	if !(q.Shortcut == "") {
		return false
	}
	if !(q.TopMessage == 0) {
		return false
	}
	if !(q.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (q *QuickReply) String() string {
	if q == nil {
		return "QuickReply(nil)"
	}
	type Alias QuickReply
	return fmt.Sprintf("QuickReply%+v", Alias(*q))
}

// FillFrom fills QuickReply from given interface.
func (q *QuickReply) FillFrom(from interface {
	GetShortcutID() (value int)
	GetShortcut() (value string)
	GetTopMessage() (value int)
	GetCount() (value int)
}) {
	q.ShortcutID = from.GetShortcutID()
	q.Shortcut = from.GetShortcut()
	q.TopMessage = from.GetTopMessage()
	q.Count = from.GetCount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*QuickReply) TypeID() uint32 {
	return QuickReplyTypeID
}

// TypeName returns name of type in TL schema.
func (*QuickReply) TypeName() string {
	return "quickReply"
}

// TypeInfo returns info about TL type.
func (q *QuickReply) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "quickReply",
		ID:   QuickReplyTypeID,
	}
	if q == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShortcutID",
			SchemaName: "shortcut_id",
		},
		{
			Name:       "Shortcut",
			SchemaName: "shortcut",
		},
		{
			Name:       "TopMessage",
			SchemaName: "top_message",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (q *QuickReply) Encode(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't encode quickReply#697102b as nil")
	}
	b.PutID(QuickReplyTypeID)
	return q.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (q *QuickReply) EncodeBare(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't encode quickReply#697102b as nil")
	}
	b.PutInt(q.ShortcutID)
	b.PutString(q.Shortcut)
	b.PutInt(q.TopMessage)
	b.PutInt(q.Count)
	return nil
}

// Decode implements bin.Decoder.
func (q *QuickReply) Decode(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't decode quickReply#697102b to nil")
	}
	if err := b.ConsumeID(QuickReplyTypeID); err != nil {
		return fmt.Errorf("unable to decode quickReply#697102b: %w", err)
	}
	return q.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (q *QuickReply) DecodeBare(b *bin.Buffer) error {
	if q == nil {
		return fmt.Errorf("can't decode quickReply#697102b to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode quickReply#697102b: field shortcut_id: %w", err)
		}
		q.ShortcutID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode quickReply#697102b: field shortcut: %w", err)
		}
		q.Shortcut = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode quickReply#697102b: field top_message: %w", err)
		}
		q.TopMessage = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode quickReply#697102b: field count: %w", err)
		}
		q.Count = value
	}
	return nil
}

// GetShortcutID returns value of ShortcutID field.
func (q *QuickReply) GetShortcutID() (value int) {
	if q == nil {
		return
	}
	return q.ShortcutID
}

// GetShortcut returns value of Shortcut field.
func (q *QuickReply) GetShortcut() (value string) {
	if q == nil {
		return
	}
	return q.Shortcut
}

// GetTopMessage returns value of TopMessage field.
func (q *QuickReply) GetTopMessage() (value int) {
	if q == nil {
		return
	}
	return q.TopMessage
}

// GetCount returns value of Count field.
func (q *QuickReply) GetCount() (value int) {
	if q == nil {
		return
	}
	return q.Count
}