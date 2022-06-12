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

type FuncArg struct {
	q url.Values
	p string
}

type Handler func(http.ResponseWriter, *http.Request, FuncArg)
type FuncDefine struct {
	Pref, Re string
	Fn Handler
}

var(
	Defs = []FuncDefine {
		{urlpath.RootPrefix, "^$", Root},
		{urlpath.ViewPostPrefix, "^[0-9]*$", ViewPost},
		{urlpath.TypePostPrefix, "^[0-9]*$", TypePost},
		{urlpath.GetTestPrefix, "", GetTest},
		{urlpath.PostTestPrefix, "", PostTest},
	}
)


func
MakeHttpHandleFunc(pref string, re *regexp.Regexp, fn Handler) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
	var(
		a FuncArg
		e error
	)

	a.p = r.URL.Path[len(pref):]
	if !urlpath.Validify(a.p, re) {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET" :
		a.q, e = url.ParseQuery(r.URL.RawQuery)
	case "POST" :
		r.ParseForm()
	}

	if e != nil {
	}
	
	fn(w, r, a)
} }

func
Root(w http.ResponseWriter, r *http.Request, a FuncArg) {
	//tmpl.Root.ExecuteTemplate(w, "root", nil)
	http.Redirect(w, r,
		urlpath.ViewPostPrefix+"0",
		http.StatusFound)
}
	
func
ViewPost(w http.ResponseWriter, r *http.Request, a FuncArg){
	id, _ := strconv.Atoi(a.p)
	pst, err := post.GetById(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	buf := md.Process([]byte(pst.Content))
	pst.Content = string(sanitize.Sanitize(buf))

	tmpl.ViewPost.ExecuteTemplate(w, "viewpost", struct{
			Id string
			Post post.Post
			HTMLContent template.HTML
		}{a.p, pst, template.HTML(pst.Content)})
}

/* Both edit and write new. */
func
TypePost(w http.ResponseWriter, r *http.Request, a FuncArg) {
	switch r.Method {
	case "GET" :
		var pst post.Post
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if a.p == "" {
			tmpl.TypePost.ExecuteTemplate(w, "typepost", struct{Post post.Post}{post.Post{}})
			return
		}

		id, _ := strconv.Atoi(a.p)
		pst, err := post.GetById(id)
		pst.Hash = ""
		if err != nil {
			http.NotFound(w, r)
			return
		}

		tmpl.TypePost.ExecuteTemplate(w, "typepost", struct{
				Post post.Post
				Id int}{pst, id})
	case "POST" :
		pass := r.Form.Get("pass")
		hsh, _ := post.Hash(pass)
		pst := post.Post{
			Content : r.Form.Get("text"),
			Title : r.Form.Get("title"),
			Hash : hsh}
		if a.p == "" { /* Creating new post if the path is empty. */
			id, _ := post.WriteNew(pst)
			ids := strconv.Itoa(id)
			http.Redirect(w, r,
				urlpath.ViewPostPrefix+ids,
				http.StatusFound)
		} else {
			id, _ := strconv.Atoi(a.p)
			if !post.CheckPass(pass, id) {
				http.NotFound(w, r)
				return
			}
			post.WriteById(pst, id)
			http.Redirect(w, r,
				urlpath.ViewPostPrefix+a.p,
			http.StatusFound)
		}
	}
}

func
PostTest(w http.ResponseWriter, r *http.Request, a FuncArg) {
	switch r.Method {
	case "GET" :
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.PostTest.ExecuteTemplate(w, "posttest", nil)
	case "POST" :
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Method: %s\n", r.Method)
		fmt.Fprintf(w, "Post data:\n%v\n", r.PostForm)
	}
}

func
GetTest(w http.ResponseWriter, r *http.Request, a FuncArg){
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Path: '%s'\nRawQuery:'%s'\n", r.URL.Path, r.URL.RawQuery)
	fmt.Fprintf(w, "a.p: '%s'\n", a.p)
	fmt.Fprintf(w, "a.q:\n")
	for k, v := range a.q {
		fmt.Fprintf(w, "\t'%s':\n", k)
		for _, s := range v {
			fmt.Fprintf(w, "\t\t'%s'\n", s)
		}
	}

}

