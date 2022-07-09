package main

import (
	"fmt"
	"link_parser"
	"regexp"
	"strings"
)

var exHtml = `<html>
<body>
<h1>Hello!</h1>
<a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {

	parsed_string := regexp.MustCompile(`[\t\r\n]+`).ReplaceAllString(exHtml, "")
	r := strings.NewReader(parsed_string)
	links, err := link_parser.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
