package purchase

import (
	"coffeeco/internal"
	"coffeeco/internal/payment"
	"coffeeco/internal/store"
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"time"
)

type Purchase struct {
	id                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []internal.Product
	total              money.Money
	PaymentMeans       payment.Means
	timeOfPurchase     time.Time
}
