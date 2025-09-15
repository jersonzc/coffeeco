package store

import (
	"coffeeco/internal"
	"github.com/google/uuid"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []internal.Product
}
