package creational

import (
	"fmt"
)

type Amount int
type PaymentType int

// PaymentMethod defines a way of paying in the shop.
type PaymentMethod interface {
	Pay(amount Amount) string
}

// Payment types
const (
	Cash PaymentType = iota
	DebitCard
)

// fmt.Stringer interface
func (pt PaymentType) String() (result string) {
	switch pt {
	case Cash:
		result = "cash"
	case DebitCard:
		result = "debit card"
	}
	return
}

// We use "new" operator to return the pointer but we could also use &Type{}.
func GetPaymentMethod(pt PaymentType) (result PaymentMethod) {
	switch pt {
	case Cash:
		result = new(CashPM)
	case DebitCard:
		result = new(NewDebitCardPM)
  default:
    panic("Invalid payment type.")
	}
	return
}

type CashPM struct{}
type DebitCardPM struct{}
type NewDebitCardPM struct{}

func (c *CashPM) Pay(amount Amount) string {
	return fmt.Sprintf("%d payed using cash.\n", amount)
}

func (c *DebitCardPM) Pay(amount Amount) string {
	return fmt.Sprintf("%d payed using debit card.\n", amount)
}

func (d *NewDebitCardPM) Pay(amount Amount) string {
	return fmt.Sprintf("%d payed using debit card (new).\n", amount)
}
