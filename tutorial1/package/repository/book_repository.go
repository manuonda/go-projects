package repository
import (
	"gorm.io/gorm"
	"github.com/manuonda/go-projects/tree/master/tutorial1/package/model"
)

type Book interface{
	Get(isbn string) ([] model.Book, error)
}

type book struct{
	db *gorm.DB
}

func( b *book) Get(isbn string) ([model.Book,error]){
	var books []model.Book

	q := b.db.Model()
}


func NewBookRepository(db *gorm.DB) Book {
  if  db == nil {
	  panic("db has not been provided")
  }
  return &book{}
}