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

// EditProxyRequest represents TL type `editProxy#a0482853`.
type EditProxyRequest struct {
	// Proxy identifier
	ProxyID int32
	// Proxy server IP address
	Server string
	// Proxy server port
	Port int32
	// Pass true to immediately enable the proxy
	Enable bool
	// Proxy type
	Type ProxyTypeClass
}

// EditProxyRequestTypeID is TL type id of EditProxyRequest.
const EditProxyRequestTypeID = 0xa0482853

// Ensuring interfaces in compile-time for EditProxyRequest.
var (
	_ bin.Encoder     = &EditProxyRequest{}
	_ bin.Decoder     = &EditProxyRequest{}
	_ bin.BareEncoder = &EditProxyRequest{}
	_ bin.BareDecoder = &EditProxyRequest{}
)

func (e *EditProxyRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ProxyID == 0) {
		return false
	}
	if !(e.Server == "") {
		return false
	}
	if !(e.Port == 0) {
		return false
	}
	if !(e.Enable == false) {
		return false
	}
	if !(e.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditProxyRequest) String() string {
	if e == nil {
		return "EditProxyRequest(nil)"
	}
	type Alias EditProxyRequest
	return fmt.Sprintf("EditProxyRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditProxyRequest) TypeID() uint32 {
	return EditProxyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditProxyRequest) TypeName() string {
	return "editProxy"
}

// TypeInfo returns info about TL type.
func (e *EditProxyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editProxy",
		ID:   EditProxyRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ProxyID",
			SchemaName: "proxy_id",
		},
		{
			Name:       "Server",
			SchemaName: "server",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
		{
			Name:       "Enable",
			SchemaName: "enable",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditProxyRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editProxy#a0482853 as nil")
	}
	b.PutID(EditProxyRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditProxyRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editProxy#a0482853 as nil")
	}
	b.PutInt32(e.ProxyID)
	b.PutString(e.Server)
	b.PutInt32(e.Port)
	b.PutBool(e.Enable)
	if e.Type == nil {
		return fmt.Errorf("unable to encode editProxy#a0482853: field type is nil")
	}
	if err := e.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editProxy#a0482853: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditProxyRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editProxy#a0482853 to nil")
	}
	if err := b.ConsumeID(EditProxyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editProxy#a0482853: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditProxyRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editProxy#a0482853 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode editProxy#a0482853: field proxy_id: %w", err)
		}
		e.ProxyID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode editProxy#a0482853: field server: %w", err)
		}
		e.Server = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode editProxy#a0482853: field port: %w", err)
		}
		e.Port = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode editProxy#a0482853: field enable: %w", err)
		}
		e.Enable = value
	}
	{
		value, err := DecodeProxyType(b)
		if err != nil {
			return fmt.Errorf("unable to decode editProxy#a0482853: field type: %w", err)
		}
		e.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditProxyRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editProxy#a0482853 as nil")
	}
	b.ObjStart()
	b.PutID("editProxy")
	b.Comma()
	b.FieldStart("proxy_id")
	b.PutInt32(e.ProxyID)
	b.Comma()
	b.FieldStart("server")
	b.PutString(e.Server)
	b.Comma()
	b.FieldStart("port")
	b.PutInt32(e.Port)
	b.Comma()
	b.FieldStart("enable")
	b.PutBool(e.Enable)
	b.Comma()
	b.FieldStart("type")
	if e.Type == nil {
		return fmt.Errorf("unable to encode editProxy#a0482853: field type is nil")
	}
	if err := e.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editProxy#a0482853: field type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditProxyRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editProxy#a0482853 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editProxy"); err != nil {
				return fmt.Errorf("unable to decode editProxy#a0482853: %w", err)
			}
		case "proxy_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode editProxy#a0482853: field proxy_id: %w", err)
			}
			e.ProxyID = value
		case "server":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode editProxy#a0482853: field server: %w", err)
			}
			e.Server = value
		case "port":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode editProxy#a0482853: field port: %w", err)
			}
			e.Port = value
		case "enable":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode editProxy#a0482853: field enable: %w", err)
			}
			e.Enable = value
		case "type":
			value, err := DecodeTDLibJSONProxyType(b)
			if err != nil {
				return fmt.Errorf("unable to decode editProxy#a0482853: field type: %w", err)
			}
			e.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetProxyID returns value of ProxyID field.
func (e *EditProxyRequest) GetProxyID() (value int32) {
	if e == nil {
		return
	}
	return e.ProxyID
}

// GetServer returns value of Server field.
func (e *EditProxyRequest) GetServer() (value string) {
	if e == nil {
		return
	}
	return e.Server
}

// GetPort returns value of Port field.
func (e *EditProxyRequest) GetPort() (value int32) {
	if e == nil {
		return
	}
	return e.Port
}

// GetEnable returns value of Enable field.
func (e *EditProxyRequest) GetEnable() (value bool) {
	if e == nil {
		return
	}
	return e.Enable
}

// GetType returns value of Type field.
func (e *EditProxyRequest) GetType() (value ProxyTypeClass) {
	if e == nil {
		return
	}
	return e.Type
}

// EditProxy invokes method editProxy#a0482853 returning error if any.
func (c *Client) EditProxy(ctx context.Context, request *EditProxyRequest) (*Proxy, error) {
	var result Proxy

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
