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

// AccountWallPapersNotModified represents TL type `account.wallPapersNotModified#1c199183`.
// No new wallpapers were found
//
// See https://core.telegram.org/constructor/account.wallPapersNotModified for reference.
type AccountWallPapersNotModified struct {
}

// AccountWallPapersNotModifiedTypeID is TL type id of AccountWallPapersNotModified.
const AccountWallPapersNotModifiedTypeID = 0x1c199183

func (w *AccountWallPapersNotModified) Zero() bool {
	if w == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (w *AccountWallPapersNotModified) String() string {
	if w == nil {
		return "AccountWallPapersNotModified(nil)"
	}
	type Alias AccountWallPapersNotModified
	return fmt.Sprintf("AccountWallPapersNotModified%+v", Alias(*w))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountWallPapersNotModified) TypeID() uint32 {
	return AccountWallPapersNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountWallPapersNotModified) TypeName() string {
	return "account.wallPapersNotModified"
}

// TypeInfo returns info about TL type.
func (w *AccountWallPapersNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.wallPapersNotModified",
		ID:   AccountWallPapersNotModifiedTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (w *AccountWallPapersNotModified) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode account.wallPapersNotModified#1c199183 as nil")
	}
	b.PutID(AccountWallPapersNotModifiedTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *AccountWallPapersNotModified) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode account.wallPapersNotModified#1c199183 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *AccountWallPapersNotModified) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode account.wallPapersNotModified#1c199183 to nil")
	}
	if err := b.ConsumeID(AccountWallPapersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode account.wallPapersNotModified#1c199183: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *AccountWallPapersNotModified) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode account.wallPapersNotModified#1c199183 to nil")
	}
	return nil
}

// construct implements constructor of AccountWallPapersClass.
func (w AccountWallPapersNotModified) construct() AccountWallPapersClass { return &w }

// Ensuring interfaces in compile-time for AccountWallPapersNotModified.
var (
	_ bin.Encoder     = &AccountWallPapersNotModified{}
	_ bin.Decoder     = &AccountWallPapersNotModified{}
	_ bin.BareEncoder = &AccountWallPapersNotModified{}
	_ bin.BareDecoder = &AccountWallPapersNotModified{}

	_ AccountWallPapersClass = &AccountWallPapersNotModified{}
)

// AccountWallPapers represents TL type `account.wallPapers#702b65a9`.
// Installed wallpapers
//
// See https://core.telegram.org/constructor/account.wallPapers for reference.
type AccountWallPapers struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// Wallpapers
	Wallpapers []WallPaperClass
}

// AccountWallPapersTypeID is TL type id of AccountWallPapers.
const AccountWallPapersTypeID = 0x702b65a9

func (w *AccountWallPapers) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Hash == 0) {
		return false
	}
	if !(w.Wallpapers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *AccountWallPapers) String() string {
	if w == nil {
		return "AccountWallPapers(nil)"
	}
	type Alias AccountWallPapers
	return fmt.Sprintf("AccountWallPapers%+v", Alias(*w))
}

// FillFrom fills AccountWallPapers from given interface.
func (w *AccountWallPapers) FillFrom(from interface {
	GetHash() (value int)
	GetWallpapers() (value []WallPaperClass)
}) {
	w.Hash = from.GetHash()
	w.Wallpapers = from.GetWallpapers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountWallPapers) TypeID() uint32 {
	return AccountWallPapersTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountWallPapers) TypeName() string {
	return "account.wallPapers"
}

// TypeInfo returns info about TL type.
func (w *AccountWallPapers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.wallPapers",
		ID:   AccountWallPapersTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Wallpapers",
			SchemaName: "wallpapers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *AccountWallPapers) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode account.wallPapers#702b65a9 as nil")
	}
	b.PutID(AccountWallPapersTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *AccountWallPapers) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode account.wallPapers#702b65a9 as nil")
	}
	b.PutInt(w.Hash)
	b.PutVectorHeader(len(w.Wallpapers))
	for idx, v := range w.Wallpapers {
		if v == nil {
			return fmt.Errorf("unable to encode account.wallPapers#702b65a9: field wallpapers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.wallPapers#702b65a9: field wallpapers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (w *AccountWallPapers) GetHash() (value int) {
	return w.Hash
}

// GetWallpapers returns value of Wallpapers field.
func (w *AccountWallPapers) GetWallpapers() (value []WallPaperClass) {
	return w.Wallpapers
}

// MapWallpapers returns field Wallpapers wrapped in WallPaperClassArray helper.
func (w *AccountWallPapers) MapWallpapers() (value WallPaperClassArray) {
	return WallPaperClassArray(w.Wallpapers)
}

// Decode implements bin.Decoder.
func (w *AccountWallPapers) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode account.wallPapers#702b65a9 to nil")
	}
	if err := b.ConsumeID(AccountWallPapersTypeID); err != nil {
		return fmt.Errorf("unable to decode account.wallPapers#702b65a9: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *AccountWallPapers) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode account.wallPapers#702b65a9 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.wallPapers#702b65a9: field hash: %w", err)
		}
		w.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.wallPapers#702b65a9: field wallpapers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeWallPaper(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.wallPapers#702b65a9: field wallpapers: %w", err)
			}
			w.Wallpapers = append(w.Wallpapers, value)
		}
	}
	return nil
}

