package models

import (
	"github.com/Moral-Authority/backend/graph/model"
)

// CalculateOffset computes the offset based on the page number and items per page.
func CalculateOffset(p model.PaginationInput) int {
	if *p.Page <= 0 {
		page := 1
		p.Page = &page
	}
	return (*p.Page - 1) * *p.Items
}

// GetSortOptions returns the SQL ORDER BY clause based on SortBy and SortOrder.
// Defaults to "created_at DESC" if no sorting options are provided.
func GetSortOptions(s model.SortByInput) string {
	// Default sort by "created_at" in "DESC" order if SortBy is nil or empty
	if s.SortBy == nil || *s.SortBy == "" {
		defaultSortBy := "created_at"
		s.SortBy = &defaultSortBy
	}

	// Default to "DESC" order if SortOrder is nil or empty
	order := "DESC"
	if s.SortOrder != nil && *s.SortOrder == "ASC" {
		order = "ASC"
	}

	return *s.SortBy + " " + order
}