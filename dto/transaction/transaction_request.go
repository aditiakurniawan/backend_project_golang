package transactiondto

import "time"

type TransactionRequest struct {
	ID        int       `json:"id"`
	StartDate time.Time `json:"-"`
	DueDate   time.Time `json:"-"`
	UserId    int       `json:"UserId"`
	Attache   string    `json:"attache" gorm:"type: varchar(255)"`
	Status    string    `json:"Status" gorm:"type: varchar(255)"`
}
