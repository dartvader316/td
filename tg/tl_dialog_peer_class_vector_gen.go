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

// DialogPeerClassVector is a box for Vector<DialogPeer>
type DialogPeerClassVector struct {
	// Elements of Vector<DialogPeer>
	Elems []DialogPeerClass
}

// Encode implements bin.Encoder.
func (vec *DialogPeerClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<DialogPeer> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<DialogPeer>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<DialogPeer>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *DialogPeerClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<DialogPeer> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<DialogPeer>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDialogPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<DialogPeer>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for DialogPeerClassVector.
var (
	_ bin.Encoder = &DialogPeerClassVector{}
	_ bin.Decoder = &DialogPeerClassVector{}
)