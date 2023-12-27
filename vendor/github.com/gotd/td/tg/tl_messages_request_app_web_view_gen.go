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

// MessagesRequestAppWebViewRequest represents TL type `messages.requestAppWebView#8c5a3b3c`.
// Open a bot mini app¹ from a named Mini App deep link², sending over user information
// after user confirmation.
// After calling this method, until the user closes the webview, messages
// prolongWebView¹ must be called every 60 seconds.
//
// Links:
//  1. https://core.telegram.org/bots/webapps
//  2. https://core.telegram.org/api/links#named-mini-app-links
//  3. https://core.telegram.org/method/messages.prolongWebView
//
// See https://core.telegram.org/method/messages.requestAppWebView for reference.
type MessagesRequestAppWebViewRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set this flag if the bot is asking permission to send messages to the user as
	// specified in the named Mini App deep link¹ docs, and the user agreed.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#named-mini-app-links
	WriteAllowed bool
	// If the client has clicked on the link in a Telegram chat, pass the chat's peer
	// information; otherwise pass the bot's peer information, instead.
	Peer InputPeerClass
	// The app obtained by invoking messages.getBotApp¹ as specified in the named Mini App
	// deep link² docs.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.getBotApp
	//  2) https://core.telegram.org/api/links#named-mini-app-links
	App InputBotAppClass
	// If the startapp query string parameter is present in the named Mini App deep link¹,
	// pass it to start_param.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#named-mini-app-links
	//
	// Use SetStartParam and GetStartParam helpers.
	StartParam string
	// Theme parameters »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/webapps#theme-parameters
	//
	// Use SetThemeParams and GetThemeParams helpers.
	ThemeParams DataJSON
	// Short name of the application; 0-64 English letters, digits, and underscores
	Platform string
}

// MessagesRequestAppWebViewRequestTypeID is TL type id of MessagesRequestAppWebViewRequest.
const MessagesRequestAppWebViewRequestTypeID = 0x8c5a3b3c

// Ensuring interfaces in compile-time for MessagesRequestAppWebViewRequest.
var (
	_ bin.Encoder     = &MessagesRequestAppWebViewRequest{}
	_ bin.Decoder     = &MessagesRequestAppWebViewRequest{}
	_ bin.BareEncoder = &MessagesRequestAppWebViewRequest{}
	_ bin.BareDecoder = &MessagesRequestAppWebViewRequest{}
)

func (r *MessagesRequestAppWebViewRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.WriteAllowed == false) {
		return false
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.App == nil) {
		return false
	}
	if !(r.StartParam == "") {
		return false
	}
	if !(r.ThemeParams.Zero()) {
		return false
	}
	if !(r.Platform == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesRequestAppWebViewRequest) String() string {
	if r == nil {
		return "MessagesRequestAppWebViewRequest(nil)"
	}
	type Alias MessagesRequestAppWebViewRequest
	return fmt.Sprintf("MessagesRequestAppWebViewRequest%+v", Alias(*r))
}

// FillFrom fills MessagesRequestAppWebViewRequest from given interface.
func (r *MessagesRequestAppWebViewRequest) FillFrom(from interface {
	GetWriteAllowed() (value bool)
	GetPeer() (value InputPeerClass)
	GetApp() (value InputBotAppClass)
	GetStartParam() (value string, ok bool)
	GetThemeParams() (value DataJSON, ok bool)
	GetPlatform() (value string)
}) {
	r.WriteAllowed = from.GetWriteAllowed()
	r.Peer = from.GetPeer()
	r.App = from.GetApp()
	if val, ok := from.GetStartParam(); ok {
		r.StartParam = val
	}

	if val, ok := from.GetThemeParams(); ok {
		r.ThemeParams = val
	}

	r.Platform = from.GetPlatform()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesRequestAppWebViewRequest) TypeID() uint32 {
	return MessagesRequestAppWebViewRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesRequestAppWebViewRequest) TypeName() string {
	return "messages.requestAppWebView"
}

