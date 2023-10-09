package render

import (
	"bytes"
	"github.com/h1mmeister/bookings_go/pkg/config"
	"github.com/h1mmeister/bookings_go/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// var templateCache = make(map[string]*template.Template)

var app *config.AppConfig

func NewTemplates(appConfig *config.AppConfig) {
	app = appConfig
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// function to render templates
func RenderTemplate(writer http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// create template cache
	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Println("Not able to start the app", err)
	// 	log.Fatal(err)
	// }

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Didn't find the requested template")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)

	if err != nil {
		log.Println("Not able to execute the template via buffer", err)
	}

	// render the template
	_, err = buf.WriteTo(writer)
	if err != nil {
		log.Println("Not able to write to ResponseWriter", err)
	}

	// parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	// if err != nil {
	// 	fmt.Println("Something went wrong!", err)
	// }

	// error := parsedTemplate.Execute(writer, nil)
	// if err != nil {
	// 	fmt.Println("Error parsing template", error)
	// 	return
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		log.Println("Not able to get the files", err)
		return myCache, err
	}

	// range through all the files ending with *.page.tmpl
	for _, page := range pages {
		fileName := filepath.Base(page)
		ts, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			log.Println("Error parsing the page", err)
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println("Not able to get the layout template", err)
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Println("Not able to add layout to template set", err)
				return myCache, err
			}
		}

		myCache[fileName] = ts
	}
	return myCache, nil
}

// func RenderTemplate(writer http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have the template in our cache
// 	_, inMap := templateCache[t]
// 	log.Println("Value of inMap", inMap)
// 	if !inMap {
// 		// we need to create the template
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println("Error creating the template", err)
// 		}
// 	} else {
// 		log.Println("Using template from Cache!")
// 	}

// 	tmpl = templateCache[t]
// 	err = tmpl.Execute(writer, nil)
// 	if err != nil {
// 		log.Println("Error", err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
// 	}
// 	log.Println("Template formation ->", templates)

// 	// parse the templates
// 	parsedTemplate, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		log.Println("Error parsing the templates", err)
// 		return err
// 	}

// 	templateCache[t] = parsedTemplate
// 	return nil
// }
