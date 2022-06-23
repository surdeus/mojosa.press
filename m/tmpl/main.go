package tmpl

import(
	"github.com/k1574/mojosa.press/m/path"
	"html/template"
	"io/ioutil"
	"io"
	"reflect"
	//"fmt"
)

var(
	Templates map[string] *template.Template
)

func Execute(w io.Writer, t string, v interface{}) {
	Templates[t].ExecuteTemplate(w, t, v)
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
