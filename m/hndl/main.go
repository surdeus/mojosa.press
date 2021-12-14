package hndl

import(
	"net/http"
	"mojosa/press/m/render"
)

func
Root(w http.ResponseWriter, r *http.Request){
	render.WriteTemplate(w, "root", nil)
}
	

func
ReadPost(w http.ResponseWriter, r *http.Request){
}

