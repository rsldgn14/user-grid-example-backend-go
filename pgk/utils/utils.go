package utils

import (
	"encoding/json"
	_ "encoding/json"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
	_ "net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
