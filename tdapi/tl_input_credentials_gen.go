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

// InputCredentialsSaved represents TL type `inputCredentialsSaved#86bdbe2c`.
type InputCredentialsSaved struct {
	// Identifier of the saved credentials
	SavedCredentialsID string
}

// InputCredentialsSavedTypeID is TL type id of InputCredentialsSaved.
const InputCredentialsSavedTypeID = 0x86bdbe2c

// construct implements constructor of InputCredentialsClass.
func (i InputCredentialsSaved) construct() InputCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputCredentialsSaved.
var (
	_ bin.Encoder     = &InputCredentialsSaved{}
	_ bin.Decoder     = &InputCredentialsSaved{}
	_ bin.BareEncoder = &InputCredentialsSaved{}
	_ bin.BareDecoder = &InputCredentialsSaved{}

	_ InputCredentialsClass = &InputCredentialsSaved{}
)

func (i *InputCredentialsSaved) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.SavedCredentialsID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputCredentialsSaved) String() string {
	if i == nil {
		return "InputCredentialsSaved(nil)"
	}
	type Alias InputCredentialsSaved
	return fmt.Sprintf("InputCredentialsSaved%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputCredentialsSaved) TypeID() uint32 {
	return InputCredentialsSavedTypeID
}

// TypeName returns name of type in TL schema.
func (*InputCredentialsSaved) TypeName() string {
	return "inputCredentialsSaved"
}

// TypeInfo returns info about TL type.
func (i *InputCredentialsSaved) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputCredentialsSaved",
		ID:   InputCredentialsSavedTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SavedCredentialsID",
			SchemaName: "saved_credentials_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputCredentialsSaved) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsSaved#86bdbe2c as nil")
	}
	b.PutID(InputCredentialsSavedTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputCredentialsSaved) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsSaved#86bdbe2c as nil")
	}
	b.PutString(i.SavedCredentialsID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputCredentialsSaved) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsSaved#86bdbe2c to nil")
	}
	if err := b.ConsumeID(InputCredentialsSavedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputCredentialsSaved#86bdbe2c: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputCredentialsSaved) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsSaved#86bdbe2c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputCredentialsSaved#86bdbe2c: field saved_credentials_id: %w", err)
		}
		i.SavedCredentialsID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputCredentialsSaved) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsSaved#86bdbe2c as nil")
	}
	b.ObjStart()
	b.PutID("inputCredentialsSaved")
	b.FieldStart("saved_credentials_id")
	b.PutString(i.SavedCredentialsID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputCredentialsSaved) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsSaved#86bdbe2c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputCredentialsSaved"); err != nil {
				return fmt.Errorf("unable to decode inputCredentialsSaved#86bdbe2c: %w", err)
			}
		case "saved_credentials_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputCredentialsSaved#86bdbe2c: field saved_credentials_id: %w", err)
			}
			i.SavedCredentialsID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSavedCredentialsID returns value of SavedCredentialsID field.
func (i *InputCredentialsSaved) GetSavedCredentialsID() (value string) {
	return i.SavedCredentialsID
}

// InputCredentialsNew represents TL type `inputCredentialsNew#ce8bf12a`.
type InputCredentialsNew struct {
	// JSON-encoded data with the credential identifier from the payment provider
	Data string
	// True, if the credential identifier can be saved on the server side
	AllowSave bool
}

// InputCredentialsNewTypeID is TL type id of InputCredentialsNew.
const InputCredentialsNewTypeID = 0xce8bf12a

// construct implements constructor of InputCredentialsClass.
func (i InputCredentialsNew) construct() InputCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputCredentialsNew.
var (
	_ bin.Encoder     = &InputCredentialsNew{}
	_ bin.Decoder     = &InputCredentialsNew{}
	_ bin.BareEncoder = &InputCredentialsNew{}
	_ bin.BareDecoder = &InputCredentialsNew{}

	_ InputCredentialsClass = &InputCredentialsNew{}
)