// construct implements constructor of AccountWallPapersClass.
func (w AccountWallPapers) construct() AccountWallPapersClass { return &w }

// Ensuring interfaces in compile-time for AccountWallPapers.
var (
	_ bin.Encoder     = &AccountWallPapers{}
	_ bin.Decoder     = &AccountWallPapers{}
	_ bin.BareEncoder = &AccountWallPapers{}
	_ bin.BareDecoder = &AccountWallPapers{}

	_ AccountWallPapersClass = &AccountWallPapers{}
)

// AccountWallPapersClass represents account.WallPapers generic type.
//
// See https://core.telegram.org/type/account.WallPapers for reference.
//
// Example:
//  g, err := tg.DecodeAccountWallPapers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.AccountWallPapersNotModified: // account.wallPapersNotModified#1c199183
//  case *tg.AccountWallPapers: // account.wallPapers#702b65a9
//  default: panic(v)
//  }
type AccountWallPapersClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AccountWallPapersClass

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

	// AsModified tries to map AccountWallPapersClass to AccountWallPapers.
	AsModified() (*AccountWallPapers, bool)
}

// AsModified tries to map AccountWallPapersNotModified to AccountWallPapers.
func (w *AccountWallPapersNotModified) AsModified() (*AccountWallPapers, bool) {
	return nil, false
}

// AsModified tries to map AccountWallPapers to AccountWallPapers.
func (w *AccountWallPapers) AsModified() (*AccountWallPapers, bool) {
	return w, true
}

// DecodeAccountWallPapers implements binary de-serialization for AccountWallPapersClass.
func DecodeAccountWallPapers(buf *bin.Buffer) (AccountWallPapersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AccountWallPapersNotModifiedTypeID:
		// Decoding account.wallPapersNotModified#1c199183.
		v := AccountWallPapersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountWallPapersClass: %w", err)
		}
		return &v, nil
	case AccountWallPapersTypeID:
		// Decoding account.wallPapers#702b65a9.
		v := AccountWallPapers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountWallPapersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AccountWallPapersClass: %w", bin.NewUnexpectedID(id))
	}
}

// AccountWallPapers boxes the AccountWallPapersClass providing a helper.
type AccountWallPapersBox struct {
	WallPapers AccountWallPapersClass
}

// Decode implements bin.Decoder for AccountWallPapersBox.
func (b *AccountWallPapersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AccountWallPapersBox to nil")
	}
	v, err := DecodeAccountWallPapers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.WallPapers = v
	return nil
}

// Encode implements bin.Encode for AccountWallPapersBox.
func (b *AccountWallPapersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.WallPapers == nil {
		return fmt.Errorf("unable to encode AccountWallPapersClass as nil")
	}
	return b.WallPapers.Encode(buf)
}

// AccountWallPapersClassArray is adapter for slice of AccountWallPapersClass.
type AccountWallPapersClassArray []AccountWallPapersClass

// Sort sorts slice of AccountWallPapersClass.
func (s AccountWallPapersClassArray) Sort(less func(a, b AccountWallPapersClass) bool) AccountWallPapersClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AccountWallPapersClass.
func (s AccountWallPapersClassArray) SortStable(less func(a, b AccountWallPapersClass) bool) AccountWallPapersClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AccountWallPapersClass.
func (s AccountWallPapersClassArray) Retain(keep func(x AccountWallPapersClass) bool) AccountWallPapersClassArray {
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
func (s AccountWallPapersClassArray) First() (v AccountWallPapersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AccountWallPapersClassArray) Last() (v AccountWallPapersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AccountWallPapersClassArray) PopFirst() (v AccountWallPapersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AccountWallPapersClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AccountWallPapersClassArray) Pop() (v AccountWallPapersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsAccountWallPapers returns copy with only AccountWallPapers constructors.
func (s AccountWallPapersClassArray) AsAccountWallPapers() (to AccountWallPapersArray) {
	for _, elem := range s {
		value, ok := elem.(*AccountWallPapers)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s AccountWallPapersClassArray) AppendOnlyModified(to []*AccountWallPapers) []*AccountWallPapers {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s AccountWallPapersClassArray) AsModified() (to []*AccountWallPapers) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s AccountWallPapersClassArray) FirstAsModified() (v *AccountWallPapers, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s AccountWallPapersClassArray) LastAsModified() (v *AccountWallPapers, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *AccountWallPapersClassArray) PopFirstAsModified() (v *AccountWallPapers, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *AccountWallPapersClassArray) PopAsModified() (v *AccountWallPapers, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// AccountWallPapersArray is adapter for slice of AccountWallPapers.
type AccountWallPapersArray []AccountWallPapers

// Sort sorts slice of AccountWallPapers.
func (s AccountWallPapersArray) Sort(less func(a, b AccountWallPapers) bool) AccountWallPapersArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AccountWallPapers.
func (s AccountWallPapersArray) SortStable(less func(a, b AccountWallPapers) bool) AccountWallPapersArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AccountWallPapers.
func (s AccountWallPapersArray) Retain(keep func(x AccountWallPapers) bool) AccountWallPapersArray {
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
func (s AccountWallPapersArray) First() (v AccountWallPapers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AccountWallPapersArray) Last() (v AccountWallPapers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AccountWallPapersArray) PopFirst() (v AccountWallPapers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AccountWallPapers
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AccountWallPapersArray) Pop() (v AccountWallPapers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
