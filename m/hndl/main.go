package hndl

import(
	"html/template"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"mojosa/press/m/urlpath"
	"mojosa/press/m/post"
	"mojosa/press/m/sanitize"
	"mojosa/press/m/md"
	"mojosa/press/m/tmpl"
	//"io/ioutil"
	"fmt"
)

type Handler func(http.ResponseWriter, *http.Request, url.Values, string)
type FuncDefine struct {
	Pref, Re string
	Fn Handler
}

var(
	Defs = []FuncDefine {
		{urlpath.RootPrefix, "^$", Root},
		{urlpath.ViewPostPrefix, "^[0-9]+$", ViewPost},
		{urlpath.TypePostPrefix, "^$", TypePost},
		{urlpath.TypePostHndlPrefix, "^$", TypePostHndl},
		{urlpath.GetTestPrefix, "", GetTest},
		{urlpath.PostTestPrefix, "", PostTest},
	}
)


func
MakeHttpHandleFunc(pref string, re *regexp.Regexp, fn Handler) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path[len(pref):]
	if !urlpath.Validify(p, re) {
		http.NotFound(w, r)
		return
	}
	q, e := url.ParseQuery(r.URL.RawQuery)

	if e != nil {
	}
	
	fn(w, r, q, p)
} }

func
Root(w http.ResponseWriter, r *http.Request, q url.Values, p string) {
	tmpl.Root.ExecuteTemplate(w, "root", nil)
}
	
func
ViewPost(w http.ResponseWriter, r *http.Request, q url.Values, p string){
	id, _ := strconv.Atoi(p)
	pst, err := post.GetById(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	buf := md.Process([]byte(pst.Content))
	pst.Content = template.HTML(sanitize.Sanitize(buf))

	tmpl.ViewPost.ExecuteTemplate(w, "viewpost", pst)
}

func
TypePost(w http.ResponseWriter, r *http.Request, q url.Values, p string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.TypePost.ExecuteTemplate(w, "typepost", nil)
}

func
TypePostHndl(w http.ResponseWriter, r *http.Request, q url.Values, p string) {
	//post.WriteNew()
	http.Redirect(w, r, urlpath.RootPrefix, http.StatusFound)
}

func
PostTest(w http.ResponseWriter, r *http.Request, q url.Values, p string) {
	switch r.Method {
	case "GET" :
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.PostTest.ExecuteTemplate(w, "posttest", nil)
	case "POST" :
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Method: %s\n", r.Method)
		r.ParseForm()
		fmt.Fprintf(w, "Post data:\n%v\n", r.PostForm)
	}
}

func
GetTest(w http.ResponseWriter, r *http.Request, q url.Values, p string){
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

