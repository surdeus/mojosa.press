package main

import(
	"os"
	"flag"
	"log"
	"net/http"
	"github.com/surdeus/mojosa.press/src/hndl"
	"github.com/surdeus/mojosa.press/src/tmpl"
	"github.com/surdeus/mojosa.press/src/path"
	"github.com/surdeus/mojosa.press/src/urlpath"
	"github.com/surdeus/mojosa.press/src/tempconfig"
	"regexp"
)

var(
)

func
main(){
	var err error

	AddrStr := flag.String("a", ":8080", "Adress string")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		os.Exit(1)
	}
	tmpl.Templates = tmpl.ParseSepTemplates()

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

	log.Printf("%s: running on '%s'\n", os.Args[0], *AddrStr)
	log.Fatal(http.ListenAndServe(*AddrStr, nil))
}
