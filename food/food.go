package food

import "time"

// Represents an ingredient
type Ingredient struct {
	Name     string    `json:"name"`
	Metadata string    `json:"metadata,omitempty"`
	Expiry   time.Time `json:"expiry,omitempty"`
}
