package handler

import (
	db "fileStoreServer/db"
	"fileStoreServer/util"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	pwd_salt = "@131#"
)

//	处理用户注册请求
func SiginUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		file, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(file)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	fmt.Println(username, password)
	if len(username) < 3 || len(password) < 5 {
		w.Write([]byte("Invailed parameter"))
		return
	}
	enc_pwd := util.Sha1([]byte(password + pwd_salt))
	fmt.Println(enc_pwd + "11111111111111111111")
	suc := db.UserSignUp(username, enc_pwd)
	fmt.Println(suc)
	fmt.Println("1111111111112222222222222222222222222222")
	//fmt.Println(suc + "222222222222222222222222222222")
	if suc {
		fmt.Println("成功")
		w.Write([]byte("SUCCESS"))
		w.WriteHeader(http.StatusAccepted)
	} else {
		fmt.Println("失败")
		w.Write([]byte("FAILED"))
	}
}
