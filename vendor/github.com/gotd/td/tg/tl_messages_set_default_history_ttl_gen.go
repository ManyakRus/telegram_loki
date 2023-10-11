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

// MessagesSetDefaultHistoryTTLRequest represents TL type `messages.setDefaultHistoryTTL#9eb51445`.
// Changes the default value of the Time-To-Live setting, applied to all new chats.
//
// See https://core.telegram.org/method/messages.setDefaultHistoryTTL for reference.
type MessagesSetDefaultHistoryTTLRequest struct {
	// The new default Time-To-Live of all messages sent in new chats.
	Period int
}

// MessagesSetDefaultHistoryTTLRequestTypeID is TL type id of MessagesSetDefaultHistoryTTLRequest.
const MessagesSetDefaultHistoryTTLRequestTypeID = 0x9eb51445

// Ensuring interfaces in compile-time for MessagesSetDefaultHistoryTTLRequest.
var (
	_ bin.Encoder     = &MessagesSetDefaultHistoryTTLRequest{}
	_ bin.Decoder     = &MessagesSetDefaultHistoryTTLRequest{}
	_ bin.BareEncoder = &MessagesSetDefaultHistoryTTLRequest{}
	_ bin.BareDecoder = &MessagesSetDefaultHistoryTTLRequest{}
)

func (s *MessagesSetDefaultHistoryTTLRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Period == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetDefaultHistoryTTLRequest) String() string {
	if s == nil {
		return "MessagesSetDefaultHistoryTTLRequest(nil)"
	}
	type Alias MessagesSetDefaultHistoryTTLRequest
	return fmt.Sprintf("MessagesSetDefaultHistoryTTLRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetDefaultHistoryTTLRequest from given interface.
func (s *MessagesSetDefaultHistoryTTLRequest) FillFrom(from interface {
	GetPeriod() (value int)
}) {
	s.Period = from.GetPeriod()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetDefaultHistoryTTLRequest) TypeID() uint32 {
	return MessagesSetDefaultHistoryTTLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetDefaultHistoryTTLRequest) TypeName() string {
	return "messages.setDefaultHistoryTTL"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetDefaultHistoryTTLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setDefaultHistoryTTL",
		ID:   MessagesSetDefaultHistoryTTLRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Period",
			SchemaName: "period",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetDefaultHistoryTTLRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setDefaultHistoryTTL#9eb51445 as nil")
	}
	b.PutID(MessagesSetDefaultHistoryTTLRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetDefaultHistoryTTLRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setDefaultHistoryTTL#9eb51445 as nil")
	}
	b.PutInt(s.Period)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSetDefaultHistoryTTLRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setDefaultHistoryTTL#9eb51445 to nil")
	}
	if err := b.ConsumeID(MessagesSetDefaultHistoryTTLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setDefaultHistoryTTL#9eb51445: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetDefaultHistoryTTLRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setDefaultHistoryTTL#9eb51445 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setDefaultHistoryTTL#9eb51445: field period: %w", err)
		}
		s.Period = value
	}
	return nil
}

// GetPeriod returns value of Period field.
func (s *MessagesSetDefaultHistoryTTLRequest) GetPeriod() (value int) {
	if s == nil {
		return
	}
	return s.Period
}

// MessagesSetDefaultHistoryTTL invokes method messages.setDefaultHistoryTTL#9eb51445 returning error if any.
// Changes the default value of the Time-To-Live setting, applied to all new chats.
//
// See https://core.telegram.org/method/messages.setDefaultHistoryTTL for reference.
// Can be used by bots.
func (c *Client) MessagesSetDefaultHistoryTTL(ctx context.Context, period int) (bool, error) {
	var result BoolBox

	request := &MessagesSetDefaultHistoryTTLRequest{
		Period: period,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
