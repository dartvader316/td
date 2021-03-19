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

// HelpDeepLinkInfoEmpty represents TL type `help.deepLinkInfoEmpty#66afa166`.
// Deep link info empty
//
// See https://core.telegram.org/constructor/help.deepLinkInfoEmpty for reference.
type HelpDeepLinkInfoEmpty struct {
}

// HelpDeepLinkInfoEmptyTypeID is TL type id of HelpDeepLinkInfoEmpty.
const HelpDeepLinkInfoEmptyTypeID = 0x66afa166

func (d *HelpDeepLinkInfoEmpty) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *HelpDeepLinkInfoEmpty) String() string {
	if d == nil {
		return "HelpDeepLinkInfoEmpty(nil)"
	}
	type Alias HelpDeepLinkInfoEmpty
	return fmt.Sprintf("HelpDeepLinkInfoEmpty%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpDeepLinkInfoEmpty) TypeID() uint32 {
	return HelpDeepLinkInfoEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpDeepLinkInfoEmpty) TypeName() string {
	return "help.deepLinkInfoEmpty"
}

// TypeInfo returns info about TL type.
func (d *HelpDeepLinkInfoEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.deepLinkInfoEmpty",
		ID:   HelpDeepLinkInfoEmptyTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (d *HelpDeepLinkInfoEmpty) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode help.deepLinkInfoEmpty#66afa166 as nil")
	}
	b.PutID(HelpDeepLinkInfoEmptyTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *HelpDeepLinkInfoEmpty) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode help.deepLinkInfoEmpty#66afa166 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *HelpDeepLinkInfoEmpty) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode help.deepLinkInfoEmpty#66afa166 to nil")
	}
	if err := b.ConsumeID(HelpDeepLinkInfoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode help.deepLinkInfoEmpty#66afa166: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *HelpDeepLinkInfoEmpty) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode help.deepLinkInfoEmpty#66afa166 to nil")
	}
	return nil
}

// construct implements constructor of HelpDeepLinkInfoClass.
func (d HelpDeepLinkInfoEmpty) construct() HelpDeepLinkInfoClass { return &d }

// Ensuring interfaces in compile-time for HelpDeepLinkInfoEmpty.
var (
	_ bin.Encoder     = &HelpDeepLinkInfoEmpty{}
	_ bin.Decoder     = &HelpDeepLinkInfoEmpty{}
	_ bin.BareEncoder = &HelpDeepLinkInfoEmpty{}
	_ bin.BareDecoder = &HelpDeepLinkInfoEmpty{}

	_ HelpDeepLinkInfoClass = &HelpDeepLinkInfoEmpty{}
)

// HelpDeepLinkInfo represents TL type `help.deepLinkInfo#6a4ee832`.
// Deep linking info
//
// See https://core.telegram.org/constructor/help.deepLinkInfo for reference.
type HelpDeepLinkInfo struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// An update of the app is required to parse this link
	UpdateApp bool
	// Message to show to the user
	Message string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
}

// HelpDeepLinkInfoTypeID is TL type id of HelpDeepLinkInfo.
const HelpDeepLinkInfoTypeID = 0x6a4ee832

func (d *HelpDeepLinkInfo) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.UpdateApp == false) {
		return false
	}
	if !(d.Message == "") {
		return false
	}
	if !(d.Entities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *HelpDeepLinkInfo) String() string {
	if d == nil {
		return "HelpDeepLinkInfo(nil)"
	}
	type Alias HelpDeepLinkInfo
	return fmt.Sprintf("HelpDeepLinkInfo%+v", Alias(*d))
}

