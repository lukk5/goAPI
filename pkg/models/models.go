package models

// Item represents an item that can be retrieved
// swagger:model Item
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
