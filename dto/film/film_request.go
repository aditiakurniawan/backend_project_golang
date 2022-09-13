package filmdto

type FilmRequest struct {
	ID            int    `json:"id" `
	Title         string `json:"title" from:"title"  gorm:"type: varchar(255)" `
	Thumbnailfilm string `json:"thumbnailfilm" gorm:"type:varchar(255)"`
	Year          string `json:"year" form:"year" gorm:"type: string"`
	Description   string `json:"description" gorm:"type: text" `
	CategoryID    int    `json:"categoryid"`
}
