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
	TempConfig = Data+"/temp-config.json"
)

func
PostById(id int) string {
	return Post+"/"+strconv.Itoa(id)
}

