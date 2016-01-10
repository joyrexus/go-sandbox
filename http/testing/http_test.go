package http_test

import (
	"fmt"
    "net/http"
    "net/http/httptest"
    "testing"
)

func Hi(w http.ResponseWriter, r *http.Request) {
    // http.Error(w, "something failed", http.StatusInternalServerError)
	fmt.Fprint(w, "Hello world!")
}

func Test_Hi(t *testing.T) {
    req, err := http.NewRequest("GET", "http://example.com/foo", nil)
    if err != nil {
        t.Fatal(err)
    }

    rec := httptest.NewRecorder()
    Hi(rec, req)

    want, got := "Hello world!", rec.Body.String()
    if want != got {
        t.Fatalf("want %s, got %s", want, got)
    }
}
