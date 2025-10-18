package entities

import "atlas/util"

// CryptoType represents a cryptocurrency type that can be associated with a festival card.
type CryptoType struct {
	// ID is the unique identifier of the cryptocurrency
	ID int64 `json:"id"`

	// Symbol is the short symbol of the cryptocurrency (e.g., BTC, ETH)
	Symbol string `json:"symbol"`

	// Name is the full name of the cryptocurrency
	Name string `json:"name"`

	// StatusCode indicates the current status of the crypto (e.g., active, inactive)
	StatusCode int64 `json:"status_code"`

	// CreatedAt is the timestamp when the crypto type was created
	CreatedAt util.DateTime `json:"created_at"`

	// ModifiedAt is the timestamp when the crypto type was last modified
	ModifiedAt util.DateTime `json:"modified_at"`
}
