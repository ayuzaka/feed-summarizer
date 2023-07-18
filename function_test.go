package function

import (
	"net/http/httptest"
	"testing"
)

func TestSummarizeFeed(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	SummarizeFeed(rr, req)

	want := "Sent Successfully!!"
	if got := rr.Body.String(); got != want {
		t.Errorf("SummarizeFeed() = %q, want %q", got, want)
	}
}
