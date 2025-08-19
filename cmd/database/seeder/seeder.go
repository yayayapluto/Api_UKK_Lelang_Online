package seeder

import (
	"fmt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {

	ObjectTypeSeeder(db, 50)

	fmt.Println("Seeding done")
	return nil
}
