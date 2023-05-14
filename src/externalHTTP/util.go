package external_http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var pr = fmt.Println

func readJSONResponse(res *http.Response, concreteType any) error {
	content := make([]byte, res.ContentLength)

	if _, err := res.Body.Read(content); err != io.EOF {
		pr(content)
		return err
	}

	pr("CONTENT!: ", string(content))

	err := json.Unmarshal(content, &concreteType)
	if err != nil {
		return err
	}

	return nil
}
