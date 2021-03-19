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

// ChatParticipant represents TL type `chatParticipant#c8d7493e`.
// Group member.
//
// See https://core.telegram.org/constructor/chatParticipant for reference.
type ChatParticipant struct {
	// Member user ID
	UserID int
	// ID of the user that added the member to the group
	InviterID int
	// Date added to the group
	Date int
}

// ChatParticipantTypeID is TL type id of ChatParticipant.
const ChatParticipantTypeID = 0xc8d7493e

func (c *ChatParticipant) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.InviterID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipant) String() string {
	if c == nil {
		return "ChatParticipant(nil)"
	}
	type Alias ChatParticipant
	return fmt.Sprintf("ChatParticipant%+v", Alias(*c))
}

// FillFrom fills ChatParticipant from given interface.
func (c *ChatParticipant) FillFrom(from interface {
	GetUserID() (value int)
	GetInviterID() (value int)
	GetDate() (value int)
}) {
	c.UserID = from.GetUserID()
	c.InviterID = from.GetInviterID()
	c.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatParticipant) TypeID() uint32 {
	return ChatParticipantTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatParticipant) TypeName() string {
	return "chatParticipant"
}

// TypeInfo returns info about TL type.
func (c *ChatParticipant) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatParticipant",
		ID:   ChatParticipantTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "InviterID",
			SchemaName: "inviter_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatParticipant) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipant#c8d7493e as nil")
	}
	b.PutID(ChatParticipantTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatParticipant) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipant#c8d7493e as nil")
	}
	b.PutInt(c.UserID)
	b.PutInt(c.InviterID)
	b.PutInt(c.Date)
	return nil
}

// GetUserID returns value of UserID field.
func (c *ChatParticipant) GetUserID() (value int) {
	return c.UserID
}

// GetInviterID returns value of InviterID field.
func (c *ChatParticipant) GetInviterID() (value int) {
	return c.InviterID
}

// GetDate returns value of Date field.
func (c *ChatParticipant) GetDate() (value int) {
	return c.Date
}

// Decode implements bin.Decoder.
func (c *ChatParticipant) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipant#c8d7493e to nil")
	}
	if err := b.ConsumeID(ChatParticipantTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipant#c8d7493e: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatParticipant) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipant#c8d7493e to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipant#c8d7493e: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipant#c8d7493e: field inviter_id: %w", err)
		}
		c.InviterID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipant#c8d7493e: field date: %w", err)
		}
		c.Date = value
	}
	return nil
}

// construct implements constructor of ChatParticipantClass.
func (c ChatParticipant) construct() ChatParticipantClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipant.
var (
	_ bin.Encoder     = &ChatParticipant{}
	_ bin.Decoder     = &ChatParticipant{}
	_ bin.BareEncoder = &ChatParticipant{}
	_ bin.BareDecoder = &ChatParticipant{}

	_ ChatParticipantClass = &ChatParticipant{}
)

// ChatParticipantCreator represents TL type `chatParticipantCreator#da13538a`.
// Represents the creator of the group
//
// See https://core.telegram.org/constructor/chatParticipantCreator for reference.
type ChatParticipantCreator struct {
	// ID of the user that created the group
	UserID int
}

// ChatParticipantCreatorTypeID is TL type id of ChatParticipantCreator.
const ChatParticipantCreatorTypeID = 0xda13538a

func (c *ChatParticipantCreator) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipantCreator) String() string {
	if c == nil {
		return "ChatParticipantCreator(nil)"
	}
	type Alias ChatParticipantCreator
	return fmt.Sprintf("ChatParticipantCreator%+v", Alias(*c))
}

// FillFrom fills ChatParticipantCreator from given interface.
func (c *ChatParticipantCreator) FillFrom(from interface {
	GetUserID() (value int)
}) {
	c.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatParticipantCreator) TypeID() uint32 {
	return ChatParticipantCreatorTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatParticipantCreator) TypeName() string {
	return "chatParticipantCreator"
}

