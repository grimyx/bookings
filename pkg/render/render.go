package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/grimyx/bookings/models"
	"github.com/grimyx/bookings/pkg/config"
)

var app *config.AppConfig

// mapa funkcija koje ce se koristiti u templejtu
var functions = template.FuncMap{}

// Sets the config for template
func NewTemplate(a *config.AppConfig) {
	app = a
}
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// moze i ovako, samo dodam base layout na kraj
	// parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")

	var tc map[string]*template.Template

	// ako je Use_cache true main kreira kes i to se koristi
	// ako nije svaki put se poziva CreateTemplateCache i to je za development, jer ce svaki put da cita fajlove sa diska i mogu da se uoce promene odmah
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// ok je bool , i to je mogucnost mape da u slucaju da ne nadje vrednost vrati false, a ukoliko nadje vrati true
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Template file not found !!!!")
	}

	// kreira byte bufer
	buf := new(bytes.Buffer)

	// ubacuje templejt u bufer
	_ = t.Execute(buf, td)

	// bute bufer se prebacuje u ResponseWriter
	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to the browser,", err)
	}

	// stari kod koji ne podrzava layout

	/*
		parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
		err = parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error parsing template: ", err)
			return
		}
	*/
}

// CreateTemplateCache - creates template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	// kreira mapu [String][Template] sa imenima stranica i templajtima za nju
	myCache := map[string]*template.Template{}

	// pages sadrzi niz svih fajlova koji se nalaze u direktorijumu templates i imaju ekstenziju .pages.html
	// svaki element je cela putanja do fajla, na pr ./templates/home.pages.html
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		// izdvaja samo ime iz kompletne putanje ka fajlu templejta
		name := filepath.Base(page)

		// ovo kreira templejt , dodeljuje mu neke funkcije to ce da vidimo sta je i parsuje fajl koji je
		// vezan za taj templejt
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		// ovo pretrazuje direktorijum templates za sve fajlove koji se zavrsavaju na layout.html
		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
