package main

import "io/ioutil"
import "fmt"
import "log"
import "strings"
import "strconv"
import "net/http"
import "code.google.com/p/gcfg"

type Config struct {
	Server struct {
		Port int
	}
	Fossil struct {
		Path string
	}
}

func handler(w http.ResponseWriter, r *http.Request, c *Config) {
	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<header>")
	fmt.Fprintf(w, "<title>Fossil Repositories</title>")
	fmt.Fprintf(w, "</header>")
	fmt.Fprintf(w, "<body>")
	fmt.Fprintf(w, "<h1>Fossil Repositories</h1>")

	finfos, err := ioutil.ReadDir(c.Fossil.Path)
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
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "fossilhub.gcfg")
	if err != nil {
		log.Fatalf("invalid config file: %s", err)
	}
	fmt.Println(cfg)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, &cfg)
	})
	http.ListenAndServe(":"+strconv.Itoa(cfg.Server.Port), nil)
}
