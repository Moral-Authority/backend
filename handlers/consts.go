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
		return "home-and-garden"
	case HealthBathBeautyProductDepartment:
		return  "clothing-and-accessories"
	case ClothingAccessoriesProductDepartment:
		return "health-bath-and-beauty"
	case ToysKidsBabiesProductDepartment:
		return "toys-and-kids-and-babies"
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

	normalizedDepartment := strings.TrimSpace(department)

	switch normalizedDepartment {
	case "home-and-garden":
		return HomeGardenProductDepartment, true
	case "clothing-and-accessories":
		return ClothingAccessoriesProductDepartment, true
	case "health-bath-and-beauty":
		return HealthBathBeautyProductDepartment, true
	case "toys-and-kids-and-babies":
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
		return "cleaning-supplies"
	case HomeGardenKitchenDining:
		return "kitchen-and-dining"
	case HomeGardenFurnitureArt:
		return "furniture-and-art"
	case HomeGardenGarden:
		return "garden"
	case HomeGardenBeddingLinens:
		return "bedding-and-linens"
	case HomeGardenLightingLamps:
		return "lighting-and-lamps"
	case HomeGardenBathLaundry:
		return "bath-and-laundry"
	default:
		return "Unknown"
	}
}

// ToInt converts HomeGardenSubDep to int
func (p HomeGardenSubDep) ToInt() int {
	return int(p)
}


func IsStringValidHomeGardenSubDep(subdep string) (HomeGardenSubDep, bool) {

	normalizedSubDep := strings.TrimSpace(subdep)

	switch normalizedSubDep {
	case "cleaning-supplies":
		return HomeGardenCleaningSupplies, true
	case "kitchen-and-dining":
		return HomeGardenKitchenDining, true
	case "furniture-and-art":
		return HomeGardenFurnitureArt, true
	case "garden":
		return HomeGardenGarden, true
    case "bedding-and-linens":
        return HomeGardenBeddingLinens, true
    case "lighting-and-lamps":
        return HomeGardenLightingLamps, true
    case "bath-and-laundry":
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
		return "clothing"
	case ClothingAccessoriesAccessories:
		return "accessories"
	case ClothingAccessoriesBagsPurses:
		return "bags-and-purses"
	default:
		return "Unknown"
	}
}


func (p ClothingAccessoriesSubDep) ToInt() int {
	return int(p)
}

func IsStringValidClothingAccessoriesSubDep(subdep string) (ClothingAccessoriesSubDep, bool) {

    normalizedSubDep := strings.TrimSpace(subdep)

    switch normalizedSubDep {
    case "clothing":
        return ClothingAccessoriesClothing, true
    case "accessories":
        return ClothingAccessoriesAccessories, true
    case "bags-and-purses":
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
		return "beauty"
	case HealthBathBeautyBathBody:
		return "bath-and-body"
	case HealthBathBeautyHealth:
		return "health"
	default:
		return "Unknown"
	}
}


func (p HealthBathBeautySubDep) ToInt() int {
	return int(p)
}


func IsStringValidHealthBathBeautySubDep(subdep string) (HealthBathBeautySubDep, bool) {

    normalizedSubDep := strings.TrimSpace(subdep)

    switch normalizedSubDep {
    case "beauty":
        return HealthBathBeautyBeauty, true
    case "bath-and-body":
        return HealthBathBeautyBathBody, true
    case "health":
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
		return "clothing"
	case ToysKidsBabiesShoes:
		return "shoes"
	case ToysKidsBabiesToys:
		return "toys"
	default:
		return "Unknown"
	}
}


func (p ToysKidsBabiesSubDep) ToInt() int {
	return int(p)
}


func IsStringValidToysKidsBabiesSubDep(subdep string) (ToysKidsBabiesSubDep, bool) {

	normalizedSubDep := strings.TrimSpace(subdep)

    switch normalizedSubDep {
    case "clothing":
        return ToysKidsBabiesClothing, true
    case "shoes":
        return ToysKidsBabiesShoes, true
    case "toys":
        return ToysKidsBabiesToys, true
    default:
        return 0, false
    }

}
