package database

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func StringToUint(s string) (uint, error) {
	ans, err := strconv.ParseUint(s, 10, 32)
	return uint(ans), err
}


// ApplyFilters applies the provided filters to the GORM query builder.
func ApplyFilters(query *gorm.DB, filters map[string]interface{}) *gorm.DB {
	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(strings.ToLower(key)+" ILIKE ?", "%"+value.(string)+"%")
		}
	}
	return query
}

