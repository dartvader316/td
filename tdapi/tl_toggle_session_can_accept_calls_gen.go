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

// ToggleSessionCanAcceptCallsRequest represents TL type `toggleSessionCanAcceptCalls#6c6c2708`.
type ToggleSessionCanAcceptCallsRequest struct {
	// Session identifier
	SessionID int64
	// Pass true to allow accepting incoming calls by the session; pass false otherwise
	CanAcceptCalls bool
}

// ToggleSessionCanAcceptCallsRequestTypeID is TL type id of ToggleSessionCanAcceptCallsRequest.
const ToggleSessionCanAcceptCallsRequestTypeID = 0x6c6c2708

// Ensuring interfaces in compile-time for ToggleSessionCanAcceptCallsRequest.
var (
	_ bin.Encoder     = &ToggleSessionCanAcceptCallsRequest{}
	_ bin.Decoder     = &ToggleSessionCanAcceptCallsRequest{}
	_ bin.BareEncoder = &ToggleSessionCanAcceptCallsRequest{}
	_ bin.BareDecoder = &ToggleSessionCanAcceptCallsRequest{}
)

func (t *ToggleSessionCanAcceptCallsRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.SessionID == 0) {
		return false
	}
	if !(t.CanAcceptCalls == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleSessionCanAcceptCallsRequest) String() string {
	if t == nil {
		return "ToggleSessionCanAcceptCallsRequest(nil)"
	}
	type Alias ToggleSessionCanAcceptCallsRequest
	return fmt.Sprintf("ToggleSessionCanAcceptCallsRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleSessionCanAcceptCallsRequest) TypeID() uint32 {
	return ToggleSessionCanAcceptCallsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleSessionCanAcceptCallsRequest) TypeName() string {
	return "toggleSessionCanAcceptCalls"
}

// TypeInfo returns info about TL type.
func (t *ToggleSessionCanAcceptCallsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleSessionCanAcceptCalls",
		ID:   ToggleSessionCanAcceptCallsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SessionID",
			SchemaName: "session_id",
		},
		{
			Name:       "CanAcceptCalls",
			SchemaName: "can_accept_calls",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleSessionCanAcceptCallsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSessionCanAcceptCalls#6c6c2708 as nil")
	}
	b.PutID(ToggleSessionCanAcceptCallsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleSessionCanAcceptCallsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSessionCanAcceptCalls#6c6c2708 as nil")
	}
	b.PutLong(t.SessionID)
	b.PutBool(t.CanAcceptCalls)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleSessionCanAcceptCallsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSessionCanAcceptCalls#6c6c2708 to nil")
	}
	if err := b.ConsumeID(ToggleSessionCanAcceptCallsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleSessionCanAcceptCalls#6c6c2708: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleSessionCanAcceptCallsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSessionCanAcceptCalls#6c6c2708 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSessionCanAcceptCalls#6c6c2708: field session_id: %w", err)
		}
		t.SessionID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSessionCanAcceptCalls#6c6c2708: field can_accept_calls: %w", err)
		}
		t.CanAcceptCalls = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleSessionCanAcceptCallsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSessionCanAcceptCalls#6c6c2708 as nil")
	}
	b.ObjStart()
	b.PutID("toggleSessionCanAcceptCalls")
	b.Comma()
	b.FieldStart("session_id")
	b.PutLong(t.SessionID)
	b.Comma()
	b.FieldStart("can_accept_calls")
	b.PutBool(t.CanAcceptCalls)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleSessionCanAcceptCallsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSessionCanAcceptCalls#6c6c2708 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleSessionCanAcceptCalls"); err != nil {
				return fmt.Errorf("unable to decode toggleSessionCanAcceptCalls#6c6c2708: %w", err)
			}
		case "session_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSessionCanAcceptCalls#6c6c2708: field session_id: %w", err)
			}
			t.SessionID = value
		case "can_accept_calls":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSessionCanAcceptCalls#6c6c2708: field can_accept_calls: %w", err)
			}
			t.CanAcceptCalls = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSessionID returns value of SessionID field.
func (t *ToggleSessionCanAcceptCallsRequest) GetSessionID() (value int64) {
	if t == nil {
		return
	}
	return t.SessionID
}

// GetCanAcceptCalls returns value of CanAcceptCalls field.
func (t *ToggleSessionCanAcceptCallsRequest) GetCanAcceptCalls() (value bool) {
	if t == nil {
		return
	}
	return t.CanAcceptCalls
}

// ToggleSessionCanAcceptCalls invokes method toggleSessionCanAcceptCalls#6c6c2708 returning error if any.
func (c *Client) ToggleSessionCanAcceptCalls(ctx context.Context, request *ToggleSessionCanAcceptCallsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
