package employee

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// validate token
	token := r.Header.Get("GovernMint-token")
	if token != "pa$$word" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(http.StatusText(http.StatusForbidden)))
		return
	}

	// validate query
	if r.URL.RawQuery == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	// split query
	param := strings.Split(r.URL.RawQuery, "=")
	str := fmt.Sprintf(`{"query": {"match": {"%s": "%s"}}}`, param[0], param[1])

	// build request
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9200/_search", bytes.NewBuffer([]byte(str)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	req.Header.Add("content-type", "application/json")

	// send request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	defer res.Body.Close()

	// read body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	// write response
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
