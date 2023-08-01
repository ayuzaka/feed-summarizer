package function

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/ayuzaka/feed-summarizer/feed"
	"github.com/ayuzaka/feed-summarizer/mail"
)

func init() {
	functions.HTTP("SummarizeFeed", SummarizeFeed)
}

func SummarizeFeed(w http.ResponseWriter, r *http.Request) {
	url := "https://gist.githubusercontent.com/ayuzaka/e61d8176572eef041c41262d4c041c89/raw"

	urlList, err := feed.FetchURLList(r.Context(), url)
	if err != nil {
		panic(err)
	}

	entries := feed.FindEntries(urlList)

	if len(entries) == 0 {
		return
	}

	subject := "Today's RSS Feed Summary"

	var body string

	for _, entry := range entries {
		body += fmt.Sprintf("%s（%s）\n%s\n-----------------\n\n", entry.Title, entry.Published.Format("2006-01-02 15:04"), entry.URL)
	}

	err = mail.SendMail(subject, body)

	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, "Sent Successfully!!")
}
