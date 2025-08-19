package entities

type Organizer struct {
	ID            uint   `gorm:"primaryKey; autoIncrement" json:"id"`
	Name          string `gorm:"unique; not null" json:"name"`
	Address       string `gorm:"not null" json:"address"`
	BankName      string `gorm:"not null" json:"bank_name"`
	AccountNumber string `gorm:"not null" json:"account_number"`
	AccountName   string `gorm:"not null" json:"account_name"`
	Timestamp
}
