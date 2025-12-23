package domain

import "gorm.io/gorm"

type User struct {
	ID           uint           `gorm:"primaryKey"`
	CreatedAt    int64          `gorm:"autoCreateTime"`
	UpdatedAt    int64          `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Email        string         `gorm:"uniqueIndex;not null"`
	PasswordHash string         `gorm:"not null"`
	GlobalName   string         `gorm:"not null"`
	Memberships  []Membership   `gorm:"foreignKey:UserID"`
}

type Household struct {
	ID        uint     `gorm:"primaryKey"`
	CreatedAt int64    `gorm:"autoCreateTime"`
	Name      string   `gorm:"not null"`
	Timezone  string   `gorm:"not null"`
	Recipies  []Recipe `gorm:"foreignKey:HouseholdID"`
}

type Membership struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null;index"`
	HouseholdID uint   `gorm:"not null;index"`
	Role        string `gorm:"default:'regular'"`
	IsActive    bool   `gorm:"default:true"`
}

type CatalogItem struct {
	ID            uint `gorm:"primaryKey"`
	HouseholdID   uint `gorm:"not null;index"`
	CanonicalName string
	Category      string `gorm:"not null";default:'other'`
}

type Recipe struct {
	ID          uint   `gorm:"primaryKey"`
	HouseholdID uint   `gorm:"not null;index"`
	Title       string `gorm:"not null"`
	Description string
	SourceURL   string
	Ingredients []RecipeIngredient `gorm:"foreignKey:RecipeID"`
}

type RecipeIngredient struct {
	ID        uint `gorm:"primaryKey"`
	RecipeID  uint `gorm:"not null;index"`
	CatalogID uint `gorm:"not null;index"`
	Quantity  string
	Unit      string
}
