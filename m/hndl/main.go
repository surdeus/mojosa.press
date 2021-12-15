package hndl

import(
	"net/http"
	"mojosa/press/m/render"
	"mojosa/press/m/uri"
	//"mojosa/press/m/post"
	"fmt"
)

func
Root(w http.ResponseWriter, r *http.Request){
	render.WriteTemplate(w, "root", nil)
}
	

func
ViewPost(w http.ResponseWriter, r *http.Request){
	if uri.Validify(r.URL.Path) == false {
		http.NotFound(w, r)
	}
	
	fmt.Fprintf(w, "'%s'", r.URL.Path)
}

