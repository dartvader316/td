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

// AuthExportAuthorizationRequest represents TL type `auth.exportAuthorization#e5bfffcd`.
// Returns data for copying authorization to another data-centre.
//
// See https://core.telegram.org/method/auth.exportAuthorization for reference.
type AuthExportAuthorizationRequest struct {
	// Number of a target data-centre
	DCID int
}

// AuthExportAuthorizationRequestTypeID is TL type id of AuthExportAuthorizationRequest.
const AuthExportAuthorizationRequestTypeID = 0xe5bfffcd

func (e *AuthExportAuthorizationRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.DCID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *AuthExportAuthorizationRequest) String() string {
	if e == nil {
		return "AuthExportAuthorizationRequest(nil)"
	}
	type Alias AuthExportAuthorizationRequest
	return fmt.Sprintf("AuthExportAuthorizationRequest%+v", Alias(*e))
}

// FillFrom fills AuthExportAuthorizationRequest from given interface.
func (e *AuthExportAuthorizationRequest) FillFrom(from interface {
	GetDCID() (value int)
}) {
	e.DCID = from.GetDCID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthExportAuthorizationRequest) TypeID() uint32 {
	return AuthExportAuthorizationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthExportAuthorizationRequest) TypeName() string {
	return "auth.exportAuthorization"
}

// TypeInfo returns info about TL type.
func (e *AuthExportAuthorizationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.exportAuthorization",
		ID:   AuthExportAuthorizationRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *AuthExportAuthorizationRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode auth.exportAuthorization#e5bfffcd as nil")
	}
	b.PutID(AuthExportAuthorizationRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *AuthExportAuthorizationRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode auth.exportAuthorization#e5bfffcd as nil")
	}
	b.PutInt(e.DCID)
	return nil
}

// GetDCID returns value of DCID field.
func (e *AuthExportAuthorizationRequest) GetDCID() (value int) {
	return e.DCID
}

// Decode implements bin.Decoder.
func (e *AuthExportAuthorizationRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode auth.exportAuthorization#e5bfffcd to nil")
	}
	if err := b.ConsumeID(AuthExportAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.exportAuthorization#e5bfffcd: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *AuthExportAuthorizationRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode auth.exportAuthorization#e5bfffcd to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportAuthorization#e5bfffcd: field dc_id: %w", err)
		}
		e.DCID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthExportAuthorizationRequest.
var (
	_ bin.Encoder     = &AuthExportAuthorizationRequest{}
	_ bin.Decoder     = &AuthExportAuthorizationRequest{}
	_ bin.BareEncoder = &AuthExportAuthorizationRequest{}
	_ bin.BareDecoder = &AuthExportAuthorizationRequest{}
)

// AuthExportAuthorization invokes method auth.exportAuthorization#e5bfffcd returning error if any.
// Returns data for copying authorization to another data-centre.
//
// Possible errors:
//  400 DC_ID_INVALID: The provided DC ID is invalid
//
// See https://core.telegram.org/method/auth.exportAuthorization for reference.
// Can be used by bots.
func (c *Client) AuthExportAuthorization(ctx context.Context, dcid int) (*AuthExportedAuthorization, error) {
	var result AuthExportedAuthorization

	request := &AuthExportAuthorizationRequest{
		DCID: dcid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
