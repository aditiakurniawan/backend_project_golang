package transactiondto

// import "time"

type TransactionResponse struct {
	ID        int    `json:"id"`
	StartDate string `json:"startdate" gorm:"type: varchar(255)"`
	DueDate   string `json:"duedate" gorm:"type: varchar(255)"`
	UserId    int    `json:"UserId"`
	Attache   string `json:"attache" gorm:"type: varchar(255)"`
	Status    string `json:"Status" gorm:"type: varchar(255)"`
}
