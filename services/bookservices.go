package bookservices

import (
	"fake.com/anshuman/models"
	"fake.com/anshuman/server"
)

func SaveBook(book string, author string, pages int64, isbn int64) {
	bookToAdd := models.Book{Name: book, Author: author, Pages: pages, ISBN: isbn}
	_ = server.Db.Select("Name","Author","Pages","ISBN").Create(&bookToAdd)
}

func GetBook(book string) models.Book {
	var bookToGet models.Book
	server.Db.First(&bookToGet, "Name = ?", book)
	return bookToGet
}

func UpdateBook(bookToUpdate, book string, author string, pages int64, isbn int64) models.Book {
	var updateBook models.Book

	server.Db.First(&updateBook, "Name = ?", bookToUpdate)
	updateBook.Name = book
	updateBook.Author = author
	updateBook.Pages = pages
	updateBook.ISBN = isbn
	server.Db.Save(&updateBook)

	return updateBook
}

func DeleteBook(book string) {
	var bookToDelete models.Book

	// soft delete 
	//server.Db.First(&bookToDelete, "Name = ?", book)  

	server.Db.Unscoped().Delete(&bookToDelete)

}
