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
			case int:
				query = query.Where(fmt.Sprintf("%s = ?", key), v)
			case *int:
				if v != nil {
					query = query.Where(fmt.Sprintf("%s = ?", key), *v)
				}
			case float64:
				query = query.Where(fmt.Sprintf("%s = ?", key), v)
			case *float64:
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

func ApplyProductFilters(query *gorm.DB, filters map[string]interface{}, subDept int, productTable string) *gorm.DB {

    // Ensure the sub_department filter is applied early
    query = query.Where("sub_department = ?", subDept)

    if query == nil {
        logrus.Error("Query is nil")
        return query
    }

    logrus.Infof("Applying filters: %v", filters)

    // Track if the purchase_info table join is already applied
    purchaseInfoJoined := false

    // Always join the `companies` table if filtering by company name or certifications
    if filters["companyCertifications"] != nil || filters["companies"] != nil {
        logrus.Infof("Applying company certifications join")
        query = query.
            InnerJoins("INNER JOIN companies ON companies.id = home_garden_products.company_id").
            InnerJoins("INNER JOIN company_certifications ON company_certifications.company_id = companies.id").
            InnerJoins("INNER JOIN certifications AS company_certs ON company_certifications.certification_id = company_certs.id")
    }

    // Always join the `product_certifications` table if filtering by product certifications
    if filters["productCertifications"] != nil {
        logrus.Infof("Applying product certifications join")
        query = query.
            InnerJoins(fmt.Sprintf("INNER JOIN product_certifications ON product_certifications.product_id = %s.id", productTable)).
            InnerJoins("INNER JOIN certifications AS product_certs ON product_certifications.certification_id = product_certs.id")
    }

    // Now, apply the filters
    for key, value := range filters {
        if value != nil {
            switch key {
            case "priceMin", "priceMax":
                if !purchaseInfoJoined {
                    logrus.Infof("Joining purchase_infos table for price filter")
                    query = query.Joins(fmt.Sprintf("JOIN purchase_infos ON purchase_infos.product_id = %s.id", productTable))
                    purchaseInfoJoined = true
                }

                if key == "priceMin" {
                    if v, ok := value.(float64); ok {
                        query = query.Where("purchase_infos.price >= ?", v)
                    }
                }

                if key == "priceMax" {
                    if v, ok := value.(float64); ok {
                        query = query.Where("purchase_infos.price <= ?", v)
                    }
                }

            case "companyCertifications":
                if certs, ok := value.([]string); ok && len(certs) > 0 {
                    logrus.Infof("Applying company certifications filter: %v", certs)
                    query = query.Where("company_certs.name IN (?)", certs)

                    if err := query.Error; err != nil {
                        logrus.Error("Error applying company certifications filter: ", err)
                        return nil
                    }
                } else {
                    logrus.Warn("Invalid companyCertifications filter")
                }

            case "productCertifications":
                if certs, ok := value.([]string); ok && len(certs) > 0 {
                    logrus.Infof("Applying product certifications filter: %v", certs)
                    query = query.Where("product_certs.name IN (?)", certs)
                }

            case "companies":
                if companies, ok := value.([]string); ok && len(companies) > 0 {
                    logrus.Infof("Applying company name filter: %v", companies)
                    query = query.Where("companies.name IN (?)", companies)
                }

            default:
                logrus.Warnf("Unsupported filter type for key %s: %T", key, value)
            }
        }
    }

    logrus.Infof("Final SQL Query: %s", query.Statement.SQL.String()) // Ensure the actual query is logged
    return query
}
