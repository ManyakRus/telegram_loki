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

// AccountUpdateConnectedBotRequest represents TL type `account.updateConnectedBot#43d8521d`.
//
// See https://core.telegram.org/method/account.updateConnectedBot for reference.
type AccountUpdateConnectedBotRequest struct {
	// Flags field of AccountUpdateConnectedBotRequest.
	Flags bin.Fields
	// CanReply field of AccountUpdateConnectedBotRequest.
	CanReply bool
	// Deleted field of AccountUpdateConnectedBotRequest.
	Deleted bool
	// Bot field of AccountUpdateConnectedBotRequest.
	Bot InputUserClass
	// Recipients field of AccountUpdateConnectedBotRequest.
	Recipients InputBusinessBotRecipients
}

// AccountUpdateConnectedBotRequestTypeID is TL type id of AccountUpdateConnectedBotRequest.
const AccountUpdateConnectedBotRequestTypeID = 0x43d8521d

// Ensuring interfaces in compile-time for AccountUpdateConnectedBotRequest.
var (
	_ bin.Encoder     = &AccountUpdateConnectedBotRequest{}
	_ bin.Decoder     = &AccountUpdateConnectedBotRequest{}
	_ bin.BareEncoder = &AccountUpdateConnectedBotRequest{}
	_ bin.BareDecoder = &AccountUpdateConnectedBotRequest{}
)

func (u *AccountUpdateConnectedBotRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.CanReply == false) {
		return false
	}
	if !(u.Deleted == false) {
		return false
	}
	if !(u.Bot == nil) {
		return false
	}
	if !(u.Recipients.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateConnectedBotRequest) String() string {
	if u == nil {
		return "AccountUpdateConnectedBotRequest(nil)"
	}
	type Alias AccountUpdateConnectedBotRequest
	return fmt.Sprintf("AccountUpdateConnectedBotRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateConnectedBotRequest from given interface.
func (u *AccountUpdateConnectedBotRequest) FillFrom(from interface {
	GetCanReply() (value bool)
	GetDeleted() (value bool)
	GetBot() (value InputUserClass)
	GetRecipients() (value InputBusinessBotRecipients)
}) {
	u.CanReply = from.GetCanReply()
	u.Deleted = from.GetDeleted()
	u.Bot = from.GetBot()
	u.Recipients = from.GetRecipients()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateConnectedBotRequest) TypeID() uint32 {
	return AccountUpdateConnectedBotRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateConnectedBotRequest) TypeName() string {
	return "account.updateConnectedBot"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateConnectedBotRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateConnectedBot",
		ID:   AccountUpdateConnectedBotRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CanReply",
			SchemaName: "can_reply",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "Deleted",
			SchemaName: "deleted",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "Recipients",
			SchemaName: "recipients",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *AccountUpdateConnectedBotRequest) SetFlags() {
	if !(u.CanReply == false) {
		u.Flags.Set(0)
	}
	if !(u.Deleted == false) {
		u.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (u *AccountUpdateConnectedBotRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateConnectedBot#43d8521d as nil")
	}
	b.PutID(AccountUpdateConnectedBotRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateConnectedBotRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateConnectedBot#43d8521d as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateConnectedBot#43d8521d: field flags: %w", err)
	}
	if u.Bot == nil {
		return fmt.Errorf("unable to encode account.updateConnectedBot#43d8521d: field bot is nil")
	}
	if err := u.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateConnectedBot#43d8521d: field bot: %w", err)
	}
	if err := u.Recipients.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateConnectedBot#43d8521d: field recipients: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateConnectedBotRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateConnectedBot#43d8521d to nil")
	}
	if err := b.ConsumeID(AccountUpdateConnectedBotRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateConnectedBot#43d8521d: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateConnectedBotRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateConnectedBot#43d8521d to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateConnectedBot#43d8521d: field flags: %w", err)
		}
	}
	u.CanReply = u.Flags.Has(0)
	u.Deleted = u.Flags.Has(1)
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.updateConnectedBot#43d8521d: field bot: %w", err)
		}
		u.Bot = value
	}
	{
		if err := u.Recipients.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateConnectedBot#43d8521d: field recipients: %w", err)
		}
	}
	return nil
}

// SetCanReply sets value of CanReply conditional field.
func (u *AccountUpdateConnectedBotRequest) SetCanReply(value bool) {
	if value {
		u.Flags.Set(0)
		u.CanReply = true
	} else {
		u.Flags.Unset(0)
		u.CanReply = false
	}
}

// GetCanReply returns value of CanReply conditional field.
func (u *AccountUpdateConnectedBotRequest) GetCanReply() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// SetDeleted sets value of Deleted conditional field.
func (u *AccountUpdateConnectedBotRequest) SetDeleted(value bool) {
	if value {
		u.Flags.Set(1)
		u.Deleted = true
	} else {
		u.Flags.Unset(1)
		u.Deleted = false
	}
}

// GetDeleted returns value of Deleted conditional field.
func (u *AccountUpdateConnectedBotRequest) GetDeleted() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(1)
}

// GetBot returns value of Bot field.
func (u *AccountUpdateConnectedBotRequest) GetBot() (value InputUserClass) {
	if u == nil {
		return
	}
	return u.Bot
}

// GetRecipients returns value of Recipients field.
func (u *AccountUpdateConnectedBotRequest) GetRecipients() (value InputBusinessBotRecipients) {
	if u == nil {
		return
	}
	return u.Recipients
}

// AccountUpdateConnectedBot invokes method account.updateConnectedBot#43d8521d returning error if any.
//
// See https://core.telegram.org/method/account.updateConnectedBot for reference.
func (c *Client) AccountUpdateConnectedBot(ctx context.Context, request *AccountUpdateConnectedBotRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
