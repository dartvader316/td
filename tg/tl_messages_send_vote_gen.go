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

// MessagesSendVoteRequest represents TL type `messages.sendVote#10ea6184`.
type MessagesSendVoteRequest struct {
	// Peer field of MessagesSendVoteRequest.
	Peer InputPeerClass
	// MsgID field of MessagesSendVoteRequest.
	MsgID int
	// Options field of MessagesSendVoteRequest.
	Options [][]byte
}

// MessagesSendVoteRequestTypeID is TL type id of MessagesSendVoteRequest.
const MessagesSendVoteRequestTypeID = 0x10ea6184

// Encode implements bin.Encoder.
func (s *MessagesSendVoteRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendVote#10ea6184 as nil")
	}
	b.PutID(MessagesSendVoteRequestTypeID)
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendVote#10ea6184: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendVote#10ea6184: field peer: %w", err)
	}
	b.PutInt(s.MsgID)
	b.PutVectorHeader(len(s.Options))
	for _, v := range s.Options {
		b.PutBytes(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendVoteRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendVote#10ea6184 to nil")
	}
	if err := b.ConsumeID(MessagesSendVoteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendVote#10ea6184: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendVote#10ea6184: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendVote#10ea6184: field msg_id: %w", err)
		}
		s.MsgID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendVote#10ea6184: field options: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode messages.sendVote#10ea6184: field options: %w", err)
			}
			s.Options = append(s.Options, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendVoteRequest.
var (
	_ bin.Encoder = &MessagesSendVoteRequest{}
	_ bin.Decoder = &MessagesSendVoteRequest{}
)

// MessagesSendVote invokes method messages.sendVote#10ea6184 returning error if any.
func (c *Client) MessagesSendVote(ctx context.Context, request *MessagesSendVoteRequest) (UpdatesClass, error) {
	var result UpdatesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}