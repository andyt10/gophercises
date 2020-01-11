package bookserver

import (
	"cor_gophercises/adventure/pkg/logger"
	book "cor_gophercises/adventure/pkg/storybook"
	"html/template"
	"net/http"
)

const defaultPage = "intro"
const queryParam = "page"

func StartServer(bookData book.BookFile) {

	logger.Info.Println("Starting server on :8080")
	http.ListenAndServe(":8080", buildHandlers(bookData))

}

//Book Pages Stuff
var bookPageTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<title>Choose Your Own Adventure</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Story}}
			<p>{{.}}</p>
		{{end}}
		{{if .Options}}
			{{range .Options}}
				<li><a href="/book?page={{.Arc}}">{{.Text}}</a></li>
			{{end}}
		{{else}}
			<h3><a href="/">The....End(?)</a></h3>
		{{end}}
	</body>
</html>	
`

func buildBookPageHandler(bookData book.BookFile) http.HandlerFunc {

	pageData := func(w http.ResponseWriter, r *http.Request) {
		queryParam := r.URL.Query()[queryParam][0]

		if _, ok := bookData[queryParam]; ok {

			t, err := template.New("home").Parse(bookPageTemplate)

			if err != nil {
				logger.Error.Println("Error creating default template")
			}
			t.Execute(w, bookData[queryParam])

		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
		return
	}

	return pageData
}

// Default and Home Page Stuff

func buildHandlers(bookData book.BookFile) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultPageData)
	mux.HandleFunc("/book", buildBookPageHandler(bookData))
	return mux
}

func defaultPageData(w http.ResponseWriter, r *http.Request) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<title>Choose Your Own Adventure</title>
	</head>
	<body>
		<h1>Welcome To The CYOA Web Page</h1>
		<a href="/book?page=intro">Click Here to Begin!</a>
	</body>
</html>	
`
	t, err := template.New("home").Parse(tpl)

	if err != nil {
		logger.Error.Println("Error creating default template")
	}
	t.Execute(w, nil)
}
