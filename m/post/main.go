package post

import(
	"fs"
	"log"
	"ioutil"
	"strconv"
	"encoding/json"
	"mojosa.press/m/path"
)

type Post struct {
	Username string
	Content string
}

var(
	lastId int
)

func
init(){
	buf, err = ioutil.ReadFile(path.LastPostIdFile)
	if err != nil {
			ioutil.WriteFile(path.LastPostIdFile, []byte("0"), 0755)
	}

	lastId, err = strconv.Atoi(string(buf))
	if err != nil {
		log.Fatal(err)
	}
}

func
ById(id int) Post, error {
	var p Post

	buf := ioutil.ReadFile(path.PostById(id))
	err := json.Unmarshal(buf, &p)
	if err != nil {
		nil, err
	}

	return p, nil
}

func
WriteNew(p Post) error {
	return nil
}

func
writeById(p Post, id int) error {
	return nil
}
