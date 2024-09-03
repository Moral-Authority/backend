package handlers

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
	BathBeautyProductDepartment
	ClothingAccessoriesProductDepartment
	ToysKidsBabiesProductDepartment
)

// ToString converts ProductDepartment to its string representation
func (p ProductDepartment) ToString() string {
	switch p {
	case HomeGardenProductDepartment:
		return "HomeGardenProduct"
	case BathBeautyProductDepartment:
		return "BathBeautyProduct"
	case ClothingAccessoriesProductDepartment:
		return "ClothingAccessoriesProduct"
	case ToysKidsBabiesProductDepartment:
		return "ToysKidsBabiesProduct"
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
    case HomeGardenProductDepartment, BathBeautyProductDepartment, ClothingAccessoriesProductDepartment, ToysKidsBabiesProductDepartment:
        return ProductDepartment(department), true
    default:
        return 0, false
    }
}

// IsStringValidProductDepartment checks if the department string is valid and returns the corresponding ProductDepartment type if true
func IsStringValidProductDepartment(department string) (ProductDepartment, bool) {
    switch department {
    case "HomeGardenProduct":
        return HomeGardenProductDepartment, true
    case "BathBeautyProduct":
        return BathBeautyProductDepartment, true
    case "ClothingAccessoriesProduct":
        return ClothingAccessoriesProductDepartment, true
    case "ToysKidsBabiesProduct":
        return ToysKidsBabiesProductDepartment, true
    default:
        return 0, false
    }
}
