package entities

type ObjectType struct {
	ID   uint   `gorm:"primaryKey; autoIncrement" json:"id"`
	Name string `gorm:"uniqueIndex, not null" json:"name"`
	Timestamp
}
