package path

import(
	"strconv"
)

var(
	Data = "dat"
	Tmpl = Data+"/tmpl"
	Temp = Data+"/tmp"
	Pub = Data+"/pub"
	Static = Pub+"/s"
	Database = Pub+"/db"
	Post = Pub+"/p"
	LastPostIdFile = Temp+"/lastpost"
)

func
PostById(id int) string {
	return Post+"/"+strconv.Itoa(id)
}
