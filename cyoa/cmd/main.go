package main

import (
	"flag"
	"fmt"
	"gophercises/cyoa"
	"log"
	"net/http"
	"os"
)

func main() {
	var fileName = flag.String("file", "story.json", "provide the path to json file containing story")
	var port = flag.Int("port", 3030, "provide port to run the program")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	story, err := cyoa.ReadStory(file)

	tmpl := template.Must(template.New("").Parse(storyTmpl))
	h := cyoa.NewHandler(story,
		cyoa.WithTemplate(tmpl),
		cyoa.WithPathFunc(pathFunc))

	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Running on port :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

// pathFunc is custom path parsing function which returns chapter for json
// story
func pathFunc(req *http.Request) string {
	path := strings.TrimSpace(req.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

// custom html template (not much difference from default one, just for demo)
var storyTmpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<section class="page">
			<h1>{{.Title}}</h1>
			{{range .Paragraphs}}
			<p>{{.}}</p>
			{{end}}
			<br>
			<ul>
				{{range .Options}}
				<li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
				{{end}}
			</ul>
		</section>
		<style>
			body {
				font-family: Source Code Pro;
			}
			h1 {
				text-align: center;
				position: relative;
			}
			.page {
				width: 80%;
				max-width: 750px;
				margin: auto;
				margin-top: 40px;
				margin-bottom: 40px;
				padding: 40px;
				background: #e2ddff;
				border: 1px solid #eee;
				box-shadow: 5px 10px #333;
			}
			ul {
				border-top: 1px dotted #ccc;
				padding: 10px 0 0 0;
				-webkit-padding-start: 0;
			}
			li {
				padding: 10px;
			}
			a,
			a:visited {
				text-decoration: none;
				color: #659ef8;
			}
			a,
			a:hover {
				color: #319aff;
			}
			p {
				text-indent: 1em;
			}
		</style>
	</body>
</html>`
