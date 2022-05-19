package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.useragentstring.com/pages/useragentstring.php?name=All")
	if err != nil {
		panic(err)
	}

	htmlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	html := string(htmlBytes)

	patt := regexp.MustCompile(`<a href='/index.php\?id=\d+'>([^<]+)</a>`)
	l := patt.FindAllStringSubmatch(html, -1)
	var lines []string
	for _, ss := range l {
		ua := ss[1]
		lines = append(lines, fmt.Sprintf(`		%s,`, strconv.Quote(ua)))
	}

	t := `
package sdhttpua
var (
	rawUserAgents = []string{
%s
	}
)
`
	goFile := fmt.Sprintf(t, strings.Join(lines, "\n"))
	fmt.Println(goFile)
}
