package render

import(
	"html/template"
	"net/http"
	"mojosa/press/m/path"
)

var(
	Templates *template.Template = template.Must(
		template.ParseGlob(path.Tmpl+"/*" ))
)

func
WriteTemplate(w http.ResponseWriter, tmpl string, v interface{} ){
	e := Templates.ExecuteTemplate(w, tmpl, v)
	if e != nil {
		http.Error(w, e.Error(),
			http.StatusInternalServerError)
	}
}
