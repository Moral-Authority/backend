package handlers

import "strings"

const DefaultDescription = "No description available -- Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,"
const DefaultPrice = "0.00"
const DefaultLink = "https://www.google.com"
const DefaultCity = "Unknown City"
const DefaultState = "Unknown State"
const DefaultCountry = "Unknown Country"
const DefaultZipcode = "00000"

//  WIP
func IsStringValidProductSubDepartmentFORSEED(department ProductDepartment, subDepartment string)(int, bool) {
    switch department {
    case HomeGardenProductDepartment:
        subDept, isSubDept := IsStringValidHomeGardenSubDep(subDepartment)
        if !isSubDept {
            return 0, false
        }
        return subDept.ToInt(), true

    case HealthBathBeautyProductDepartment:
        subDept, isSubDept :=  IsStringValidHealthBathBeautySubDep(subDepartment)
        if !isSubDept {
            return 0, false
        }
        return subDept.ToInt(), true
    case ClothingAccessoriesProductDepartment:
        subDept, isSubDept := IsStringValidClothingAccessoriesSubDep(subDepartment)
        if !isSubDept {
            return 0, false
        }
        return subDept.ToInt(), true
    case ToysKidsBabiesProductDepartment:
        subDept, isSubDept := IsStringValidToysKidsBabiesSubDep(subDepartment)
        if !isSubDept {
            return 0, false
        }
        return subDept.ToInt(), true
    default:
        return 0, false
    }
}


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

func IsValidProductDepartment(department int) (ProductDepartment, bool) {
	switch ProductDepartment(department) {
	case HomeGardenProductDepartment, HealthBathBeautyProductDepartment, ClothingAccessoriesProductDepartment, ToysKidsBabiesProductDepartment:
		return ProductDepartment(department), true
	default:
		return 0, false
	}
}

func IsStringValidProductDepartment(department string) (ProductDepartment, bool) {

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

// ================== SUB DEPARTMENTS  ==================

type HomeGardenSubDep int

const (
	HomeGardenCleaningSupplies HomeGardenSubDep = iota
	HomeGardenKitchenDining
	HomeGardenFurnitureArt
	HomeGardenGarden
	HomeGardenBeddingLinens
	HomeGardenLightingLamps
	HomeGardenBathLaundry
)

// ToString converts ProductDepartment to its string representation
func (p HomeGardenSubDep) ToString() string {
	switch p {
	case HomeGardenCleaningSupplies:
		return "Cleaning Supplies"
	case HomeGardenKitchenDining:
		return "Kitchen & Dining"
	case HomeGardenFurnitureArt:
		return "Furniture & Art"
	case HomeGardenGarden:
		return "Garden"
	case HomeGardenBeddingLinens:
		return "Bedding & Linens"
	case HomeGardenLightingLamps:
		return "Lighting & Lamps"
	case HomeGardenBathLaundry:
		return "Bath & Laundry"
	default:
		return "Unknown"
	}
}

// ToInt converts HomeGardenSubDep to int
func (p HomeGardenSubDep) ToInt() int {
	return int(p)
}


func IsStringValidHomeGardenSubDep(subdep string) (HomeGardenSubDep, bool) {

	normalizedSubDep := strings.ReplaceAll(subdep, "&", "and")
	normalizedSubDep = strings.TrimSpace(normalizedSubDep)

	switch normalizedSubDep {
	case "Cleaning Supplies":
		return HomeGardenCleaningSupplies, true
	case "Kitchen and Dining":
		return HomeGardenKitchenDining, true
	case "Furniture and Art":
		return HomeGardenFurnitureArt, true
	case "Garden":
		return HomeGardenGarden, true
    case "Bedding and Linens":
        return HomeGardenBeddingLinens, true
    case "Lighting and Lamps":
        return HomeGardenLightingLamps, true
    case "Bath and Laundry":
        return HomeGardenBathLaundry, true
	default:
		return 0, false
	}
}

// ================== CLOTHING AND ACCESSORIES  ==================

type ClothingAccessoriesSubDep int

const (
	ClothingAccessoriesClothing ClothingAccessoriesSubDep = iota
	ClothingAccessoriesAccessories
	ClothingAccessoriesBagsPurses
)

func (p ClothingAccessoriesSubDep) ToString() string {
	switch p {
	case ClothingAccessoriesClothing:
		return "Clothing"
	case ClothingAccessoriesAccessories:
		return "Accessories"
	case ClothingAccessoriesBagsPurses:
		return "Bags & Purses"
	default:
		return "Unknown"
	}
}


func (p ClothingAccessoriesSubDep) ToInt() int {
	return int(p)
}

func IsStringValidClothingAccessoriesSubDep(subdep string) (ClothingAccessoriesSubDep, bool) {

    normalizedSubDep := strings.ReplaceAll(subdep, "&", "and")
    normalizedSubDep = strings.TrimSpace(normalizedSubDep)

    switch normalizedSubDep {
    case "Clothing":
        return ClothingAccessoriesClothing, true
    case "Accessories":
        return ClothingAccessoriesAccessories, true
    case "Bags and Purses":
        return ClothingAccessoriesBagsPurses, true
    default:
        return 0, false

    }
}



// ================== HEALTH BATH AND BEAUTY  ==================

type HealthBathBeautySubDep int

const (
	HealthBathBeautyBeauty HealthBathBeautySubDep = iota
	HealthBathBeautyBathBody
	HealthBathBeautyHealth
)


func (p HealthBathBeautySubDep) ToString() string {
	switch p {
	case HealthBathBeautyBeauty:
		return "Beauty"
	case HealthBathBeautyBathBody:
		return "Bath & Body"
	case HealthBathBeautyHealth:
		return "Health"
	default:
		return "Unknown"
	}
}


func (p HealthBathBeautySubDep) ToInt() int {
	return int(p)
}


func IsStringValidHealthBathBeautySubDep(subdep string) (HealthBathBeautySubDep, bool) {

    normalizedSubDep := strings.ReplaceAll(subdep, "&", "and")
    normalizedSubDep = strings.TrimSpace(normalizedSubDep)

    switch normalizedSubDep {
    case "Beauty":
        return HealthBathBeautyBeauty, true
    case "Bath and Body":
        return HealthBathBeautyBathBody, true
    case "Health":
        return HealthBathBeautyHealth, true
    default:
        return 0, false
        
    }

}


// ================== TOYS, KIDS AND BABIES  ==================

type ToysKidsBabiesSubDep int

const (
	ToysKidsBabiesClothing ToysKidsBabiesSubDep = iota
	ToysKidsBabiesShoes
	ToysKidsBabiesToys
)

func (p ToysKidsBabiesSubDep) ToString() string {

	switch p {
	case ToysKidsBabiesClothing:
		return "Clothing"
	case ToysKidsBabiesShoes:
		return "Shoes"
	case ToysKidsBabiesToys:
		return "Toys"
	default:
		return "Unknown"
	}
}


func (p ToysKidsBabiesSubDep) ToInt() int {
	return int(p)
}


func IsStringValidToysKidsBabiesSubDep(subdep string) (ToysKidsBabiesSubDep, bool) {

    normalizedSubDep := strings.ReplaceAll(subdep, "&", "and")
    normalizedSubDep = strings.TrimSpace(normalizedSubDep)

    switch normalizedSubDep {
    case "Clothing":
        return ToysKidsBabiesClothing, true
    case "Shoes":
        return ToysKidsBabiesShoes, true
    case "Toys":
        return ToysKidsBabiesToys, true
    default:
        return 0, false
    }

}
