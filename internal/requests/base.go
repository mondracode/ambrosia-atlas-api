package requests

import (
	"bytes"
	"encoding/json"
)

func ToJSON(s interface{}) []byte {
	j, _ := json.Marshal(s)
	return j
}

func ToJSONBuffer(s interface{}) *bytes.Buffer {
	return bytes.NewBuffer(ToJSON(s))
}
