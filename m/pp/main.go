package pp

import (
	"github.com/k1574/mojosa.press/m/md"
	"github.com/k1574/mojosa.press/m/sanitize"
	"github.com/k1574/mojosa.press/m/post"
	"html/template"
)

/* Preprocess post for output in templates. */
func Preprocess(p post.Post) post.PostHTML {
	var pret post.PostHTML

	buf := md.Process([]byte(p.Desc))
	pret.Desc = template.HTML(string(sanitize.Sanitize(buf)))

	buf = md.Process([]byte(p.Title))
	pret.Title = template.HTML(string(sanitize.Sanitize(buf)))

	buf = md.Process([]byte(p.Content))
	pret.Content = template.HTML(string(sanitize.Sanitize(buf)))

	return pret
}
