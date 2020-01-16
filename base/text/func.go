package text

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/zyc737347123/goPractice/base/web"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// Report1 is a func
func Report1(result *web.IssuesSearchResult) {
	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(htmlTempl))

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
