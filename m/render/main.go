package render

import(
	"html/template"
	"net/http"
	"mojosa/press/m/path"
)

var(
	Templates *template.Template = template.Must(
		template.ParseGlob(
			path.Tmpl+"/*"
		)
	)
)

func
WriteTemplate(w http.ResponseWrite, tmpl string, p *post.Post){
	e := Templates.ExecuteTemplate(w, tmpl, p)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternal)
	}
}
