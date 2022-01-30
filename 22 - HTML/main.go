package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates *template.Template

type User struct {
	Name  string
	Email string
}

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)

	return err == nil
}

func getTemplates(templateName string) *template.Template {
	if !FileExists(templateName) {
		log.Fatal("Template informed does not exists!")
	}

	templates = template.Must(template.ParseGlob(templateName))

	return templates
}

func home(rw http.ResponseWriter, r *http.Request) {
	templateName := "index.html"
	homeTemplate := getTemplates(templateName)

	user := User{
		"Gustavo Ramos",
		"gustavo.ramos@email.com",
	}

	homeTemplate.ExecuteTemplate(rw, templateName, user)
}

func main() {
	http.HandleFunc("/home", home)

	fmt.Println("Runnning server at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