func (i *InputCredentialsNew) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Data == "") {
		return false
	}
	if !(i.AllowSave == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputCredentialsNew) String() string {
	if i == nil {
		return "InputCredentialsNew(nil)"
	}
	type Alias InputCredentialsNew
	return fmt.Sprintf("InputCredentialsNew%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputCredentialsNew) TypeID() uint32 {
	return InputCredentialsNewTypeID
}

// TypeName returns name of type in TL schema.
func (*InputCredentialsNew) TypeName() string {
	return "inputCredentialsNew"
}

// TypeInfo returns info about TL type.
func (i *InputCredentialsNew) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputCredentialsNew",
		ID:   InputCredentialsNewTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Data",
			SchemaName: "data",
		},
		{
			Name:       "AllowSave",
			SchemaName: "allow_save",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputCredentialsNew) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsNew#ce8bf12a as nil")
	}
	b.PutID(InputCredentialsNewTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputCredentialsNew) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsNew#ce8bf12a as nil")
	}
	b.PutString(i.Data)
	b.PutBool(i.AllowSave)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputCredentialsNew) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsNew#ce8bf12a to nil")
	}
	if err := b.ConsumeID(InputCredentialsNewTypeID); err != nil {
		return fmt.Errorf("unable to decode inputCredentialsNew#ce8bf12a: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputCredentialsNew) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsNew#ce8bf12a to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputCredentialsNew#ce8bf12a: field data: %w", err)
		}
		i.Data = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode inputCredentialsNew#ce8bf12a: field allow_save: %w", err)
		}
		i.AllowSave = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputCredentialsNew) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsNew#ce8bf12a as nil")
	}
	b.ObjStart()
	b.PutID("inputCredentialsNew")
	b.FieldStart("data")
	b.PutString(i.Data)
	b.FieldStart("allow_save")
	b.PutBool(i.AllowSave)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputCredentialsNew) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsNew#ce8bf12a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputCredentialsNew"); err != nil {
				return fmt.Errorf("unable to decode inputCredentialsNew#ce8bf12a: %w", err)
			}
		case "data":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputCredentialsNew#ce8bf12a: field data: %w", err)
			}
			i.Data = value
		case "allow_save":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode inputCredentialsNew#ce8bf12a: field allow_save: %w", err)
			}
			i.AllowSave = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetData returns value of Data field.
func (i *InputCredentialsNew) GetData() (value string) {
	return i.Data
}

// GetAllowSave returns value of AllowSave field.
func (i *InputCredentialsNew) GetAllowSave() (value bool) {
	return i.AllowSave
}

// InputCredentialsApplePay represents TL type `inputCredentialsApplePay#b5b2d6d1`.
type InputCredentialsApplePay struct {
	// JSON-encoded data with the credential identifier
	Data string
}

// InputCredentialsApplePayTypeID is TL type id of InputCredentialsApplePay.
const InputCredentialsApplePayTypeID = 0xb5b2d6d1

// construct implements constructor of InputCredentialsClass.
func (i InputCredentialsApplePay) construct() InputCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputCredentialsApplePay.
var (
	_ bin.Encoder     = &InputCredentialsApplePay{}
	_ bin.Decoder     = &InputCredentialsApplePay{}
	_ bin.BareEncoder = &InputCredentialsApplePay{}
	_ bin.BareDecoder = &InputCredentialsApplePay{}

	_ InputCredentialsClass = &InputCredentialsApplePay{}
)

