package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCountHandlerSync(t *testing.T) {
	incCount := func(i int) {
		req1 := httptest.NewRequest(http.MethodGet, "/path1", nil)
		w1 := httptest.NewRecorder()

		countHandler(w1, req1)
		res := w1.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("expected error to be nil, got: %v", err)
		}

		expected := fmt.Sprintf("counts[/path1]=%d\n", i+1)

		if string(data) != expected {
			t.Fatalf("expected %s, got: %v", expected, string(data))
		}
	}

	for i := 0; i < 20; i++ {
		incCount(i)
	}
}

func TestCountHandlerAsync(t *testing.T) {
	incCount := func() {
		req1 := httptest.NewRequest(http.MethodGet, "/path1", nil)
		w1 := httptest.NewRecorder()

		countHandler(w1, req1)
		res := w1.Result()
		defer func(body io.ReadCloser) {
			_ = body.Close()
		}(res.Body)
	}

	for i := 0; i < 100; i++ {
		go incCount()
	}
}
