package entities

import (
	"atlas/util"
	"github.com/shopspring/decimal"
)

// FestivalCards represents a festival card owned by a user.
// It includes references to the user, crypto information, pricing, and timestamps.
type FestivalCards struct {
	// ID is the unique identifier of the festival card
	ID int64 `json:"id"`

	// UserInfo contains the information of the card owner
	UserInfo UserInfo `json:"user"`

	// Balance is the current balance on the card, using decimal.Decimal for precise monetary values
	Balance decimal.Decimal `json:"balance"`

	// CryptoType represents the type of cryptocurrency associated with the card
	CryptoType CryptoType `json:"crypto_type"`

	// CryptoPrice is the price of the cryptocurrency at the time of card creation or sale
	CryptoPrice decimal.Decimal `json:"crypto_price"`

	// SoldAt is the timestamp when the card was sold (nullable)
	SoldAt *util.DateTime `json:"sold_at"`

	// StatusCode represents the current status of the card (e.g., active, sold)
	StatusCode int64 `json:"status_code"`

	// CreatedAt is the timestamp when the card was created
	CreatedAt util.DateTime `json:"created_at"`

	// ModifiedAt is the timestamp when the card was last modified
	ModifiedAt util.DateTime `json:"modified_at"`
}
