package hndl

import(
	//"html/template"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"github.com/k1574/mojosa.press/m/urlpath"
	"github.com/k1574/mojosa.press/m/post"
	"github.com/k1574/mojosa.press/m/tmpl"
	"github.com/k1574/mojosa.press/m/tempconfig"
	"github.com/k1574/mojosa.press/m/pp"
	//"io/ioutil"
	"fmt"
	"reflect"
)

type HndlArg struct {
	q url.Values
	p string
}

type Handler func(http.ResponseWriter, *http.Request, HndlArg)
type FuncDefine struct {
	Pref, Re string
	Fn Handler
}

var(
	PageSize int = 15
	Defs = []FuncDefine {
		{urlpath.RootPrefix, "^$", Root},
		{urlpath.ViewPostPrefix, "^[0-9]*$", ViewPost},
		{urlpath.TypePostPrefix, "^[0-9]*$", TypePost},
		{urlpath.GetTestPrefix, "", GetTest},
		{urlpath.PostTestPrefix, "", PostTest},
		{urlpath.ListPostsPrefix, "^[0-9]+$", ListPosts},
	}
)

func clamp(a, b, c int) int {
	if b < a {
		return a
	}
	if b > c {
		return c
	}
	return b
}

func MakeHttpHandleFunc(pref string, re *regexp.Regexp, fn Handler) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
	var(
		a HndlArg
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

func Root(w http.ResponseWriter, r *http.Request, a HndlArg) {
	//tmpl.Root.ExecuteTemplate(w, "root", nil)
	http.Redirect(w, r,
		urlpath.ViewPostPrefix+"0",
		http.StatusFound)
}


func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func ListPosts(w http.ResponseWriter, r *http.Request, a HndlArg) {
	switch(r.Method){
	case "GET" :
		var htmlPosts []pp.PostHTML
		lpId := tempconfig.TmpCfg.LastPostId
		pageId, _ := strconv.Atoi(a.p)
		firstId := lpId - (pageId+1)*(PageSize) + 1
		lastId := firstId + PageSize - 1

		if firstId <= -PageSize  {
			http.NotFound(w, r)
			return
		}

		firstId = clamp(1, firstId, lpId)
		lastId = clamp(1, lastId, lpId)

		posts, err := post.GetList(firstId, lastId)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		ReverseSlice(posts)

		for i, v := range posts {
			htmlPosts = append(htmlPosts, pp.Preprocess(v, lastId-i))
		}

		tmpl.Execute(w, "listposts", struct{
			Posts []pp.PostHTML
			Page int
			FirstId int
			LastId int
			}{ htmlPosts, pageId, firstId, lastId})
	case "POST" :
		http.NotFound(w, r)
		return
	}
}
	
func ViewPost(w http.ResponseWriter, r *http.Request, a HndlArg){
	id, _ := strconv.Atoi(a.p)
	pst, err := post.GetById(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl.Execute(w, "viewpost", struct{
			Id string
			Post pp.PostHTML
		}{a.p, pp.Preprocess(pst, id)})
}

/* Both edit and write new. */
func TypePost(w http.ResponseWriter, r *http.Request, a HndlArg) {
	switch r.Method {
	case "GET" :
		var pst post.Post
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if a.p == "" {
			tmpl.Execute(w, "typepost", struct{}{})
			return
		}

		id, _ := strconv.Atoi(a.p)
		pst, err := post.GetById(id)
		pst.Hash = ""
		if err != nil {
			http.NotFound(w, r)
			return
		}

		tmpl.Execute(w, "typepost", struct{
				Post post.Post
				Id int}{pst, id})
	case "POST" :
		pass := r.Form.Get("pass")
		hsh, _ := post.Hash(pass)
		pst := post.Post{
			Content : r.Form.Get("text"),
			Title : r.Form.Get("title"),
			Desc : r.Form.Get("desc"),
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

func PostTest(w http.ResponseWriter, r *http.Request, a HndlArg) {
	switch r.Method {
	case "GET" :
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, "posttest", nil)
	case "POST" :
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Method: %s\n", r.Method)
		fmt.Fprintf(w, "Post data:\n%v\n", r.PostForm)
	}
}

func GetTest(w http.ResponseWriter, r *http.Request, a HndlArg){
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


