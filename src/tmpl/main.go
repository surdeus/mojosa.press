package tmpl

import(
	"github.com/surdeus/mojosa.press/src/path"
	"html/template"
	"io/ioutil"
	"io"
	"reflect"
	"log"
	//"fmt"
)

var(
	Templates map[string] *template.Template
)

func Execute(w io.Writer, t string, v interface{}) {
	err := Templates[t].ExecuteTemplate(w, t, v)
	if err != nil {
		log.Println(err)
	}
}

func ParseSepTemplates() map[string] *template.Template {
	ret := make(map[string] *template.Template)

	files, err := ioutil.ReadDir(path.TmplSep)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		ret[f.Name()] =
			MustParse(f.Name())
	}

	return ret
}

func MustParse(t string) *template.Template {
	lfs := []string{path.TmplSep+"/"+t}

	files, _ := ioutil.ReadDir(path.TmplGen)
	for _, f := range files {
		lfs = append(lfs, path.TmplGen+"/"+f.Name())
	}

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"hasField" : hasField,
			"sum" : sum,
			"neg" : neg,
		}).ParseFiles(lfs...)
	if err != nil {
		panic(err)
	}

	return tmpl
}

func
hasField(v interface{}, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}

func
sum(a, b int) int {
	return a + b
}

func neg(a int) int {
	return -a
}

