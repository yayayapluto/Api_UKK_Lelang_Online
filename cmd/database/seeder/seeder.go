package seeder

import (
	"fmt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	_ = db
	fmt.Println("Seeding done")
	return nil
}
