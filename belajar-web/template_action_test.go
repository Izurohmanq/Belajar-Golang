package belajarweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateActionIft(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template data map",
		Name:  "ASEPPPPP",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "htttp://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIft(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
func TemplateOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template data map",
		"FinalValue": 90,
	})
}

func TestTemplateOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "htttp://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template data map",
		"Hobbies": []string{
			"Mbut", "mek", "kon",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "htttp://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(writer, "with.gohtml", map[string]interface{}{
		"Title": "Template data map",
		"Address": map[string]interface{}{
			"Street": "jalan konto",
			"City":   "babandinadiand",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "htttp://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