// TypeInfo returns info about TL type.
func (r *MessagesRequestAppWebViewRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.requestAppWebView",
		ID:   MessagesRequestAppWebViewRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "WriteAllowed",
			SchemaName: "write_allowed",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "App",
			SchemaName: "app",
		},
		{
			Name:       "StartParam",
			SchemaName: "start_param",
			Null:       !r.Flags.Has(1),
		},
		{
			Name:       "ThemeParams",
			SchemaName: "theme_params",
			Null:       !r.Flags.Has(2),
		},
		{
			Name:       "Platform",
			SchemaName: "platform",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *MessagesRequestAppWebViewRequest) SetFlags() {
	if !(r.WriteAllowed == false) {
		r.Flags.Set(0)
	}
	if !(r.StartParam == "") {
		r.Flags.Set(1)
	}
	if !(r.ThemeParams.Zero()) {
		r.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (r *MessagesRequestAppWebViewRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestAppWebView#8c5a3b3c as nil")
	}
	b.PutID(MessagesRequestAppWebViewRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesRequestAppWebViewRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestAppWebView#8c5a3b3c as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestAppWebView#8c5a3b3c: field flags: %w", err)
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.requestAppWebView#8c5a3b3c: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestAppWebView#8c5a3b3c: field peer: %w", err)
	}
	if r.App == nil {
		return fmt.Errorf("unable to encode messages.requestAppWebView#8c5a3b3c: field app is nil")
	}
	if err := r.App.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestAppWebView#8c5a3b3c: field app: %w", err)
	}
	if r.Flags.Has(1) {
		b.PutString(r.StartParam)
	}
	if r.Flags.Has(2) {
		if err := r.ThemeParams.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.requestAppWebView#8c5a3b3c: field theme_params: %w", err)
		}
	}
	b.PutString(r.Platform)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRequestAppWebViewRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestAppWebView#8c5a3b3c to nil")
	}
	if err := b.ConsumeID(MessagesRequestAppWebViewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.requestAppWebView#8c5a3b3c: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesRequestAppWebViewRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestAppWebView#8c5a3b3c to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.requestAppWebView#8c5a3b3c: field flags: %w", err)
		}
	}
	r.WriteAllowed = r.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestAppWebView#8c5a3b3c: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := DecodeInputBotApp(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestAppWebView#8c5a3b3c: field app: %w", err)
		}
		r.App = value
	}
	if r.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestAppWebView#8c5a3b3c: field start_param: %w", err)
		}
		r.StartParam = value
	}
	if r.Flags.Has(2) {
		if err := r.ThemeParams.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.requestAppWebView#8c5a3b3c: field theme_params: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestAppWebView#8c5a3b3c: field platform: %w", err)
		}
		r.Platform = value
	}
	return nil
}

// SetWriteAllowed sets value of WriteAllowed conditional field.
func (r *MessagesRequestAppWebViewRequest) SetWriteAllowed(value bool) {
	if value {
		r.Flags.Set(0)
		r.WriteAllowed = true
	} else {
		r.Flags.Unset(0)
		r.WriteAllowed = false
	}
}

// GetWriteAllowed returns value of WriteAllowed conditional field.
func (r *MessagesRequestAppWebViewRequest) GetWriteAllowed() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (r *MessagesRequestAppWebViewRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// GetApp returns value of App field.
func (r *MessagesRequestAppWebViewRequest) GetApp() (value InputBotAppClass) {
	if r == nil {
		return
	}
	return r.App
}

// SetStartParam sets value of StartParam conditional field.
func (r *MessagesRequestAppWebViewRequest) SetStartParam(value string) {
	r.Flags.Set(1)
	r.StartParam = value
}

// GetStartParam returns value of StartParam conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestAppWebViewRequest) GetStartParam() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(1) {
		return value, false
	}
	return r.StartParam, true
}

// SetThemeParams sets value of ThemeParams conditional field.
func (r *MessagesRequestAppWebViewRequest) SetThemeParams(value DataJSON) {
	r.Flags.Set(2)
	r.ThemeParams = value
}

// GetThemeParams returns value of ThemeParams conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestAppWebViewRequest) GetThemeParams() (value DataJSON, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(2) {
		return value, false
	}
	return r.ThemeParams, true
}

// GetPlatform returns value of Platform field.
func (r *MessagesRequestAppWebViewRequest) GetPlatform() (value string) {
	if r == nil {
		return
	}
	return r.Platform
}

// MessagesRequestAppWebView invokes method messages.requestAppWebView#8c5a3b3c returning error if any.
// Open a bot mini app¹ from a named Mini App deep link², sending over user information
// after user confirmation.
// After calling this method, until the user closes the webview, messages
// prolongWebView¹ must be called every 60 seconds.
//
// Links:
//  1. https://core.telegram.org/bots/webapps
//  2. https://core.telegram.org/api/links#named-mini-app-links
//  3. https://core.telegram.org/method/messages.prolongWebView
//
// See https://core.telegram.org/method/messages.requestAppWebView for reference.
func (c *Client) MessagesRequestAppWebView(ctx context.Context, request *MessagesRequestAppWebViewRequest) (*AppWebViewResultURL, error) {
	var result AppWebViewResultURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
