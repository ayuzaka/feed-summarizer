package function

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSummarizeFeed(t *testing.T) {
	t.Skip()
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	SummarizeFeed(rr, req)

	want := "Sent Successfully!!"
	if got := rr.Body.String(); got != want {
		t.Errorf("SummarizeFeed() = %q, want %q", got, want)
	}
}
