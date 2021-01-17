package repository

type Book interface{
	Get(isbn string) ([] model.Book, error)
}