/**
 * Example for 4.1
 * https://github.com/astaxie/build-web-application-with-golang
 *
 * @desc 学习文件上传。
 *
 * @author Neo Lin
 * @email
 * @date 2019/9/3
 */

package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if "GET" == r.Method {
		// 生成Server端验证的token，防止多次提交。
		st := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(st, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Printf("生成token：%s\n", token)

		// 执行模板，并传递token。
		tpl, _ := template.ParseFiles("upload.tpl")
		tpl.Execute(w, token)
	} else {
		// 根据Form的enctype属性调用
		r.ParseMultipartForm(32 << 20)
		// 防止多次提交
		token := r.Form.Get("token") // token 为隐藏input的名称
		if "" != token {
			// 验证token是否合法
			fmt.Printf("网页token：%s\n", token)
		} else {
			// 非法的数据提交
		}
		// get file
		f, h, err := r.FormFile("upfile") // 参数名为Form中属性为file的input的名称
		if nil != err {
			fmt.Fprintf(w, "上传文件错误")
			return
		}
		defer f.Close()

		fmt.Fprintf(w, "%v\n", h.Header)
		file, err := os.OpenFile("./upload/"+h.Filename, os.O_WRONLY|os.O_CREATE, 0666) // upload目录存在
		if nil != err {
			fmt.Println(err)
			return
		}
		defer file.Close()
		io.Copy(file, f)
	}
}

func main() {
	http.HandleFunc("/up", upload)

	http.ListenAndServe(":80", nil)
}
