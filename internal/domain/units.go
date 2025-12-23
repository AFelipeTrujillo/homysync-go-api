package domain

const (
	UnitGram       = "g"
	UnitKilogram   = "kg"
	UnitMilliliter = "ml"
	UnitLiters     = "l"
	UnitPiece      = "unit"
	UnitTablespoon = "tbsp"
)

func GetStandardUnits() []string {
	return []string{
		UnitGram,
		UnitKilogram,
		UnitMilliliter,
		UnitLiters,
		UnitPiece,
		UnitTablespoon,
	}
}

func IsValidUnit(unit string) bool {
	for _, u := range GetStandardUnits() {
		if u == unit {
			return true
		}
	}
	return false
}
