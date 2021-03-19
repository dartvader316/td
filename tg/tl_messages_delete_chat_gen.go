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

// MessagesDeleteChatRequest represents TL type `messages.deleteChat#83247d11`.
//
// See https://core.telegram.org/method/messages.deleteChat for reference.
type MessagesDeleteChatRequest struct {
	// ChatID field of MessagesDeleteChatRequest.
	ChatID int
}

// MessagesDeleteChatRequestTypeID is TL type id of MessagesDeleteChatRequest.
const MessagesDeleteChatRequestTypeID = 0x83247d11

func (d *MessagesDeleteChatRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeleteChatRequest) String() string {
	if d == nil {
		return "MessagesDeleteChatRequest(nil)"
	}
	type Alias MessagesDeleteChatRequest
	return fmt.Sprintf("MessagesDeleteChatRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeleteChatRequest from given interface.
func (d *MessagesDeleteChatRequest) FillFrom(from interface {
	GetChatID() (value int)
}) {
	d.ChatID = from.GetChatID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDeleteChatRequest) TypeID() uint32 {
	return MessagesDeleteChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDeleteChatRequest) TypeName() string {
	return "messages.deleteChat"
}

// TypeInfo returns info about TL type.
func (d *MessagesDeleteChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.deleteChat",
		ID:   MessagesDeleteChatRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDeleteChatRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteChat#83247d11 as nil")
	}
	b.PutID(MessagesDeleteChatRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDeleteChatRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteChat#83247d11 as nil")
	}
	b.PutInt(d.ChatID)
	return nil
}

// GetChatID returns value of ChatID field.
func (d *MessagesDeleteChatRequest) GetChatID() (value int) {
	return d.ChatID
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteChatRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteChat#83247d11 to nil")
	}
	if err := b.ConsumeID(MessagesDeleteChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteChat#83247d11: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDeleteChatRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteChat#83247d11 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteChat#83247d11: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesDeleteChatRequest.
var (
	_ bin.Encoder     = &MessagesDeleteChatRequest{}
	_ bin.Decoder     = &MessagesDeleteChatRequest{}
	_ bin.BareEncoder = &MessagesDeleteChatRequest{}
	_ bin.BareDecoder = &MessagesDeleteChatRequest{}
)

// MessagesDeleteChat invokes method messages.deleteChat#83247d11 returning error if any.
//
// See https://core.telegram.org/method/messages.deleteChat for reference.
func (c *Client) MessagesDeleteChat(ctx context.Context, chatid int) (bool, error) {
	var result BoolBox

	request := &MessagesDeleteChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
