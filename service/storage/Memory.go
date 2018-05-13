package storage

import "github.com/oswaldl/cell/service/storage/surfacemap"

/*
	记忆，建立在HardDisk之上
	有缓存
  */
type Memory struct {
	// 长久记录
	HardDisk HardDisk

	// 一级缓存区：最多预加载9个, 读操作
	CurrentLoadSurfaces []surfacemap.Content

	/*
		二级修改暂存区，也叫替换区,
		在写操作时，先把文件读取缓存起来，然后替换hardDisk文件，如此避免中间态的情况下，读取有脏数据
		在读操作时，除了优先读取CurrentLoadSurfaces，还会优先读取SecondModifyingSurfaces

	    append(s, 0)
	 */
	SecondModifyingSurfaces []surfacemap.Content
}
