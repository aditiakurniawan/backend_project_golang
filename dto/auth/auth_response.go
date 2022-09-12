package authdto

type LoginResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullName" `
	Email    string `gorm:"type: varchar(255)" json:"email" from:"email" `
	Password string `gorm:"type: varchar(255)" json:"password" `
	Token    string `gorm:"type: varchar(255)" json:"token" `
}

// type RegisterResponse struct {
// 	Email string `gorm:"type: varchar(255)" json:"email" `
// 	Token string `gorm:"type: varchar(255)" json:"Token" `
// }
