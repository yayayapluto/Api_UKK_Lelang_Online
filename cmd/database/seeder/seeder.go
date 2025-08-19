package seeder

import (
	"fmt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {

	// data master
	ObjectTypeSeeder(db, 5)
	OrganizerSeeder(db, 10)

	fmt.Println("Seeding done")
	return nil
}
