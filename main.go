package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// fileServer := http.FileServer(http.Dir("./pages"))
	// http.Handle("/", fileServer)

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "Hello there!\n")
	// })

	http.HandleFunc("/", process)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

		http.ServeFile(w, r, "./pages/form.html")
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		name := r.FormValue("name")
		occupation := r.FormValue("occupation")

		fmt.Fprintf(w, "%s is a %s\n", name, occupation)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
