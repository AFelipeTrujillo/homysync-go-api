package domain

const (
	CategoryDairy     = "dairy"
	CategoryProduce   = "produce"
	CategoryMeat      = "meat"
	CategoryGrains    = "grains"
	CategoryPantry    = "pantry"
	CategoryBeverages = "beverages"
	CategoryCleaning  = "cleaning"
	CategoryFrozen    = "frozen"
	CategoryPersonal  = "personal"
	CategoryOther     = "other"
)

func ListAllCategories() []string {
	return []string{
		CategoryDairy,
		CategoryProduce,
		CategoryMeat,
		CategoryGrains,
		CategoryPantry,
		CategoryBeverages,
		CategoryCleaning,
		CategoryFrozen,
		CategoryPersonal,
		CategoryOther,
	}
}

func IsValidCategory(category string) bool {
	for _, c := range ListAllCategories() {
		if c == category {
			return true
		}
	}
	return false
}
