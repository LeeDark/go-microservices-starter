package data

import "testing"

func TestProductValidate(t *testing.T) {
	p := &Product{
		Name:  "Latte",
		Price: 35.00,
		SKU:   "abc-abcd-abcdef",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
