package storage

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"github.com/oswaldl/cell/service/storage/surfacemap"
	"github.com/oswaldl/cell/log"
)

// 初始化
func (hardDisk *HardDisk) InitRegions() {

	// 初始化文件夹
	os.MkdirAll(hardDisk.DiskRoot, 0777)

	// 2个循环，从1-100，开始把现在的地址数据读过来，并写到文件
	for x := 1; x <= 100; x++ {
		for y := 1; y <= 100; y++ {
			url := fmt.Sprintf("http://game.obwork.site/getRegionTerrain?regionX=%d&regionY=%d", x, y)

			resp, err := http.Get(url)

			if err == nil {
				writeRegionFile(resp, hardDisk, x, y)
				resp.Body.Close()
			} else {
				log.Debug("fetch data by http error", err)
			}
		}
	}

}

func writeRegionFile(resp *http.Response, hardDisk *HardDisk, x int, y int) {

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		// save file: region_x_y.d
		writeDisk(body, hardDisk.DiskRoot, x, y)
	} else {
		log.Debug("read error", err)
	}
}

func (hardDisk *HardDisk) WriteRegion(surface *surfacemap.Content) {
	jsonStr, _ := json.Marshal(surface)
	writeDisk(jsonStr, hardDisk.DiskRoot, surface.RegionX, surface.RegionY)
}

// save file: region_x_y.d
func writeDisk(data []byte, diskRoot string, x int, y int) {

	fileName := fmt.Sprintf("%s/region_%d_%d.d", diskRoot, x, y)
	log.Debug("trying to write disk:", fileName)
	fout, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		// create one
		fout, err := os.Create(fileName)
		if err != nil {
			log.Debug("error while create file:%s, error is %v", fileName, err)
			return
		} else {
			fout.Write(data)
			fout.Close()
		}
	} else {
		fout.Write(data)
		fout.Close()
	}

}

// 从HardDisk中读取信息
func (hardDisk *HardDisk) ReadRegion(x int, y int) (surfacemap.Content) {
	fileName := fmt.Sprintf("%s/region_%d_%d.d", hardDisk.DiskRoot, x, y)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Debug("%v", err)
	} else {
		log.Debug("conf file loaded")
	}

	content := surfacemap.Content{}
	err = json.Unmarshal(data, &content)
	return content
}
