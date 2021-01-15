package main

import (
	"bufio"
	"errors"
	"net/http"
	"os"
	"strings"
)

func getLinks() (m map[string]string, err error) {
	m = map[string]string{}
	f, err := os.Open("./links.txt")
	if err != nil {
		return
	}
	defer f.Close()

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		line := scn.Text()
		args := strings.Split(line, " ")
		if len(args) != 2 {
			return m, errors.New("invalid line format: '" + line + "'")
		}
		m[args[0]] = args[1]
	}

	err = scn.Err()
	if err != nil {
		return
	}

	return
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		links, err := getLinks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		link, ok := links[r.URL.Path]
		if !ok {
			http.Error(w, "link not found", http.StatusBadRequest)
		}
		http.Redirect(w, r, link, http.StatusSeeOther)
	})
	http.ListenAndServe(":80", nil)
}
