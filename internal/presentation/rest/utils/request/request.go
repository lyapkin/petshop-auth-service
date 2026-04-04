package request

import (
	"encoding/json"
	"io"
)

func ParseBody(body io.Reader, v any) error {
	decoder := json.NewDecoder(body)

	return decoder.Decode(v)
}
