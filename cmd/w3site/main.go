package main

import(
	"os"
	"log"
	"net/http"
	"mojosa/press/m/hndl"
	"mojosa/press/m/path"
	"mojosa/press/m/urlpath"
	"mojosa/press/m/tempconfig"
	"regexp"
)

var(
	AddrStr = ":8080"
)

func
main(){
	var err error
	os.Mkdir(path.Data, 0755)
	tempconfig.TmpCfg, err = tempconfig.Read(path.TempConfig)
	os.Mkdir(path.Post, 0755)
	if err != nil { panic(err) }
	fs := http.FileServer(http.Dir(path.Static))
	http.Handle(urlpath.StaticPrefix,
		http.StripPrefix(
			urlpath.StaticPrefix,
			fs) )

	for _, v := range hndl.Defs {
		http.HandleFunc(v.Pref,
			hndl.MakeHttpHandleFunc(v.Pref, regexp.MustCompile(v.Re), v.Fn))
	}

	log.Printf("%s: running on '%s'\n", os.Args[0], AddrStr)
	log.Fatal(http.ListenAndServe(AddrStr, nil))
}
