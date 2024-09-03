package handlers

import "strings"

const DefaultDescription = "No description available -- Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,"
const DefaultPrice = "0.00"
const DefaultLink = "https://www.google.com"
const DefaultCity = "Unknown City"
const DefaultState = "Unknown State"
const DefaultCountry = "Unknown Country"
const DefaultZipcode = "00000"


type ProductDepartment int

const (
	HomeGardenProductDepartment ProductDepartment = iota
	HealthBathBeautyProductDepartment
	ClothingAccessoriesProductDepartment
	ToysKidsBabiesProductDepartment
)

// ToString converts ProductDepartment to its string representation
func (p ProductDepartment) ToString() string {
	switch p {
	case HomeGardenProductDepartment:
		return "Home & Garden"
	case HealthBathBeautyProductDepartment:
		return "Clothing & Accessories"
	case ClothingAccessoriesProductDepartment:
		return "Health, Bath & Beauty"
	case ToysKidsBabiesProductDepartment:
		return "Toys, Kids & Babies"
	default:
		return "Unknown"
	}
}


// ToInt converts ProductDepartment to int
func (p ProductDepartment) ToInt() int {
	return int(p)
}

// IsValidProductDepartment checks if the department int is valid and returns the corresponding ProductDepartment type if true
func IsValidProductDepartment(department int) (ProductDepartment, bool) {
    switch ProductDepartment(department) {
    case HomeGardenProductDepartment, HealthBathBeautyProductDepartment, ClothingAccessoriesProductDepartment, ToysKidsBabiesProductDepartment:
        return ProductDepartment(department), true
    default:
        return 0, false
    }
}

func IsStringValidProductDepartment(department string) (ProductDepartment, bool) {
    // Normalize the string
    normalizedDepartment := strings.ReplaceAll(department, "&", "and")
    normalizedDepartment = strings.TrimSpace(normalizedDepartment)

    switch normalizedDepartment {
    case "Home and Garden":
        return HomeGardenProductDepartment, true
    case "Clothing and Accessories":
        return ClothingAccessoriesProductDepartment, true
    case "Health, Bath and Beauty":
        return HealthBathBeautyProductDepartment, true
    case "Toys, Kids and Babies":
        return ToysKidsBabiesProductDepartment, true
    default:
        return 0, false
    }
}


// // IsStringValidProductDepartment checks if the department string is valid and returns the corresponding ProductDepartment type if true
// func IsStringValidProductDepartment(department string) (ProductDepartment, bool) {
//     switch department {
//     case "Home & Garden":
//         return HomeGardenProductDepartment, true
//     case "Clothing & Accessories":
//         return BathBeautyProductDepartment, true
//     case "Health, Bath & Beauty":
//         return ClothingAccessoriesProductDepartment, true
//     case "Toys, Kids & Babies":
//         return ToysKidsBabiesProductDepartment, true
//     default:
//         return 0, false
//     }
// }

