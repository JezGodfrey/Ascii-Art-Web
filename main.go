package main

import (
	ascii "ascii/asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type AsciiReq struct {
	Text  string
	Style string
	Art   string
}

// Generic handler for pages in the aaw directory
func handler(w http.ResponseWriter, r *http.Request) {
	// Check if files exists - if not then 404, else 500
	expected := []string{"/", "/ascii-art.html", "/bad-request.html", "/index.html", "/not-found.html"}
	found := false
	for _, e := range expected {
		if r.URL.Path == e {
			found = true
		}
	}

	// If web page isn't found - try to display 404 error page
	if !found {
		p, err := template.ParseFiles("aaw/not-found.html")
		// Otherwise default to Internal Server Error
		if err != nil {
			fmt.Printf("500 - Internal Server Error - %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Printf("%v - Page not found\n", http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		p.Execute(w, "not-found.html")
		return
	}

	// If no address path (default), display index page
	if r.URL.Path == "/" {
		r.URL.Path = "/index.html"
	}

	p, err := template.ParseFiles("aaw" + r.URL.Path)
	if err != nil {
		fmt.Printf("500 - Internal Server Error - %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// After parsing the html file above, write the status code to the header and execute the page
	fmt.Printf("%v - Status OK\n", http.StatusOK)
	w.WriteHeader(http.StatusOK)
	p.Execute(w, r.URL.Path)
}

// Handler for generating ascii-art and displaying ascii-art page
func asciiHandler(w http.ResponseWriter, r *http.Request) {
	var ar AsciiReq

	p, err := template.ParseFiles("aaw/ascii-art.html")
	if err != nil {
		fmt.Printf("500 - Internal Server Error - %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get form values from txt and style ('name' attributes)
	ar.Text = r.FormValue("txt")
	ar.Style = r.FormValue("style")
	ar.Art = ascii.AsciiArt(ar.Text, ar.Style)

	// Error handling in AsciiArt function returns "error" making it bad request
	if ar.Art == "error" {
		p, err = template.ParseFiles("aaw/bad-request.html")
		if err != nil {
			fmt.Printf("500 - Internal Server Error - %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Printf("%v - Bad Request\n", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		p.Execute(w, "bad-request.html")
		return
	}

	fmt.Printf("%v - Status OK\n", http.StatusOK)
	w.WriteHeader(http.StatusOK)
	p.ExecuteTemplate(w, "ascii-art.html", ar)
}

func main() {
	// HandleFunc adds route handlers to the web server - how to handle incoming requests
	// The first argument accepts path to listen for "/"
	// Second argument takes a function that holds logic to correctly respond to the request
	// By default, the function is func(w http.ResponseWriter, r *http.Request) - writer sending response back
	// For an object (type) to respond to HTTP requests, use http.Handle
	// For a function to respond to HTTP requests, use http.HandleFunc
	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.HandleFunc("/ascii-art.html", asciiHandler)

	// If failure to start web server on port 8080, fatal log
	log.Fatal(http.ListenAndServe(":8080", nil))
}
