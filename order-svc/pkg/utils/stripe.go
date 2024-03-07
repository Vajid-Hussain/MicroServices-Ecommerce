package utils_order

import (
	"fmt"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	config_order_svc "github.com/vajid-hussain/mobile-mart-microservice-order/pkg/config"
)

func StripePaymentInitiation(data int, config *config_order_svc.Config) (string, error) {
	stripe.Key = config.StripeSecretKey

	// product_params := &stripe.ProductParams{
	// 	Name:        stripe.String("first try"),
	// 	Description: stripe.String("try after cli done"),
	// }

	// starter_product, _ := product.New(product_params)

	// price_params := &stripe.PriceParams{
	// 	Currency:   stripe.String(string(stripe.CurrencyINR)),
	// 	Product:    stripe.String(string(starter_product.ID)),
	// 	UnitAmount: stripe.Int64(786),
	// }

	// starter_price, _ := price.New(price_params)

	params := &stripe.PaymentIntentParams{
		Amount:      stripe.Int64(int64(data)),
		Currency:    stripe.String(string(stripe.CurrencyINR)),
		Description: stripe.String("Description of the transaction"),
		Shipping: &stripe.ShippingDetailsParams{
			Name: stripe.String("Jenny Rosen"),
			Address: &stripe.AddressParams{
				Line1:      stripe.String("510 Townsend St"),
				PostalCode: stripe.String("98140"),
				City:       stripe.String("San Francisco"),
				State:      stripe.String("CA"),
				Country:    stripe.String("US"),
			},
		},
	}

	intent, err := paymentintent.New(params)
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	fmt.Println("***&*", intent.ID, intent.ClientSecret)

	return intent.ClientSecret, nil
}
