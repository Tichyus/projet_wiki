package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)


// We use a form to enter data in order to create an article
// The way it is done, we check for a potential error at each step, making sure to make a return if an error is defined at any point.
// We also use packr in order to simplify the handling of templates.
func GetData(w http.ResponseWriter, r *http.Request) {
	htmlTemplate := template.New("getData.html")
	getData, error := box.FindString("getData.html")
	if error != nil {
		log.Print(error)
		return
	}
	template, error := htmlTemplate.Parse(getData) // We parse the template file, replacing values
	if error != nil {
		log.Print(error)
		return
	}
	error = template.Execute(w, nil)
	if error != nil {
		log.Print(error)
		return
	}
}

func ReceiveData(w http.ResponseWriter, r *http.Request) {

	error := r.ParseForm()
	if error != nil {
		log.Printf("error: %v", error)
		return
	}
	test_name := r.FormValue("test_name")
    test_city := r.FormValue("test_city")

	fmt.Fprintf(w, "Entered name = %s\n", test_name)
    fmt.Fprintf(w, "Entered city = %s\n", test_city)
}