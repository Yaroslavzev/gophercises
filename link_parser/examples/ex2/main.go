package main

import (
	"fmt"
	"link_parser"
	"regexp"
	"strings"
)

var exHtml = `<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
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
	_ = links
	fmt.Printf("%+v\n", links)
}
