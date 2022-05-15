package tmpl

import(
	"mojosa/press/m/path"
	"html/template"
	"io/ioutil"
)

var(
	PostTest = MustParse("posttest")
	ViewPost = MustParse("viewpost")
	TypePost = MustParse("typepost")
	EditPost = MustParse("editpost")
	Root = MustParse("root")
)

func
MustParse(t string) *template.Template {
	lfs := []string{path.TmplSep+"/"+t}

	files, _ := ioutil.ReadDir(path.TmplGen)
	for _, f := range files {
		lfs = append(lfs, path.TmplGen+"/"+f.Name())
	}

	return template.Must(template.ParseFiles(lfs...))
}
