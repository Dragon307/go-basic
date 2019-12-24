package main

import (
	"context"
	"fmt"
	"golang.org/x/net/webdav"
	"log"
	"net/http"
	"os"
)

// 浏览器视图
func main() {
	fs := &webdav.Handler{
		FileSystem: webdav.Dir("."),
		LockSystem: webdav.NewMemLS(),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" && handleDirList(fs.FileSystem, w, req) {
			return
		}
		fs.ServeHTTP(w, req)
	})

	http.ListenAndServe(":8088", nil)
}

func handleDirList(fs webdav.FileSystem, w http.ResponseWriter, req *http.Request) bool {
	ctx := context.Background()

	f, err := fs.OpenFile(ctx, req.URL.Path, os.O_RDONLY, 0)
	if err != nil {
		return false
	}
	defer f.Close()

	if fi, _ := f.Stat(); fi != nil && !fi.IsDir() {
		return false
	}

	dirs, err := f.Readdir(-1)
	if err != nil {
		log.Print(w, "Error reading directory", http.StatusInternalServerError)
		return false
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<pre>\n")
	for _, d := range dirs {
		name := d.Name()
		if d.IsDir() {
			name += "/"
		}
		fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", name, name)
	}
	fmt.Fprintf(w, "</pre>\n")
	return true
}

// https://chai2010.cn/post/2018/webdav/

/*
WebDAV （Web-based Distributed Authoring and Versioning） 是一种基于 HTTP 1.1协议的通信协议。
它扩展了HTTP 1.1，在GET、POST、HEAD等几个HTTP标准方法以外添加了一些新的方法，使应用程序可对Web Server直接读写，
并支持写文件锁定(Locking)及解锁(Unlock)，还可以支持文件的版本控制。
*/