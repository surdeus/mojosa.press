package main

import(
	"log"
	"net/http"
	"mojosa/press/m/hndl"
	"mojosa/press/m/path"
)

var(
	AddrStr = ":8080"
)

func
main(){
	http.HandleFunc("/", hndl.Root)
	fs := http.FileServer(http.Dir(path.Static))
	http.Handle("/s/",
		http.StripPrefix(
			"/s/",
			fs) )
	log.Fatal(http.ListenAndServe(AddrStr, nil))
}
