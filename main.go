package main

import (
	ascii "myascii/AsciiHelper"
	"html/template"
	"net/http"
)

type TemplateData struct {
    Result string
}


func homePage(w http.ResponseWriter, r *http.Request) {
	// Parse the template file
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Render the template
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func asciiWeb(w http.ResponseWriter, r *http.Request) {
	// Only process POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Get the form values
	text := r.FormValue("text")
	templateChoice := r.FormValue("template")

	// Here you can process the input (e.g., ASCII art conversion)
	// For this example, we'll just display the received values
	result := ascii.Transform(text, templateChoice)
	tmpl, err := template.ParseFiles("templates/ascii-web.html")
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }

    // Create a TemplateData instance to pass to the template
    data := TemplateData{
        Result: result,
    }

    // Render the template with the result
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    }
	
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ascii-web", asciiWeb)
	http.ListenAndServe(":8080", nil)
}