// TypeInfo returns info about TL type.
func (c *ChatParticipantCreator) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatParticipantCreator",
		ID:   ChatParticipantCreatorTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatParticipantCreator) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantCreator#da13538a as nil")
	}
	b.PutID(ChatParticipantCreatorTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatParticipantCreator) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantCreator#da13538a as nil")
	}
	b.PutInt(c.UserID)
	return nil
}

// GetUserID returns value of UserID field.
func (c *ChatParticipantCreator) GetUserID() (value int) {
	return c.UserID
}

// Decode implements bin.Decoder.
func (c *ChatParticipantCreator) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantCreator#da13538a to nil")
	}
	if err := b.ConsumeID(ChatParticipantCreatorTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipantCreator#da13538a: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatParticipantCreator) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantCreator#da13538a to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantCreator#da13538a: field user_id: %w", err)
		}
		c.UserID = value
	}
	return nil
}

// construct implements constructor of ChatParticipantClass.
func (c ChatParticipantCreator) construct() ChatParticipantClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipantCreator.
var (
	_ bin.Encoder     = &ChatParticipantCreator{}
	_ bin.Decoder     = &ChatParticipantCreator{}
	_ bin.BareEncoder = &ChatParticipantCreator{}
	_ bin.BareDecoder = &ChatParticipantCreator{}

	_ ChatParticipantClass = &ChatParticipantCreator{}
)

// ChatParticipantAdmin represents TL type `chatParticipantAdmin#e2d6e436`.
// Chat admin
//
// See https://core.telegram.org/constructor/chatParticipantAdmin for reference.
type ChatParticipantAdmin struct {
	// ID of a group member that is admin
	UserID int
	// ID of the user that added the member to the group
	InviterID int
	// Date when the user was added
	Date int
}

// ChatParticipantAdminTypeID is TL type id of ChatParticipantAdmin.
const ChatParticipantAdminTypeID = 0xe2d6e436

func (c *ChatParticipantAdmin) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.InviterID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipantAdmin) String() string {
	if c == nil {
		return "ChatParticipantAdmin(nil)"
	}
	type Alias ChatParticipantAdmin
	return fmt.Sprintf("ChatParticipantAdmin%+v", Alias(*c))
}

// FillFrom fills ChatParticipantAdmin from given interface.
func (c *ChatParticipantAdmin) FillFrom(from interface {
	GetUserID() (value int)
	GetInviterID() (value int)
	GetDate() (value int)
}) {
	c.UserID = from.GetUserID()
	c.InviterID = from.GetInviterID()
	c.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatParticipantAdmin) TypeID() uint32 {
	return ChatParticipantAdminTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatParticipantAdmin) TypeName() string {
	return "chatParticipantAdmin"
}

// TypeInfo returns info about TL type.
func (c *ChatParticipantAdmin) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatParticipantAdmin",
		ID:   ChatParticipantAdminTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "InviterID",
			SchemaName: "inviter_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatParticipantAdmin) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantAdmin#e2d6e436 as nil")
	}
	b.PutID(ChatParticipantAdminTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatParticipantAdmin) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantAdmin#e2d6e436 as nil")
	}
	b.PutInt(c.UserID)
	b.PutInt(c.InviterID)
	b.PutInt(c.Date)
	return nil
}

// GetUserID returns value of UserID field.
func (c *ChatParticipantAdmin) GetUserID() (value int) {
	return c.UserID
}

// GetInviterID returns value of InviterID field.
func (c *ChatParticipantAdmin) GetInviterID() (value int) {
	return c.InviterID
}

// GetDate returns value of Date field.
func (c *ChatParticipantAdmin) GetDate() (value int) {
	return c.Date
}

