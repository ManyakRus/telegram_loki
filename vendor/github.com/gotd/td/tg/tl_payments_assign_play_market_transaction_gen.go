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

// PaymentsAssignPlayMarketTransactionRequest represents TL type `payments.assignPlayMarketTransaction#dffd50d3`.
// Informs server about a purchase made through the Play Store: for official applications
// only.
//
// See https://core.telegram.org/method/payments.assignPlayMarketTransaction for reference.
type PaymentsAssignPlayMarketTransactionRequest struct {
	// Receipt
	Receipt DataJSON
	// Payment purpose
	Purpose InputStorePaymentPurposeClass
}

// PaymentsAssignPlayMarketTransactionRequestTypeID is TL type id of PaymentsAssignPlayMarketTransactionRequest.
const PaymentsAssignPlayMarketTransactionRequestTypeID = 0xdffd50d3

// Ensuring interfaces in compile-time for PaymentsAssignPlayMarketTransactionRequest.
var (
	_ bin.Encoder     = &PaymentsAssignPlayMarketTransactionRequest{}
	_ bin.Decoder     = &PaymentsAssignPlayMarketTransactionRequest{}
	_ bin.BareEncoder = &PaymentsAssignPlayMarketTransactionRequest{}
	_ bin.BareDecoder = &PaymentsAssignPlayMarketTransactionRequest{}
)

func (a *PaymentsAssignPlayMarketTransactionRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Receipt.Zero()) {
		return false
	}
	if !(a.Purpose == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *PaymentsAssignPlayMarketTransactionRequest) String() string {
	if a == nil {
		return "PaymentsAssignPlayMarketTransactionRequest(nil)"
	}
	type Alias PaymentsAssignPlayMarketTransactionRequest
	return fmt.Sprintf("PaymentsAssignPlayMarketTransactionRequest%+v", Alias(*a))
}

// FillFrom fills PaymentsAssignPlayMarketTransactionRequest from given interface.
func (a *PaymentsAssignPlayMarketTransactionRequest) FillFrom(from interface {
	GetReceipt() (value DataJSON)
	GetPurpose() (value InputStorePaymentPurposeClass)
}) {
	a.Receipt = from.GetReceipt()
	a.Purpose = from.GetPurpose()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsAssignPlayMarketTransactionRequest) TypeID() uint32 {
	return PaymentsAssignPlayMarketTransactionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsAssignPlayMarketTransactionRequest) TypeName() string {
	return "payments.assignPlayMarketTransaction"
}

// TypeInfo returns info about TL type.
func (a *PaymentsAssignPlayMarketTransactionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.assignPlayMarketTransaction",
		ID:   PaymentsAssignPlayMarketTransactionRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Receipt",
			SchemaName: "receipt",
		},
		{
			Name:       "Purpose",
			SchemaName: "purpose",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *PaymentsAssignPlayMarketTransactionRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode payments.assignPlayMarketTransaction#dffd50d3 as nil")
	}
	b.PutID(PaymentsAssignPlayMarketTransactionRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *PaymentsAssignPlayMarketTransactionRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode payments.assignPlayMarketTransaction#dffd50d3 as nil")
	}
	if err := a.Receipt.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.assignPlayMarketTransaction#dffd50d3: field receipt: %w", err)
	}
	if a.Purpose == nil {
		return fmt.Errorf("unable to encode payments.assignPlayMarketTransaction#dffd50d3: field purpose is nil")
	}
	if err := a.Purpose.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.assignPlayMarketTransaction#dffd50d3: field purpose: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *PaymentsAssignPlayMarketTransactionRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode payments.assignPlayMarketTransaction#dffd50d3 to nil")
	}
	if err := b.ConsumeID(PaymentsAssignPlayMarketTransactionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.assignPlayMarketTransaction#dffd50d3: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *PaymentsAssignPlayMarketTransactionRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode payments.assignPlayMarketTransaction#dffd50d3 to nil")
	}
	{
		if err := a.Receipt.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.assignPlayMarketTransaction#dffd50d3: field receipt: %w", err)
		}
	}
	{
		value, err := DecodeInputStorePaymentPurpose(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.assignPlayMarketTransaction#dffd50d3: field purpose: %w", err)
		}
		a.Purpose = value
	}
	return nil
}

// GetReceipt returns value of Receipt field.
func (a *PaymentsAssignPlayMarketTransactionRequest) GetReceipt() (value DataJSON) {
	if a == nil {
		return
	}
	return a.Receipt
}

// GetPurpose returns value of Purpose field.
func (a *PaymentsAssignPlayMarketTransactionRequest) GetPurpose() (value InputStorePaymentPurposeClass) {
	if a == nil {
		return
	}
	return a.Purpose
}

// PaymentsAssignPlayMarketTransaction invokes method payments.assignPlayMarketTransaction#dffd50d3 returning error if any.
// Informs server about a purchase made through the Play Store: for official applications
// only.
//
// See https://core.telegram.org/method/payments.assignPlayMarketTransaction for reference.
func (c *Client) PaymentsAssignPlayMarketTransaction(ctx context.Context, request *PaymentsAssignPlayMarketTransactionRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
