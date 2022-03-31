package post

import(
	"log"
	"io/ioutil"
	"strconv"
	"encoding/json"
	"mojosa/press/m/path"
	"html/template"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Post struct {
	Content template.HTML
	Title string
	Hash string
}

var(
	lastId int
	InitId = 0
)

func
init(){
	buf, err := ioutil.ReadFile(path.LastPostIdFile)
	if err != nil {
		ioutil.WriteFile(path.LastPostIdFile, []byte(string(InitId)), 0644)
		buf = []byte("0")
	}

	lastId, err = strconv.Atoi(string(buf))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lastId)
}

func
GetById(id int) (Post, error) {
	var p Post

	buf, err := ioutil.ReadFile(path.PostById(id))
	if err != nil {
		return Post{}, err
	}

	err = json.Unmarshal(buf, &p)
	if err != nil {
		return Post{}, err
	}

	return p, nil
}

func
incrementLastId() error {
	lastId++
	ioutil.WriteFile(path.LastPostIdFile, []byte(string(lastId)), 0644)
	return nil
}

func
WriteNew(p Post) error {
	var err error

	err = incrementLastId()
	if err != nil {
		return err
	}

	err = WriteById(p, lastId)
	if err != nil {
		return err
	}

	return nil
}

func
Hash(pass string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(b), err
}

func
CheckHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err != nil
}

func
CheckPass(pass string, id int) bool {
	p, err := GetById(id)
	if err != nil {
		return false
	}
	
	return CheckHash(pass, p.Hash)
}


func
WriteById(p Post, id int) error {
	var err error
	j, err := json.Marshal(p)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.PostById(id), j, 0755)
	if err != nil {
		return err
	}

	return nil
}
