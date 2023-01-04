package main

import (
	"cherryGoBlog/config"
	"cherryGoBlog/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func isODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "cherry blog"
	indexData.Desc = "现在是入门教程"
	t := template.New("index.html")

	// 1. 拿到当前路径
	path := config.Cfg.System.CurrentDir
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	pagination := path + "/template/layout/pagination.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println("解析模板出错： ", err)
	}
	// 页面上涉及到的所有数据，必须有定义
	//================================================
	// 假数据
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	//================================================
	t.Execute(w, hr)
}

func main() {
	// 程序入口，一个项目只能由一个入口
	// web程序  http协议
	server := http.Server{
		Addr: "127.0.0.1:8088",
	}
	http.Handle("/index.html", http.HandlerFunc(indexHtml))
	// url请求映射 resource -> public/resource
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
