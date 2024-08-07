
// product.go
package models

import (
    "encoding/json"
)

type SquareProduct struct {
    ID                  string   `json:"id"`
    Name                string   `json:"name"`
    Description         string   `json:"description"`
    PriceCents          int      `json:"price_cents"`
    Variants            []Variant `json:"variants"`
}

type Variant struct {
    Name        string `json:"name"`
    Quantity    int    `json:"quantity"`
    PriceCents  int    `json:"price_cents"`
    Sku         string `json:"sku"`
    Weight      float64 `json:"weight"`
}

func (p *SquareProduct) ToProduct() Product {
    return Product{
        ID:   p.ID,
        Name: p.Name,
        Description: p.Description,
        PriceCents:  p.PriceCents,
        Variants:   convertVariants(p.Variants),
    }
}

func convertVariants(variants []Variant) []VariantProduct {
    variantProducts := make([]VariantProduct, len(variants))
    for i, variant := range variants {
        variantProducts[i] = VariantProduct{
            Name:     variant.Name,
            Quantity: variant.Quantity,
            PriceCents: variant.PriceCents,
            Sku:      variant.Sku,
            Weight:   variant.Weight,
        }
    }
    return variantProducts
}

type Product struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Description string `json:"description"`
    PriceCents int    `json:"price_cents"`
    Variants  []VariantProduct
}

type VariantProduct struct {
    Name     string `json:"name"`
    Quantity int    `json:"quantity"`
    PriceCents int   `json:"price_cents"`
    Sku      string `json:"sku"`
    Weight   float64 `json:"weight"`
}