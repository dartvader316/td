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

// FileLocationToBeDeprecated represents TL type `fileLocationToBeDeprecated#bc7fc6cd`.
// Indicates the location of a photo, will be deprecated soon
//
// See https://core.telegram.org/constructor/fileLocationToBeDeprecated for reference.
type FileLocationToBeDeprecated struct {
	// Volume ID
	VolumeID int64
	// Local ID
	LocalID int
}

// FileLocationToBeDeprecatedTypeID is TL type id of FileLocationToBeDeprecated.
const FileLocationToBeDeprecatedTypeID = 0xbc7fc6cd

func (f *FileLocationToBeDeprecated) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.VolumeID == 0) {
		return false
	}
	if !(f.LocalID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FileLocationToBeDeprecated) String() string {
	if f == nil {
		return "FileLocationToBeDeprecated(nil)"
	}
	type Alias FileLocationToBeDeprecated
	return fmt.Sprintf("FileLocationToBeDeprecated%+v", Alias(*f))
}

// FillFrom fills FileLocationToBeDeprecated from given interface.
func (f *FileLocationToBeDeprecated) FillFrom(from interface {
	GetVolumeID() (value int64)
	GetLocalID() (value int)
}) {
	f.VolumeID = from.GetVolumeID()
	f.LocalID = from.GetLocalID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FileLocationToBeDeprecated) TypeID() uint32 {
	return FileLocationToBeDeprecatedTypeID
}

// TypeName returns name of type in TL schema.
func (*FileLocationToBeDeprecated) TypeName() string {
	return "fileLocationToBeDeprecated"
}

// TypeInfo returns info about TL type.
func (f *FileLocationToBeDeprecated) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "fileLocationToBeDeprecated",
		ID:   FileLocationToBeDeprecatedTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "VolumeID",
			SchemaName: "volume_id",
		},
		{
			Name:       "LocalID",
			SchemaName: "local_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FileLocationToBeDeprecated) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileLocationToBeDeprecated#bc7fc6cd as nil")
	}
	b.PutID(FileLocationToBeDeprecatedTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FileLocationToBeDeprecated) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fileLocationToBeDeprecated#bc7fc6cd as nil")
	}
	b.PutLong(f.VolumeID)
	b.PutInt(f.LocalID)
	return nil
}

// GetVolumeID returns value of VolumeID field.
func (f *FileLocationToBeDeprecated) GetVolumeID() (value int64) {
	return f.VolumeID
}

// GetLocalID returns value of LocalID field.
func (f *FileLocationToBeDeprecated) GetLocalID() (value int) {
	return f.LocalID
}

// Decode implements bin.Decoder.
func (f *FileLocationToBeDeprecated) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileLocationToBeDeprecated#bc7fc6cd to nil")
	}
	if err := b.ConsumeID(FileLocationToBeDeprecatedTypeID); err != nil {
		return fmt.Errorf("unable to decode fileLocationToBeDeprecated#bc7fc6cd: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FileLocationToBeDeprecated) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fileLocationToBeDeprecated#bc7fc6cd to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode fileLocationToBeDeprecated#bc7fc6cd: field volume_id: %w", err)
		}
		f.VolumeID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode fileLocationToBeDeprecated#bc7fc6cd: field local_id: %w", err)
		}
		f.LocalID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for FileLocationToBeDeprecated.
var (
	_ bin.Encoder     = &FileLocationToBeDeprecated{}
	_ bin.Decoder     = &FileLocationToBeDeprecated{}
	_ bin.BareEncoder = &FileLocationToBeDeprecated{}
	_ bin.BareDecoder = &FileLocationToBeDeprecated{}
)
