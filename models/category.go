package models

type Category struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" form:"name" validate:"required"`
	// Film   FilmResponse `json:"film" gorm:"films"`
	// FilmID int          `json:"-" form:"-" gorm:"-"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CategoryResponse) TableName() string {
	return "categorys"
}
