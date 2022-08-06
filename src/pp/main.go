package pp

import (
	"github.com/surdeus/mojosa.press/src/md"
	"github.com/surdeus/mojosa.press/src/sanitize"
	"github.com/surdeus/mojosa.press/src/post"
	"html/template"
	"strconv"
)

type PostHTML struct {
	Content, Title, Desc template.HTML
	WebTitle, Id string
}

/* Preprocess post for output in templates. */
func Preprocess(p post.Post, id int) PostHTML {
	var pret PostHTML

	buf := md.Process([]byte(p.Desc))
	pret.Desc = template.HTML(string(sanitize.Sanitize(buf)))

	pret.WebTitle = string(sanitize.Sanitize([]byte(p.Title)))
	
	buf = md.Process([]byte(p.Title))
	pret.Title = template.HTML(string(sanitize.Sanitize(buf)))

	buf = md.Process([]byte(p.Content))
	pret.Content = template.HTML(string(sanitize.Sanitize(buf)))

	pret.Id = strconv.Itoa(id)

	return pret
}
