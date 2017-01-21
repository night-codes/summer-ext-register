package main

import (
	"fmt"
	"gopkg.in/night-codes/summer.v1"
)

type (
	obj map[string]interface{}
	arr []interface{}
)

var (
	panel = summer.Create(summer.Settings{
		Title:             "Training project implements external registration in the Summer",
		Port:              8888,
		AuthPrefix:        "cabinet",
		AuthSalt:          "71A8F-117-t57",
		DefaultPage:       "/home",
		UsersCollection:   "webmasters", // collection for panel's users
		NotifyCollection:  "webmastersNotify",
		DisableFirstStart: true,
		Path:              "", // application path
		DBName:            "external-register-project",
		Views:             "templates/main",
		ViewsDoT:          "templates/dot", // doT.js templates
		Files:             "files",         // static files dir
	})
)

func main() {
	fmt.Println("Application started at http://localhost:8888/")
	summer.Wait()
}
