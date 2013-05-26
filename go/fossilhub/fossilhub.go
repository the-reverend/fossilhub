package main

import "io/ioutil"
import "fmt"
import "strings"
import "strconv"
import "net/http"
import "flag"

func handler(w http.ResponseWriter, r *http.Request, path string) {
	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<header>")
	fmt.Fprintf(w, "<title>Fossil Repositories</title>")
	fmt.Fprintf(w, "</header>")
	fmt.Fprintf(w, "<body>")
	fmt.Fprintf(w, "<h1>Fossil Repositories</h1>")

	finfos, err := ioutil.ReadDir(path)
	if err == nil {
		fmt.Fprintf(w, "<ul>")
		for _, fi := range finfos {
			parts := strings.Split(fi.Name(), ".")
			if parts[len(parts)-1] == "fossil" {
				front := strings.Join(parts[0:len(parts)-1], ".")
				fmt.Fprintf(w, "<li><a href='http://fossil.ronwilson.org/%s'>%s</a></li>", front, fi.Name())
			}
		}
		fmt.Fprintf(w, "</ul>")
	} else {
		fmt.Fprintf(w, "<p>none</p>")
	}

	fmt.Fprintf(w, "</body>")
	fmt.Fprintf(w, "</html>")
}

func main() {
	var port = flag.Int("p",8081,"Port to listen on")
	var path = flag.String("r","/home/rev/fossil/","Repositories folder")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, *path)
	})
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
