package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
	"path/filepath"
	"os"
)

var CellConfig struct {
	LogLevel    string
	LogPath     string
	LogFlag  int
	SurfaceRootPath      string
	MaxConnNum  int
	ConsolePort int
	SkipHardDisk bool
}

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("%v", err)
	}
	log.Debug(dir)

	data, err := ioutil.ReadFile("conf/config.json")
	if err != nil {
		log.Fatal("%v", err)
	}else{
		log.Debug("conf file loaded")
	}
	err = json.Unmarshal(data, &CellConfig)
	if err != nil {
		log.Fatal("%v", err)
	}else{
		log.Debug("conf data loaded")
	}


}
