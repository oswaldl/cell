package storage

import (
	"github.com/oswaldl/cell/service/storage/surfacemap"
	"github.com/oswaldl/cell/conf"
	"github.com/oswaldl/cell/log"
)

// 初始化
func (memory *Memory) InitRegions() {
	// 初始化硬盘数据
	if (!conf.CellConfig.SkipHardDisk) {
		memory.HardDisk.InitRegions()
	} else {
		log.Debug("skip network reload to disk")
	}

	// 初始化缓存区
	loadRegionOnDisk(memory, 1, 1)

}

// 从记忆中读取信息
func (memory *Memory) ReadRegion(regionX int, regionY int) surfacemap.Content {

	// 首先读一级缓存区
	value1, result1 := readRegionFromFirstCache(memory, regionX, regionY)
	if (result1) {
		return *value1;
	}

	// 其次读二级缓存区
	value2, result2 := readRegionFromSecondCache(memory, regionX, regionY)
	if (result2) {
		return *value2;
	}

	// 未命中，从硬盘读取
	return loadRegionOnDisk(memory, regionX, regionY)
}

// 插入一条记录
func (memory *Memory) AddRegion(surface *surfacemap.Content) {

	// 先磁盘中到数据预加载到二级缓存
	diskSurface := memory.HardDisk.ReadRegion(surface.RegionX, surface.RegionY)

	if (diskSurface.Terrain != nil) {
		log.Debug("add new region data:", diskSurface.Terrain)
		memory.SecondModifyingSurfaces = append(memory.SecondModifyingSurfaces, diskSurface)
	}

	// 再覆盖磁盘
	memory.HardDisk.WriteRegion(surface)

	// 把二级缓存数据移除掉, 失效了
	go resetCacheRegion(memory, surface)

}
func resetCacheRegion(memory *Memory, content *surfacemap.Content) {
	resetFirstCacheRegion(memory, content)
	removeSecondCacheRegion(memory, content)
}
func resetFirstCacheRegion(memory *Memory, content *surfacemap.Content) {
	for index, value := range memory.CurrentLoadSurfaces {
		if (content.RegionX == value.RegionX && content.RegionY == value.RegionY) {
			log.Debug("replace frist cache now:", index, content)
			memory.CurrentLoadSurfaces[index] = *content
			log.Debug("after replace frist cache:", memory.CurrentLoadSurfaces[index])
		}
	}
}

func removeSecondCacheRegion(memory *Memory, content *surfacemap.Content) {
	for index, value := range memory.SecondModifyingSurfaces {
		if (content.RegionX == value.RegionX && content.RegionY == value.RegionY) {
			memory.SecondModifyingSurfaces = remove(memory.SecondModifyingSurfaces, index)
			log.Debug("remove second cache now", index)
			return
		}
	}
}

func remove(s []surfacemap.Content, i int) []surfacemap.Content {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func readRegionFromFirstCache(memory *Memory, regionX int, regionY int) (*surfacemap.Content, bool) {

	return readRegionFromCache(memory.CurrentLoadSurfaces, regionX, regionY)
}

func readRegionFromSecondCache(memory *Memory, regionX int, regionY int) (*surfacemap.Content, bool) {

	return readRegionFromCache(memory.SecondModifyingSurfaces, regionX, regionY)
}

func readRegionFromCache(surfaces []surfacemap.Content, regionX int, regionY int) (*surfacemap.Content, bool) {
	for index, value := range surfaces {
		if (value.RegionX == regionX && value.RegionY == regionY) {
			return &surfaces[index], true
		}
	}

	return nil, false
}

// 触发一次从硬盘加载某个region的操作
func loadRegionOnDisk(memory *Memory, regionX int, regionY int) surfacemap.Content {

	// 读硬盘
	content := memory.HardDisk.ReadRegion(regionX, regionY)

	// 触发异步任务，重现刷新缓存区加载
	go resetCurrentSurface(&content, memory)

	// 返回结果
	return content

}

// 以当前region数据为中心，加载附近的region
func resetCurrentSurface(content *surfacemap.Content, memory *Memory) {
	var tempSurfaces []surfacemap.Content

	// x,y
	tempSurfaces = resetRegion(memory, tempSurfaces, content.RegionX, content.RegionY)

	// x,y-1
	tempSurfaces = resetRegion(memory, tempSurfaces, content.RegionX, content.RegionY-1)

	// x-1,y
	tempSurfaces = resetRegion(memory, tempSurfaces, content.RegionX-1, content.RegionY)

	// x+1,y
	tempSurfaces = resetRegion(memory, tempSurfaces, content.RegionX+1, content.RegionY)

	// x,y+1
	tempSurfaces = resetRegion(memory, tempSurfaces, content.RegionX, content.RegionY+1)

	// x-1,y-1
	tempSurfaces = resetRegion(memory, tempSurfaces, content.RegionX-1, content.RegionY-1)

	// x+1,y-1
	tempSurfaces = resetRegion(memory, tempSurfaces, content.RegionX+1, content.RegionY-1)

	// x-1,y+1
	tempSurfaces = resetRegion(memory, tempSurfaces, content.RegionX-1, content.RegionY+1)

	// x+1, y+1
	tempSurfaces = resetRegion(memory, tempSurfaces, content.RegionX+1, content.RegionY+1)

	memory.CurrentLoadSurfaces = tempSurfaces;
	log.Debug("first cache loaded:", content.RegionX, content.RegionY)
}

func resetRegion(memory *Memory, tempSurfaces []surfacemap.Content, x int, y int) ([]surfacemap.Content) {
	if (isValid(x, y)) {
		newContent := memory.HardDisk.ReadRegion(x, y)
		tempSurfaces = append(tempSurfaces, newContent)
	}

	return tempSurfaces
}

func isValid(x int, y int) bool {
	if (x > 0 && y > 0) {
		return true
	} else {
		return false
	}

}
