package seeder

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/yayayapluto/api-ukk-online/entities"
	"gorm.io/gorm"
)

func ObjectTypeSeeder(db *gorm.DB, q int) {
	if err := db.Exec("TRUNCATE table object_types RESTART IDENTITY").Error; err != nil {
		panic(err)
	}
	for i := 0; i < q; i++ {
		data := &entities.ObjectType{
			Name: gofakeit.LoremIpsumWord(),
		}
		if err := db.Create(data).Error; err != nil {
			fmt.Println("skipped")
			continue
		}
	}
	fmt.Println("Seeding object type done")
}
