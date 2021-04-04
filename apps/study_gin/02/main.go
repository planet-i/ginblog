package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}) //注册一个路由及这个路由的handler到DefaultServeMux中

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("start http server fail:", err)
	}
}
