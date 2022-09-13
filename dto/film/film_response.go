package filmdto

type FilmResponse struct {
	ID            int    `json:"id" `
	Title         string `json:"title" from:"title"  gorm:"type: varchar(255)" `
	Thumbnailfilm string `json:"thumbnailfilm" gorm:"type:varchar(255)"`
	Year          string `json:"year" form:"year" gorm:"type: text"`
	Description   string `json:"description" gorm:"type: text" `
	Category      int    `json:"category"`
}
