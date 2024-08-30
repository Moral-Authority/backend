package handlers

import (
	"context"

	"github.com/Moral-Authority/backend/graph/model"
)

var DefaultColors = map[string]string{
	"BLACK":      "#242424",
	"GREY":       "#cccfd8",
	"BROWN":      "#8b6b61",
	"OLIVE":      "#a3a47f",
	"RED":        "#b34030",
	"PINK":       "#e5b2b8",
	"ORANGE":     "#f5a941",
	"YELLOW":     "#f9e29c",
	"GREEN":      "#6db048",
	"TEAL":       "#4b9395",
	"BLUE":       "#324a77",
	"PURPLE":     "#7d4ac4",
	"GOLD":       "#d5b05d",
	"MULTICOLOR": "#7b4934",
	"WHITE":      "#ffffff",
}

// GetSubDepartmentFilters is the resolver for the getSubDepartmentFilters field.
func (s ProductService) GetSubDepartmentFiltersHandler(ctx context.Context, input string) (*model.Filters, error) {

	colors := FormatStringListToColorStructList(DefaultColors)

	result := model.Filters{
		Colors:                colors,
		Sizes:                 []*string{ConvertString("SM"), ConvertString("M"), ConvertString("L"), ConvertString("XL"), ConvertString("XXL")},
		Companies:             []*string{ConvertString("Adidas"), ConvertString("Nike"), ConvertString("Puma")},
		CompanyCertifications: []*string{ConvertString("BCorp"), ConvertString("Women Owned"), ConvertString("Black Owned"), ConvertString("Indigenous Owned"), ConvertString("Asian Pacific Islander"), ConvertString("LGBTQ+ Owned"), ConvertString("Political Donations?")},
		ProductCertifications: []*string{ConvertString("Fair Trade"), ConvertString("Leaping Bunny"), ConvertString("Plastic - Free"), ConvertString("Vegan"), ConvertString("Organic"), ConvertString("Made in America")},
	}

	return &result, nil
}
