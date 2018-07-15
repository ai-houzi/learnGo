package main

import (
	"net/http"
	"learnGo/errhandler/filelisting"
	"os"
	_ "net/http/pprof"
	"log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errorWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)

		if err != nil {
			log.Printf("error occureed handing request :%s", err.Error())
		}

		code := http.StatusOK
		if err != nil {
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound //文件未找到
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errorWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":7788", nil)

	if err != nil {
		panic(err)
	}
}
