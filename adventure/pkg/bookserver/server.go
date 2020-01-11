package bookserver

import (
	"cor_gophercises/adventure/pkg/logger"
	"html/template"
	"net/http"
)

func StartServer() {
	svr := defaultHandler()

	logger.Info.Println("Starting server on :8080")
	http.ListenAndServe(":8080", svr)

}

func defaultHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultPageData)
	return mux
}

func defaultPageData(w http.ResponseWriter, r *http.Request) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<title>ZZZ</title>
	</head>
	<body>
		<h1>Hello World!</h1>
	</body>
</html>	
`
	t, err := template.New("home").Parse(tpl)

	if err != nil {
		logger.Error.Println("Error creating template")
	}
	t.Execute(w, nil)
}
