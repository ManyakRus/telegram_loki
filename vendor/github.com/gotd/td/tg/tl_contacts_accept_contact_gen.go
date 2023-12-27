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

// ContactsAcceptContactRequest represents TL type `contacts.acceptContact#f831a20f`.
// If the add contact action bar is active¹, add that user as contact
//
// Links:
//  1. https://core.telegram.org/api/action-bar#add-contact
//
// See https://core.telegram.org/method/contacts.acceptContact for reference.
type ContactsAcceptContactRequest struct {
	// The user to add as contact
	ID InputUserClass
}

// ContactsAcceptContactRequestTypeID is TL type id of ContactsAcceptContactRequest.
const ContactsAcceptContactRequestTypeID = 0xf831a20f

// Ensuring interfaces in compile-time for ContactsAcceptContactRequest.
var (
	_ bin.Encoder     = &ContactsAcceptContactRequest{}
	_ bin.Decoder     = &ContactsAcceptContactRequest{}
	_ bin.BareEncoder = &ContactsAcceptContactRequest{}
	_ bin.BareDecoder = &ContactsAcceptContactRequest{}
)

func (a *ContactsAcceptContactRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *ContactsAcceptContactRequest) String() string {
	if a == nil {
		return "ContactsAcceptContactRequest(nil)"
	}
	type Alias ContactsAcceptContactRequest
	return fmt.Sprintf("ContactsAcceptContactRequest%+v", Alias(*a))
}

// FillFrom fills ContactsAcceptContactRequest from given interface.
func (a *ContactsAcceptContactRequest) FillFrom(from interface {
	GetID() (value InputUserClass)
}) {
	a.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsAcceptContactRequest) TypeID() uint32 {
	return ContactsAcceptContactRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsAcceptContactRequest) TypeName() string {
	return "contacts.acceptContact"
}

// TypeInfo returns info about TL type.
func (a *ContactsAcceptContactRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.acceptContact",
		ID:   ContactsAcceptContactRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *ContactsAcceptContactRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode contacts.acceptContact#f831a20f as nil")
	}
	b.PutID(ContactsAcceptContactRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *ContactsAcceptContactRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode contacts.acceptContact#f831a20f as nil")
	}
	if a.ID == nil {
		return fmt.Errorf("unable to encode contacts.acceptContact#f831a20f: field id is nil")
	}
	if err := a.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.acceptContact#f831a20f: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *ContactsAcceptContactRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode contacts.acceptContact#f831a20f to nil")
	}
	if err := b.ConsumeID(ContactsAcceptContactRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.acceptContact#f831a20f: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *ContactsAcceptContactRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode contacts.acceptContact#f831a20f to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.acceptContact#f831a20f: field id: %w", err)
		}
		a.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (a *ContactsAcceptContactRequest) GetID() (value InputUserClass) {
	if a == nil {
		return
	}
	return a.ID
}

// ContactsAcceptContact invokes method contacts.acceptContact#f831a20f returning error if any.
// If the add contact action bar is active¹, add that user as contact
//
// Links:
//  1. https://core.telegram.org/api/action-bar#add-contact
//
// Possible errors:
//
//	400 CONTACT_ADD_MISSING: Contact to add is missing.
//	400 CONTACT_ID_INVALID: The provided contact ID is invalid.
//	400 CONTACT_REQ_MISSING: Missing contact request.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//
// See https://core.telegram.org/method/contacts.acceptContact for reference.
func (c *Client) ContactsAcceptContact(ctx context.Context, id InputUserClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ContactsAcceptContactRequest{
		ID: id,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
