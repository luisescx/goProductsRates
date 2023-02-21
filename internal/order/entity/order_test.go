package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGiven_NewOrder_EmptyId_HasError(t *testing.T) {
	order := Order{}
	assert.EqualError(t, order.IsValid(), "invalid id")
}

func TestGiven_NewOrder_EmptyPrice_HasError(t *testing.T) {
	order := Order{ID: "123", Tax: 10.0}
	assert.EqualError(t, order.IsValid(), "invalid price")
}

func TestGiven_NewOrder_EmptyTax_HasError(t *testing.T) {
	order := Order{ID: "123", Price: 12.3 }
	assert.EqualError(t, order.IsValid(), "invalid tax")
}

func TestGiven_AllValidParams(t *testing.T) {
	order := Order{
		ID: "123",
		Price: 10.0,
		Tax: 2.0,
	}
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGiven_PriceAndTax_CallCalculateTax_ShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}