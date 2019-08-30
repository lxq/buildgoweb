/**
 * Example for 4.1
 * https://github.com/astaxie/build-web-application-with-golang
 *
 * @description 重点学习Form内容提交，Form数据通过post方法传输。
 *
 * @author Neo Lin
 * @email
 * @date 2019/8/30
 */

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析url请求信息，对于POST则解析 r.Body
	r.ParseForm()
	// 以下关于FORM的调用基于先调用 r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Paht: ", r.URL.Path)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Printf("%s = %s \n", k, strings.Join(v, ", "))
	}

	fmt.Fprintf(w, "已解析URL请求参数。")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("login.gtpl") // ignore error
		tmpl.Execute(w, nil)
	} else {
		// POST数据必须先解析，否则Form取值为空。
		r.ParseForm()
		fmt.Printf("%s 密码是： %s\n", r.Form["username"], r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":80", nil)
	if nil != err {
		log.Fatal("http.ListenAndServe: ", err)
	}
}
