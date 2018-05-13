package surfacemap

// surfacemap的内容，最终会转为json格式返回上层
type Content struct {
	// id
	Id string

	// regionX, 从1开始的正数
	RegionX int

	// regionY, 从1开始的正数
	RegionY int

	// width
	Width int

	// height
	Height int

	// terrain
	Terrain [][]float64
}
