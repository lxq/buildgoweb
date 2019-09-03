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
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if "GET" == r.Method {
		tpl, _ := template.ParseFiles("upload.tpl")
		tpl.Execute(w, nil)
	} else {
		// 根据Form的enctype属性调用
		r.ParseMultipartForm(32 << 20)
		// get file
		f, h, err := r.FormFile("upfile") // 参数名为Form中属性为file的input的名称
		if nil != err {
			fmt.Fprintf(w, "上传文件错误")
			return
		}
		defer f.Close()

		fmt.Fprintf(w, "文件头信息：%s\n", h.Header)
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