func (i *InputCredentialsApplePay) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Data == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputCredentialsApplePay) String() string {
	if i == nil {
		return "InputCredentialsApplePay(nil)"
	}
	type Alias InputCredentialsApplePay
	return fmt.Sprintf("InputCredentialsApplePay%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputCredentialsApplePay) TypeID() uint32 {
	return InputCredentialsApplePayTypeID
}

// TypeName returns name of type in TL schema.
func (*InputCredentialsApplePay) TypeName() string {
	return "inputCredentialsApplePay"
}

// TypeInfo returns info about TL type.
func (i *InputCredentialsApplePay) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputCredentialsApplePay",
		ID:   InputCredentialsApplePayTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputCredentialsApplePay) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsApplePay#b5b2d6d1 as nil")
	}
	b.PutID(InputCredentialsApplePayTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputCredentialsApplePay) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsApplePay#b5b2d6d1 as nil")
	}
	b.PutString(i.Data)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputCredentialsApplePay) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsApplePay#b5b2d6d1 to nil")
	}
	if err := b.ConsumeID(InputCredentialsApplePayTypeID); err != nil {
		return fmt.Errorf("unable to decode inputCredentialsApplePay#b5b2d6d1: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputCredentialsApplePay) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsApplePay#b5b2d6d1 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputCredentialsApplePay#b5b2d6d1: field data: %w", err)
		}
		i.Data = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputCredentialsApplePay) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsApplePay#b5b2d6d1 as nil")
	}
	b.ObjStart()
	b.PutID("inputCredentialsApplePay")
	b.FieldStart("data")
	b.PutString(i.Data)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputCredentialsApplePay) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsApplePay#b5b2d6d1 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputCredentialsApplePay"); err != nil {
				return fmt.Errorf("unable to decode inputCredentialsApplePay#b5b2d6d1: %w", err)
			}
		case "data":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputCredentialsApplePay#b5b2d6d1: field data: %w", err)
			}
			i.Data = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetData returns value of Data field.
func (i *InputCredentialsApplePay) GetData() (value string) {
	return i.Data
}

// InputCredentialsGooglePay represents TL type `inputCredentialsGooglePay#32544764`.
type InputCredentialsGooglePay struct {
	// JSON-encoded data with the credential identifier
	Data string
}

// InputCredentialsGooglePayTypeID is TL type id of InputCredentialsGooglePay.
const InputCredentialsGooglePayTypeID = 0x32544764

// construct implements constructor of InputCredentialsClass.
func (i InputCredentialsGooglePay) construct() InputCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputCredentialsGooglePay.
var (
	_ bin.Encoder     = &InputCredentialsGooglePay{}
	_ bin.Decoder     = &InputCredentialsGooglePay{}
	_ bin.BareEncoder = &InputCredentialsGooglePay{}
	_ bin.BareDecoder = &InputCredentialsGooglePay{}

	_ InputCredentialsClass = &InputCredentialsGooglePay{}
)

func (i *InputCredentialsGooglePay) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Data == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputCredentialsGooglePay) String() string {
	if i == nil {
		return "InputCredentialsGooglePay(nil)"
	}
	type Alias InputCredentialsGooglePay
	return fmt.Sprintf("InputCredentialsGooglePay%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputCredentialsGooglePay) TypeID() uint32 {
	return InputCredentialsGooglePayTypeID
}

// TypeName returns name of type in TL schema.
func (*InputCredentialsGooglePay) TypeName() string {
	return "inputCredentialsGooglePay"
}

// TypeInfo returns info about TL type.
func (i *InputCredentialsGooglePay) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputCredentialsGooglePay",
		ID:   InputCredentialsGooglePayTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputCredentialsGooglePay) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsGooglePay#32544764 as nil")
	}
	b.PutID(InputCredentialsGooglePayTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputCredentialsGooglePay) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsGooglePay#32544764 as nil")
	}
	b.PutString(i.Data)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputCredentialsGooglePay) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsGooglePay#32544764 to nil")
	}
	if err := b.ConsumeID(InputCredentialsGooglePayTypeID); err != nil {
		return fmt.Errorf("unable to decode inputCredentialsGooglePay#32544764: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputCredentialsGooglePay) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsGooglePay#32544764 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputCredentialsGooglePay#32544764: field data: %w", err)
		}
		i.Data = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputCredentialsGooglePay) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCredentialsGooglePay#32544764 as nil")
	}
	b.ObjStart()
	b.PutID("inputCredentialsGooglePay")
	b.FieldStart("data")
	b.PutString(i.Data)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputCredentialsGooglePay) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCredentialsGooglePay#32544764 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputCredentialsGooglePay"); err != nil {
				return fmt.Errorf("unable to decode inputCredentialsGooglePay#32544764: %w", err)
			}
		case "data":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputCredentialsGooglePay#32544764: field data: %w", err)
			}
			i.Data = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetData returns value of Data field.
