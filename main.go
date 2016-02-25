package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/sj14/calc/relay"
)

//Compile templates on start
var (
	templates = template.Must(template.ParseFiles("web/index.html"))
)

// Displays the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	equation := r.FormValue("calculate")

	// q := r.URL.Query()
	fmt.Printf("URL: %v\n", r.URL.RawQuery)
	// fmt.Printf("Query: %v\n", q.Values)
	fmt.Printf("Equation: %v\n", equation)

	// To match functions, which are lowercase
	equationLow := strings.ToLower(equation)
	result, steps, err := relay.Relay(equationLow)

	log.Printf("Web In: %v Result: %v Error: %v\n", equation, result, err)

	m := map[string]interface{}{"Input": equation}

	if equation != "" {
		if err == nil {
			m["Result"] = result

			if len(steps) > 1 {
				m["Steps"] = steps
			}
		} else {
			m["Error"] = err
		}
	}
	display(w, "index.html", &m)
	return
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/bootstrap/", http.StripPrefix("/bootstrap/", http.FileServer(http.Dir("web/bootstrap"))))
	log.Fatal(http.ListenAndServe(":9000", nil))
}
