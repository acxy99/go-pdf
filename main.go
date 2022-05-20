package main

import (
	"fmt"
	u "go-pdf/pdfGenerator"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/testing", testing)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("error")
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func exportHandler(w http.ResponseWriter, r *http.Request) {

	// filebytes, err := ioutil.ReadFile("storage/example.pdf")
	// if err != nil {
	// 	fmt.Println("unable to read file")
	// }
	// if err := r.ParseForm(); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// path := r.FormValue("path")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	basepath, _ := os.Getwd()

	fileLocation := filepath.Join(basepath, "/storage/example.pdf")

	Openfile, err := os.Open(fileLocation) //Open the file to be downloaded later

	if err != nil {
		http.Error(w, "File not found.", 404) //return 404 if file is not found
		return
	}
	defer Openfile.Close() //Close after function return

	// tempBuffer := make([]byte, 512) //Create a byte array to read the file later
	// Openfile.Read(tempBuffer)

	cd := fmt.Sprintf("attachment; filename=%s", "WHATEVER_YOU_WANT.pdf")
	w.Header().Set("Content-Disposition", cd)

	if _, err := io.Copy(w, Openfile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func testing(w http.ResponseWriter, r *http.Request) {
	b := u.NewRequestPdf("")

	//html template path
	templatePath := "templates/sample.html"

	//path for download pdf
	outputPath := "storage/exampl2.pdf"

	//html template data
	templateData := struct {
		Title       string
		Description string
		Company     string
		Contact     string
		Country     string
	}{
		Title:       "HTML to PDF generator",
		Description: "This is the simple HTML to PDF file.",
		Company:     "Jhon Lewis",
		Contact:     "Maria Anders",
		Country:     "Germany",
	}

	if err := b.ParseTemplate(templatePath, templateData); err == nil {
		err := b.GeneratePDF(outputPath)
		if err != nil {
			fmt.Println("unable to print pdf")
		}
	}
}
