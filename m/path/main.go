package path

import(
	"strconv"
)

var(
	Data = "dat"
	Tmpl = "tmpl"
	TmplGen = Tmpl+"/gen"
	TmplSep = Tmpl+"/sep"
	// Temp = Data+"/tmp"
	Static = Data+"/s"
	Post = Data+"/p"
	LastPostIdFile = Post+"/last"
)

func
PostById(id int) string {
	return Post+"/"+strconv.Itoa(id)
}
