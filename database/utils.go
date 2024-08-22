package database

import (
	"fmt"
	"strconv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func StringToUint(s string) (uint, error) {
	ans, err := strconv.ParseUint(s, 10, 32)
	return uint(ans), err
}


// ApplyFilters applies the provided filters to the GORM query builder.
func ApplyFilters(query *gorm.DB, filters map[string]interface{}) *gorm.DB {
    if query == nil {
        logrus.Error("Query is nil")
        return query
    }

    for key, value := range filters {
        if value != nil {
            switch v := value.(type) {
            case string:
                if v != "" {
                    query = query.Where(fmt.Sprintf("%s ILIKE ?", key), "%"+v+"%")
                }
            case bool:
                query = query.Where(fmt.Sprintf("%s = ?", key), v)
            case *string:
                if v != nil && *v != "" {
                    query = query.Where(fmt.Sprintf("%s ILIKE ?", key), "%"+*v+"%")
                }
            case *bool:
                if v != nil {
                    query = query.Where(fmt.Sprintf("%s = ?", key), *v)
                }
            default:
                logrus.Warnf("Unsupported filter type for key %s: %T", key, value)
            }
        }
    }
    return query
}
