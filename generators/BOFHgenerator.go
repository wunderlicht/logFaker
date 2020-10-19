package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func main() {
	const templName = "generators/BOFHexcuses.gotmpl"
	const generateName = "BOFHexcuses.go"
	var data []string

	response, err := http.Get("http://pages.cs.wisc.edu/~ballard/bofh/excuses")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		escapedStr := strings.ReplaceAll(scanner.Text(), "\"", "\\\"")
		data = append(data, escapedStr)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	f, err := os.Create(generateName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t := template.Must(template.ParseFiles(templName))
	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}
}
