package models

type Film struct {
	ID            int              `json:"id"`
	Title         string           `json:"title" gorm:"type: varchar(255)"`
	Thumbnailfilm string           `jspon:"thumbnailfilm" gorm:"type:varchar(255)"`
	Year          string           `json:"year" gorm:"type: text"`
	Category      CategoryResponse `json:"category" gorm:"categorys"`
	CategoryID    int              `json:"-" form:"-" gorm:"-"`
	Description   string           `json:"description" gorm:"type: text"`
	// UserId        int              `json:"UserId"`
	// User          UsersResponse    `json:"user" gorm:"users"`
}

type FilmResponse struct {
	ID            int              `json:"id"`
	Name          string           `json:"name"`
	Title         string           `json:"title"`
	Thumbnailfilm string           `jspon:"thumbnailfilm"`
	Year          string           `json:"year"`
	Category      CategoryResponse `json:"category"`
	CategoryID    int              `json:"-" form:"-" gorm:"-"`
	Description   string           `json:"description" gorm:"type: text"`
}

func (FilmResponse) TableName() string {
	return "Films"
}
