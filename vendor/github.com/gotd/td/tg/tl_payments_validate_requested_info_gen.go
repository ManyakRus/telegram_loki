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

// PaymentsValidateRequestedInfoRequest represents TL type `payments.validateRequestedInfo#b6c8f12b`.
// Submit requested order information for validation
//
// See https://core.telegram.org/method/payments.validateRequestedInfo for reference.
type PaymentsValidateRequestedInfoRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Save order information to re-use it for future orders
	Save bool
	// Invoice
	Invoice InputInvoiceClass
	// Requested order information
	Info PaymentRequestedInfo
}

// PaymentsValidateRequestedInfoRequestTypeID is TL type id of PaymentsValidateRequestedInfoRequest.
const PaymentsValidateRequestedInfoRequestTypeID = 0xb6c8f12b

// Ensuring interfaces in compile-time for PaymentsValidateRequestedInfoRequest.
var (
	_ bin.Encoder     = &PaymentsValidateRequestedInfoRequest{}
	_ bin.Decoder     = &PaymentsValidateRequestedInfoRequest{}
	_ bin.BareEncoder = &PaymentsValidateRequestedInfoRequest{}
	_ bin.BareDecoder = &PaymentsValidateRequestedInfoRequest{}
)

func (v *PaymentsValidateRequestedInfoRequest) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Flags.Zero()) {
		return false
	}
	if !(v.Save == false) {
		return false
	}
	if !(v.Invoice == nil) {
		return false
	}
	if !(v.Info.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *PaymentsValidateRequestedInfoRequest) String() string {
	if v == nil {
		return "PaymentsValidateRequestedInfoRequest(nil)"
	}
	type Alias PaymentsValidateRequestedInfoRequest
	return fmt.Sprintf("PaymentsValidateRequestedInfoRequest%+v", Alias(*v))
}

// FillFrom fills PaymentsValidateRequestedInfoRequest from given interface.
func (v *PaymentsValidateRequestedInfoRequest) FillFrom(from interface {
	GetSave() (value bool)
	GetInvoice() (value InputInvoiceClass)
	GetInfo() (value PaymentRequestedInfo)
}) {
	v.Save = from.GetSave()
	v.Invoice = from.GetInvoice()
	v.Info = from.GetInfo()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsValidateRequestedInfoRequest) TypeID() uint32 {
	return PaymentsValidateRequestedInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsValidateRequestedInfoRequest) TypeName() string {
	return "payments.validateRequestedInfo"
}

// TypeInfo returns info about TL type.
func (v *PaymentsValidateRequestedInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.validateRequestedInfo",
		ID:   PaymentsValidateRequestedInfoRequestTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Save",
			SchemaName: "save",
			Null:       !v.Flags.Has(0),
		},
		{
			Name:       "Invoice",
			SchemaName: "invoice",
		},
		{
			Name:       "Info",
			SchemaName: "info",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (v *PaymentsValidateRequestedInfoRequest) SetFlags() {
	if !(v.Save == false) {
		v.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (v *PaymentsValidateRequestedInfoRequest) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode payments.validateRequestedInfo#b6c8f12b as nil")
	}
	b.PutID(PaymentsValidateRequestedInfoRequestTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *PaymentsValidateRequestedInfoRequest) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode payments.validateRequestedInfo#b6c8f12b as nil")
	}
	v.SetFlags()
	if err := v.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.validateRequestedInfo#b6c8f12b: field flags: %w", err)
	}
	if v.Invoice == nil {
		return fmt.Errorf("unable to encode payments.validateRequestedInfo#b6c8f12b: field invoice is nil")
	}
	if err := v.Invoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.validateRequestedInfo#b6c8f12b: field invoice: %w", err)
	}
	if err := v.Info.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.validateRequestedInfo#b6c8f12b: field info: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *PaymentsValidateRequestedInfoRequest) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode payments.validateRequestedInfo#b6c8f12b to nil")
	}
	if err := b.ConsumeID(PaymentsValidateRequestedInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.validateRequestedInfo#b6c8f12b: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *PaymentsValidateRequestedInfoRequest) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode payments.validateRequestedInfo#b6c8f12b to nil")
	}
	{
		if err := v.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.validateRequestedInfo#b6c8f12b: field flags: %w", err)
		}
	}
	v.Save = v.Flags.Has(0)
	{
		value, err := DecodeInputInvoice(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.validateRequestedInfo#b6c8f12b: field invoice: %w", err)
		}
		v.Invoice = value
	}
	{
		if err := v.Info.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.validateRequestedInfo#b6c8f12b: field info: %w", err)
		}
	}
	return nil
}

// SetSave sets value of Save conditional field.
func (v *PaymentsValidateRequestedInfoRequest) SetSave(value bool) {
	if value {
		v.Flags.Set(0)
		v.Save = true
	} else {
		v.Flags.Unset(0)
		v.Save = false
	}
}

// GetSave returns value of Save conditional field.
func (v *PaymentsValidateRequestedInfoRequest) GetSave() (value bool) {
	if v == nil {
		return
	}
	return v.Flags.Has(0)
}

// GetInvoice returns value of Invoice field.
func (v *PaymentsValidateRequestedInfoRequest) GetInvoice() (value InputInvoiceClass) {
	if v == nil {
		return
	}
	return v.Invoice
}

// GetInfo returns value of Info field.
func (v *PaymentsValidateRequestedInfoRequest) GetInfo() (value PaymentRequestedInfo) {
	if v == nil {
		return
	}
	return v.Info
}

// PaymentsValidateRequestedInfo invokes method payments.validateRequestedInfo#b6c8f12b returning error if any.
// Submit requested order information for validation
//
// Possible errors:
//
//	400 MESSAGE_ID_INVALID: The provided message id is invalid.
//
// See https://core.telegram.org/method/payments.validateRequestedInfo for reference.
func (c *Client) PaymentsValidateRequestedInfo(ctx context.Context, request *PaymentsValidateRequestedInfoRequest) (*PaymentsValidatedRequestedInfo, error) {
	var result PaymentsValidatedRequestedInfo

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