// FillFrom fills HelpDeepLinkInfo from given interface.
func (d *HelpDeepLinkInfo) FillFrom(from interface {
	GetUpdateApp() (value bool)
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass, ok bool)
}) {
	d.UpdateApp = from.GetUpdateApp()
	d.Message = from.GetMessage()
	if val, ok := from.GetEntities(); ok {
		d.Entities = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpDeepLinkInfo) TypeID() uint32 {
	return HelpDeepLinkInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpDeepLinkInfo) TypeName() string {
	return "help.deepLinkInfo"
}

// TypeInfo returns info about TL type.
func (d *HelpDeepLinkInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.deepLinkInfo",
		ID:   HelpDeepLinkInfoTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UpdateApp",
			SchemaName: "update_app",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !d.Flags.Has(1),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *HelpDeepLinkInfo) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode help.deepLinkInfo#6a4ee832 as nil")
	}
	b.PutID(HelpDeepLinkInfoTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *HelpDeepLinkInfo) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode help.deepLinkInfo#6a4ee832 as nil")
	}
	if !(d.UpdateApp == false) {
		d.Flags.Set(0)
	}
	if !(d.Entities == nil) {
		d.Flags.Set(1)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.deepLinkInfo#6a4ee832: field flags: %w", err)
	}
	b.PutString(d.Message)
	if d.Flags.Has(1) {
		b.PutVectorHeader(len(d.Entities))
		for idx, v := range d.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode help.deepLinkInfo#6a4ee832: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode help.deepLinkInfo#6a4ee832: field entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// SetUpdateApp sets value of UpdateApp conditional field.
func (d *HelpDeepLinkInfo) SetUpdateApp(value bool) {
	if value {
		d.Flags.Set(0)
		d.UpdateApp = true
	} else {
		d.Flags.Unset(0)
		d.UpdateApp = false
	}
}

// GetUpdateApp returns value of UpdateApp conditional field.
func (d *HelpDeepLinkInfo) GetUpdateApp() (value bool) {
	return d.Flags.Has(0)
}

// GetMessage returns value of Message field.
func (d *HelpDeepLinkInfo) GetMessage() (value string) {
	return d.Message
}

// SetEntities sets value of Entities conditional field.
func (d *HelpDeepLinkInfo) SetEntities(value []MessageEntityClass) {
	d.Flags.Set(1)
	d.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (d *HelpDeepLinkInfo) GetEntities() (value []MessageEntityClass, ok bool) {
	if !d.Flags.Has(1) {
		return value, false
	}
	return d.Entities, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (d *HelpDeepLinkInfo) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !d.Flags.Has(1) {
		return value, false
	}
	return MessageEntityClassArray(d.Entities), true
}

// Decode implements bin.Decoder.
func (d *HelpDeepLinkInfo) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode help.deepLinkInfo#6a4ee832 to nil")
	}
	if err := b.ConsumeID(HelpDeepLinkInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode help.deepLinkInfo#6a4ee832: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *HelpDeepLinkInfo) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode help.deepLinkInfo#6a4ee832 to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.deepLinkInfo#6a4ee832: field flags: %w", err)
		}
	}
	d.UpdateApp = d.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.deepLinkInfo#6a4ee832: field message: %w", err)
		}
		d.Message = value
	}
	if d.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.deepLinkInfo#6a4ee832: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.deepLinkInfo#6a4ee832: field entities: %w", err)
			}
			d.Entities = append(d.Entities, value)
		}
	}
	return nil
}

// construct implements constructor of HelpDeepLinkInfoClass.
func (d HelpDeepLinkInfo) construct() HelpDeepLinkInfoClass { return &d }

// Ensuring interfaces in compile-time for HelpDeepLinkInfo.
var (
	_ bin.Encoder     = &HelpDeepLinkInfo{}
	_ bin.Decoder     = &HelpDeepLinkInfo{}
	_ bin.BareEncoder = &HelpDeepLinkInfo{}
	_ bin.BareDecoder = &HelpDeepLinkInfo{}

	_ HelpDeepLinkInfoClass = &HelpDeepLinkInfo{}
)

// HelpDeepLinkInfoClass represents help.DeepLinkInfo generic type.
//
// See https://core.telegram.org/type/help.DeepLinkInfo for reference.
//
// Example:
//  g, err := tg.DecodeHelpDeepLinkInfo(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.HelpDeepLinkInfoEmpty: // help.deepLinkInfoEmpty#66afa166
//  case *tg.HelpDeepLinkInfo: // help.deepLinkInfo#6a4ee832
//  default: panic(v)
//  }
type HelpDeepLinkInfoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() HelpDeepLinkInfoClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsNotEmpty tries to map HelpDeepLinkInfoClass to HelpDeepLinkInfo.
	AsNotEmpty() (*HelpDeepLinkInfo, bool)
}

// AsNotEmpty tries to map HelpDeepLinkInfoEmpty to HelpDeepLinkInfo.
func (d *HelpDeepLinkInfoEmpty) AsNotEmpty() (*HelpDeepLinkInfo, bool) {
	return nil, false
}

// AsNotEmpty tries to map HelpDeepLinkInfo to HelpDeepLinkInfo.
func (d *HelpDeepLinkInfo) AsNotEmpty() (*HelpDeepLinkInfo, bool) {
	return d, true
}

