package catalogv1_test

import (
	"testing"

	catalogv1 "github.com/LeeDark/go-microservices-starter/grpc-playground/catalog/v1"
	catalogv2 "github.com/LeeDark/go-microservices-starter/grpc-playground/catalog/v2"
	"google.golang.org/protobuf/proto"
)

func TestProductRoundTrip(t *testing.T) {
	want := &catalogv1.Product{
		Id:         "p-1",
		Name:       "Keyboard",
		PriceCents: 1999,
	}

	data, err := proto.Marshal(want)
	if err != nil {
		t.Fatalf("marshal product: %v", err)
	}

	var got catalogv1.Product
	if err := proto.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal product: %v", err)
	}

	if got.GetId() != want.GetId() {
		t.Errorf("id = %q, want %q", got.GetId(), want.GetId())
	}
	if got.GetName() != want.GetName() {
		t.Errorf("name = %q, want %q", got.GetName(), want.GetName())
	}
	if got.GetPriceCents() != want.GetPriceCents() {
		t.Errorf("price_cents = %d, want %d", got.GetPriceCents(), want.GetPriceCents())
	}
}

func TestV1MessageCanBeReadAsV2(t *testing.T) {
	v1Product := &catalogv1.Product{Id: "p-1", Name: "Keyboard", PriceCents: 1999}

	data, err := proto.Marshal(v1Product)
	if err != nil {
		t.Fatalf("marshal v1 product: %v", err)
	}

	var v2Product catalogv2.Product
	if err := proto.Unmarshal(data, &v2Product); err != nil {
		t.Fatalf("unmarshal as v2: %v", err)
	}

	if v2Product.GetId() != v1Product.GetId() {
		t.Errorf("id = %q, want %q", v2Product.GetId(), v1Product.GetId())
	}
	if v2Product.GetName() != v1Product.GetName() {
		t.Errorf("name = %q, want %q", v2Product.GetName(), v1Product.GetName())
	}
	if v2Product.GetPriceCents() != v1Product.GetPriceCents() {
		t.Errorf("price_cents = %d, want %d", v2Product.GetPriceCents(), v1Product.GetPriceCents())
	}
	if v2Product.GetDescription() != "" {
		t.Errorf("description = %q, want empty default", v2Product.GetDescription())
	}
	if len(v2Product.GetTags()) != 0 {
		t.Errorf("tags = %v, want empty default", v2Product.GetTags())
	}
	if v2Product.GetStatus() != catalogv2.ProductStatus_PRODUCT_STATUS_UNSPECIFIED {
		t.Errorf("status = %v, want unspecified", v2Product.GetStatus())
	}
}

func TestV2MessageCanBeReadAsV1(t *testing.T) {
	v2Product := &catalogv2.Product{
		Id:          "p-2",
		Name:        "Mouse",
		PriceCents:  999,
		Description: "Wireless mouse",
		Tags:        []string{"wireless", "input"},
		Status:      catalogv2.ProductStatus_PRODUCT_STATUS_ACTIVE,
	}

	data, err := proto.Marshal(v2Product)
	if err != nil {
		t.Fatalf("marshal v2 product: %v", err)
	}

	var v1Product catalogv1.Product
	if err := proto.Unmarshal(data, &v1Product); err != nil {
		t.Fatalf("unmarshal as v1: %v", err)
	}

	if v1Product.GetId() != v2Product.GetId() {
		t.Errorf("id = %q, want %q", v1Product.GetId(), v2Product.GetId())
	}
	if v1Product.GetName() != v2Product.GetName() {
		t.Errorf("name = %q, want %q", v1Product.GetName(), v2Product.GetName())
	}
	if v1Product.GetPriceCents() != v2Product.GetPriceCents() {
		t.Errorf("price_cents = %d, want %d", v1Product.GetPriceCents(), v2Product.GetPriceCents())
	}
}
