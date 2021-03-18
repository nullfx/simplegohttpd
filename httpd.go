package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

const notFound string = `
		<html>
			<head>
				<title>Not Found</title>
			</head>
			<body>
				<h1>404 Not Found</h1>
				<div>The resource &quot;%[1]v&quot; was not found on this server</div>
			</body>
		</html>
	`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file := path.Join(".", r.URL.Path)
		_, err := os.Stat(file)
		if err == nil {
			f, err := os.Open(file)
			defer f.Close()
			if err == nil {
				bytes, err := ioutil.ReadFile(file)
				if err == nil {
					ext := path.Ext(file)
					ct := mime.TypeByExtension(ext)
					w.Header().Add("Content-Type", ct)
					fmt.Printf("[%v] [%v] [%v] - %v\n", time.Now().Format(time.RFC3339), r.Method, "(unknown)", file)
					w.Write(bytes)
					return
				}
			}
		}
		nf := []byte(fmt.Sprintf(notFound, strings.TrimLeft(file, ".")))
		fmt.Printf("[%v] [%v] [%v] - %v\n", time.Now().Format(time.RFC3339), r.Method, "(unknown)", file)
		w.Write(nf)
	})

	err := http.ListenAndServe(":82", nil)

	if err != nil {
		log.Fatal(err)
	}
}
