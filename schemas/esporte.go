package schemas

import "gorm.io/gorm"

type Esporte struct {
	gorm.Model
	Name string
	Slug string
}
