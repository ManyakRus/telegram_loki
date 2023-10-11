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

// BotsAnswerWebhookJSONQueryRequest represents TL type `bots.answerWebhookJSONQuery#e6213f4d`.
// Answers a custom query; for bots only
//
// See https://core.telegram.org/method/bots.answerWebhookJSONQuery for reference.
type BotsAnswerWebhookJSONQueryRequest struct {
	// Identifier of a custom query
	QueryID int64
	// JSON-serialized answer to the query
	Data DataJSON
}

// BotsAnswerWebhookJSONQueryRequestTypeID is TL type id of BotsAnswerWebhookJSONQueryRequest.
const BotsAnswerWebhookJSONQueryRequestTypeID = 0xe6213f4d

// Ensuring interfaces in compile-time for BotsAnswerWebhookJSONQueryRequest.
var (
	_ bin.Encoder     = &BotsAnswerWebhookJSONQueryRequest{}
	_ bin.Decoder     = &BotsAnswerWebhookJSONQueryRequest{}
	_ bin.BareEncoder = &BotsAnswerWebhookJSONQueryRequest{}
	_ bin.BareDecoder = &BotsAnswerWebhookJSONQueryRequest{}
)

func (a *BotsAnswerWebhookJSONQueryRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.QueryID == 0) {
		return false
	}
	if !(a.Data.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *BotsAnswerWebhookJSONQueryRequest) String() string {
	if a == nil {
		return "BotsAnswerWebhookJSONQueryRequest(nil)"
	}
	type Alias BotsAnswerWebhookJSONQueryRequest
	return fmt.Sprintf("BotsAnswerWebhookJSONQueryRequest%+v", Alias(*a))
}

// FillFrom fills BotsAnswerWebhookJSONQueryRequest from given interface.
func (a *BotsAnswerWebhookJSONQueryRequest) FillFrom(from interface {
	GetQueryID() (value int64)
	GetData() (value DataJSON)
}) {
	a.QueryID = from.GetQueryID()
	a.Data = from.GetData()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsAnswerWebhookJSONQueryRequest) TypeID() uint32 {
	return BotsAnswerWebhookJSONQueryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsAnswerWebhookJSONQueryRequest) TypeName() string {
	return "bots.answerWebhookJSONQuery"
}

// TypeInfo returns info about TL type.
func (a *BotsAnswerWebhookJSONQueryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.answerWebhookJSONQuery",
		ID:   BotsAnswerWebhookJSONQueryRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *BotsAnswerWebhookJSONQueryRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode bots.answerWebhookJSONQuery#e6213f4d as nil")
	}
	b.PutID(BotsAnswerWebhookJSONQueryRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *BotsAnswerWebhookJSONQueryRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode bots.answerWebhookJSONQuery#e6213f4d as nil")
	}
	b.PutLong(a.QueryID)
	if err := a.Data.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.answerWebhookJSONQuery#e6213f4d: field data: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *BotsAnswerWebhookJSONQueryRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode bots.answerWebhookJSONQuery#e6213f4d to nil")
	}
	if err := b.ConsumeID(BotsAnswerWebhookJSONQueryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.answerWebhookJSONQuery#e6213f4d: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *BotsAnswerWebhookJSONQueryRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode bots.answerWebhookJSONQuery#e6213f4d to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode bots.answerWebhookJSONQuery#e6213f4d: field query_id: %w", err)
		}
		a.QueryID = value
	}
	{
		if err := a.Data.Decode(b); err != nil {
			return fmt.Errorf("unable to decode bots.answerWebhookJSONQuery#e6213f4d: field data: %w", err)
		}
	}
	return nil
}

// GetQueryID returns value of QueryID field.
func (a *BotsAnswerWebhookJSONQueryRequest) GetQueryID() (value int64) {
	if a == nil {
		return
	}
	return a.QueryID
}

// GetData returns value of Data field.
func (a *BotsAnswerWebhookJSONQueryRequest) GetData() (value DataJSON) {
	if a == nil {
		return
	}
	return a.Data
}

// BotsAnswerWebhookJSONQuery invokes method bots.answerWebhookJSONQuery#e6213f4d returning error if any.
// Answers a custom query; for bots only
//
// Possible errors:
//
//	400 DATA_JSON_INVALID: The provided JSON data is invalid.
//	400 QUERY_ID_INVALID: The query ID is invalid.
//	403 USER_BOT_INVALID: User accounts must provide the bot method parameter when calling this method. If there is no such method parameter, this method can only be invoked by bot accounts.
//
// See https://core.telegram.org/method/bots.answerWebhookJSONQuery for reference.
// Can be used by bots.
func (c *Client) BotsAnswerWebhookJSONQuery(ctx context.Context, request *BotsAnswerWebhookJSONQueryRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
