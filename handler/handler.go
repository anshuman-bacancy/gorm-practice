package handler

import (
	"net/http"
	"strings"
	"strconv"
	"html/template"
	"encoding/json"

	"fake.com/anshuman/services"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	tpl = template.Must(template.ParseGlob("assets/*.html"))
	tpl.ExecuteTemplate(res, "index.html", nil)
}


func AddBookHandler(res http.ResponseWriter, req *http.Request) {
	book := req.FormValue("name")
	author := req.FormValue("author")
	pages, _ := strconv.ParseInt(req.FormValue("pages")[0:], 10, 64)
	isbn, _ := strconv.ParseInt(req.FormValue("isbn")[0:], 10, 64)

	bookservices.SaveBook(book, author, pages, isbn)
}

func GetBookHandler(res http.ResponseWriter, req *http.Request) {
	name := mux.Vars(req)["name"]
	nameFormatted := strings.Replace(name, "_", " ", len(name))

	book := bookservices.GetBook(nameFormatted)

	bookJson, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "application/json")
	res.Write(bookJson)
}


func UpdateBookHandler(res http.ResponseWriter, req *http.Request) {
	bookToUpdate := mux.Vars(req)["name"]
	nameFormatted := strings.Replace(bookToUpdate, "_", " ", len(bookToUpdate))

	book := req.FormValue("name")
	author := req.FormValue("author")
	pages, _ := strconv.ParseInt(req.FormValue("pages")[0:], 10, 64)
	isbn, _ := strconv.ParseInt(req.FormValue("isbn")[0:], 10, 64)

	updatedBook := bookservices.UpdateBook(nameFormatted, book, author, pages, isbn)
	bookJson, _ := json.Marshal(updatedBook)
	res.Header().Set("Content-Type", "application/json")
	res.Write(bookJson)
}


func DeleteBookHandler(res http.ResponseWriter, req *http.Request) {
	bookToDelete := mux.Vars(req)["name"]
	nameFormatted := strings.Replace(bookToDelete, "_", " ", len(bookToDelete))

	bookservices.DeleteBook(nameFormatted)
}
