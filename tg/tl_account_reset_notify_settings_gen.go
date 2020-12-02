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

// AccountResetNotifySettingsRequest represents TL type `account.resetNotifySettings#db7e1747`.
type AccountResetNotifySettingsRequest struct {
}

// AccountResetNotifySettingsRequestTypeID is TL type id of AccountResetNotifySettingsRequest.
const AccountResetNotifySettingsRequestTypeID = 0xdb7e1747

// Encode implements bin.Encoder.
func (r *AccountResetNotifySettingsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetNotifySettings#db7e1747 as nil")
	}
	b.PutID(AccountResetNotifySettingsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetNotifySettingsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetNotifySettings#db7e1747 to nil")
	}
	if err := b.ConsumeID(AccountResetNotifySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetNotifySettings#db7e1747: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountResetNotifySettingsRequest.
var (
	_ bin.Encoder = &AccountResetNotifySettingsRequest{}
	_ bin.Decoder = &AccountResetNotifySettingsRequest{}
)

// AccountResetNotifySettings invokes method account.resetNotifySettings#db7e1747 returning error if any.
func (c *Client) AccountResetNotifySettings(ctx context.Context, request *AccountResetNotifySettingsRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}