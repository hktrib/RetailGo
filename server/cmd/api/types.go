package server

import "github.com/hktrib/RetailGo/internal/ent"

type UpdatedFields struct {
	Name         bool
	Photo        bool
	Quantity     bool
	Price        bool
	CategoryName bool
}

type ItemChange struct {
	Item          ent.Item
	Mode          string
	UpdatedFields UpdatedFields
}
