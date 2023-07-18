package feed

import (
	"sync"
	"time"

	"github.com/mmcdole/gofeed"
)

type Entry struct {
	Title       string
	URL         string
	Description string
	Published   *time.Time
}

func parseEntries(url string) ([]Entry, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)

	if err != nil {
		return nil, err
	}

	now := time.Now()

	var entries []Entry
	for _, item := range feed.Items {
		duration := now.Sub(*item.PublishedParsed)
		if duration.Hours() < 24 {
			entries = append(entries, Entry{
				Title:       item.Title,
				URL:         item.Link,
				Description: item.Description,
				Published:   item.PublishedParsed,
			})
		}

	}

	return entries, nil
}

func FindEntries(urlList []string) []Entry {
	limit := make(chan struct{}, 5)

	var wg sync.WaitGroup

	var allEntries []Entry
	for _, url := range urlList {
		wg.Add(1)
		go func(url string) {
			limit <- struct{}{}
			defer wg.Done()

			entries, _ := parseEntries(url)
			if len(entries) != 0 {
				allEntries = append(allEntries, entries...)
			}
			<-limit
		}(url)
	}

	wg.Wait()

	return allEntries
}
