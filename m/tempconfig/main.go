package tempconfig

import(
	"encoding/json"
	"os"
	"io/fs"
	"mojosa/press/m/path"
)

type TempConfig struct {
	LastPostId int
}

var(
	TmpCfg TempConfig
)

func
Read(p string) (TempConfig, error) {
	var cfg TempConfig
	buf, err := os.ReadFile(p)
	if err != nil { return TempConfig{}, err }

	err = json.Unmarshal(buf, &cfg)
	if err != nil { return TempConfig{}, err }
	return cfg, nil
}

func
Write(p string, cfg TempConfig, mode fs.FileMode) error {
	jcfg, err := json.Marshal(cfg)
	if err != nil { return err }

	err = os.WriteFile(p, jcfg, mode)
	if err != nil { return err }

	return nil
}

func
Update() error {
	return Write(path.TempConfig, TmpCfg, 0644)
}

func
IncrementLastPostId() error {
	TmpCfg.LastPostId++
	return Update()
}

func
LastPostId() int {
	return TmpCfg.LastPostId
}
