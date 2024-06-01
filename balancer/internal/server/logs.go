package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Logs(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("logs.txt")
	if err != nil {
		fmt.Printf("os.Open: %v", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Printf("file.Close: %v", err)
		}
	}()

	data, err := io.ReadAll(file)

	_, err = w.Write(data)
	if err != nil {
		fmt.Printf("w.Write: %v\n", err)
	}

	//t, _ := template.ParseFiles("static/template/index.html")
	//err := t.Execute(w, "")
	//if err != nil {
	//	log.Printf("t.Execute: %v", err)
	//}
}
