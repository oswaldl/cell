## 配置文件实例
{
  "LogLevel": "debug",
  "LogPath": "/Users/oswaldl/temp/log",
  "LogFlag": 3,
  "SurfaceRootPath": "/Users/oswaldl/temp/surfaceMap",
  "MaxConnNum": 200,
  "ConsolePort": 8887,
  "SkipHardDisk": true
}

## 操作

### 初始化
http://localhost:8888/regionTerrain/init
如果不需要硬盘初始化，设置配置文件为
"SkipHardDisk": true

### 查询
http://localhost:8887/regionTerrain/1/1
第一个1是regionX
第二个1是regionY

### 新增
post请求
http://localhost:8887/regionTerrain
body数据示例：
{"id":"region_1_1","regionX":1,"regionY":1,"width":100,"height":100,"terrain":[[0,0.1284603349415927,0.2412596407166909,]]}

### 查询缓存
一级
http://localhost:8888/regionTerrain/showFirstCache
二级
http://localhost:8887/regionTerrain/showSecondCache


