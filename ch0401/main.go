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
	"regexp"
	"strconv"
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

// 验证Client端输入
func validate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, _ := template.ParseFiles("validate.gtpl")
		tpl.Execute(w, nil)
	} else {
		r.ParseForm()

		msg := ""
		// 必填
		if len(r.Form["username"]) == 0 {
			msg += "用户名必填，不能为空！\n"
		}
		// 数字
		_, err := strconv.Atoi(r.Form.Get("age"))
		if nil != err {
			msg += "年龄必须填写数字。\n"
		}
		// 中文
		if m, _ := regexp.MatchString(`^\\p{Han}+$`, r.Form.Get("realname")); !m {
			msg += "名字必做是中文！\n"
		}
		// 英文
		if m, _ := regexp.MatchString(`^[a-zA-Z]+$`, r.Form.Get("english")); !m {
			msg += "英文名必须是英文字母.\n"
		}
		// email
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, r.Form.Get("email")); !m {
			msg += "Email格式不正确。\n"
		}
		// 下拉菜单验证
		fruits := []string{"apple", "pear", "banana"}
		tmp := r.Form.Get("fruit")
		tag := false
		for _, e := range fruits {
			if tmp == e {
				tag = true
				break
			}
		}
		if !tag {
			msg += "水果选择不正确。\n"
		}
		// write to browser
		fmt.Fprintf(w, msg)
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/val", validate)

	err := http.ListenAndServe(":80", nil)
	if nil != err {
		log.Fatal("http.ListenAndServe: ", err)
	}
}