// Decode implements bin.Decoder.
func (c *ChatParticipantAdmin) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantAdmin#e2d6e436 to nil")
	}
	if err := b.ConsumeID(ChatParticipantAdminTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipantAdmin#e2d6e436: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatParticipantAdmin) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantAdmin#e2d6e436 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantAdmin#e2d6e436: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantAdmin#e2d6e436: field inviter_id: %w", err)
		}
		c.InviterID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantAdmin#e2d6e436: field date: %w", err)
		}
		c.Date = value
	}
	return nil
}

// construct implements constructor of ChatParticipantClass.
func (c ChatParticipantAdmin) construct() ChatParticipantClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipantAdmin.
var (
	_ bin.Encoder     = &ChatParticipantAdmin{}
	_ bin.Decoder     = &ChatParticipantAdmin{}
	_ bin.BareEncoder = &ChatParticipantAdmin{}
	_ bin.BareDecoder = &ChatParticipantAdmin{}

	_ ChatParticipantClass = &ChatParticipantAdmin{}
)

// ChatParticipantClass represents ChatParticipant generic type.
//
// See https://core.telegram.org/type/ChatParticipant for reference.
//
// Example:
//  g, err := tg.DecodeChatParticipant(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.ChatParticipant: // chatParticipant#c8d7493e
//  case *tg.ChatParticipantCreator: // chatParticipantCreator#da13538a
//  case *tg.ChatParticipantAdmin: // chatParticipantAdmin#e2d6e436
//  default: panic(v)
//  }
type ChatParticipantClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatParticipantClass

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

	// Member user ID
	GetUserID() (value int)
}

