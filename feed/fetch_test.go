package feed

import (
	"context"
	"testing"
)

func TestFetchURLList(t *testing.T) {
	t.Parallel()

	url := "https://gist.githubusercontent.com/ayuzaka/e61d8176572eef041c41262d4c041c89/raw"
	ctx := context.TODO()

	urlList, err := FetchURLList(ctx, url)
	if err != nil {
		t.Fatal(err)
	}

	if len(urlList) == 0 {
		t.Fatal("urlList is empty")
	}
}
