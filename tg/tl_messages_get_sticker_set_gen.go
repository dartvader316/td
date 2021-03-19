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
)

// MessagesGetStickerSetRequest represents TL type `messages.getStickerSet#2619a90e`.
// Get info about a stickerset
//
// See https://core.telegram.org/method/messages.getStickerSet for reference.
type MessagesGetStickerSetRequest struct {
	// Stickerset
	Stickerset InputStickerSetClass
}

// MessagesGetStickerSetRequestTypeID is TL type id of MessagesGetStickerSetRequest.
const MessagesGetStickerSetRequestTypeID = 0x2619a90e

func (g *MessagesGetStickerSetRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Stickerset == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetStickerSetRequest) String() string {
	if g == nil {
		return "MessagesGetStickerSetRequest(nil)"
	}
	type Alias MessagesGetStickerSetRequest
	return fmt.Sprintf("MessagesGetStickerSetRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetStickerSetRequest from given interface.
func (g *MessagesGetStickerSetRequest) FillFrom(from interface {
	GetStickerset() (value InputStickerSetClass)
}) {
	g.Stickerset = from.GetStickerset()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetStickerSetRequest) TypeID() uint32 {
	return MessagesGetStickerSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetStickerSetRequest) TypeName() string {
	return "messages.getStickerSet"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetStickerSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getStickerSet",
		ID:   MessagesGetStickerSetRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Stickerset",
			SchemaName: "stickerset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetStickerSetRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getStickerSet#2619a90e as nil")
	}
	b.PutID(MessagesGetStickerSetRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetStickerSetRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getStickerSet#2619a90e as nil")
	}
	if g.Stickerset == nil {
		return fmt.Errorf("unable to encode messages.getStickerSet#2619a90e: field stickerset is nil")
	}
	if err := g.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getStickerSet#2619a90e: field stickerset: %w", err)
	}
	return nil
}

// GetStickerset returns value of Stickerset field.
func (g *MessagesGetStickerSetRequest) GetStickerset() (value InputStickerSetClass) {
	return g.Stickerset
}

// Decode implements bin.Decoder.
func (g *MessagesGetStickerSetRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getStickerSet#2619a90e to nil")
	}
	if err := b.ConsumeID(MessagesGetStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getStickerSet#2619a90e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetStickerSetRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getStickerSet#2619a90e to nil")
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getStickerSet#2619a90e: field stickerset: %w", err)
		}
		g.Stickerset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetStickerSetRequest.
var (
	_ bin.Encoder     = &MessagesGetStickerSetRequest{}
	_ bin.Decoder     = &MessagesGetStickerSetRequest{}
	_ bin.BareEncoder = &MessagesGetStickerSetRequest{}
	_ bin.BareDecoder = &MessagesGetStickerSetRequest{}
)

// MessagesGetStickerSet invokes method messages.getStickerSet#2619a90e returning error if any.
// Get info about a stickerset
//
// Possible errors:
//  400 STICKERSET_INVALID: The provided sticker set is invalid
//
// See https://core.telegram.org/method/messages.getStickerSet for reference.
// Can be used by bots.
func (c *Client) MessagesGetStickerSet(ctx context.Context, stickerset InputStickerSetClass) (*MessagesStickerSet, error) {
	var result MessagesStickerSet

	request := &MessagesGetStickerSetRequest{
		Stickerset: stickerset,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
