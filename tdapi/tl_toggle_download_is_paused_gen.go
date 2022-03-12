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

// ToggleDownloadIsPausedRequest represents TL type `toggleDownloadIsPaused#c7866715`.
type ToggleDownloadIsPausedRequest struct {
	// Identifier of the downloaded file
	FileID int32
	// Pass true if the download is paused
	IsPaused bool
}

// ToggleDownloadIsPausedRequestTypeID is TL type id of ToggleDownloadIsPausedRequest.
const ToggleDownloadIsPausedRequestTypeID = 0xc7866715

// Ensuring interfaces in compile-time for ToggleDownloadIsPausedRequest.
var (
	_ bin.Encoder     = &ToggleDownloadIsPausedRequest{}
	_ bin.Decoder     = &ToggleDownloadIsPausedRequest{}
	_ bin.BareEncoder = &ToggleDownloadIsPausedRequest{}
	_ bin.BareDecoder = &ToggleDownloadIsPausedRequest{}
)

func (t *ToggleDownloadIsPausedRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.FileID == 0) {
		return false
	}
	if !(t.IsPaused == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleDownloadIsPausedRequest) String() string {
	if t == nil {
		return "ToggleDownloadIsPausedRequest(nil)"
	}
	type Alias ToggleDownloadIsPausedRequest
	return fmt.Sprintf("ToggleDownloadIsPausedRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleDownloadIsPausedRequest) TypeID() uint32 {
	return ToggleDownloadIsPausedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleDownloadIsPausedRequest) TypeName() string {
	return "toggleDownloadIsPaused"
}

// TypeInfo returns info about TL type.
func (t *ToggleDownloadIsPausedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleDownloadIsPaused",
		ID:   ToggleDownloadIsPausedRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileID",
			SchemaName: "file_id",
		},
		{
			Name:       "IsPaused",
			SchemaName: "is_paused",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleDownloadIsPausedRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleDownloadIsPaused#c7866715 as nil")
	}
	b.PutID(ToggleDownloadIsPausedRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleDownloadIsPausedRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleDownloadIsPaused#c7866715 as nil")
	}
	b.PutInt32(t.FileID)
	b.PutBool(t.IsPaused)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleDownloadIsPausedRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleDownloadIsPaused#c7866715 to nil")
	}
	if err := b.ConsumeID(ToggleDownloadIsPausedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleDownloadIsPaused#c7866715: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleDownloadIsPausedRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleDownloadIsPaused#c7866715 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode toggleDownloadIsPaused#c7866715: field file_id: %w", err)
		}
		t.FileID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleDownloadIsPaused#c7866715: field is_paused: %w", err)
		}
		t.IsPaused = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleDownloadIsPausedRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleDownloadIsPaused#c7866715 as nil")
	}
	b.ObjStart()
	b.PutID("toggleDownloadIsPaused")
	b.Comma()
	b.FieldStart("file_id")
	b.PutInt32(t.FileID)
	b.Comma()
	b.FieldStart("is_paused")
	b.PutBool(t.IsPaused)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleDownloadIsPausedRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleDownloadIsPaused#c7866715 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleDownloadIsPaused"); err != nil {
				return fmt.Errorf("unable to decode toggleDownloadIsPaused#c7866715: %w", err)
			}
		case "file_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode toggleDownloadIsPaused#c7866715: field file_id: %w", err)
			}
			t.FileID = value
		case "is_paused":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleDownloadIsPaused#c7866715: field is_paused: %w", err)
			}
			t.IsPaused = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFileID returns value of FileID field.
func (t *ToggleDownloadIsPausedRequest) GetFileID() (value int32) {
	if t == nil {
		return
	}
	return t.FileID
}

// GetIsPaused returns value of IsPaused field.
func (t *ToggleDownloadIsPausedRequest) GetIsPaused() (value bool) {
	if t == nil {
		return
	}
	return t.IsPaused
}

// ToggleDownloadIsPaused invokes method toggleDownloadIsPaused#c7866715 returning error if any.
func (c *Client) ToggleDownloadIsPaused(ctx context.Context, request *ToggleDownloadIsPausedRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
