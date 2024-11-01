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

// MessagesInvitedUsers represents TL type `messages.invitedUsers#7f5defa6`.
// Contains info about successfully or unsuccessfully invited »¹ users.
//
// Links:
//  1. https://core.telegram.org/api/invites#direct-invites
//
// See https://core.telegram.org/constructor/messages.invitedUsers for reference.
type MessagesInvitedUsers struct {
	// List of updates about successfully invited users (and eventually info about the
	// created group)
	Updates UpdatesClass
	// A list of users that could not be invited, along with the reason why they couldn't be
	// invited.
	MissingInvitees []MissingInvitee
}

// MessagesInvitedUsersTypeID is TL type id of MessagesInvitedUsers.
const MessagesInvitedUsersTypeID = 0x7f5defa6

// Ensuring interfaces in compile-time for MessagesInvitedUsers.
var (
	_ bin.Encoder     = &MessagesInvitedUsers{}
	_ bin.Decoder     = &MessagesInvitedUsers{}
	_ bin.BareEncoder = &MessagesInvitedUsers{}
	_ bin.BareDecoder = &MessagesInvitedUsers{}
)

func (i *MessagesInvitedUsers) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Updates == nil) {
		return false
	}
	if !(i.MissingInvitees == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *MessagesInvitedUsers) String() string {
	if i == nil {
		return "MessagesInvitedUsers(nil)"
	}
	type Alias MessagesInvitedUsers
	return fmt.Sprintf("MessagesInvitedUsers%+v", Alias(*i))
}

// FillFrom fills MessagesInvitedUsers from given interface.
func (i *MessagesInvitedUsers) FillFrom(from interface {
	GetUpdates() (value UpdatesClass)
	GetMissingInvitees() (value []MissingInvitee)
}) {
	i.Updates = from.GetUpdates()
	i.MissingInvitees = from.GetMissingInvitees()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesInvitedUsers) TypeID() uint32 {
	return MessagesInvitedUsersTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesInvitedUsers) TypeName() string {
	return "messages.invitedUsers"
}

// TypeInfo returns info about TL type.
func (i *MessagesInvitedUsers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.invitedUsers",
		ID:   MessagesInvitedUsersTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Updates",
			SchemaName: "updates",
		},
		{
			Name:       "MissingInvitees",
			SchemaName: "missing_invitees",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *MessagesInvitedUsers) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.invitedUsers#7f5defa6 as nil")
	}
	b.PutID(MessagesInvitedUsersTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *MessagesInvitedUsers) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.invitedUsers#7f5defa6 as nil")
	}
	if i.Updates == nil {
		return fmt.Errorf("unable to encode messages.invitedUsers#7f5defa6: field updates is nil")
	}
	if err := i.Updates.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.invitedUsers#7f5defa6: field updates: %w", err)
	}
	b.PutVectorHeader(len(i.MissingInvitees))
	for idx, v := range i.MissingInvitees {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.invitedUsers#7f5defa6: field missing_invitees element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *MessagesInvitedUsers) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.invitedUsers#7f5defa6 to nil")
	}
	if err := b.ConsumeID(MessagesInvitedUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.invitedUsers#7f5defa6: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *MessagesInvitedUsers) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.invitedUsers#7f5defa6 to nil")
	}
	{
		value, err := DecodeUpdates(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.invitedUsers#7f5defa6: field updates: %w", err)
		}
		i.Updates = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.invitedUsers#7f5defa6: field missing_invitees: %w", err)
		}

		if headerLen > 0 {
			i.MissingInvitees = make([]MissingInvitee, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MissingInvitee
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.invitedUsers#7f5defa6: field missing_invitees: %w", err)
			}
			i.MissingInvitees = append(i.MissingInvitees, value)
		}
	}
	return nil
}

// GetUpdates returns value of Updates field.
func (i *MessagesInvitedUsers) GetUpdates() (value UpdatesClass) {
	if i == nil {
		return
	}
	return i.Updates
}

// GetMissingInvitees returns value of MissingInvitees field.
func (i *MessagesInvitedUsers) GetMissingInvitees() (value []MissingInvitee) {
	if i == nil {
		return
	}
	return i.MissingInvitees
}