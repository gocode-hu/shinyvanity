package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type context struct {
	Domain string
	Path   string
	Repo   string
}

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta name="go-import" content="{{.Domain}}/{{.Path}} git {{.Repo}}">
</head>
<body>
	<p>This is a Go vanity import for <code>{{.Domain}}/{{.Path}}</code>.</p>
</body>
</html>
`

func VanityHandler(w http.ResponseWriter, r *http.Request) {
	modulePath := strings.Trim(r.URL.Path, "/")
	domain := os.Getenv("VANITY_DOMAIN")
	organization := os.Getenv("VANITY_ORGANIZATION")
	repo := fmt.Sprintf("https://github.com/%s/%s", organization, modulePath)

	// Only respond to go-get=1 requests
	if r.URL.Query().Get("go-get") != "1" {
		http.Redirect(w, r, repo, http.StatusTemporaryRedirect)
		return
	}

	tmpl, err := template.New("meta").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "template error", 500)
		return
	}

	ctx := context{
		Domain: domain,
		Path:   modulePath,
		Repo:   repo,
	}
	_ = tmpl.Execute(w, ctx)
}
