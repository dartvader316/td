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

// SetBackgroundRequest represents TL type `setBackground#c2487387`.
type SetBackgroundRequest struct {
	// The input background to use; pass null to create a new filled backgrounds or to remove
	// the current background
	Background InputBackgroundClass
	// Background type; pass null to use the default type of the remote background or to
	// remove the current background
	Type BackgroundTypeClass
	// Pass true if the background is changed for a dark theme
	ForDarkTheme bool
}

// SetBackgroundRequestTypeID is TL type id of SetBackgroundRequest.
const SetBackgroundRequestTypeID = 0xc2487387

// Ensuring interfaces in compile-time for SetBackgroundRequest.
var (
	_ bin.Encoder     = &SetBackgroundRequest{}
	_ bin.Decoder     = &SetBackgroundRequest{}
	_ bin.BareEncoder = &SetBackgroundRequest{}
	_ bin.BareDecoder = &SetBackgroundRequest{}
)

func (s *SetBackgroundRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Background == nil) {
		return false
	}
	if !(s.Type == nil) {
		return false
	}
	if !(s.ForDarkTheme == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetBackgroundRequest) String() string {
	if s == nil {
		return "SetBackgroundRequest(nil)"
	}
	type Alias SetBackgroundRequest
	return fmt.Sprintf("SetBackgroundRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetBackgroundRequest) TypeID() uint32 {
	return SetBackgroundRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetBackgroundRequest) TypeName() string {
	return "setBackground"
}

// TypeInfo returns info about TL type.
func (s *SetBackgroundRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setBackground",
		ID:   SetBackgroundRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Background",
			SchemaName: "background",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "ForDarkTheme",
			SchemaName: "for_dark_theme",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetBackgroundRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBackground#c2487387 as nil")
	}
	b.PutID(SetBackgroundRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetBackgroundRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBackground#c2487387 as nil")
	}
	if s.Background == nil {
		return fmt.Errorf("unable to encode setBackground#c2487387: field background is nil")
	}
	if err := s.Background.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setBackground#c2487387: field background: %w", err)
	}
	if s.Type == nil {
		return fmt.Errorf("unable to encode setBackground#c2487387: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setBackground#c2487387: field type: %w", err)
	}
	b.PutBool(s.ForDarkTheme)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetBackgroundRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBackground#c2487387 to nil")
	}
	if err := b.ConsumeID(SetBackgroundRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setBackground#c2487387: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetBackgroundRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBackground#c2487387 to nil")
	}
	{
		value, err := DecodeInputBackground(b)
		if err != nil {
			return fmt.Errorf("unable to decode setBackground#c2487387: field background: %w", err)
		}
		s.Background = value
	}
	{
		value, err := DecodeBackgroundType(b)
		if err != nil {
			return fmt.Errorf("unable to decode setBackground#c2487387: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode setBackground#c2487387: field for_dark_theme: %w", err)
		}
		s.ForDarkTheme = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetBackgroundRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setBackground#c2487387 as nil")
	}
	b.ObjStart()
	b.PutID("setBackground")
	b.Comma()
	b.FieldStart("background")
	if s.Background == nil {
		return fmt.Errorf("unable to encode setBackground#c2487387: field background is nil")
	}
	if err := s.Background.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setBackground#c2487387: field background: %w", err)
	}
	b.Comma()
	b.FieldStart("type")
	if s.Type == nil {
		return fmt.Errorf("unable to encode setBackground#c2487387: field type is nil")
	}
	if err := s.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setBackground#c2487387: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("for_dark_theme")
	b.PutBool(s.ForDarkTheme)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetBackgroundRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setBackground#c2487387 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setBackground"); err != nil {
				return fmt.Errorf("unable to decode setBackground#c2487387: %w", err)
			}
		case "background":
			value, err := DecodeTDLibJSONInputBackground(b)
			if err != nil {
				return fmt.Errorf("unable to decode setBackground#c2487387: field background: %w", err)
			}
			s.Background = value
		case "type":
			value, err := DecodeTDLibJSONBackgroundType(b)
			if err != nil {
				return fmt.Errorf("unable to decode setBackground#c2487387: field type: %w", err)
			}
			s.Type = value
		case "for_dark_theme":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode setBackground#c2487387: field for_dark_theme: %w", err)
			}
			s.ForDarkTheme = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBackground returns value of Background field.
func (s *SetBackgroundRequest) GetBackground() (value InputBackgroundClass) {
	if s == nil {
		return
	}
	return s.Background
}

// GetType returns value of Type field.
func (s *SetBackgroundRequest) GetType() (value BackgroundTypeClass) {
	if s == nil {
		return
	}
	return s.Type
}

// GetForDarkTheme returns value of ForDarkTheme field.
func (s *SetBackgroundRequest) GetForDarkTheme() (value bool) {
	if s == nil {
		return
	}
	return s.ForDarkTheme
}

// SetBackground invokes method setBackground#c2487387 returning error if any.
func (c *Client) SetBackground(ctx context.Context, request *SetBackgroundRequest) (*Background, error) {
	var result Background

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
