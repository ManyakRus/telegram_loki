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

// PhoneGetGroupCallRequest represents TL type `phone.getGroupCall#41845db`.
// Get info about a group call
//
// See https://core.telegram.org/method/phone.getGroupCall for reference.
type PhoneGetGroupCallRequest struct {
	// The group call
	Call InputGroupCall
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// PhoneGetGroupCallRequestTypeID is TL type id of PhoneGetGroupCallRequest.
const PhoneGetGroupCallRequestTypeID = 0x41845db

// Ensuring interfaces in compile-time for PhoneGetGroupCallRequest.
var (
	_ bin.Encoder     = &PhoneGetGroupCallRequest{}
	_ bin.Decoder     = &PhoneGetGroupCallRequest{}
	_ bin.BareEncoder = &PhoneGetGroupCallRequest{}
	_ bin.BareDecoder = &PhoneGetGroupCallRequest{}
)

func (g *PhoneGetGroupCallRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Call.Zero()) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhoneGetGroupCallRequest) String() string {
	if g == nil {
		return "PhoneGetGroupCallRequest(nil)"
	}
	type Alias PhoneGetGroupCallRequest
	return fmt.Sprintf("PhoneGetGroupCallRequest%+v", Alias(*g))
}

// FillFrom fills PhoneGetGroupCallRequest from given interface.
func (g *PhoneGetGroupCallRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
	GetLimit() (value int)
}) {
	g.Call = from.GetCall()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneGetGroupCallRequest) TypeID() uint32 {
	return PhoneGetGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneGetGroupCallRequest) TypeName() string {
	return "phone.getGroupCall"
}

// TypeInfo returns info about TL type.
func (g *PhoneGetGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.getGroupCall",
		ID:   PhoneGetGroupCallRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PhoneGetGroupCallRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.getGroupCall#41845db as nil")
	}
	b.PutID(PhoneGetGroupCallRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PhoneGetGroupCallRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.getGroupCall#41845db as nil")
	}
	if err := g.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.getGroupCall#41845db: field call: %w", err)
	}
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *PhoneGetGroupCallRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.getGroupCall#41845db to nil")
	}
	if err := b.ConsumeID(PhoneGetGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.getGroupCall#41845db: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PhoneGetGroupCallRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.getGroupCall#41845db to nil")
	}
	{
		if err := g.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.getGroupCall#41845db: field call: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.getGroupCall#41845db: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetCall returns value of Call field.
func (g *PhoneGetGroupCallRequest) GetCall() (value InputGroupCall) {
	if g == nil {
		return
	}
	return g.Call
}

// GetLimit returns value of Limit field.
func (g *PhoneGetGroupCallRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// PhoneGetGroupCall invokes method phone.getGroupCall#41845db returning error if any.
// Get info about a group call
//
// Possible errors:
//
//	403 GROUPCALL_FORBIDDEN: The group call has already ended.
//	400 GROUPCALL_INVALID: The specified group call is invalid.
//
// See https://core.telegram.org/method/phone.getGroupCall for reference.
func (c *Client) PhoneGetGroupCall(ctx context.Context, request *PhoneGetGroupCallRequest) (*PhoneGroupCall, error) {
	var result PhoneGroupCall

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
