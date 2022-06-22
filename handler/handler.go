package handler

import (
	"encoding/json"
	"fileStoreServer/meta"
	"fileStoreServer/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传页面
		file, err := ioutil.ReadFile("./static/view/upload.html")
		if err != nil {
			io.WriteString(w, "接口错误")
			return
		}
		io.WriteString(w, string(file))
	} else if r.Method == "POST" {
		//接受文件流並保存到本地文件中
		//从请求里拿到file文件
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Failed to get data,err=%s", err.Error())
			return
		}
		defer file.Close()
		fileMeta := meta.FileMeta{
			FileName: header.Filename,
			Location: "F:/wangpanProject/tmp" + header.Filename,
			UploadAt: time.Now().Format("2022-01-02 15.03"),
		}

		newFile, err := os.Create(fileMeta.Location)
		//fmt.Println(newFile)
		if err != nil {
			fmt.Println("Failed to create newFile", err.Error())
			return
		}
		defer newFile.Close()
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("copy 失败，err=%s", err.Error())
			return
		}
		newFile.Seek(0, 0)
		sha1 := util.FileSha1(newFile)
		fileMeta.FileSha1 = sha1
		meta.UpdateFileMetaDb(fileMeta)
		//meta.UpdateFileMeta(fileMeta)
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}

}
func UploadScHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload has finshed")
}

//获取客户端的相对应的文件
func GetFileHander(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form["filehash"][0]
	metas := meta.GetFileMetas(fileHash)
	marshal, err := json.Marshal(metas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(marshal)
}
