// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// PaymentResult represents TL type `paymentResult#d00fe85d`.
type PaymentResult struct {
	// True, if the payment request was successful; otherwise the verification_url will be
	// non-empty
	Success bool
	// URL for additional payment credentials verification
	VerificationURL string
}

// PaymentResultTypeID is TL type id of PaymentResult.
const PaymentResultTypeID = 0xd00fe85d

// Ensuring interfaces in compile-time for PaymentResult.
var (
	_ bin.Encoder     = &PaymentResult{}
	_ bin.Decoder     = &PaymentResult{}
	_ bin.BareEncoder = &PaymentResult{}
	_ bin.BareDecoder = &PaymentResult{}
)

func (p *PaymentResult) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Success == false) {
		return false
	}
	if !(p.VerificationURL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentResult) String() string {
	if p == nil {
		return "PaymentResult(nil)"
	}
	type Alias PaymentResult
	return fmt.Sprintf("PaymentResult%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentResult) TypeID() uint32 {
	return PaymentResultTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentResult) TypeName() string {
	return "paymentResult"
}

// TypeInfo returns info about TL type.
func (p *PaymentResult) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentResult",
		ID:   PaymentResultTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Success",
			SchemaName: "success",
		},
		{
			Name:       "VerificationURL",
			SchemaName: "verification_url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentResult) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentResult#d00fe85d as nil")
	}
	b.PutID(PaymentResultTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentResult) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentResult#d00fe85d as nil")
	}
	b.PutBool(p.Success)
	b.PutString(p.VerificationURL)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentResult) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentResult#d00fe85d to nil")
	}
	if err := b.ConsumeID(PaymentResultTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentResult#d00fe85d: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentResult) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentResult#d00fe85d to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode paymentResult#d00fe85d: field success: %w", err)
		}
		p.Success = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentResult#d00fe85d: field verification_url: %w", err)
		}
		p.VerificationURL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PaymentResult) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentResult#d00fe85d as nil")
	}
	b.ObjStart()
	b.PutID("paymentResult")
	b.FieldStart("success")
	b.PutBool(p.Success)
	b.FieldStart("verification_url")
	b.PutString(p.VerificationURL)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PaymentResult) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentResult#d00fe85d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("paymentResult"); err != nil {
				return fmt.Errorf("unable to decode paymentResult#d00fe85d: %w", err)
			}
		case "success":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode paymentResult#d00fe85d: field success: %w", err)
			}
			p.Success = value
		case "verification_url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode paymentResult#d00fe85d: field verification_url: %w", err)
			}
			p.VerificationURL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSuccess returns value of Success field.
func (p *PaymentResult) GetSuccess() (value bool) {
	return p.Success
}

// GetVerificationURL returns value of VerificationURL field.
func (p *PaymentResult) GetVerificationURL() (value string) {
	return p.VerificationURL
}
