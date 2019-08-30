/**
 * Example for 3.4
 * https://github.com/astaxie/build-web-application-with-golang
 *
 * @description 重点学习http包中的ServeMux内容
 *
 * @author Neo Lin
 * @email
 * @date 2019/8/30
 */

package main

import (
	"fmt"
	"net/http"
)

// MyMux 实现Handler接口
type MyMux struct{}

// ServeHTTP 实现Handler接口的方法
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}
	http.NotFound(w, r)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "自定义的路由处理。")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":80", mux)
}
