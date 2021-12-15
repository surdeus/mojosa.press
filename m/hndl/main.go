package hndl

import(
	"net/http"
	"net/url"
	"regexp"
	"mojosa/press/m/render"
	"mojosa/press/m/urlpath"
	//"mojosa/press/m/post"
	"fmt"
)

type Handler func(http.ResponseWriter, *http.Request, url.Values, string)
type FuncDefine struct {
	Pref string
	Re *regexp.Regexp
	Fn Handler	
}

var(
	Defs = []FuncDefine {
		{urlpath.RootPrefix, nil, Root},
		{urlpath.ViewPostPrefix, nil, ViewPost},
		{urlpath.TestPrefix, nil, Test},
	}
)


func
MakeHttpHandleFunc(pref string, re *regexp.Regexp, fn Handler) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path[len(pref):]
	if urlpath.ValidifyPath(p, re) != true {
		http.NotFound(w, r)
	}
	q, e := url.ParseQuery(r.URL.RawQuery)

	if e != nil {
	}
	
	fn(w, r, q, p)
} }

func
Root(w http.ResponseWriter, r *http.Request, q url.Values, p string) {
	render.WriteTemplate(w, "root", nil)
}
	

func
ViewPost(w http.ResponseWriter, r *http.Request, q url.Values, p string){
}

func
Test(w http.ResponseWriter, r *http.Request, q url.Values, p string){
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Path: '%s'\nRawQuery:'%s'\n", r.URL.Path, r.URL.RawQuery)
	fmt.Fprintf(w, "p: '%s'\n", p)
	fmt.Fprintf(w, "q:\n")
	for k, v := range q {
		fmt.Fprintf(w, "\t'%s':\n", k)
		for _, s := range v {
			fmt.Fprintf(w, "\t\t'%s'\n", s)
		}
	}
}
