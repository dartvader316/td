// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// MessagesGetDiscussionMessageRequest represents TL type `messages.getDiscussionMessage#446972fd`.
type MessagesGetDiscussionMessageRequest struct {
	// Peer field of MessagesGetDiscussionMessageRequest.
	Peer InputPeerClass
	// MsgID field of MessagesGetDiscussionMessageRequest.
	MsgID int
}

// MessagesGetDiscussionMessageRequestTypeID is TL type id of MessagesGetDiscussionMessageRequest.
const MessagesGetDiscussionMessageRequestTypeID = 0x446972fd

// Encode implements bin.Encoder.
func (g *MessagesGetDiscussionMessageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDiscussionMessage#446972fd as nil")
	}
	b.PutID(MessagesGetDiscussionMessageRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getDiscussionMessage#446972fd: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getDiscussionMessage#446972fd: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDiscussionMessageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDiscussionMessage#446972fd to nil")
	}
	if err := b.ConsumeID(MessagesGetDiscussionMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDiscussionMessage#446972fd: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDiscussionMessage#446972fd: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDiscussionMessage#446972fd: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDiscussionMessageRequest.
var (
	_ bin.Encoder = &MessagesGetDiscussionMessageRequest{}
	_ bin.Decoder = &MessagesGetDiscussionMessageRequest{}
)

// MessagesGetDiscussionMessage invokes method messages.getDiscussionMessage#446972fd returning error if any.
func (c *Client) MessagesGetDiscussionMessage(ctx context.Context, request *MessagesGetDiscussionMessageRequest) (*MessagesDiscussionMessage, error) {
	var result MessagesDiscussionMessage
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}