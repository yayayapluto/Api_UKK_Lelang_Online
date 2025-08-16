package migrate

import (
	"fmt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	_ = db
	fmt.Println("Migration done")
	return nil
}
