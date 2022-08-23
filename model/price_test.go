package model

import (
	"fmt"
	"testing"
)

func TestGetPrices(t *testing.T) {
	p := GetPrices()

	fmt.Println(p)
}

func TestUpdatePrices(t *testing.T) {
	UpdatePrices()

	fmt.Println(GetPrices())
}
func TestCreatePrices(t *testing.T) {
	CreatePrices()

	fmt.Println(GetPrices())
}
