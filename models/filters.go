package models

type Filters struct {
	Price                 *PriceRange
	Rating                int
	Companies             []string
	CompanyCertifications []string
	ProductCertifications []string
}

type PriceRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
