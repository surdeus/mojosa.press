package main

import(
	"os"
	"flag"
	"log"
	"net/http"
	"github.com/k1574/mojosa.press/m/hndl"
	"github.com/k1574/mojosa.press/m/path"
	"github.com/k1574/mojosa.press/m/urlpath"
	"github.com/k1574/mojosa.press/m/tempconfig"
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