// DecodeHelpDeepLinkInfo implements binary de-serialization for HelpDeepLinkInfoClass.
func DecodeHelpDeepLinkInfo(buf *bin.Buffer) (HelpDeepLinkInfoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpDeepLinkInfoEmptyTypeID:
		// Decoding help.deepLinkInfoEmpty#66afa166.
		v := HelpDeepLinkInfoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpDeepLinkInfoClass: %w", err)
		}
		return &v, nil
	case HelpDeepLinkInfoTypeID:
		// Decoding help.deepLinkInfo#6a4ee832.
		v := HelpDeepLinkInfo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpDeepLinkInfoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpDeepLinkInfoClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpDeepLinkInfo boxes the HelpDeepLinkInfoClass providing a helper.
type HelpDeepLinkInfoBox struct {
	DeepLinkInfo HelpDeepLinkInfoClass
}

// Decode implements bin.Decoder for HelpDeepLinkInfoBox.
func (b *HelpDeepLinkInfoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpDeepLinkInfoBox to nil")
	}
	v, err := DecodeHelpDeepLinkInfo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DeepLinkInfo = v
	return nil
}

// Encode implements bin.Encode for HelpDeepLinkInfoBox.
func (b *HelpDeepLinkInfoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DeepLinkInfo == nil {
		return fmt.Errorf("unable to encode HelpDeepLinkInfoClass as nil")
	}
	return b.DeepLinkInfo.Encode(buf)
}

// HelpDeepLinkInfoClassArray is adapter for slice of HelpDeepLinkInfoClass.
type HelpDeepLinkInfoClassArray []HelpDeepLinkInfoClass

// Sort sorts slice of HelpDeepLinkInfoClass.
func (s HelpDeepLinkInfoClassArray) Sort(less func(a, b HelpDeepLinkInfoClass) bool) HelpDeepLinkInfoClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpDeepLinkInfoClass.
func (s HelpDeepLinkInfoClassArray) SortStable(less func(a, b HelpDeepLinkInfoClass) bool) HelpDeepLinkInfoClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpDeepLinkInfoClass.
func (s HelpDeepLinkInfoClassArray) Retain(keep func(x HelpDeepLinkInfoClass) bool) HelpDeepLinkInfoClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s HelpDeepLinkInfoClassArray) First() (v HelpDeepLinkInfoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpDeepLinkInfoClassArray) Last() (v HelpDeepLinkInfoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpDeepLinkInfoClassArray) PopFirst() (v HelpDeepLinkInfoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpDeepLinkInfoClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpDeepLinkInfoClassArray) Pop() (v HelpDeepLinkInfoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsHelpDeepLinkInfo returns copy with only HelpDeepLinkInfo constructors.
func (s HelpDeepLinkInfoClassArray) AsHelpDeepLinkInfo() (to HelpDeepLinkInfoArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpDeepLinkInfo)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s HelpDeepLinkInfoClassArray) AppendOnlyNotEmpty(to []*HelpDeepLinkInfo) []*HelpDeepLinkInfo {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s HelpDeepLinkInfoClassArray) AsNotEmpty() (to []*HelpDeepLinkInfo) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s HelpDeepLinkInfoClassArray) FirstAsNotEmpty() (v *HelpDeepLinkInfo, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s HelpDeepLinkInfoClassArray) LastAsNotEmpty() (v *HelpDeepLinkInfo, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *HelpDeepLinkInfoClassArray) PopFirstAsNotEmpty() (v *HelpDeepLinkInfo, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *HelpDeepLinkInfoClassArray) PopAsNotEmpty() (v *HelpDeepLinkInfo, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// HelpDeepLinkInfoArray is adapter for slice of HelpDeepLinkInfo.
type HelpDeepLinkInfoArray []HelpDeepLinkInfo

// Sort sorts slice of HelpDeepLinkInfo.
func (s HelpDeepLinkInfoArray) Sort(less func(a, b HelpDeepLinkInfo) bool) HelpDeepLinkInfoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpDeepLinkInfo.
func (s HelpDeepLinkInfoArray) SortStable(less func(a, b HelpDeepLinkInfo) bool) HelpDeepLinkInfoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpDeepLinkInfo.
func (s HelpDeepLinkInfoArray) Retain(keep func(x HelpDeepLinkInfo) bool) HelpDeepLinkInfoArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s HelpDeepLinkInfoArray) First() (v HelpDeepLinkInfo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpDeepLinkInfoArray) Last() (v HelpDeepLinkInfo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpDeepLinkInfoArray) PopFirst() (v HelpDeepLinkInfo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpDeepLinkInfo
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpDeepLinkInfoArray) Pop() (v HelpDeepLinkInfo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
