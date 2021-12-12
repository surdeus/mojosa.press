package path

var(
	Data = "dat"
	Tmpl = Data+"/tmpl"
	Pub = Data+"/pub"
	Static = Pub+"/s"
	Database = Pub+"/db"
	Post = Pub+"/p"
)

func
PostById(id int) string {
	return Post+"/"+string(id)
}
