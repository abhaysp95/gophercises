package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultStoryTmpl))
}

var tpl *template.Template

var defaultStoryTmpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
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
			<li><a href="/{{.Arc}}">{{.Text}}</a></li>
			{{end}}
		</ul>
	</body>
</html>`

func NewHandler(story Story) http.Handler {
	return handler{story}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// there's always going to be an "intro" in the story
	if err := tpl.Execute(w, h.s["intro"]); err != nil {
		panic(err)
	}
}

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
