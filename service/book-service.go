package service

import (
	"fmt"
	"log"

	"github.com/ibnu-a/Golang-RESTful-API/dto"
	"github.com/ibnu-a/Golang-RESTful-API/entity"
	"github.com/ibnu-a/Golang-RESTful-API/repository"
	"github.com/mashingan/smapping"
)

type BookService interface {
	Insert(b dto.BookCreateDTO) entity.Book
	Update(b dto.BookUpdateDTO) entity.Book
	Delete(b entity.Book)
	All() []entity.Book
	FindByID(bookID uint64) entity.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (db *bookService) Insert(b dto.BookCreateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	fmt.Println(book) //cek
	res := db.bookRepository.InsertBook(book)
	return res
}

func (db *bookService) Update(b dto.BookUpdateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := db.bookRepository.UpdateBook(book)
	return res
}

func (db *bookService) Delete(b entity.Book) {
	db.bookRepository.DeleteBook(b)
}

func (db *bookService) All() []entity.Book {
	return db.bookRepository.AllBooks()
}

func (db *bookService) FindByID(bookID uint64) entity.Book {
	return db.bookRepository.FindBookByID(bookID)
}

func (db *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := db.bookRepository.FindBookByID(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
