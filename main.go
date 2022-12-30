package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "cherry blog"
	indexData.Desc = "现在是入门教程"
	t := template.New("index.html")
	// 1. 拿到当前路径
	path, _ := os.Getwd()
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	pagination := path + "/template/layout/pagination.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post.html"
	t, _ = t.ParseFiles(path+"/11/index.html", home, header, footer, personal, post, pagination)
	t.Execute(w, indexData)
}

func main() {
	// 程序入口，一个项目只能由一个入口
	// web程序  http协议
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/index.html", http.HandlerFunc(indexHtml))
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
