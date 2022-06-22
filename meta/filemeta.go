package meta

//定义文件的元信息结构体
//如何进行文件的存储呢？
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

//全局的map存储方式
var fileMetas map[string]FileMeta

//文件元信息初始化
func init() {
	fileMetas = make(map[string]FileMeta)
}

//更新filemeeta元数据
func UpdateFileMeta(meta FileMeta) {
	fileMetas[meta.FileSha1] = meta
}

//更新文件元信息到到mysql数据库中
func UpdateFileMetaDb(meta FileMeta) bool {
	//忘记怎么去写了？？？
	return true
	//filehash string, filename string, filesize int64, fileArr string
}

//根据sha1获取文件原信息
func GetFileMetas(sha1 string) FileMeta {
	meta := fileMetas[sha1]
	return meta
}
