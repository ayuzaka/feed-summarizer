package function_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	function "github.com/ayuzaka/feed-summarizer"
)

func TestSummarizeFeed(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	function.SummarizeFeed(rr, req)

	want := "Sent Successfully!!"
	if got := rr.Body.String(); got != want {
		t.Errorf("SummarizeFeed() = %q, want %q", got, want)
	}
}
