package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
// 	if err != nil {
// 		log.Fatal("Unable to parse from template :", err)
// 	}
// 	parsedTemplate.Execute(w, nil)
// }

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Template cachemizde var mı diye check etmemiz gerekiyor.
	_, inMap := tc[t]
	if !inMap {
		// template yaratmamız lazım
		log.Println("Creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//cache kullanan templateimiz var
		log.Println("using cached template.")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache which is a map
	tc[t] = tmpl
	return nil
}
