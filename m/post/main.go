package post

import(
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/k1574/mojosa.press/m/path"
	"github.com/k1574/mojosa.press/m/tempconfig"
	//"html/template"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var(
	TmpCfg tempconfig.TempConfig
)

type Post struct {
	Content string
	Title string
	Hash string
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
WriteNew(p Post) (int, error) {
	var err error

	err = tempconfig.IncrementLastPostId()
	if err != nil {
		fmt.Println("yes")
		return 0, err
	}

	err = WriteById(p, tempconfig.LastPostId())
	if err != nil {
		return 0, err
	}

	return tempconfig.LastPostId(), nil
}

func
Hash(pass string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(b), err
}

func
CheckHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
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

func
Exists(id int) bool {
	_, err := os.Stat(path.PostById(id))
	if err != nil {
		return false
	}

	return true
}
