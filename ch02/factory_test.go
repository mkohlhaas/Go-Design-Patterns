package creational

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment := GetPaymentMethod(Cash)
	msg := payment.Pay(10)

	if !strings.Contains(msg, "payed using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment := GetPaymentMethod(DebitCard)
	msg := payment.Pay(22)

	if !strings.Contains(msg, "payed using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	defer func() { _ = recover() }()

  // invalid payment method should panic
	GetPaymentMethod(3)

	// Never reaches here if `OtherFunctionThatPanics` panics.
	t.Errorf("did not panic")
}
