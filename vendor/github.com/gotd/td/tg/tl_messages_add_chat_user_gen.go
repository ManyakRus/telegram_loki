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

// MessagesAddChatUserRequest represents TL type `messages.addChatUser#cbc6d107`.
// Adds a user to a chat and sends a service message on it.
//
// See https://core.telegram.org/method/messages.addChatUser for reference.
type MessagesAddChatUserRequest struct {
	// Chat ID
	ChatID int64
	// User ID to be added
	UserID InputUserClass
	// Number of last messages to be forwarded
	FwdLimit int
}

// MessagesAddChatUserRequestTypeID is TL type id of MessagesAddChatUserRequest.
const MessagesAddChatUserRequestTypeID = 0xcbc6d107

// Ensuring interfaces in compile-time for MessagesAddChatUserRequest.
var (
	_ bin.Encoder     = &MessagesAddChatUserRequest{}
	_ bin.Decoder     = &MessagesAddChatUserRequest{}
	_ bin.BareEncoder = &MessagesAddChatUserRequest{}
	_ bin.BareDecoder = &MessagesAddChatUserRequest{}
)

func (a *MessagesAddChatUserRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ChatID == 0) {
		return false
	}
	if !(a.UserID == nil) {
		return false
	}
	if !(a.FwdLimit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *MessagesAddChatUserRequest) String() string {
	if a == nil {
		return "MessagesAddChatUserRequest(nil)"
	}
	type Alias MessagesAddChatUserRequest
	return fmt.Sprintf("MessagesAddChatUserRequest%+v", Alias(*a))
}

// FillFrom fills MessagesAddChatUserRequest from given interface.
func (a *MessagesAddChatUserRequest) FillFrom(from interface {
	GetChatID() (value int64)
	GetUserID() (value InputUserClass)
	GetFwdLimit() (value int)
}) {
	a.ChatID = from.GetChatID()
	a.UserID = from.GetUserID()
	a.FwdLimit = from.GetFwdLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesAddChatUserRequest) TypeID() uint32 {
	return MessagesAddChatUserRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesAddChatUserRequest) TypeName() string {
	return "messages.addChatUser"
}

// TypeInfo returns info about TL type.
func (a *MessagesAddChatUserRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.addChatUser",
		ID:   MessagesAddChatUserRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "FwdLimit",
			SchemaName: "fwd_limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *MessagesAddChatUserRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.addChatUser#cbc6d107 as nil")
	}
	b.PutID(MessagesAddChatUserRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *MessagesAddChatUserRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.addChatUser#cbc6d107 as nil")
	}
	b.PutLong(a.ChatID)
	if a.UserID == nil {
		return fmt.Errorf("unable to encode messages.addChatUser#cbc6d107: field user_id is nil")
	}
	if err := a.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.addChatUser#cbc6d107: field user_id: %w", err)
	}
	b.PutInt(a.FwdLimit)
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesAddChatUserRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.addChatUser#cbc6d107 to nil")
	}
	if err := b.ConsumeID(MessagesAddChatUserRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.addChatUser#cbc6d107: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *MessagesAddChatUserRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.addChatUser#cbc6d107 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.addChatUser#cbc6d107: field chat_id: %w", err)
		}
		a.ChatID = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.addChatUser#cbc6d107: field user_id: %w", err)
		}
		a.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.addChatUser#cbc6d107: field fwd_limit: %w", err)
		}
		a.FwdLimit = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (a *MessagesAddChatUserRequest) GetChatID() (value int64) {
	if a == nil {
		return
	}
	return a.ChatID
}

// GetUserID returns value of UserID field.
func (a *MessagesAddChatUserRequest) GetUserID() (value InputUserClass) {
	if a == nil {
		return
	}
	return a.UserID
}

// GetFwdLimit returns value of FwdLimit field.
func (a *MessagesAddChatUserRequest) GetFwdLimit() (value int) {
	if a == nil {
		return
	}
	return a.FwdLimit
}

// MessagesAddChatUser invokes method messages.addChatUser#cbc6d107 returning error if any.
// Adds a user to a chat and sends a service message on it.
//
// Possible errors:
//
//	400 BOT_GROUPS_BLOCKED: This bot can't be added to groups.
//	403 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_ID_INVALID: The provided chat id is invalid.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example).
//	400 USER_ALREADY_PARTICIPANT: The user is already in the group.
//	400 USER_ID_INVALID: The provided user ID is invalid.
//	400 USER_IS_BLOCKED: You were blocked by this user.
//	403 USER_NOT_MUTUAL_CONTACT: The provided user is not a mutual contact.
//	403 USER_PRIVACY_RESTRICTED: The user's privacy settings do not allow you to do this.
//	400 YOU_BLOCKED_USER: You blocked this user.
//
// See https://core.telegram.org/method/messages.addChatUser for reference.
func (c *Client) MessagesAddChatUser(ctx context.Context, request *MessagesAddChatUserRequest) (*MessagesInvitedUsers, error) {
	var result MessagesInvitedUsers

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