func (i *InputCredentialsGooglePay) GetData() (value string) {
	return i.Data
}

// InputCredentialsClass represents InputCredentials generic type.
//
// Example:
//  g, err := tdapi.DecodeInputCredentials(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.InputCredentialsSaved: // inputCredentialsSaved#86bdbe2c
//  case *tdapi.InputCredentialsNew: // inputCredentialsNew#ce8bf12a
//  case *tdapi.InputCredentialsApplePay: // inputCredentialsApplePay#b5b2d6d1
//  case *tdapi.InputCredentialsGooglePay: // inputCredentialsGooglePay#32544764
//  default: panic(v)
//  }
type InputCredentialsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputCredentialsClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeInputCredentials implements binary de-serialization for InputCredentialsClass.
func DecodeInputCredentials(buf *bin.Buffer) (InputCredentialsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputCredentialsSavedTypeID:
		// Decoding inputCredentialsSaved#86bdbe2c.
		v := InputCredentialsSaved{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCredentialsClass: %w", err)
		}
		return &v, nil
	case InputCredentialsNewTypeID:
		// Decoding inputCredentialsNew#ce8bf12a.
		v := InputCredentialsNew{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCredentialsClass: %w", err)
		}
		return &v, nil
	case InputCredentialsApplePayTypeID:
		// Decoding inputCredentialsApplePay#b5b2d6d1.
		v := InputCredentialsApplePay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCredentialsClass: %w", err)
		}
		return &v, nil
	case InputCredentialsGooglePayTypeID:
		// Decoding inputCredentialsGooglePay#32544764.
		v := InputCredentialsGooglePay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCredentialsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputCredentialsClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONInputCredentials implements binary de-serialization for InputCredentialsClass.
func DecodeTDLibJSONInputCredentials(buf tdjson.Decoder) (InputCredentialsClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "inputCredentialsSaved":
		// Decoding inputCredentialsSaved#86bdbe2c.
		v := InputCredentialsSaved{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCredentialsClass: %w", err)
		}
		return &v, nil
	case "inputCredentialsNew":
		// Decoding inputCredentialsNew#ce8bf12a.
		v := InputCredentialsNew{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCredentialsClass: %w", err)
		}
		return &v, nil
	case "inputCredentialsApplePay":
		// Decoding inputCredentialsApplePay#b5b2d6d1.
		v := InputCredentialsApplePay{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCredentialsClass: %w", err)
		}
		return &v, nil
	case "inputCredentialsGooglePay":
		// Decoding inputCredentialsGooglePay#32544764.
		v := InputCredentialsGooglePay{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCredentialsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputCredentialsClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// InputCredentials boxes the InputCredentialsClass providing a helper.
type InputCredentialsBox struct {
	InputCredentials InputCredentialsClass
}

// Decode implements bin.Decoder for InputCredentialsBox.
func (b *InputCredentialsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputCredentialsBox to nil")
	}
	v, err := DecodeInputCredentials(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputCredentials = v
	return nil
}

// Encode implements bin.Encode for InputCredentialsBox.
func (b *InputCredentialsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputCredentials == nil {
		return fmt.Errorf("unable to encode InputCredentialsClass as nil")
	}
	return b.InputCredentials.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for InputCredentialsBox.
func (b *InputCredentialsBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputCredentialsBox to nil")
	}
	v, err := DecodeTDLibJSONInputCredentials(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputCredentials = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for InputCredentialsBox.
func (b *InputCredentialsBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.InputCredentials == nil {
		return fmt.Errorf("unable to encode InputCredentialsClass as nil")
	}
	return b.InputCredentials.EncodeTDLibJSON(buf)
}
