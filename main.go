package main

import(
	"log"
	"net/http"
	"mojosa/press/m/hndl"
	"mojosa/press/m/path"
	"mojosa/press/m/uri"
)

var(
	AddrStr = ":8080"
)

func
main(){
	http.HandleFunc("/", hndl.Root)
	fs := http.FileServer(http.Dir(path.Static))
	http.Handle(uri.StaticPrefix,
		http.StripPrefix(
			uri.StaticPrefix,
			fs) )

	http.HandleFunc(uri.ViewPostPrefix, hndl.ViewPost)
	//http.Handle(uri.ViewPostPrefix, http.StripPrefix(
	//	uri.ViewPostPrefix, hndl.ViewPost) )

	log.Fatal(http.ListenAndServe(AddrStr, nil))
}
