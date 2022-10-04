package clients

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/mondracode/ambrosia-atlas-api/internal/apperrors"
)

func UnmarshalBody(body io.ReadCloser, v interface{}) error {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, body)
	if err != nil {
		return apperrors.NewUnexpectedAppError(err)
	}

	return json.Unmarshal(buf.Bytes(), &v)
}
