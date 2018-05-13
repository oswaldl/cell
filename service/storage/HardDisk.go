package storage

// 硬盘
type HardDisk struct {
	// 根目录
	DiskRoot string

	// 磁盘占用上线设置, 单位 G, 目前未使用
	// MaxOccupy int

	/**
	 *	gorebill's Game Surface map file
	 *  one region pre file
	 */
	SurfaceMapFiles []string
}
