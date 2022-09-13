package models

// import "time"

type Transaction struct {
	ID        int           `json:"id"`
	StartDate string        `json:"-"`
	DueDate   string        `json:"-"`
	UserId    int           `json:"UserId"`
	User      UsersResponse `json:"user" gorm:"users"`
	Attache   string        `json:"attache" gorm:"type: varchar(255)"`
	Status    string        `json:"Status" gorm:"type: varchar(255)"`
	// CreatedAt time.Time `json:"-"`
	// UpdatedAt time.Time `json:"-"`
}

type TransactionResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
}

func (TransactionResponse) TableName() string {
	return "users"
}
