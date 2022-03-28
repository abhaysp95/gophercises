package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
)

var defaultStoryTmpl = `
<!DOCTYPE html5>
<html>
	<head>
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
		<p>{{.}}</p>
		{{end}}
		<br>
		<ul>
			{{range .Options}}
			<li><a href="/{{.Arc}}">.Text</a></li>
			{{end}}
		</ul>
	</body>
</html>`

type Chapter struct {
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

type Story map[string]Chapter

func ReadStory(data io.Reader) (Story, error) {
	decoder := json.NewDecoder(data)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
