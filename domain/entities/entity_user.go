package entities

import "atlas/util"

// UserInfo represents the information of a user who owns festival cards.
type UserInfo struct {
	// ID is the unique identifier of the user
	ID int64 `json:"id"`

	// Username is the user's name or handle
	Username string `json:"username"`

	// WalletAddress is the blockchain wallet address of the user
	WalletAddress string `json:"wallet_address"`

	// StatusCode indicates the current status of the user (e.g., active, banned)
	StatusCode int64 `json:"status_code"`

	// CreatedAt is the timestamp when the user was created
	CreatedAt util.DateTime `json:"created_at"`

	// ModifiedAt is the timestamp when the user was last modified
	ModifiedAt util.DateTime `json:"modified_at"`
}
