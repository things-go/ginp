package yaml

import (
	"io"

	"gopkg.in/yaml.v3"

	"github.com/things-go/dyn/encoding/codec"
)

// Codec is a Codec implementation with yaml.
type Codec struct{}

// ContentType always Returns "application/x-yaml".
func (*Codec) ContentType(_ interface{}) string {
	return "application/x-yaml"
}
func (*Codec) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}
func (*Codec) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
func (*Codec) NewEncoder(w io.Writer) codec.Encoder {
	return yaml.NewEncoder(w)
}
func (*Codec) NewDecoder(r io.Reader) codec.Decoder {
	return yaml.NewDecoder(r)
}
