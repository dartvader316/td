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

// FileHash represents TL type `fileHash#6242c773`.
type FileHash struct {
	// Offset field of FileHash.
	Offset int
	// Limit field of FileHash.
	Limit int
	// Hash field of FileHash.
	Hash []byte
}

// FileHashTypeID is TL type id of FileHash.
const FileHashTypeID = 0x6242c773

// Encode implements bin.Encoder.
func (f *FileHash) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileHash#6242c773 as nil")
	}
	b.PutID(FileHashTypeID)
	b.PutInt(f.Offset)
	b.PutInt(f.Limit)
	b.PutBytes(f.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (f *FileHash) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileHash#6242c773 to nil")
	}
	if err := b.ConsumeID(FileHashTypeID); err != nil {
		return fmt.Errorf("unable to decode fileHash#6242c773: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode fileHash#6242c773: field offset: %w", err)
		}
		f.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode fileHash#6242c773: field limit: %w", err)
		}
		f.Limit = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode fileHash#6242c773: field hash: %w", err)
		}
		f.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for FileHash.
var (
	_ bin.Encoder = &FileHash{}
	_ bin.Decoder = &FileHash{}
)