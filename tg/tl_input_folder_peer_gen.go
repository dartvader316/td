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

// InputFolderPeer represents TL type `inputFolderPeer#fbd2c296`.
type InputFolderPeer struct {
	// Peer field of InputFolderPeer.
	Peer InputPeerClass
	// FolderID field of InputFolderPeer.
	FolderID int
}

// InputFolderPeerTypeID is TL type id of InputFolderPeer.
const InputFolderPeerTypeID = 0xfbd2c296

// Encode implements bin.Encoder.
func (i *InputFolderPeer) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFolderPeer#fbd2c296 as nil")
	}
	b.PutID(InputFolderPeerTypeID)
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputFolderPeer#fbd2c296: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputFolderPeer#fbd2c296: field peer: %w", err)
	}
	b.PutInt(i.FolderID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputFolderPeer) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFolderPeer#fbd2c296 to nil")
	}
	if err := b.ConsumeID(InputFolderPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode inputFolderPeer#fbd2c296: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputFolderPeer#fbd2c296: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputFolderPeer#fbd2c296: field folder_id: %w", err)
		}
		i.FolderID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputFolderPeer.
var (
	_ bin.Encoder = &InputFolderPeer{}
	_ bin.Decoder = &InputFolderPeer{}
)