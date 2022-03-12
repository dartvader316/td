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

// GetGroupCallStreamsRequest represents TL type `getGroupCallStreams#9f7c9164`.
type GetGroupCallStreamsRequest struct {
	// Group call identifier
	GroupCallID int32
}

// GetGroupCallStreamsRequestTypeID is TL type id of GetGroupCallStreamsRequest.
const GetGroupCallStreamsRequestTypeID = 0x9f7c9164

// Ensuring interfaces in compile-time for GetGroupCallStreamsRequest.
var (
	_ bin.Encoder     = &GetGroupCallStreamsRequest{}
	_ bin.Decoder     = &GetGroupCallStreamsRequest{}
	_ bin.BareEncoder = &GetGroupCallStreamsRequest{}
	_ bin.BareDecoder = &GetGroupCallStreamsRequest{}
)

func (g *GetGroupCallStreamsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.GroupCallID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetGroupCallStreamsRequest) String() string {
	if g == nil {
		return "GetGroupCallStreamsRequest(nil)"
	}
	type Alias GetGroupCallStreamsRequest
	return fmt.Sprintf("GetGroupCallStreamsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetGroupCallStreamsRequest) TypeID() uint32 {
	return GetGroupCallStreamsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetGroupCallStreamsRequest) TypeName() string {
	return "getGroupCallStreams"
}

// TypeInfo returns info about TL type.
func (g *GetGroupCallStreamsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getGroupCallStreams",
		ID:   GetGroupCallStreamsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetGroupCallStreamsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCallStreams#9f7c9164 as nil")
	}
	b.PutID(GetGroupCallStreamsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetGroupCallStreamsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCallStreams#9f7c9164 as nil")
	}
	b.PutInt32(g.GroupCallID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetGroupCallStreamsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCallStreams#9f7c9164 to nil")
	}
	if err := b.ConsumeID(GetGroupCallStreamsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getGroupCallStreams#9f7c9164: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetGroupCallStreamsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCallStreams#9f7c9164 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getGroupCallStreams#9f7c9164: field group_call_id: %w", err)
		}
		g.GroupCallID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetGroupCallStreamsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupCallStreams#9f7c9164 as nil")
	}
	b.ObjStart()
	b.PutID("getGroupCallStreams")
	b.Comma()
	b.FieldStart("group_call_id")
	b.PutInt32(g.GroupCallID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetGroupCallStreamsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupCallStreams#9f7c9164 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getGroupCallStreams"); err != nil {
				return fmt.Errorf("unable to decode getGroupCallStreams#9f7c9164: %w", err)
			}
		case "group_call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getGroupCallStreams#9f7c9164: field group_call_id: %w", err)
			}
			g.GroupCallID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGroupCallID returns value of GroupCallID field.
func (g *GetGroupCallStreamsRequest) GetGroupCallID() (value int32) {
	if g == nil {
		return
	}
	return g.GroupCallID
}

// GetGroupCallStreams invokes method getGroupCallStreams#9f7c9164 returning error if any.
func (c *Client) GetGroupCallStreams(ctx context.Context, groupcallid int32) (*GroupCallStreams, error) {
	var result GroupCallStreams

	request := &GetGroupCallStreamsRequest{
		GroupCallID: groupcallid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
