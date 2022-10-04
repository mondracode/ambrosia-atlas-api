package clients

import (
	"bytes"
	"encoding/json"
	"io"
)

func UnmarshalBody(body io.ReadCloser, v interface{}) error {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf.Bytes(), &v)
}
