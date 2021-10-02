package routes

import (
	"fmt"
	"github.com/Sh4yy/ImageRender/utils"
	"github.com/gorilla/mux"
	"github.com/joncalhoun/qson"
	"html/template"
	"net/http"
	"path"
)

// Render renders html file from the templates directory
func Render(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	templateName := vars["template"]
	query := r.URL.RawQuery
	data := map[string]interface{}{}
	err := qson.Unmarshal(&data, query)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp := path.Join(utils.Config.Directory.Templates, fmt.Sprintf("%s.html", templateName))
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}