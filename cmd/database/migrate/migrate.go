package migrate

import (
	"fmt"
	"github.com/yayayapluto/api-ukk-online/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&entities.ObjectType{}); err != nil {
		panic(err.Error())
	}
	fmt.Println("Migration done")
	return nil
}
