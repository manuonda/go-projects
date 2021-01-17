package model 

//orm library
import (
	"gorm.io/gorm"
)

type Book struct{
	gorm.Model //composition in go
	Title string
	Description string
	ISBM string
	Author string
}