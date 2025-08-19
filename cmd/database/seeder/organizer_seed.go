package seeder

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/yayayapluto/api-ukk-online/entities"
	"gorm.io/gorm"
)

func OrganizerSeeder(db *gorm.DB, q int) {
	if err := db.Exec("TRUNCATE table organizers RESTART IDENTITY").Error; err != nil {
		panic(err)
	}
	for i := 0; i < q; i++ {
		data := &entities.Organizer{
			Name:          gofakeit.LoremIpsumWord(),
			Address:       gofakeit.Address().Address,
			BankName:      gofakeit.BankName(),
			AccountNumber: gofakeit.Numerify("############"),
			AccountName:   gofakeit.MinecraftAnimal(),
		}
		if err := db.Create(data).Error; err != nil {
			fmt.Println("skipped")
			continue
		}
	}
	fmt.Println("Seeding organizer done")
}
