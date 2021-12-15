package main

import(
	"log"
	"net/http"
	"mojosa/press/m/hndl"
	"mojosa/press/m/path"
	"mojosa/press/m/urlpath"
)

var(
	AddrStr = ":8080"
)

func
main(){

	fs := http.FileServer(http.Dir(path.Static))
	http.Handle(urlpath.StaticPrefix,
		http.StripPrefix(
			urlpath.StaticPrefix,
			fs) )

	for _, v := range hndl.Defs {
		http.HandleFunc(v.Pref,
			hndl.MakeHttpHandleFunc(v.Pref, v.Re, v.Fn))
	}

	log.Fatal(http.ListenAndServe(AddrStr, nil))
}