// DecodeChatParticipant implements binary de-serialization for ChatParticipantClass.
func DecodeChatParticipant(buf *bin.Buffer) (ChatParticipantClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatParticipantTypeID:
		// Decoding chatParticipant#c8d7493e.
		v := ChatParticipant{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", err)
		}
		return &v, nil
	case ChatParticipantCreatorTypeID:
		// Decoding chatParticipantCreator#da13538a.
		v := ChatParticipantCreator{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", err)
		}
		return &v, nil
	case ChatParticipantAdminTypeID:
		// Decoding chatParticipantAdmin#e2d6e436.
		v := ChatParticipantAdmin{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatParticipant boxes the ChatParticipantClass providing a helper.
type ChatParticipantBox struct {
	ChatParticipant ChatParticipantClass
}

// Decode implements bin.Decoder for ChatParticipantBox.
func (b *ChatParticipantBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatParticipantBox to nil")
	}
	v, err := DecodeChatParticipant(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatParticipant = v
	return nil
}

// Encode implements bin.Encode for ChatParticipantBox.
func (b *ChatParticipantBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatParticipant == nil {
		return fmt.Errorf("unable to encode ChatParticipantClass as nil")
	}
	return b.ChatParticipant.Encode(buf)
}

// ChatParticipantClassArray is adapter for slice of ChatParticipantClass.
type ChatParticipantClassArray []ChatParticipantClass

// Sort sorts slice of ChatParticipantClass.
func (s ChatParticipantClassArray) Sort(less func(a, b ChatParticipantClass) bool) ChatParticipantClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipantClass.
func (s ChatParticipantClassArray) SortStable(less func(a, b ChatParticipantClass) bool) ChatParticipantClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipantClass.
func (s ChatParticipantClassArray) Retain(keep func(x ChatParticipantClass) bool) ChatParticipantClassArray {
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
func (s ChatParticipantClassArray) First() (v ChatParticipantClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantClassArray) Last() (v ChatParticipantClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantClassArray) PopFirst() (v ChatParticipantClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipantClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantClassArray) Pop() (v ChatParticipantClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsChatParticipant returns copy with only ChatParticipant constructors.
func (s ChatParticipantClassArray) AsChatParticipant() (to ChatParticipantArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatParticipant)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsChatParticipantCreator returns copy with only ChatParticipantCreator constructors.
func (s ChatParticipantClassArray) AsChatParticipantCreator() (to ChatParticipantCreatorArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatParticipantCreator)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsChatParticipantAdmin returns copy with only ChatParticipantAdmin constructors.
func (s ChatParticipantClassArray) AsChatParticipantAdmin() (to ChatParticipantAdminArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatParticipantAdmin)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ChatParticipantArray is adapter for slice of ChatParticipant.
type ChatParticipantArray []ChatParticipant

// Sort sorts slice of ChatParticipant.
func (s ChatParticipantArray) Sort(less func(a, b ChatParticipant) bool) ChatParticipantArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipant.
func (s ChatParticipantArray) SortStable(less func(a, b ChatParticipant) bool) ChatParticipantArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipant.
func (s ChatParticipantArray) Retain(keep func(x ChatParticipant) bool) ChatParticipantArray {
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
func (s ChatParticipantArray) First() (v ChatParticipant, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantArray) Last() (v ChatParticipant, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantArray) PopFirst() (v ChatParticipant, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipant
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantArray) Pop() (v ChatParticipant, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of ChatParticipant by Date.
func (s ChatParticipantArray) SortByDate() ChatParticipantArray {
	return s.Sort(func(a, b ChatParticipant) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of ChatParticipant by Date.
func (s ChatParticipantArray) SortStableByDate() ChatParticipantArray {
	return s.SortStable(func(a, b ChatParticipant) bool {
		return a.GetDate() < b.GetDate()
	})
}

// ChatParticipantCreatorArray is adapter for slice of ChatParticipantCreator.
type ChatParticipantCreatorArray []ChatParticipantCreator

// Sort sorts slice of ChatParticipantCreator.
func (s ChatParticipantCreatorArray) Sort(less func(a, b ChatParticipantCreator) bool) ChatParticipantCreatorArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipantCreator.
func (s ChatParticipantCreatorArray) SortStable(less func(a, b ChatParticipantCreator) bool) ChatParticipantCreatorArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipantCreator.
func (s ChatParticipantCreatorArray) Retain(keep func(x ChatParticipantCreator) bool) ChatParticipantCreatorArray {
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
func (s ChatParticipantCreatorArray) First() (v ChatParticipantCreator, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantCreatorArray) Last() (v ChatParticipantCreator, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantCreatorArray) PopFirst() (v ChatParticipantCreator, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipantCreator
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantCreatorArray) Pop() (v ChatParticipantCreator, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ChatParticipantAdminArray is adapter for slice of ChatParticipantAdmin.
type ChatParticipantAdminArray []ChatParticipantAdmin

// Sort sorts slice of ChatParticipantAdmin.
func (s ChatParticipantAdminArray) Sort(less func(a, b ChatParticipantAdmin) bool) ChatParticipantAdminArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatParticipantAdmin.
func (s ChatParticipantAdminArray) SortStable(less func(a, b ChatParticipantAdmin) bool) ChatParticipantAdminArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatParticipantAdmin.
func (s ChatParticipantAdminArray) Retain(keep func(x ChatParticipantAdmin) bool) ChatParticipantAdminArray {
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
func (s ChatParticipantAdminArray) First() (v ChatParticipantAdmin, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatParticipantAdminArray) Last() (v ChatParticipantAdmin, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatParticipantAdminArray) PopFirst() (v ChatParticipantAdmin, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatParticipantAdmin
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatParticipantAdminArray) Pop() (v ChatParticipantAdmin, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of ChatParticipantAdmin by Date.
func (s ChatParticipantAdminArray) SortByDate() ChatParticipantAdminArray {
	return s.Sort(func(a, b ChatParticipantAdmin) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of ChatParticipantAdmin by Date.
func (s ChatParticipantAdminArray) SortStableByDate() ChatParticipantAdminArray {
	return s.SortStable(func(a, b ChatParticipantAdmin) bool {
		return a.GetDate() < b.GetDate()
	})
}
