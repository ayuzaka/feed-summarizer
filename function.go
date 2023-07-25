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

func SummarizeFeed(w http.ResponseWriter, _ *http.Request) {
	urlList := []string{
		"https://blog.cloudflare.com/rss/",
		"https://blog.jxck.io/feeds/atom.xml",
		"https://blog.langchain.dev/rss/",
		//	"https://blog.uhy.ooo/rss.xml",
		"https://coliss.com/feed/",
		"https://deno.com/feed",
		"https://devblogs.microsoft.com/typescript/feed/",
		"https://engineering.mercari.com/blog/feed.xml",
		"https://github.blog/feed/",
		"https://ishadeed.com/feed.xml",
		"https://kentcdodds.com/blog/rss.xml",
		"https://moderncss.dev/feed/",
		"https://nextjs.org/feed.xml",
		"https://storybook.js.org/blog/rss/",
		"https://svelte.dev/blog/rss.xml",
		"https://tkdodo.eu/blog/rss.xml",
		"https://web.dev/feed.xml",
		"https://www.builder.io/blog/feed.xml",
		"https://www.figma.com/blog/feed/atom.xml",
		"https://www.joshwcomeau.com/rss.xml",
		"https://www.publickey1.jp/atom.xml",
		"https://www.totaltypescript.com/rss.xml",
		"https://vercel.com/atom",
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

	err := mail.SendMail(subject, body)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, "Sent Successfully!!")
}
